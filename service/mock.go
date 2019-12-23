package service

import (
	pb "github.com/meateam/mock-service/proto"
)

// Mock is an interface of a mock object.
type Mock interface {
	GetID() string
	SetID(id string) error

	GetMockID() string
	SetMockID(mockID string) error

	MarshalProto(mock *pb.MockObject) error
}
