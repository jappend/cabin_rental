package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Initializing our Environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := fiber.Config{
		AppName:      "Cabin Rental",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	app := fiber.New(config)
	app.Use(logger.New(logger.Config{
		Format: "IP+PORT: [${ip}]:${port} | STATUS: ${status} | METHOD: ${method} | PATH: ${path}\n",
	}))

	// Routes
	app.Get("/hello", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK).JSON("hello world")
		return nil
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
