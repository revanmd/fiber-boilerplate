package main

import (
	"api-ipms/config"
	"api-ipms/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load environtment configuration
	config.Load()

	// Connect to Databases
	// database.Connect()

	// Initilaize Fiber
	app := fiber.New()

	// Setup fiber middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup router
	routes.SetupRoutes(app)

	// Listen to connections
	app.Listen(":" + config.Cfg.ServerPort)

}
