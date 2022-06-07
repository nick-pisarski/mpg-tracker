package fill_up_controller

import (
	"net/http"
	"testing"
)

type MockResponseWriter struct{}

func (w MockResponseWriter) Header() http.Header {
	return http.Header{}
}
func (w MockResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}
func (w MockResponseWriter) WriteHeader(statusCode int) {}

func TestMain(m *testing.M) {
	m.Run()
}

func TestIndex(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
	})
}

func TestPostFillUp(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
	})
}

func TestGetFillUpById(t *testing.T) {
	t.Run("should run", func(t *testing.T) {

	})
}
func TestGetListFillUp(t *testing.T) {
	t.Run("should run", func(t *testing.T) {

	})
}
