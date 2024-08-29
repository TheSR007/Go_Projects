package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/contrib/websocket"
)

func SetupWebSocket(app *fiber.App) {
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
        defer c.Close()

        for {
            messageType, msg, err := c.ReadMessage()
            if err != nil {
                log.Println("Error reading message:", err)
                return
            }

            if string(msg) == "ping" {
                err := c.WriteMessage(messageType, []byte("pong"))
                if err != nil {
                    log.Println("Error writing message:", err)
                    return
                }
            }
        }
	}))
}
