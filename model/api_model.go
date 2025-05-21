package model

type APIResponse struct {
	Status  string      `json:"status"` // "success" or "error"
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // optional, only included for success
	Error   string      `json:"error,omitempty"` // optional, only included for error
}
