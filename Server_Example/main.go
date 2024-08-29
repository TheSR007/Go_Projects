package main

import (
	"context"
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "Server_Example/routes"
)

func main() {
    ssl_state := false // Set this to false to use HTTP

    app := fiber.New()

	// Middleware
    app.Use(recover.New())

	// Setting UP index for Get and Post Request
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Server is Running."})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Server is Running."})
	})
    
	// Setting UP websocket
    SetupWebSocket(app)

	// Initialize database connection
    client := InitDB()
    defer client.Disconnect(context.Background())

	// Setting up routes with database connection
    routes.SetupRoutes(app, client)

	// Start Enet Server
	go SetupEnetServer()

    // Start TCP server
    var port string
    if ssl_state {
        port = ":443"
        log.Println("TCP Server Running on", port)
        if err := app.ListenTLS(port, "certs/cert.pem", "certs/key.pem"); err != nil {
            log.Fatal(err)
        }
    } else {
        port = ":80"
        log.Println("TCP Server Running on", port)
        if err := app.Listen(port); err != nil {
            log.Fatal(err)
        }
    }
}