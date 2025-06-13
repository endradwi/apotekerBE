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

		head := ctx.GetHeader("Authorization")

		if head == "" {
			ctx.JSON(http.StatusNotFound, controllers.Response{
				Success: false,
				Message: "Token not found",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]

		tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
		if err != nil {
			fmt.Println("Error parsing token:", err)
		}

		out := make(map[string]interface{})

		godotenv.Load()
		err = tok.Claims([]byte(lib.GetMD5hash(os.Getenv("JWT_SECRET"))), &out)

		ctx.Set("userId", int(out["userId"].(float64)))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized",
			})

			ctx.Abort()
		}

		ctx.Next()
	}
}
