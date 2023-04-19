package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var fhirMock = map[string]any{
	"resourceType": "Bundle",
	"entry": []map[string]any{
		{
			"resource": map[string]any{
				"resourceType": "Patient",
			},
		},
	},
}

type spy struct {
}

func (s spy) Compile(_ map[string]any) []byte {
	ba, _ := json.Marshal(fhirMock)
	return ba
}

func (s spy) Scrape(_ []byte) map[string]any {
	return fhirMock
}

func TestServerBehaviour(t *testing.T) {
	t.Run("handles post request", func(t *testing.T) {
		spyIns := spy{}
		s := api.NewServer(spyIns, spyIns)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, _ := writer.CreateFormFile("prescription", "prescription.pdf")
		file, _ := os.Open("prescription.pdf")
		io.Copy(fw, file)
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "api/v1/fhir", bytes.NewReader(body.Bytes()))
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		got, _ := io.ReadAll(response.Body)

		want, _ := json.Marshal(fhirMock)

		assert.Equal(t, want, got)

	})
}
