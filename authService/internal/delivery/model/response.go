package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessfullResponse struct {
	Message string `json:"message"`
}
