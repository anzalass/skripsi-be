package routes

import (
	"testskripsi/module/feature/auth"
	"testskripsi/module/feature/chatbot"
	"testskripsi/module/feature/faq"
	"testskripsi/module/feature/invoice"
	"testskripsi/module/feature/pelanggan"
	"testskripsi/module/feature/pengaduan"

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
	e.GET("/all-pembayaran", h.GetAllPembayaran())
	// e.GET("/transaksi-bulanan", h.GetTransaksiBulanan())
}

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface) {
	e.POST("/register", h.Register())
	e.POST("/login", h.LoginAdmin())
	e.GET("/user", h.GetUserFomCookies())
}
func RouteChat(e *echo.Echo, h chatbot.ChatHandlerInterface) {
	e.POST("/answer", h.CreateAnswer())
	e.POST("/question", h.CreateQuestion())
	e.GET("/chat/:email", h.GetChatByEmail())

}
func RoutePelanggan(e *echo.Echo, h pelanggan.HandlerPelanggan) {
	e.POST("/create-pelanggan", h.CreatePelanggan())
	e.PUT("/update-pelanggan/:id", h.UpdatePelanggan())
	e.GET("/all-pelanggan", h.GetAllPelanggan())
	e.GET("/pelanggan/:id", h.GetPelangganByID())
	e.GET("/detail-pelanggan/:id", h.GetAllDetailPelanggan())
	e.DELETE("/delete-pelanggan/:id", h.DeletePelanggan())
	e.GET("/cekiduser/:email", h.CheckIdUserByEmail())
	e.POST("/insertidakun", h.InsertIdUserByEmail())
}

func RoutePengaduan(e *echo.Echo, h pengaduan.HandlerPengaduanInterface) {
	e.GET("/pengaduan", h.GetAllPengaduan())
	e.GET("/pengaduan/:id", h.GetPengaduanByID())
	e.GET("/pelanggan-pengaduan/:email", h.GetPengaduanByEmailPelanggan())
	e.GET("/pelangganpengaduan/:id", h.GetPelangganByID())
	e.PUT("/editstatuspengaduan/:id", h.EditStatusPengaduan())
	e.POST("/create-pengaduan", h.CreatePengaduan())
	e.GET("/statuspengaduan/:id", h.GetStatusPengaduan())
}

func RouteFaq(e *echo.Echo, h faq.FaqHandlerInterface) {
	e.POST("/create-faq", h.CreateFaq())
	e.PUT("/update-faq/:id", h.UpdateFaqById())
	e.GET("/faq", h.GetAllFaq())
	e.GET("/faq/:id", h.GetFaqById())
	e.DELETE("/faq/:id", h.DeleteFaqById())
	e.PUT("/increment-faq/:id", h.IncrementViewsFaq())
}
