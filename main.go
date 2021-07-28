// GO POMODORO
//  There are many like this one, but this one is my own.
/*
	TODO:
		- implement cobra
		- handle errors to keep users from putting in float values as time flags.
		- convert to 60 second count UPS rather than count downs, less anxiety inducing. We are aiming to PUT IN TIME, not watch it slip away.

	FIXME:
		- 7/27/21: loop doesnt seem to catch the Start fn on the second loop.
*/

package main

import (
	"fmt"

	timer "github.com/brendisurfs/go-pomodoro/Timer"
)

func main() {
	fmt.Println("this is a pomodoro app for your terminal.")
	// gotta figure a way to stop users from putting in floats for args HERE
	userTime := timer.NewTimeframe(1, 1, 4)

	userTime.BeginPomo()
}
