package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
	"github.com/summit/summit-api/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(ur *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: ur, jwtSecret: jwtSecret}
}

func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, apperror.Unauthorized("invalid email or password")
	}

	if !user.IsActive {
		return nil, apperror.Unauthorized("account is disabled")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, apperror.Unauthorized("invalid email or password")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, apperror.Internal("failed to generate token")
	}

	return &models.AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user, err := s.userRepo.Create(ctx, req.Email, string(hash), req.EmployeeID)
	if err != nil {
		return nil, apperror.Conflict("user with this email already exists")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, apperror.Internal("failed to generate token")
	}

	return &models.AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) ValidateToken(tokenString string) (int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	userID := int(claims["user_id"].(float64))
	role := claims["role"].(string)
	return userID, role, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
