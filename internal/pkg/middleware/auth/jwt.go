package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"go-web-demo/internal/pkg/middleware"
)

// AuthzAudience defines the value of jwt audience field.
const AuthzAudience = "go-web-demo"

// JWTStrategy defines jwt bearer authentication strategy.
type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ middleware.AuthStrategy = &JWTStrategy{}

// NewJWTStrategy create jwt bearer strategy with GinJWTMiddleware.
func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

// AuthFunc defines jwt bearer strategy as the gin authentication middleware.
func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}