package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "github.com/meateam/mock-service/proto"
)

// Controller is the mock service business logic implementation.
type Controller struct {
}

// NewController returns a new controller.
func NewController() (Controller, error) {

	return Controller{}, nil
}

// HealthCheck runs store's healthcheck and returns true if healthy, otherwise returns false
// and any error if occured.
func (c Controller) HealthCheck(ctx context.Context) (bool, error) {
	return true, nil
}

// DBGetMockByID returns a mock.
func (c Controller) DBGetMockByID(ctx context.Context, mockID string) (pb.MockObject, error) {
	mocks := [10]string{
		"I fart in your general direction. Your mother was a hamster and your father smelt of elderberries.",
		"You know, you are a classic example of the inverse ratio between the size of the mouth and the size of the brain.",
		"To call you stupid would be an insult to stupid people! I've known sheep that could outwit you. I've worn pants with higher IQs.",
		"You bowl like your momma. Unless of course she bowls well, in which case you bowl nothing like her.",
		"In the short time we've been together, you have demonstrated every loathsome characteristic of the male personality and even discovered a few new ones. You are physically repulsive, intellectually retarded, you're morally reprehensible, vulgar, insensitive, selfish, stupid, you have no taste, a lousy sense of humor and you smell. You're not even interesting enough to make me sick.",
		"Some cause happiness wherever they go; others, whenever they go.",
		"[He] may look like an idiot and talk like an idiot but don't let that fool you. He really is an idiot.",
		"I'll explain and I'll use small words so that you'll be sure to understand, you warthog-faced buffoon.",
		"Well, I'll tell you something that should be of vital interest to you. That you, sir, are a NITWIT!",
		"You are about one bit short of a byte.",
	}

	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	mockMessage := mocks[rand.Intn(len(mocks))]
	fmt.Println(mockMessage)
	mockRes := pb.MockObject{MockID: mockID, Mocker: mockMessage}

	return mockRes, nil
}
