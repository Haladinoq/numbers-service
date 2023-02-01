package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors sets the middleware and allowed origins separated by comma.
func Cors() gin.HandlerFunc {
	cc := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type,Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(cc)
}
