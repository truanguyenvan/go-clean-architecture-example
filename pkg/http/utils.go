package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

func defaultRetryCondition(r *resty.Response, err error) bool {
	statusCode := r.StatusCode()
	return statusCode == http.StatusRequestTimeout || statusCode >= http.StatusInternalServerError || statusCode == http.StatusTooManyRequests
}
