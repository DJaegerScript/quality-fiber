package authentication

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service Service
}

func NewHandler(db *sql.DB) *Handler {
	service := NewService(db)

	return &Handler{
		Service: service,
	}
}

func (h *Handler) Register(ctx *fiber.Ctx) error {
	body := new(AuthenticationDTO)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "FAILED",
			"message": "Bad Request",
			"code":    fiber.StatusBadRequest,
		})
	}

	err, status := h.Service.Registration(body.Username, body.Password)
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "FAILED",
			"code":    status,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "Created",
		"code":    fiber.StatusCreated,
	})
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	body := new(AuthenticationDTO)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "FAILED",
			"message": "Bad Request",
			"code":    fiber.StatusBadRequest,
		})
	}

	err, status, token := h.Service.Login(body.Username, body.Password)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(status).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "FAILED",
			"code":    status,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "OK",
		"code":    fiber.StatusOK,
		"content": fiber.Map{
			"token": token,
		},
	})
}
