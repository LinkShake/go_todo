package helpers

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/LinkShake/go_todo/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GenerateSid() string {
	uuidPart := uuid.New()

	// Generate a random 16-byte array
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle error if the random number generation fails
		panic(err)
	}

	// Convert the random bytes to a hexadecimal string
	randomPart := hex.EncodeToString(randomBytes)

	// Combine the UUID and the random part to form the session ID
	sid := uuidPart.String() + randomPart

	return sid
}

func CheckLoggedIn(c *fiber.Ctx) bool {
	sid := c.Cookies("sid")
	if sid == "" {
		return false
	}
	_, err := redis.GetUserId(sid)
	if err != nil {
		if err.Error() == "invalid session id" {
			return false
		}
		panic(err)
	}
	return true
}