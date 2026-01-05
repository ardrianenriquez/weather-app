package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestHandleResponse(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		body           string
		expectedOutput string
		expectedError  bool
	}{
		{
			name:       "Success - 200",
			statusCode: http.StatusOK,
			body: `{
				"weather": [{"description": "cloudy"}],
				"main": {"temp": 298.15, "humidity": 70},
				"sys": {"country": "PH"},
				"name": "Manila"
			}`,
			expectedOutput: "The current weather now in Manila, PH is cloudy. The temperature is 25.00°C and the humidity is 70\n",
			expectedError:  false,
		}, {
			name:           "Not found - 404",
			statusCode:     http.StatusNotFound,
			body:           `{"message": "city not found"}`,
			expectedOutput: "Not Found: city not found",
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock response
			mockResponse := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(bytes.NewBufferString(tt.body)),
			}

			output, err := handleResponse(mockResponse)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if output != tt.expectedOutput {
					t.Errorf("Expected output: %v but got %v", tt.expectedOutput, output)
				}
			}
		})
	}
}
