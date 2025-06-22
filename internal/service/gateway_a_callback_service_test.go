package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGatewayACallbackService_HandleCallback_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		UpdateStatus("tx1", constants.StatusSuccess).
		Return(nil)

	svc := NewGatewayACallbackService(mockTx)
	req := dtos.HandleCallbackRequest{
		TransactionID: "tx1",
		Status:        string(constants.StatusSuccess),
		GatewayRef:    "ref1",
		Amount:        100,
		Currency:      "USD",
	}

	resp, err := svc.HandleCallback(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status success, got %s", resp.Status)
	}
}

func TestGatewayACallbackService_HandleCallback_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	svc := NewGatewayACallbackService(mockTx)
	req := dtos.HandleCallbackRequest{} // missing required fields

	_, err := svc.HandleCallback(req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestGatewayACallbackService_HandleCallback_UpdateStatusError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		UpdateStatus("tx1", constants.StatusFailed).
		Return(errors.New("update error"))

	svc := NewGatewayACallbackService(mockTx)
	req := dtos.HandleCallbackRequest{
		TransactionID: "tx1",
		Status:        string(constants.StatusFailed),
		GatewayRef:    "ref1",
		Amount:        100,
		Currency:      "USD",
	}

	_, err := svc.HandleCallback(req)
	if err == nil {
		t.Fatal("expected update error, got nil")
	}
}
