package robotgo

import (
	rg "github.com/go-vgo/robotgo"

	"fmt"
)

type Controller struct{}

func (Controller) Move(x, y int) {
	rg.Move(x, y)
}

func (Controller) Configure() {
	fmt.Println("Controller configured")
}
