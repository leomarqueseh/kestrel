package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
)

var ErrInvalidToken = errors.New("auth: invalid or expired token")

// Claims carried inside every Kestrel JWT.
type Claims struct {
	UserID    string `json:"user_id"`
	Role      Role   `json:"role"`
	TokenType string `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

type TokenIssuer struct {
	secret []byte
}

func NewTokenIssuer(secret string) *TokenIssuer {
	return &TokenIssuer{secret: []byte(secret)}
}

func (t *TokenIssuer) generate(userID string, role Role, tokenType string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID:    userID,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.secret)
}

func (t *TokenIssuer) GenerateAccessToken(userID string, role Role) (string, error) {
	return t.generate(userID, role, "access", accessTokenTTL)
}

func (t *TokenIssuer) GenerateRefreshToken(userID string, role Role) (string, error) {
	return t.generate(userID, role, "refresh", refreshTokenTTL)
}

// Parse validates a token's signature and expiry and returns its claims.
func (t *TokenIssuer) Parse(rawToken string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
