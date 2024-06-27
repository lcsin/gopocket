package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lcsin/ginlayout/internal/domain"
	"github.com/lcsin/gopocket/util/ginx"
	"github.com/spf13/viper"
)

type JwtBuilder struct {
	paths map[string]bool
}

func NewJwtBuilder() *JwtBuilder {
	return &JwtBuilder{}
}

func (j *JwtBuilder) IgnorePaths(paths ...string) *JwtBuilder {
	j.paths = make(map[string]bool, len(paths))
	for _, path := range paths {
		j.paths[path] = true
	}
	return j
}

func (j *JwtBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要登录校验
		path := c.Request.URL.Path
		if j.paths[path] {
			return
		}

		header := c.GetHeader("Authorization")
		if header == "" {
			ginx.ResponseError(c, ginx.ErrUnauthorized)
			c.Abort()
			return
		}

		segment := strings.Split(header, " ")
		if len(segment) != 2 {
			ginx.ResponseError(c, ginx.ErrUnauthorized)
			c.Abort()
			return
		}
		tokenStr := segment[1]
		var claims domain.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.Get("jwt.key").(string)), nil
		})
		if err != nil {
			ginx.ResponseError(c, ginx.ErrUnauthorized)
			c.Abort()
			return
		}
		if !token.Valid || claims.UserAgent != c.GetHeader("User-Agent") {
			ginx.ResponseError(c, ginx.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Set("uid", claims.UID)
	}
}
