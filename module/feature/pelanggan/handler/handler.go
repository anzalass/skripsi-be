package handler

import (
	"net/http"
	"strconv"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"testskripsi/module/feature/pelanggan/dto"

	"github.com/labstack/echo/v4"
)

type PelangganHandler struct {
	service pelanggan.ServicePelanggan
}

func NewPelangganHandler(service pelanggan.ServicePelanggan) pelanggan.HandlerPelanggan {
	return &PelangganHandler{
		service: service,
	}
}

func (h *PelangganHandler) CreatePelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {
		pelangganRequest := new(dto.CreatePelangganRequest)
		if err := c.Bind(pelangganRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		value := &entities.UserModels{
			Name:           pelangganRequest.Name,
			Alamat:         pelangganRequest.Alamat,
			PaketLangganan: pelangganRequest.PaketLangganan,
			HargaLangganan: pelangganRequest.HargaLangganan,
		}

		res, err := h.service.CreatePelanggan(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal membuat pelanggan",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"status":  true,
			"data":    res,
		})

	}
}

func (h *PelangganHandler) GetPelangganByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res, err := h.service.GetPelangganByID(int(idparse))
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

func (h *PelangganHandler) UpdatePelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {
		pelangganRequest := new(dto.CreatePelangganRequest)
		if err := c.Bind(pelangganRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		value := &entities.UserModels{
			Name:           pelangganRequest.Name,
			Status:         pelangganRequest.Status,
			Alamat:         pelangganRequest.Alamat,
			PaketLangganan: pelangganRequest.PaketLangganan,
			HargaLangganan: pelangganRequest.HargaLangganan,
		}

		res, err := h.service.UpdatePelanggan(int(idparse), value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal update pelanggan",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"status":  true,
			"data":    res,
		})
	}
}
func (h *PelangganHandler) DeletePelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {

		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res, err := h.service.DeletePelanggan(int(idparse))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal menghapus pelanggan ",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil menghapus",
			"status":  res,
		})

	}
}
func (h *PelangganHandler) GetAllPelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllPelanggan()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal mendapatkan data pelanggan",
				"error":   err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil mendapatkan data pelanggan",
			"data":    res,
		})

	}
}