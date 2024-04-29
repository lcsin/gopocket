package jwt

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestNew(t *testing.T) {
	tk, err := New([]byte("hello,world!"))
	if err != nil {
		panic(err)
	}
	fmt.Println(tk)

	tk, err = NewWithClaims([]byte("hello,world!"), map[string]interface{}{
		"username": "zhangsan",
		"role":     "admin",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(tk)
}

func TestValid(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VybmFtZSI6InpoYW5nc2FuIn0.ehbYTpFx1w9BwUbmZp9sd3x4cEux3QXfZmkrmjYJjiM"
	valid := Valid(tk, []byte("hello,world!"))
	fmt.Println(valid)
}

func TestParse(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VybmFtZSI6InpoYW5nc2FuIn0.ehbYTpFx1w9BwUbmZp9sd3x4cEux3QXfZmkrmjYJjiM"
	token, err := Parse(tk, []byte("hello,world!"))
	if err != nil {
		panic(err)
	}
	fmt.Println(token.Valid)
	fmt.Println(token.Claims.(jwt.MapClaims)["username"])
	fmt.Println(token.Claims.(jwt.MapClaims)["role"])
}
