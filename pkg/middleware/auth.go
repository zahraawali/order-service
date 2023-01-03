package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/protos/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAuthenticated(c) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func isAuthenticated(c *gin.Context) bool {
	tokenString := c.GetHeader("Authorization")
	sep := strings.Split(tokenString, " ")
	if len(sep) != 2 {
		return false
	}
	var token string = sep[1]

	user, err := db.Auth.GetUser(
		context.Background(),
		&auth.JsonWebToken{
			Jwt: token,
		},
	)

	if err != nil {
		return false
	}

	c.Set("user_id", user.GetXId())
	return true
}
