package dto

type AkunRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
