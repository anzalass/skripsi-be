package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pengaduan"
	"testskripsi/module/feature/pengaduan/dto"

	"github.com/labstack/echo/v4"
)

type PengaduanHandler struct {
	service pengaduan.ServicePengaduanInterface
}

func NewPengaduanHandler(service pengaduan.ServicePengaduanInterface) pengaduan.HandlerPengaduanInterface {
	return &PengaduanHandler{
		service: service,
	}
}

func (h *PengaduanHandler) GetPelangganByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.GetPelangganByID(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data pelanggan",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
func (h *PengaduanHandler) CreatePengaduan() echo.HandlerFunc {
	return func(c echo.Context) error {
		pengaduanReuest := new(dto.CreateSPengaduan)
		if err := c.Bind(pengaduanReuest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		file, _ := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}

		newData := &entities.PengaduanModel{
			IDPelanggan: pengaduanReuest.IDPelanggan,
			Email:       pengaduanReuest.Email,
			Nama:        pengaduanReuest.Nama,
			Deskripsi:   pengaduanReuest.Deskripsi,
			Alamat:      pengaduanReuest.Alamat,
			NoWhatsapp:  pengaduanReuest.NoWhatsapp,
		}

		res, err := h.service.CreatePengaduan(newData, file.Filename, src, pengaduanReuest.WaktuKunjungan)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal membuat pengaduan",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})

	}
}
func (h *PengaduanHandler) GetAllPengaduan() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllPengaduan()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data pengaduan",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
func (h *PengaduanHandler) EditStatusPengaduan() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		EditStatusPengaduanRequest := new(dto.EditStatusPengaduan)
		if err := c.Bind(EditStatusPengaduanRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}
		err := h.service.EditStatusPengaduan(idparse, EditStatusPengaduanRequest.Status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
		})

	}
}
func (h *PengaduanHandler) GetPengaduanByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.GetPengaduanByID(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data pengaduan",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
func (h *PengaduanHandler) GetPengaduanByEmailPelanggan() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		res, err := h.service.GetPengaduanByEmailPelanggan(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data pengaduan",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
func (h *PengaduanHandler) GetStatusPengaduan() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.GetStatusPengaduan(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data status pengaduan",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
