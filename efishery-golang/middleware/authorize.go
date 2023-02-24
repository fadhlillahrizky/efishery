package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Authorize(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
		authData, valid := validateJwt(tokenString)

		if valid && roles != nil && authData["role"] != nil {
			authorized := containRole(roles, authData["role"].(string))
			if authorized {
				c.Set("authData", authData)
				c.Next()
				return
			}
		}

		response := defaultUnauthorizedResponse()
		response.Message = "Authorization failed."
		response.ResponseTime = fmt.Sprint(time.Since(start).Milliseconds(), " ms.")
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
	}
}

func containRole(roles []string, role string) bool {
	for i := range roles {
		if roles[i] == role {
			return true
		}
	}

	return false
}
