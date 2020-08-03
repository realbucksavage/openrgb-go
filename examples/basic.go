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

	fmt.Println("Connected to server")
	time.Sleep(time.Second * 2)
	c.Close()
}
