package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

func retryCondition(r *resty.Response, err error) bool {
	statusCode := r.StatusCode()
	return statusCode == http.StatusRequestTimeout || statusCode >= http.StatusInternalServerError || statusCode == http.StatusTooManyRequests
}
