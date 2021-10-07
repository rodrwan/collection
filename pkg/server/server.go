package server

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/collection/services"
)

var (
	ErrServiceCannotBeNil = errors.New("service cannot be nil")
)

type Server struct {
	collectionService *services.CollectionService
}

func NewServer(collectionService *services.CollectionService) (Server, error) {
	if collectionService == nil {
		return Server{}, ErrServiceCannotBeNil
	}

	return Server{
		collectionService: collectionService,
	}, nil
}

func (srv Server) CreateRecord(c *fiber.Ctx) error {
	params := new(struct {
		Name string
		Kind string
	})

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	record, err := srv.collectionService.AddRecord(params.Name, params.Kind)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"ok":     true,
		"record": record,
	})
}

func (srv Server) GetRecords(c *fiber.Ctx) error {
	records, err := srv.collectionService.FindAllRecord()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"records": records,
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

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

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
