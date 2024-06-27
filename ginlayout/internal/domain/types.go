// Package domain
// 业务领域对象
package domain

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	UID       int64  `json:"uid"`
	UserAgent string `json:"userAgent"`
}

type Hello struct {
}
