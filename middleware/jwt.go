package middleware

import (
	"gofiber-clean-architecture/configuration"
	"gofiber-clean-architecture/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(model.APIResponse{
				Status:  "error",
				Message: "Unauthorized",
				Error:   "Missing or malformed token",
			})
		}

		tokenStr := authHeader[7:]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(model.APIResponse{
				Status:  "error",
				Message: "Unauthorized",
				Error:   "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["userid"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(model.APIResponse{
				Status:  "error",
				Message: "Unauthorized",
				Error:   "Invalid token claims",
			})
		}

		// Set user_id only for this request
		c.Locals("userid", claims["userid"].(string))
		return c.Next()
	}
}

func GenerateJWT(userID, username string) (string, error) {
	jwtSecret := []byte(configuration.Get("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userID,
		"username": username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})
	return token.SignedString(jwtSecret)
}
