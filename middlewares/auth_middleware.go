package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/helpers"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// MENGAMBIL TOKEN
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized due no token"})
			return
		}

		// PARSING TOKEN
		_, token, err := helpers.ParseToken(tokenString)
		// JIKA ERROR
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorExpired:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, token expired!"})
				return
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				return
			}
		}
		// JIKA TOKEN INVALID
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.Next()
	}
}