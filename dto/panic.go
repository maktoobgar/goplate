package dto

type PanicResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Errors  any    `json:"errors"`
}
