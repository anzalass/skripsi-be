package handler

import (
	"net/http"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	service chatbot.ChatServiceInterface
}

func NewChatHandler(service chatbot.ChatServiceInterface) chatbot.ChatHandlerInterface {
	return &ChatHandler{service: service}
}

func (h *ChatHandler) CreateAnswer() echo.HandlerFunc {
	return func(c echo.Context) error {
		chat := new(entities.Chat)

		if err := c.Bind(chat); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		res, err := h.service.CreateAnswer(*chat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

func (h *ChatHandler) CreateQuestion() echo.HandlerFunc {
	return func(c echo.Context) error {
		chat := new(entities.Chat)

		if err := c.Bind(chat); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		res, err := h.service.CreateAnswer(*chat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}
