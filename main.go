package main

import (
	"fmt"
	"net/http"
	hAuth "testskripsi/module/feature/auth/handler"
	rAuth "testskripsi/module/feature/auth/repository"
	sAuth "testskripsi/module/feature/auth/service"
	hInvoice "testskripsi/module/feature/invoice/handler"
	rInvoice "testskripsi/module/feature/invoice/repository"
	sInvoice "testskripsi/module/feature/invoice/service"
	hPelanggan "testskripsi/module/feature/pelanggan/handler"
	rPelanggan "testskripsi/module/feature/pelanggan/repository"
	sPelanggan "testskripsi/module/feature/pelanggan/service"

	"testskripsi/routes"
	"testskripsi/utils"
	"testskripsi/utils/database"
	"testskripsi/utils/midtrans"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hello")

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

	routes.RouteInvoice(e, InvoiceHandler)
	routes.RouteAuth(e, AuthHandler)
	routes.RoutePelanggan(e, pelangganHandler)

	e.Logger.Fatal(e.Start(":8000"))

}
