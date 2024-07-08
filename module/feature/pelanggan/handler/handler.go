package handler

import (
	"net/http"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"testskripsi/module/feature/pelanggan/dto"
	"testskripsi/utils"

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
		if err := utils.ValidateStruct(pelangganRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}
		value := &entities.UserModels{
			ID:             pelangganRequest.ID,
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
		idparse := c.Param("id")

		res, err := h.service.GetPelangganByID(idparse)
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

		idparse := c.Param("id")

		if err := utils.ValidateStruct(pelangganRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "data tidak boleh kosong",
			})
		}

		value := &entities.UserModels{
			ID:             idparse,
			Name:           pelangganRequest.Name,
			Status:         pelangganRequest.Status,
			Alamat:         pelangganRequest.Alamat,
			PaketLangganan: pelangganRequest.PaketLangganan,
			HargaLangganan: pelangganRequest.HargaLangganan,
		}

		res, err := h.service.UpdatePelanggan(idparse, value)
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

		res, err := h.service.DeletePelanggan(c.Param("id"))
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
func (h *PelangganHandler) GetAllDetailPelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse := c.Param("id")

		res, err := h.service.GetAllDetailPelanggan(idparse)
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

func (h *PelangganHandler) CheckIdUserByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		res, err := h.service.CheckIdUserByEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal mendapatkan id user",
				"error":   err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}
}

func (h *PelangganHandler) InsertIdUserByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		pelangganRequest := new(dto.InsertIDAkun)
		if err := c.Bind(pelangganRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		res, err := h.service.InsertIdUserByEmail(pelangganRequest.Email, pelangganRequest.IdUser)
		if res == false {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}
}
