package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/collection/services"
)

type Server struct {
	collectionService *services.CollectionService
}

func NewServer(collectionService *services.CollectionService) Server {
	return Server{
		collectionService: collectionService,
	}
}

func (srv Server) CreateRecord(c *fiber.Ctx) error {
	params := new(struct {
		Name string
		Kind string
	})

	c.BodyParser(&params)

	record, err := srv.collectionService.AddRecord(params.Name, params.Kind)
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
}

func (srv Server) GetRecords(c *fiber.Ctx) error {
	records, err := srv.collectionService.FindAllRecord()
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
}

func (srv Server) GetRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	record, err := srv.collectionService.FindRecord(id)
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
}

func (srv Server) AddSongToRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	record, err := srv.collectionService.FindRecord(id)
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

	if err := srv.collectionService.AddSongToRecord(record.ToRecord(), params.Name, params.Length); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ok":     true,
		"record": record,
	})
}
