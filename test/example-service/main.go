package main

import (
	"example-service/handler"
	example "example-service/proto"

	"xinhari.com/xinhari"
	"xinhari.com/xinhari/util/log"
)

func main() {
	// New Service
	service := xinhari.NewService(
		xinhari.Name("go.micro.service.example"),
		xinhari.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
