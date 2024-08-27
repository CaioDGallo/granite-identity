package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	pb "github.com/CaioDGallo/granite-identity/internal/grpc"
	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/CaioDGallo/granite-identity/internal/service"
	"github.com/google/uuid"
)

func (s *GraniteGRPCServer) GetAccountByID(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := service.NewAccountService().GetAccountByID(in.GetId())
	if err != nil {
		logger.GetLogger().Error("failed to get account by id", slog.String("error", err.Error()))
		return nil, err
	}

	// Debug prints to check time values
	fmt.Printf("CreatedAt: %v\n", account.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", account.UpdatedAt)
	fmt.Printf("LastActivity: %v\n", account.LastActivity)

	return &pb.GetAccountResponse{
		Id:            account.ID.String(),
		CreatedAt:     account.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     account.UpdatedAt.Format(time.RFC3339),
		LastActivity:  account.LastActivity.Format(time.RFC3339),
		Currency:      account.Currency,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance.String(),
		AccountType:   account.AccountType.String(),
		Status:        account.Status.String(),
		UserId:        account.UserID.String(),
	}, err
}

func (s *GraniteGRPCServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	userUUID, err := uuid.Parse(in.GetUserId())
	if err != nil {
		logger.GetLogger().Error("failed to parse user ID", slog.String("error", err.Error()))
		return nil, err
	}

	req := service.CreateAccountRequest{
		Currency:    in.GetCurrency(),
		AccountType: in.GetAccountType(),
		UserID:      userUUID,
	}
	account, err := service.NewAccountService().CreateAccount(req)
	if err != nil {
		logger.GetLogger().Error("failed to create account", slog.String("error", err.Error()))
		return nil, err
	}

	return &pb.CreateAccountResponse{
		Id:            account.ID.String(),
		CreatedAt:     account.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     account.UpdatedAt.Format(time.RFC3339),
		LastActivity:  account.LastActivity.Format(time.RFC3339),
		Currency:      account.Currency,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance.String(),
		AccountType:   account.AccountType.String(),
		Status:        account.Status.String(),
		UserId:        account.UserID.String(),
	}, err
}
