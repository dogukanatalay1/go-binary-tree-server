package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "golang-binary-tree/handlers"
	types "golang-binary-tree/types"
)

func TestHandleMaxPathSum(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name: "TestCase",
			input: `{
				"tree": {
					"nodes": [
						{"id": "1", "left": "2", "right": "3", "value": 1},
						{"id": "3", "left": null, "right": null, "value": 3},
						{"id": "2", "left": null, "right": null, "value": 2}
					],
					"root": "1"
				}
			}`,
			expectedOutput: 6,
		},
		{
			name: "TestCase2",
			input: `{
				"tree": {
					"nodes": [
						{"id": "1", "left": "-10", "right": "-5", "value": 1},
						{"id": "-5", "left": "-20", "right": "-21", "value": -5},
						{"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
						{"id": "1-3", "left": null, "right": null, "value": 1},
						{"id": "100-2", "left": null, "right": null, "value": 100},
						{"id": "-20", "left": "100", "right": "2", "value": -20},
						{"id": "2", "left": null, "right": null, "value": 2},
						{"id": "100", "left": null, "right": null, "value": 100},
						{"id": "-10", "left": "30", "right": "45", "value": -10},
						{"id": "45", "left": "3", "right": "-3", "value": 45},
						{"id": "-3", "left": null, "right": null, "value": -3},
						{"id": "3", "left": null, "right": null, "value": 3},
						{"id": "30", "left": "5", "right": "1-2", "value": 30},
						{"id": "1-2", "left": null, "right": null, "value": 1},
						{"id": "5", "left": null, "right": null, "value": 5}
					],
					"root": "1"
				}
			}`,
			expectedOutput: 154,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/max-path-sum", bytes.NewBufferString(tc.input))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.HandleMaxPathSum)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			var output types.Output
			err = json.Unmarshal(rr.Body.Bytes(), &output)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if output.MaxPathSum != tc.expectedOutput {
				t.Errorf("Handler returned unexpected body: got %v want %v", output.MaxPathSum, tc.expectedOutput)
			}
		})
	}
}
