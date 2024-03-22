package handler

import (
	"net/http"
	"strconv"
	"testskripsi/module/entities"
	"testskripsi/module/feature/invoice"
	"testskripsi/module/feature/invoice/dto"
	"testskripsi/utils/midtrans"

	"github.com/labstack/echo/v4"
)

type InvoiceHandler struct {
	service  invoice.ServiceInvoice
	midtrans midtrans.MidtransServiceInterface
}

func NewInvoiceHandler(service invoice.ServiceInvoice, midtrans midtrans.MidtransServiceInterface) invoice.HandlerInvoice {
	return &InvoiceHandler{
		service:  service,
		midtrans: midtrans,
	}
}

func (h *InvoiceHandler) CreateAllInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		invoiceRequest := new(dto.CreateInvoiceRequest)
		if err := c.Bind(invoiceRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		result, err := h.service.CreateAllInvoice(invoiceRequest.Bulan, invoiceRequest.Tahun)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": result,
		})

	}
}

func (h *InvoiceHandler) GetAllInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllData()

		if err != nil {

			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data invoice",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data":    res,
			"message": "success mendapatkan data invoice",
			"success": true,
		})
	}
}

func (h *InvoiceHandler) GetInvoiceById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.GetTagihanByIdPelanggan(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data invoice",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}

func (h *InvoiceHandler) CreateTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		transaksiRequest := new(dto.TransaksiRequest)
		if err := c.Bind(transaksiRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}
		newData := &entities.TransaksiModels{
			IDPelanggan:      transaksiRequest.IDPelanggan,
			IDAkun:           transaksiRequest.IDAkun,
			Email:            transaksiRequest.Email,
			Name:             transaksiRequest.Name,
			Alamat:           transaksiRequest.Alamat,
			PaketLangganan:   transaksiRequest.PaketLangganan,
			HargaLangganan:   transaksiRequest.HargaLangganan,
			PeriodePemakaian: transaksiRequest.PeriodePemakaian,
			TotalAmount:      transaksiRequest.TotalAmount,
		}

		res, err := h.service.CreateTransaksi(newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan data snapurl",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data": res,
		})
	}
}

func (h *InvoiceHandler) AfterPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}
		if err := c.Bind(&notificationPayload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan notification payload",
			})
		}

		orderId, exists := notificationPayload["order_id"].(string)
		if !exists {
			return c.JSON(http.StatusBadRequest, map[string]any{

				"message": "gagal mendapatkan order id",
			})
		}

		success, _ := h.midtrans.VerifyPayment(c.Request().Context(), orderId)
		if success == "success" {
			_, err := h.service.ConfirmedPayment(orderId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"error":   err,
					"message": "gagal mengkonfirmasi pembayaran",
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "berhasil mengkonfirmasi pembayaran",
		})
	}
}

func (h *InvoiceHandler) GetTransaksiByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		page := c.QueryParam("page")
		pageParse, _ := strconv.ParseUint(page, 10, 64)
		res, err := h.service.GetTransaksiByEmail(email, pageParse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "berhasil mendapatkan riwayat",
			"data":    res,
		})
	}
}

func (h *InvoiceHandler) GetTransaksiByIdPembayaran() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		res, err := h.service.GetTransaksiByIdPembayaran(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "berhasil mendapatkan riwayat",
			"data":    res,
		})
	}
}
