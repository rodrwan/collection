package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/collection/config"
	"github.com/rodrwan/collection/pkg/server"
	"github.com/rodrwan/collection/services"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New(config.NewFiberConfig)

	// Use middlewares for each route
	app.Use(
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}),
		helmet.New(),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}),
	)

	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}

	api := app.Group("/api")

	handlers, err := server.NewServer(collectionService)
	if err != nil {
		log.Fatal(err)
	}

	api.Get("/getRecords", handlers.GetRecords)
	api.Post("/createRecord", handlers.CreateRecord)
	api.Get("/getRecordById/:id", handlers.GetRecordById)
	api.Post("/addSongToRecordById/:id", handlers.AddSongToRecordById)

	sig := make(chan os.Signal, 1)
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		fmt.Println("Gracefully shutting down...")

		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
	// Wait for server context to be stopped
	<-serverCtx.Done()
}
