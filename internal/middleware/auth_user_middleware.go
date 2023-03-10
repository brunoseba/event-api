package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/brunoseba/event-api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		part := strings.Split(tokenString, " ")
		if len(part) != 2 || part[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Bearer token"})
		}

		token, err := jwt.Parse(part[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(utils.Secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized no found token"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized nose que paso"})
			return
		}
		enabled := claim["userType"].(bool)
		log.Printf("ESto es enabled en AuthUser: %v", enabled)

		if !claim["userType"].(bool) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Acceso denegado"})
			return
		}
		c.Set("enabled", enabled)
		c.Next()
	}
}

func UserTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		part := strings.Split(tokenString, " ")
		if len(part) != 2 || part[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Bearer token"})
		}

		token, err := jwt.Parse(part[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(utils.Secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized no found token"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized nose que paso"})
			return
		}
		enabled := claim["userType"].(bool)

		c.Set("enabled", enabled)
		c.Next()
	}
}
