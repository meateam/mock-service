package service

import (
	"context"

	pb "github.com/meateam/mock-service/proto"
)

// Controller is an interface for the business logic of the mock-service which uses a Store.
type Controller interface {
	DBGetMockByID(ctx context.Context, mockID string) (pb.MockObject, error)
	HealthCheck(ctx context.Context) (bool, error)
}
