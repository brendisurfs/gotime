// GOTIME
//  There are many like this one, but this one is my own.
/*
	TODO:
		- implement cobra
		- handle errors to keep users from putting in float values as time flags.
		- DONE: convert to 60 second count UPS rather than count downs, less anxiety inducing.

		We are aiming to PUT IN TIME, not watch it slip away.
*/

package main

import (
	"fmt"
	"time"

	timer "github.com/brendisurfs/go-pomodoro/Timer"
)

func main() {
	fmt.Println("this is a pomodoro app for your terminal.")
	time.Sleep(2 * time.Second)
	ut := timer.NewTimeframe(25, 2, 4)
	ut.BeginPomo()
}
