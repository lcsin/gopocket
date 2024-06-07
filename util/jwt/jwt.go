package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func New(key []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString(key)
}

func NewWithClaims(key []byte, claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString(key)
}

func Valid(token string, key []byte) bool {
	if _, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}); err != nil {
		return false
	}

	return true
}

func Parse(token string, key []byte) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
}

func Claims(token string, key []byte) (map[string]interface{}, error) {
	_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	return _token.Claims.(jwt.MapClaims), nil
}
