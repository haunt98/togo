package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/haunt98/togo/internal/token"
)

const (
	userIDField  = "user_id"
	expiredField = "exp"
)

var (
	TokenNotValidError = errors.New("token not valid")
)

var _ token.Generator = (*JWTGenerator)(nil)

type JWTGenerator struct {
	key string
}

func NewGenerator(key string) *JWTGenerator {
	return &JWTGenerator{
		key: key,
	}
}

// Create token from userID
func (g *JWTGenerator) CreateToken(userID string) (string, error) {
	jwtClaims := jwt.MapClaims{}
	jwtClaims[userIDField] = userID
	// TODO add expired time with custom clock
	jwtClaims[expiredField] = time.Now().Add(time.Minute * 15).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token, err := jwtToken.SignedString([]byte(g.key))
	if err != nil {
		return "", fmt.Errorf("jwt failed to signed string: %w", err)
	}
	return token, nil
}

// Validate token and return userID if valid
func (g *JWTGenerator) ValidateToken(token string) (string, error) {
	jwtClaims := make(jwt.MapClaims)
	jwtToken, err := jwt.ParseWithClaims(token, jwtClaims, func(*jwt.Token) (interface{}, error) {
		return []byte(g.key), nil
	})
	if err != nil {
		return "", fmt.Errorf("jwt failed to parse with claims: %w", err)
	}

	if !jwtToken.Valid {
		return "", TokenNotValidError
	}

	userID, ok := jwtClaims[userIDField].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse userID")
	}

	return userID, nil
}
