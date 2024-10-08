package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtAuthenticator struct {
	signKey string
}

func NewAuth(key string) Authenticator {
	return &jwtAuthenticator{signKey: key}
}

type MyCustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (s *jwtAuthenticator) GenerateToken(tokenClaims *GenerateTokenClaimsOptions) (string, error) {
	mySigningKey := []byte(s.signKey)

	claims := MyCustomClaims{
		UserId: tokenClaims.AccountId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(560 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "droplet-api",
			Subject:   "client",
			Audience:  []string{"droplet"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *jwtAuthenticator) RefreshToken(tokenClaims *RefreshTokenClaimsOptions) (string, error) {
	mySigningKey := []byte(s.signKey)

	claims := MyCustomClaims{
		UserId: tokenClaims.AccountId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "droplet-api",
			Subject:   "client",
			Audience:  []string{"droplet"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *jwtAuthenticator) ParseToken(accessToken string) (*ParseTokenClaimsOutput, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims := token.Claims.(jwt.MapClaims)

	username := claims["username"]
	if username == nil {
		return nil, fmt.Errorf("token is not valid")
	}
	userId := claims["userId"]
	if userId == nil {
		return nil, fmt.Errorf("token is not valid")
	}

	return &ParseTokenClaimsOutput{AccountId: fmt.Sprint(userId)}, nil
}
