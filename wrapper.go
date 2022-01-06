package plugin

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

type WrappedWriter struct {
	writer http.ResponseWriter
}

func NewWrappedWriter(writer http.ResponseWriter) *WrappedWriter {
	return &WrappedWriter{
		writer: writer,
	}
}

func (w *WrappedWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *WrappedWriter) Write(buf []byte) (int, error) {
	return w.writer.Write(buf)
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
	w.writer.WriteHeader(statusCode)
}

func (w *WrappedWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.writer.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}
