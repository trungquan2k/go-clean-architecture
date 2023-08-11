package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) SignUpUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignUpInput{}

		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		createdUser, err := h.userUseCase.SignUpUser(ctx, &payload)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": createdUser, "error": nil})
	}
}
