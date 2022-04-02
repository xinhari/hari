package main

import (
	"fmt"

	"xinhari.com/xinhari"
)

func main() {
	// New Service
	service := xinhari.NewService(
		xinhari.Name("go.micro.service.config-read"),
		xinhari.Version("latest"),
	)
	service.Init()

	// create a new config
	c := service.Options().Config

	// set a value
	fmt.Println("Value of key.subkey: ", c.Get("key", "subkey").String(""))
}
