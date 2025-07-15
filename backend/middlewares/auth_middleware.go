package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/services"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Invalid username or password",
			})
			return
		}

		user, err := services.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}