package routes

import (
	"testskripsi/module/feature/auth"
	"testskripsi/module/feature/invoice"
	"testskripsi/module/feature/pelanggan"

	"github.com/labstack/echo/v4"
)

func RouteInvoice(e *echo.Echo, h invoice.HandlerInvoice) {
	e.POST("/createallinvoice", h.CreateAllInvoice())
	e.GET("/all-tagihan", h.GetAllInvoice())
	e.GET("/tagihan/:id", h.GetInvoiceById())
	e.POST("/getsnap", h.CreateTransaksi())
	e.POST("/callback-payment", h.AfterPayment())
	e.GET("/riwayat/:email", h.GetTransaksiByEmail())
	e.GET("/riwayatbyidpembayaran/:id", h.GetTransaksiByIdPembayaran())
}

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface) {
	e.POST("/register", h.Register())
	e.POST("/login", h.LoginAdmin())
	e.GET("/user", h.GetUserFomCookies())
}
func RoutePelanggan(e *echo.Echo, h pelanggan.HandlerPelanggan) {
	e.POST("/create-pelanggan", h.CreatePelanggan())
	e.PUT("/update-pelanggan/:id", h.UpdatePelanggan())
	e.GET("/all-pelanggan", h.GetAllPelanggan())
	e.GET("/pelanggan/:id", h.GetPelangganByID())
	e.DELETE("/delete-pelanggan/:id", h.DeletePelanggan())
}
