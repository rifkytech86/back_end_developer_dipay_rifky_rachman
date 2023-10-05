package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

//go:generate mockery --name IJWTRSAToken
type IJWTRSAToken interface {
	GenerateToken(userID string, username string) (string, error)
	ParserToken(tokenString string) (string, string, error)
}

type CustomClaims struct {
	UserID int `json:"user_id"`
}

type jwtRSATokenRepository struct {
	publicKey  []byte
	privateKey []byte
}

func NewJWTRSAToken(privateKey []byte, publicKey []byte) IJWTRSAToken {
	return &jwtRSATokenRepository{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (j *jwtRSATokenRepository) GenerateToken(userID string, username string) (string, error) {
	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = userID
	atClaims["username"] = username
	atClaims["token_uuid"] = uuid.New().String()
	atClaims["exp"] = time.Now().Add(time.Hour * 8766).Unix()

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(string(j.privateKey))
	if err != nil {
		return "", fmt.Errorf("could not decode token private key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse token private key: %w", err)
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtRSATokenRepository) ParserToken(tokenString string) (string, string, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(string(j.publicKey))
	if err != nil {
		return "", "", fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return "", "", fmt.Errorf("validate: parse key: %w", err)
	}
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return "", "", fmt.Errorf("validate: %w", err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", "", fmt.Errorf("validate: invalid token")
	}

	userID, err := json.Marshal(claims["sub"])
	if err != nil {
		return "", "", fmt.Errorf("validate: invalid token")
	}

	username, err := json.Marshal(claims["username"])
	if err != nil {
		return "", "", fmt.Errorf("validate: invalid token")
	}

	return string(userID), string(username), nil
}
