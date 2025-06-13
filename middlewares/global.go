package middlewares

import (
	"apotekerBE/controllers"
	"apotekerBE/lib"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/joho/godotenv"
)

func ValidationToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Missing or invalid Authorization header",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Invalid token format",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		tok, err := jwt.ParseSigned(tokenString, []jose.SignatureAlgorithm{jose.HS256})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: fmt.Sprintf("Failed to parse token: %s", err.Error()),
			})
			ctx.Abort()
			return
		}

		_ = godotenv.Load()
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			ctx.JSON(http.StatusInternalServerError, controllers.Response{
				Success: false,
				Message: "Server error: JWT secret not configured",
			})
			ctx.Abort()
			return
		}

		var claims map[string]interface{}
		err = tok.Claims([]byte(lib.GetMD5hash(secret)), &claims)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized: invalid claims",
			})
			ctx.Abort()
			return
		}

		// Convert userId safely
		userIDRaw, ok := claims["userId"]
		if !ok {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized: userId missing in token",
			})
			ctx.Abort()
			return
		}

		userIDFloat, ok := userIDRaw.(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized: invalid userId type",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userId", int(userIDFloat))
		ctx.Next()
	}
}
