package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-bookstore-api/internal/app/handlers"
	"go-bookstore-api/internal/app/middleware"
	"go-bookstore-api/internal/app/service"
	"go-bookstore-api/internal/config"
	"go-bookstore-api/internal/infrastructure/database"
	"go-bookstore-api/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Load()

	if cfg.JWTSecret == "dev-insecure-secret-change-me" {
		log.Println("WARNING: JWT_SECRET em uso com valor padrão — defina um segredo forte em produção")
	}

	gin.SetMode(gin.ReleaseMode)

	if _, err := database.InitDB(); err != nil {
		return fmt.Errorf("erro ao inicializar banco: %w", err)
	}

	bookRepo := repository.NewSQLiteBookRepository()
	userRepo := repository.NewSQLiteUserRepository()

	bookService := service.NewBookService(bookRepo)
	authService := service.NewAuthService(userRepo)

	bookHandler := handlers.NewBookHandler(bookService)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.CORS(cfg.AllowedOrigins))
	r.Use(middleware.MaxBodySize(1 << 20))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	books := r.Group("/books")
	books.Use(middleware.NewIPRateLimiter(cfg.RateLimitReq, time.Duration(cfg.RateLimitWin)*time.Second).Middleware())
	books.Use(middleware.AuthRequired())
	{
		books.GET("", bookHandler.GetBooks)
		books.GET("/:id", bookHandler.GetBookByID)
		books.POST("", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errCh := make(chan error, 1)
	go func() {
		log.Printf("servidor iniciado na porta %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("falha no servidor: %w", err)
	case sig := <-stop:
		log.Printf("sinal %s recebido, encerrando...", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("erro ao encerrar servidor: %w", err)
	}

	log.Println("servidor encerrado com sucesso")
	return nil
}