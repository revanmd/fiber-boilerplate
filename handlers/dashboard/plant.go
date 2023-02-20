package dashboard

import "github.com/gofiber/fiber/v2"

func GetPlantResult(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to Dashboard",
	})
}
