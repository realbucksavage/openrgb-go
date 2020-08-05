package main

import (
	"fmt"
	"github.com/realbucksavage/openrgb-go"
	"log"
	"time"
)

func main() {
	c, err := openrgb.Connect("localhost", 1337)
	if err != nil {
		log.Fatal(err)
	}

	count, err := c.GetControllerCount()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < count; i++ {
		controller := c.GetDeviceController(i)
		fmt.Println(controller)
	}

	time.Sleep(time.Second * 2)
	c.Close()
}
