package auth

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/torexanovich/savelife/backend/internal/database"
)

// Service handles authentication logic
type Service struct {
	userRepo *database.UserRepository
	jwtKey   []byte
}

// NewService creates a new authenticatxion service
func NewService(userRepo *database.UserRepository, jwtKey []byte) *Service {
	return &Service{userRepo, jwtKey}
}

// AuthenticateUser authenticates a user based on username and password
func (s *Service) AuthenticateUser(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if user.Password != password {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	signedToken, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the JWT token
func (s *Service) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
