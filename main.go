package plugin

import (
	"context"
	"net/http"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type Middlware struct {
	next   http.Handler
	config *Config
	name   string
}

func (h *Middlware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := NewWrappedWriter(w)
	h.next.ServeHTTP(rw, r)
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	handler := &Middlware{
		next:   next,
		config: config,
		name:   name,
	}

	return handler, nil
}
