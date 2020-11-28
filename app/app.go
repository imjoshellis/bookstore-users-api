// Package app is the entry point
package app

import (
	"context"
	"log"
	"time"
	"users/data/mongodb/usersdb"

	"github.com/gofiber/fiber/v2"
)

// StartApp sets up routes and starts a fiber server
func StartApp() {
	app := fiber.New()
	MapRoutes(app)
	defer func() {
		log.Println("Disconnecting from MongoDB...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := usersdb.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	log.Fatal(app.Listen("localhost:8080"))
}
