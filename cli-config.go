package main

import (
	"fmt"
	"github.com/ArT-Programming/dynamixelBridge/config"
)

func main() {
	config.PromptMe()
	c := &config.Config{}
	c, err := c.GetConfig("default.yaml")
	if err != nil {
		fmt.Print(err.Error())
	}
}
