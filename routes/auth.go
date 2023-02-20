package routes

import (
	"api-ipms/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func initializeAuthRoutes(app *fiber.App) {
	routerName := "/auth"
	router := app.Group(routerName)

	router.Get("/", auth.Login)

}
