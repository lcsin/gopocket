package config

import "github.com/golang-jwt/jwt/v5"

type config struct {
	DB     MySQLConfig
	Redis  RedisConfig
	JWTKey string
	Port   string
}

type MySQLConfig struct {
	DNS string
}

type RedisConfig struct {
	Addr   string
	Passwd string
}

type UserClaims struct {
	jwt.RegisteredClaims
	UID       int64  `json:"uid"`
	UserAgent string `json:"userAgent"`
}
