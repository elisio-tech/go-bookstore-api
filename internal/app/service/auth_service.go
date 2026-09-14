package service

import (
	"errors"
	"fmt"
	"time"

	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/config"
	"go-bookstore-api/internal/domain/entity"
	"go-bookstore-api/internal/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, err := s.userRepo.GetByUsername(req.Username)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("usuário já existe")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("falha ao processar senha")
	}

	user := &entity.User{
		Username:     req.Username,
		PasswordHash: string(hash),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("credenciais inválidas")
		}
		return nil, fmt.Errorf("falha ao buscar usuário")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("credenciais inválidas")
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

func generateToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Get().JWTSecret))
	if err != nil {
		return "", fmt.Errorf("falha ao gerar token")
	}
	return tokenString, nil
}