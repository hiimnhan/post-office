package types

import (
	"net/http"
	"time"
)

type Request struct {
	Request   *http.Request
	CreatedAt time.Time
	UUID      string
}
