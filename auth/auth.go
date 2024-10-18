package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/auth", public)
	app.Get("/auth/login", login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Get("/protected", protected)

	app.Listen(":3000")
}

func public(c *fiber.Ctx) error {

	return c.SendString("Public Route")
}
func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")

	if user != "teste" || password != "123" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid credentials")
	}

	claims := jwt.MapClaims{
		"user":  user,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"admin": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating token")
	}
	return c.JSON(fiber.Map{"token": t})
}

func protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Protected Route: " + name)
}
