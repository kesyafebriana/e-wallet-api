package helper

import (
	"os"
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/golang-jwt/jwt/v5"
)

type TokenHelper interface {
	CreateAndSign(userId int64, walletNumber string) (string, error)
	ParseAndVerify(signed string) (jwt.MapClaims, error)
}

type TokenImplementation struct {
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var APP_NAME string = os.Getenv("APP_NAME")

func (t *TokenImplementation) CreateAndSign(userId int64, walletNumber string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       userId,
		"wallet_number": walletNumber,
		"iat":           time.Now().Unix(),
		"iss":           APP_NAME,
		"exp":           time.Now().Add(1 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (t *TokenImplementation) ParseAndVerify(signed string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	}, jwt.WithIssuer(APP_NAME),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		return nil, apperror.StatusUnauthorized(err, constant.ExpiredTokenMsg)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, constant.ErrorUnknownClaims
	}
}
