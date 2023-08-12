package stream_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aqyuki/rbtxt/stream"
)

func TestCreateStream(t *testing.T) {
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(mockResponse.StatusCode)
	}))
	defer mockServer.Close()

	response, err := stream.CreateStream(mockServer.URL)
	if err != nil {
		t.Errorf("An error has occurred. %+v", err)
	}
	if response.StatusCode != mockResponse.StatusCode {
		t.Errorf("Expected status code %d, but got %d", mockResponse.StatusCode, response.StatusCode)
	}
}
