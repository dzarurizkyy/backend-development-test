package dto

type CreateTerminalRequest struct {
	Name string `json:"name"`
}

type CreateTerminalResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
