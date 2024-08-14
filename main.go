package main

import (
	"log"

	"github.com/Cesta-UESC/cesta-backend/plataform/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/jombas", func(c fiber.Ctx) error {
        db, err := database.MysqlConnection()
        if err != nil {
            // Return status 500 and database connection error.
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": true,
                "msg":   err.Error(),
            })
	}

        log.Print(db.Stats())

        // Send a string response to the client
        return c.SendString("Hello, World jombas!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
