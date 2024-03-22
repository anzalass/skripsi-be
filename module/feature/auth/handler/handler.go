package handler

import (
	"net/http"
	"testskripsi/module/entities"
	"testskripsi/module/feature/auth"
	"testskripsi/module/feature/auth/dto"
	"testskripsi/utils"
	_ "time"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service auth.ServiceAuthInterface
	jwt     utils.JWTInterface
}

func NewAuthHandler(service auth.ServiceAuthInterface, jwt utils.JWTInterface) auth.HandlerAuthInterface {
	return &AuthHandler{
		service: service,
		jwt:     jwt,
	}

}

func (h *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		authRequest := new(dto.AkunRequest)
		if err := c.Bind(authRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		value := &entities.AkunModel{
			Name:     authRequest.Name,
			Email:    authRequest.Email,
			Password: authRequest.Password,
		}
		result, err := h.service.Register(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    result,
		})
	}
}

func (h *AuthHandler) LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		LoginAdminRequest := new(dto.LoginAdminRequest)
		if err := c.Bind(LoginAdminRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err,
			})
		}

		res, err := h.service.LoginAdmin(LoginAdminRequest.Email, LoginAdminRequest.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "login gagal",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"token":   res,
		})
	}
}

func (h *AuthHandler) GetUserFomCookies() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("grasinet")
		if err != nil {
			return err
		}

		res, err := h.jwt.ValidateToken(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"success": false,
				"message": err,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res.Claims,
		})
	}
}
