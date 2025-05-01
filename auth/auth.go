package auth

import (
	"errors"
	"fmt"
	"time"
	"tumdum_backend/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	jwtConfig             *config.JWTConfig
)

// Initialize sets up the auth package with configuration
func Initialize(cfg *config.JWTConfig) {
	jwtConfig = cfg
}

// Claims represents the JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken generates a new JWT token
func GenerateToken(userID uint, email string) (string, error) {
	if jwtConfig == nil || jwtConfig.Secret == "" {
		return "", fmt.Errorf("JWT configuration not initialized")
	}

	// Create claims with multiple fields
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	if jwtConfig == nil || jwtConfig.Secret == "" {
		return nil, fmt.Errorf("JWT configuration not initialized")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	// Validate the token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
