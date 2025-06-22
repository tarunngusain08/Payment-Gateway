package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/mocks"
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGatewayBCallbackHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	mockCallback.EXPECT().
		HandleCallback(gomock.Any()).
		Return(&dtos.HandleCallbackResponse{Status: "success"}, nil)

	handler := NewGatewayBCallback(mockCallback)
	reqBody := dtos.HandleCallbackRequest{
		TransactionID: "tx2",
		Status:        "success",
		GatewayRef:    "ref2",
		Amount:        200,
		Currency:      "EUR",
	}
	body, _ := xml.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/callbacks/gateway-b", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/xml")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGatewayBCallbackHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	handler := NewGatewayBCallback(mockCallback)
	req := httptest.NewRequest("POST", "/callbacks/gateway-b", bytes.NewReader([]byte("invalid xml")))
	req.Header.Set("Content-Type", "application/xml")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGatewayBCallbackHandler_ServeHTTP_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := mocks.NewMockCallback(ctrl)
	mockCallback.EXPECT().
		HandleCallback(gomock.Any()).
		Return(nil, errors.New("service error"))

	handler := NewGatewayBCallback(mockCallback)
	reqBody := dtos.HandleCallbackRequest{
		TransactionID: "tx2",
		Status:        "success",
		GatewayRef:    "ref2",
		Amount:        200,
		Currency:      "EUR",
	}
	body, _ := xml.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/callbacks/gateway-b", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/xml")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
