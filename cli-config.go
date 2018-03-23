package main

import (
	"fmt"
	"github.com/denautonomepirat/dynamixel/config"
)

func main() {
	config.PromptMe()
	c := &config.Config{}
	c, err := c.GetConfig("default.yaml")
	if err != nil {
		fmt.Print(err.Error())
	}
}
