package main

import (
	"fmt"
	"AirMouse/desktop/internal/server"
	"AirMouse/desktop/internal/system/robotgo"
	rg "github.com/go-vgo/robotgo"
	"time"
)

func main() {
	fmt.Println("Hello, AirMouse!")

	server := server.New()
	server.Start()
	controller := robotgo.Controller{}
	controller.Configure()
	w, h := rg.GetScreenSize()
	fmt.Println("Center")
	controller.Move(w/2, h/2)
	time.Sleep(1 * time.Second)
	fmt.Println("Right top")
	controller.Move(w-10, h-10)
	time.Sleep(1 * time.Second)
	fmt.Println("Right bottom")
	controller.Move(w-10, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Left botton")
	controller.Move(10, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Left top")
	controller.Move(10, h-1)
	server.Stop()
}
