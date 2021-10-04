package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

	app.Post("/records", func(c *fiber.Ctx) error {
		params := new(struct {
			Name string
			Kind string
		})

		c.BodyParser(&params)

		record, err := collectionService.AddRecord(params.Name, params.Kind)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"ok":     true,
			"record": record,
		})
	})

	app.Get("/records", func(c *fiber.Ctx) error {
		records, err := collectionService.FindAllRecord()
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"ok":    false,
				"error": "Not Found",
			})
		}

		return c.JSON(fiber.Map{
			"ok":     true,
			"record": records,
		})
	})

	app.Get("/records/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		record, err := collectionService.FindRecord(id)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"ok":    false,
				"error": "Not Found",
			})
		}

		return c.JSON(fiber.Map{
			"ok":     true,
			"record": record,
		})
	})

	app.Post("/records/:id/songs", func(c *fiber.Ctx) error {
		id := c.Params("id")

		record, err := collectionService.FindRecord(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		params := new(struct {
			Name   string
			Length int64
		})

		c.BodyParser(&params)

		if err := collectionService.AddSongToRecord(record.ToRecord(), params.Name, params.Length); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"ok":     true,
			"record": record,
		})
	})

	app.Listen(":3000")
}
