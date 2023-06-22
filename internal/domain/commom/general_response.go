package commom

import "go-clean-architecture-example/internal/domain/enum"

type GeneralResponse struct {
	Code      enum.HTTPCode `json:"code"`
	ErrorCode interface{}   `json:"error_code"`
	Message   string        `json:"message"`
	Data      interface{}   `json:"data"`
}
