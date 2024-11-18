package vyos

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowResponse(t *testing.T) {

    t.Parallel()

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

			// Parse the multipart form
			err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
			if err != nil {
					t.Errorf("Error parsing multipart form: %v", err)
					return
			}

			data := r.FormValue("data")
			if data == "" {
					t.Error("No data found in the request")
			}

			key := r.FormValue("key")
			if key != "test" {
					t.Errorf("Token is %v, want %v", key, "test")
			}

			// Create a response map
			response := map[string]string{
					"data": data,
					"key":  key,
			}

			// Encode the response map to JSON
			responseJSON, err := json.Marshal(response)
			if err != nil {
					t.Errorf("Error encoding response to JSON: %v", err)
					return
			}

			// Write the JSON response
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)

		}))
		defer srv.Close()

		// Define the expected data
		expectedData := map[string]interface{}{
			"op":   "show",
			"path": []interface{}{"show", "version"},
		}

		// Convert expectedData to JSON string
		expectedDataJSON, err := json.Marshal(expectedData)
		if err != nil {
				t.Errorf("Error encoding expected data to JSON: %v", err)
		}

		c := NewClient(nil).WithURL(srv.URL).WithToken("test")
		r, _, err := c.Show.Do(context.TODO(), "show version")
		if err != nil {
				t.Errorf("Show.Do returned error: %v", err)
		}

		// Compare the JSON strings
		if r.Data != string(expectedDataJSON) {
				t.Errorf("Show.Do returned %v, want %v", r.Data, string(expectedDataJSON))
		}

}