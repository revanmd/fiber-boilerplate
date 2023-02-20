package routes

import (
	"api-ipms/handlers/dashboard"

	"github.com/gofiber/fiber/v2"
)

func initializeDashboardRoutes(app *fiber.App) {
	routerName := "/dashboard"
	router := app.Group(routerName)

	router.Get("/", dashboard.GetPlantResult)

}
