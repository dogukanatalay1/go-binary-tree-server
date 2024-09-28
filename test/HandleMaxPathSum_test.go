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
	for _, tc := range MaxPathSumTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/max-path-sum", bytes.NewBufferString(tc.Input))
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

			if output.MaxPathSum != tc.ExpectedOutput {
				t.Errorf("Handler returned unexpected body: got %v want %v", output.MaxPathSum, tc.ExpectedOutput)
			}
		})
	}
}
