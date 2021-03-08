package httpclient

import (
	"context"
	"io"
)

type RequestPayload struct {
	Context context.Context
	Method  string
	Path    string
	Query   map[string]string
	Body    io.Reader
}
