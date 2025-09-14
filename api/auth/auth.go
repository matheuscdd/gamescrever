package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/matheuscdd/gamescrever/api/environment"
)

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenIns := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	env, _ := environment.LoadEnv()
	return tokenIns.SignedString(env.JWTSecretKey)
}

// TODO  melhorar validação na response
func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	tokenIns, err := jwt.Parse(tokenStr, func(tokenIns *jwt.Token) (interface{}, error) {
		if tokenIns.Method.Alg() == jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("invalid signing method")
		}

		env, _ := environment.LoadEnv()

		return env.JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenIns.Claims.(jwt.MapClaims)
	if ok && tokenIns.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Missing authentication token"},
			)
			return
		}

		const header string = "Bearer "
		if !strings.Contains(tokenStr, header) {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid authentication token"},
			)
			return
		}

		tokenStr = strings.ReplaceAll(tokenStr, header, "")
		claims, err := VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid authentication token"},
			)
			return
		}

		// TODO lógica para verificar se existe usuário

		ctx.Set("user_id", claims["user_id"].(int))
		ctx.Next()
	}
}
