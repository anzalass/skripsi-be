package handler

import (
	"net/http"
	"strconv"
	"testskripsi/module/entities"
	"testskripsi/module/feature/faq"
	"testskripsi/module/feature/faq/dto"

	"github.com/labstack/echo/v4"
)

type FaqHandler struct {
	service faq.FaqServiceInterface
}

func NewFaqHandler(service faq.FaqServiceInterface) faq.FaqHandlerInterface {
	return &FaqHandler{
		service: service,
	}
}

func (h *FaqHandler) CreateFaq() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.FaqRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}
		value := &entities.FaqModel{
			Question: req.Question,
			Answer:   req.Answer,
		}

		res, err := h.service.CreateFaq(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "gagal",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
func (h *FaqHandler) GetAllFaq() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllFaq()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "gagal",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
func (h *FaqHandler) GetFaqById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)

		res, err := h.service.GetFaqById(idparse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "data tidak ditemukan",
				"error":   err,
				"success": false,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil ditemukan",
			"data":    res,
			"success": true,
		})
	}
}
func (h *FaqHandler) UpdateFaqById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		req := new(dto.FaqRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}
		value := &entities.FaqModel{
			Question: req.Question,
			Answer:   req.Answer,
		}

		res, err := h.service.UpdateFaqById(idparse, value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "gagal",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
func (h *FaqHandler) IncrementViewsFaq() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		err := h.service.IncrementViewsFaq(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "gagal",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
		})
	}
}
func (h *FaqHandler) DeleteFaqById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		err := h.service.DeleteFaqById(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "gagal",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
		})
	}
}
