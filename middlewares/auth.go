package middlewares

import (
	"encoding/json"
	jwtToken "github.com/dgrijalva/jwt-go"
	"github.com/erdoganakdeniz/models"
	"github.com/erdoganakdeniz/utils"
	"github.com/gofiber/fiber"
	jwt "github.com/gofiber/jwt"
	"log"
)
var secret string
var WithGuard func(*fiber.Ctx)

func init() {
	secret =utils.GoDotEnvVariable("JWT_SECRET")

	if secret=="" {
		panic("JWT_SECRET not provided")
	}
	WithGuard=jwt.New(jwt.Config{
		SigningKey: []byte(secret),
		ErrorHandler: jwtError,
		ContextKey: "payload",
	})
}
func WithUser(c *fiber.Ctx) {
	payload :=c.Locals("payload").(*jwtToken.Token)

	userPayload:=models.User{}
	p,err:=json.Marshal(payload.Claims.(jwtToken.MapClaims))
	if err != nil {
		log.Fatal(err)
	}
	err=json.Unmarshal(p,&userPayload)
	if err != nil {
		log.Fatal(err)
	}
	c.Locals("user",userPayload)
	c.Next()
}
func jwtError(c *fiber.Ctx,err error) {
	if err.Error() == "Missing or malformed JWT" {
		if err := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
			"data":    nil}); err != nil {
			c.Status(500).Send(err)
			return
		}
	} else {
		if err := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil}); err != nil {
			c.Status(500).Send(err)
			return
		}
	}

}