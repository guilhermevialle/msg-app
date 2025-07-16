package infra_services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SecretKey string
}

type ITokenService interface {
	Generate(userId string, expiresIn time.Duration) (string, error)
	Validate(tokenString string) (*Claims, error)
}

type JWTService struct {
	config Config
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

var _ ITokenService = (*JWTService)(nil)

func NewJwtTokenService(secretKey string) *JWTService {
	return &JWTService{
		config: Config{
			SecretKey: secretKey,
		},
	}
}

func (s *JWTService) Generate(userId string, expiresIn time.Duration) (string, error) {
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

func (s *JWTService) Validate(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
