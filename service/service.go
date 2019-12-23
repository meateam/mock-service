package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/meateam/mock-service/proto"
	ctrlr "github.com/meateam/mock-service/service/db"
	"github.com/sirupsen/logrus"
)

const ()

// Service is the structure used for handling
type Service struct {
	logger    *logrus.Logger
	grantType string
	audience  string
}

// HealthCheck checks the health of the service, and returns a boolean accordingly.
func (s *Service) HealthCheck(mongoClientPingTimeout time.Duration) bool {
	return true
}

// NewService creates a Service and returns it.
func NewService(logger *logrus.Logger) Service {
	s := Service{logger: logger}
	return s
}

// GetMockByID is the request handler for getting a mock (user, status) by file id.
func (s Service) GetMockByID(ctx context.Context, req *pb.GetMockByIDRequest) (*pb.GetMockByIDResponse, error) {
	controller := ctrlr.Controller{}
	mockID := req.GetMockID()
	if mockID == "" {
		return nil, fmt.Errorf("mockID is required")
	}

	mock, err := controller.DBGetMockByID(ctx, mockID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the mock %v", err)
	}

	return &pb.GetMockByIDResponse{MockID: mock.MockID, Mocker: mock.Mocker}, nil
}
