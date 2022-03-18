package service

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_client "github.com/rlc4u/health-check/test/mocks"
)

func TestHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name          string
		doMockHTTPSvc func(*mock_client.MockCustomHTTP)
		expectedCode  int
		expectedErr   error
	}{
		{
			name: "Test Case 1 | Good Flow",
			doMockHTTPSvc: func(mch *mock_client.MockCustomHTTP) {
				mch.EXPECT().Get(gomock.Any()).Return([]byte(`{"name":"michael","age":70,"count":233482}`), nil).AnyTimes()
			},
			expectedCode: http.StatusOK,
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedHttpClient := mock_client.NewMockCustomHTTP(mockCtrl)
			tt.doMockHTTPSvc(mockedHttpClient)

			gg := NewHealthSvc(mockedHttpClient)
			data, err := gg.HealthCheck()
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedCode, data.Status)
		})
	}
}
