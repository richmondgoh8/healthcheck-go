package model

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenerateResponse(t *testing.T) {
	tests := []struct {
		name         string
		inputCode    int
		inputMsg     interface{}
		expectedCode int
		expectedMsg  interface{}
	}{
		{
			name:         "Test Case 1 | Good Flow With String",
			inputCode:    http.StatusOK,
			inputMsg:     "hello world",
			expectedCode: http.StatusOK,
			expectedMsg:  "hello world",
		},
		{
			name:         "Test Case 2 | Good Flow With Int",
			inputCode:    http.StatusOK,
			inputMsg:     1,
			expectedCode: http.StatusOK,
			expectedMsg:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			curatedData := GenerateResponse(tt.inputCode, tt.inputMsg)
			assert.Equal(t, tt.expectedCode, curatedData.Status)
			assert.Equal(t, tt.expectedMsg, curatedData.Data)
		})
	}
}
