package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/collection/pkg/server"
	"github.com/rodrwan/collection/services"

	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}

	api := app.Group("/api")

	handlers := server.NewServer(collectionService)

	api.Get("/records", handlers.GetRecords)
	api.Post("/records", handlers.CreateRecord)
	api.Get("/records/:id", handlers.GetRecord)
	api.Post("/records/:id/songs", handlers.AddSongToRecord)

	app.Listen(":3000")
}
