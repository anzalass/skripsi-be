package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"
	"testskripsi/module/feature/chatbot/dto"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	service chatbot.ChatServiceInterface
}

func NewChatHandler(service chatbot.ChatServiceInterface) chatbot.ChatHandlerInterface {
	return &ChatHandler{service: service}
}

func (h *ChatHandler) CreateAnswer() echo.HandlerFunc {
	return func(c echo.Context) error {
		chat := new(entities.Chat)

		if err := c.Bind(chat); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		res, err := h.service.CreateAnswer(*chat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

func (h *ChatHandler) CreateQuestion() echo.HandlerFunc {
	return func(c echo.Context) error {
		chat := new(entities.Chat)

		if err := c.Bind(chat); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		res, err := h.service.CreateQuestion(*chat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}
func (h *ChatHandler) CreateDataset() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := new(dto.RequestDataset)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		var result []any
		var res any
		var err error
		for _, dto := range req.Request {
			value := &entities.DatasetAi{
				Role:    dto.Role,
				Content: dto.Content,
				Tipe:    dto.Tipe,
			}
			res, err = h.service.CreateDataset(value)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			result = append(result, res)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    result,
		})
	}
}
func (h *ChatHandler) UpdateDatasetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.ParseInt(id, 10, 64)
		req := new(entities.DatasetAi)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		value := &entities.DatasetAi{
			Content: req.Content,
		}

		res, err := h.service.UpdateDatasetById(int(idconv), value)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}
func (h *ChatHandler) DeleteDatasetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idconv, err := strconv.ParseInt(id, 10, 64)
		req := new(entities.DatasetAi)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err1 := h.service.DeleteDatasetById(int(idconv))
		if err1 != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
		})
	}
}

func (h *ChatHandler) GetAllDataset() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := h.service.GetAllDataset()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}
func (h *ChatHandler) GetDatasetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.ParseInt(id, 10, 64)
		res, res2, err := h.service.GetDatasetById(int(idconv))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
			"data2":   res2,
		})
	}
}

func (h *ChatHandler) GetChatByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")

		res, err := h.service.GetChatByEmail(email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}
func (h *ChatHandler) GetAllChat() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := h.service.GetAllChat()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DataRequest struct {
	Messages    []Message   `json:"messages"`
	Model       string      `json:"model"`
	Temperature int         `json:"temperature"`
	MaxTokens   int         `json:"max_tokens"`
	TopP        int         `json:"top_p"`
	Stream      bool        `json:"stream"`
	Stop        interface{} `json:"stop"`
}

func (h *ChatHandler) GroqAi() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := DataRequest{
			Messages: []Message{
				{Role: "system", Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet,dilarang menjawab selain pertanyaan tentang Perusahaan dan tentang keluhan pelanggan"},
				{Role: "user", Content: "siapa pendiri perusahaan grasi net"},
				{Role: "assistant", Content: "bapa khudori"},
				{Role: "user", Content: "alamat kantor grasi net"},
				{Role: "assistant", Content: "Kantor nya di perumahan puri rajeg jalan mangga blok d5 no 3, buka nya jam 08:00 WIB sampai jam 17:00 WIB"},
				{Role: "user", Content: c.FormValue("text")},
			},
			Model:       "llama3-8b-8192",
			Temperature: 1,
			MaxTokens:   500,
			TopP:        1,
			Stream:      false,
			Stop:        nil,
		}

		// Marshal data menjadi JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		// URL endpoint API
		url := "https://api.groq.com/openai/v1/chat/completions"

		// Membuat request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}

		// Mengatur header request
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer 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 ") // Ganti YOUR_API_KEY dengan token API Anda

		// Melakukan request ke API
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Membaca respons dari API
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}

		// Mengembalikan respons dari API
		return c.JSON(http.StatusOK, result)
	}
}

func (h *ChatHandler) GroqAi2() echo.HandlerFunc {
	return func(c echo.Context) error {
		chat := new(entities.Chat)

		if err := c.Bind(chat); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		res, err := h.service.GroqAi(*chat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": res,
		})

		// return c.JSON(http.StatusOK, map[string]any{
		// 	"success": true,
		// 	"data":    res,
		// })
	}
}
