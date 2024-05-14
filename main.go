package main

import (
	"fmt"
	"net/http"
	"testskripsi/config"
	hAuth "testskripsi/module/feature/auth/handler"
	rAuth "testskripsi/module/feature/auth/repository"
	sAuth "testskripsi/module/feature/auth/service"
	hChat "testskripsi/module/feature/chatbot/handler"
	rChat "testskripsi/module/feature/chatbot/repository"
	sChat "testskripsi/module/feature/chatbot/service"
	hFaq "testskripsi/module/feature/faq/handler"
	rFaq "testskripsi/module/feature/faq/repository"
	sFaq "testskripsi/module/feature/faq/service"
	hInvoice "testskripsi/module/feature/invoice/handler"
	rInvoice "testskripsi/module/feature/invoice/repository"
	sInvoice "testskripsi/module/feature/invoice/service"
	hPelanggan "testskripsi/module/feature/pelanggan/handler"
	rPelanggan "testskripsi/module/feature/pelanggan/repository"
	sPelanggan "testskripsi/module/feature/pelanggan/service"
	hPengaduan "testskripsi/module/feature/pengaduan/handler"
	rPengaduan "testskripsi/module/feature/pengaduan/repository"
	sPengaduan "testskripsi/module/feature/pengaduan/service"

	"testskripsi/routes"
	"testskripsi/utils"
	"testskripsi/utils/database"
	"testskripsi/utils/midtrans"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hello")
	config.InitConfig()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello gays")
	})

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	var db = database.InitDatabase()
	database.Migrate(db)
	jwtUtils := utils.NewJWT("rahasia")
	midtransUtils := midtrans.NewMidtrans()

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, jwtUtils)
	AuthHandler := hAuth.NewAuthHandler(authService, jwtUtils)

	pelangganRepo := rPelanggan.NewPelangganRepository(db)
	pelangganService := sPelanggan.NewPelangganService(pelangganRepo)
	pelangganHandler := hPelanggan.NewPelangganHandler(pelangganService)

	invoiceRepo := rInvoice.NewInvoiceRepository(db)
	InvoiceService := sInvoice.NewInvoiceService(invoiceRepo, pelangganRepo, midtransUtils)
	InvoiceHandler := hInvoice.NewInvoiceHandler(InvoiceService, midtransUtils)

	pengaduanRepo := rPengaduan.NewPengaduanRepository(db)
	pengaduanService := sPengaduan.NewPengaduanService(pengaduanRepo, pelangganRepo)
	pengaduanHandler := hPengaduan.NewPengaduanHandler(pengaduanService)

	chatRepository := rChat.NewChatbotRepository(database.ConnectMongoDB())
	chatService := sChat.NewChatService(chatRepository, pelangganRepo)
	ChatHandler := hChat.NewChatHandler(chatService)

	faqRepository := rFaq.NewFaqRepository(db)
	faqService := sFaq.NewFaqService(faqRepository)
	faqHandler := hFaq.NewFaqHandler(faqService)

	routes.RouteInvoice(e, InvoiceHandler)
	routes.RouteAuth(e, AuthHandler)
	routes.RoutePelanggan(e, pelangganHandler)
	routes.RouteChat(e, ChatHandler)
	routes.RoutePengaduan(e, pengaduanHandler)
	routes.RouteFaq(e, faqHandler)

	e.Logger.Fatal(e.Start(":8000"))

}
