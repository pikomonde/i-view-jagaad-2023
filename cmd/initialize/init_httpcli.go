package initialize

import (
	"net/http"
	"time"
)

func NewHttpCli(
	timeout time.Duration,
) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
