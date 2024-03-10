package domain

type LoginRequest struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
