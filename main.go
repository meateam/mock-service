package main

import (
	"github.com/meateam/mock-service/server"
)

func main() {
	server.NewServer(nil).Serve(nil)
}
