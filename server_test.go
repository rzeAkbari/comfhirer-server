package comfhirer_server_test

import (
	"bytes"
	"encoding/json"
	comfhirer_server "github.com/rzeAkbari/comfhirer-server"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type mockServer struct{}

func (m mockServer) Compile(i []byte) (comfhirer_server.Bundle, error) {
	return comfhirer_server.Bundle{Entry: []comfhirer_server.Entry{
			{comfhirer_server.FhirResource{ResourceType: "Medication"}}}},
		nil
}

func (m mockServer) Scrape(file []byte) ([]byte, error) {
	return nil, nil
}

func TestExtract(t *testing.T) {
	server := comfhirer_server.Server{
		Compiler: mockServer{},
		Scraper:  mockServer{},
	}

	t.Run("it returns resources related to prescription file", func(t *testing.T) {
		file, _ := os.ReadFile("./fixture/prescription.jpg")
		body := bytes.NewReader(file)
		request, _ := http.NewRequest(http.MethodPost, "api/v1/extract", body)
		response := httptest.NewRecorder()

		var got comfhirer_server.Bundle
		server.ServeHTTP(response, request)

		binaryGot := response.Body.Bytes()
		json.Unmarshal(binaryGot, &got)

		want := comfhirer_server.Bundle{Entry: []comfhirer_server.Entry{
			{Resource: comfhirer_server.FhirResource{ResourceType: "Medication"}},
		}}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
