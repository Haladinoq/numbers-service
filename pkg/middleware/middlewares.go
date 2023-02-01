package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middleware contains all middlewares used.
func Middleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(Cors())
}
