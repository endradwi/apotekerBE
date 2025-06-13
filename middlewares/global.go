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
		var tokenString string

		// 1. Coba ambil dari Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 {
				tokenString = parts[1]
			}
		}

		// 2. Kalau kosong, coba ambil dari cookie
		if tokenString == "" {
			cookie, err := ctx.Cookie("token")
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, controllers.Response{
					Success: false,
					Message: "Missing token",
				})
				ctx.Abort()
				return
			}
			tokenString = cookie
		}

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
