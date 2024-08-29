package routes

import (
    "context"
    "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, client *mongo.Client) {
    app.Get("/db", func(c *fiber.Ctx) error {
        collection := client.Database("DatabaseName").Collection("CollectionName")
        cursor, err := collection.Find(context.TODO(), bson.D{})
        if err != nil {
            return c.Status(500).SendString("Error querying database")
        }
        var results []map[string]interface{}
        if err = cursor.All(context.TODO(), &results); err != nil {
            return c.Status(500).SendString("Error decoding database results")
        }
        return c.JSON(results)
    })
}
