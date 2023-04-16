package api_test

import (
	"bytes"
	"fmt"
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServerBehaviour(t *testing.T) {
	t.Run("handles post request", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, _ := writer.CreateFormFile("prescription", "prescription.pdf")
		file, _ := os.Open("prescription.pdf")
		io.Copy(fw, file)
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "api/v1/fhir", bytes.NewReader(body.Bytes()))
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response := httptest.NewRecorder()

		s := api.Server{}
		fhirBundle := s.ServeHTTP(response, request)

		fmt.Print(fhirBundle)
	})
}
