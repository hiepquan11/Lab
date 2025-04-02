package models

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
