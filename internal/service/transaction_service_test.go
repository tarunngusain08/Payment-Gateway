package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/pkg/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var workerPool = NewWorkerPool(1, 10)

func TestCreateAndProcessDeposit_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepository(ctrl)
	mockGatewayPool := mocks.NewMockGatewayPool(ctrl)
	mockGateway := mocks.NewMockPaymentGateway(ctrl)

	depositReq := &models.DepositRequest{Account: "acc1", Amount: 100}

	mockRepo.EXPECT().CreateTransaction(gomock.Any()).Return(nil)
	mockGatewayPool.EXPECT().GetRoundRobinGateway().Return(mockGateway, nil)
	mockGateway.EXPECT().ProcessDeposit(gomock.Any()).Return(nil, nil)
	mockRepo.EXPECT().UpdateTransactionStatus(gomock.Any(), constants.StatusSuccess).Return(nil)

	svc := NewTransactionService(mockRepo, mockGatewayPool, workerPool, 1*time.Second)
	_, err := svc.CreateAndProcessDeposit(depositReq)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateAndProcessDeposit_GatewayError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepository(ctrl)
	mockGatewayPool := mocks.NewMockGatewayPool(ctrl)
	mockGateway := mocks.NewMockPaymentGateway(ctrl)

	depositReq := &models.DepositRequest{Account: "acc1", Amount: 100}

	mockRepo.EXPECT().CreateTransaction(gomock.Any()).Return(nil)
	mockGatewayPool.EXPECT().GetRoundRobinGateway().Return(mockGateway, nil)
	mockGateway.EXPECT().ProcessDeposit(gomock.Any()).Return(nil, errors.New("gateway error"))
	mockRepo.EXPECT().UpdateTransactionStatus(gomock.Any(), constants.StatusFailed).Return(nil)

	svc := NewTransactionService(mockRepo, mockGatewayPool, workerPool, 1*time.Second)
	_, err := svc.CreateAndProcessDeposit(depositReq)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCreateAndProcessWithdrawal_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepository(ctrl)
	mockGatewayPool := mocks.NewMockGatewayPool(ctrl)
	mockGateway := mocks.NewMockPaymentGateway(ctrl)

	withdrawalReq := &models.WithdrawalRequest{Account: "acc2", Amount: 50}

	mockRepo.EXPECT().CreateTransaction(gomock.Any()).Return(nil)
	mockGatewayPool.EXPECT().GetRoundRobinGateway().Return(mockGateway, nil)
	mockGateway.EXPECT().ProcessWithdrawal(gomock.Any()).Return(nil, nil)
	mockRepo.EXPECT().UpdateTransactionStatus(gomock.Any(), constants.StatusSuccess).Return(nil)

	svc := NewTransactionService(mockRepo, mockGatewayPool, workerPool, 1*time.Second)
	_, err := svc.CreateAndProcessWithdrawal(withdrawalReq)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateAndProcessWithdrawal_GatewayError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepository(ctrl)
	mockGatewayPool := mocks.NewMockGatewayPool(ctrl)
	mockGateway := mocks.NewMockPaymentGateway(ctrl)

	withdrawalReq := &models.WithdrawalRequest{Account: "acc2", Amount: 50}

	mockRepo.EXPECT().CreateTransaction(gomock.Any()).Return(nil)
	mockGatewayPool.EXPECT().GetRoundRobinGateway().Return(mockGateway, nil)
	mockGateway.EXPECT().ProcessWithdrawal(gomock.Any()).Return(nil, errors.New("gateway error"))
	mockRepo.EXPECT().UpdateTransactionStatus(gomock.Any(), constants.StatusFailed).Return(nil)

	svc := NewTransactionService(mockRepo, mockGatewayPool, workerPool, 1*time.Second)
	_, err := svc.CreateAndProcessWithdrawal(withdrawalReq)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
