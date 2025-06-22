package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGatewayBCallbackService_HandleCallback_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		UpdateStatus("tx2", constants.StatusSuccess).
		Return(nil)

	svc := &GatewayBCallbackService{transactionService: mockTx}
	req := dtos.HandleCallbackRequest{
		TransactionID: "tx2",
		Status:        string(constants.StatusSuccess),
		GatewayRef:    "ref2",
		Amount:        200,
		Currency:      "EUR",
	}

	resp, err := svc.HandleCallback(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status success, got %s", resp.Status)
	}
}

func TestGatewayBCallbackService_HandleCallback_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	svc := &GatewayBCallbackService{transactionService: mockTx}
	req := dtos.HandleCallbackRequest{} // missing required fields

	_, err := svc.HandleCallback(req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestGatewayBCallbackService_HandleCallback_UpdateStatusError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		UpdateStatus("tx2", constants.StatusFailed).
		Return(errors.New("update error"))

	svc := &GatewayBCallbackService{transactionService: mockTx}
	req := dtos.HandleCallbackRequest{
		TransactionID: "tx2",
		Status:        string(constants.StatusFailed),
		GatewayRef:    "ref2",
		Amount:        200,
		Currency:      "EUR",
	}

	_, err := svc.HandleCallback(req)
	if err == nil {
		t.Fatal("expected update error, got nil")
	}
}
