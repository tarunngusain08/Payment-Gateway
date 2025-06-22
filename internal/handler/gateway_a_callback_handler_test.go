package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGatewayACallbackHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	mockCallback.EXPECT().HandleCallback(gomock.Any()).Return(&dtos.HandleCallbackResponse{Status: "success"}, nil)

	mockCache := mocks.NewMockCacheStore(ctrl)
	mockCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, false)
	mockCache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	handler := NewGatewayACallback(mockCallback, mockCache)
	reqBody := dtos.HandleCallbackRequest{
		TransactionID: "tx1",
		Status:        "success",
		GatewayRef:    "ref1",
		Amount:        100,
		Currency:      "USD",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/callbacks/gateway-a", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// config.GetConfig().Cache.TTLSeconds = 10
	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGatewayACallbackHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	mockCache := mocks.NewMockCacheStore(ctrl)
	handler := NewGatewayACallback(mockCallback, mockCache)
	req := httptest.NewRequest("POST", "/callbacks/gateway-a", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGatewayACallbackHandler_ServeHTTP_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	mockCallback.EXPECT().HandleCallback(gomock.Any()).Return(nil, errors.New("service error"))

	mockCache := mocks.NewMockCacheStore(ctrl)
	mockCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, false)

	handler := NewGatewayACallback(mockCallback, mockCache)
	reqBody := dtos.HandleCallbackRequest{
		TransactionID: "tx1",
		Status:        "success",
		GatewayRef:    "ref1",
		Amount:        100,
		Currency:      "USD",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/callbacks/gateway-a", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
