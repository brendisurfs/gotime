package main

import (
	"fmt"

	timer "github.com/brendisurfs/go-pomodoro/Timer"
)

func main() {
	fmt.Println("this is a pomodoro app for your terminal.")

	userTime := timer.NewTimeframe(2, 3, 4)

	userTime.BeginPomo()
}
