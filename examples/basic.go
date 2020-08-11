package main

import (
	"fmt"
	"github.com/realbucksavage/openrgb-go"
	"log"
)

func main() {
	c, err := openrgb.Connect("localhost", 6742)
	if err != nil {
		log.Fatal(err)
	}

	count, err := c.GetControllerCount()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		controller, _ := c.GetDeviceController(i)

		colors := make([]openrgb.Color, len(controller.Colors))
		for i := 0; i < len(colors); i++ {
			colors[i] = openrgb.Color{0, 255, 255}
		}

		fmt.Printf("Setting color of %s to Cyan\n", controller.Name)
		if err := c.UpdateLEDs(i, colors); err != nil {
			log.Fatal(err)
		}
	}

	c.Close()
}
