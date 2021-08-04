// TIMER:
// NEED to take in the user work time, breaktime, loops.

//  NOTE: THESE ARE THE BASIC FUNCTIONS AND WILL BE MOVED TO COBRA EVENTUALLY.

/* NOTES: The crew helped me figure that the events in the beginPomo()
 might be all triggering at once. Will investigate!
THANKS CARTER, LINH, WADE! */

package timer

import (
	"fmt"
	"time"
)

type Timeframe struct {
	workTime, breakTime, loopAmount int
}

// constructs the timer function
func NewTimeframe(workTime, breakTime, loopAmount int) Timeframe {
	return Timeframe{
		workTime:   workTime,
		breakTime:  breakTime,
		loopAmount: loopAmount,
	}
}

// start tf Timeframe is how the func inherits this.
func (tf *Timeframe) startWork() {
	fmt.Println("*Pomodoro session starting*")
	i := 0
	for i < tf.workTime {
		fmt.Printf("%v minutes remaining\n", tf.workTime)
		time.Sleep(time.Second)
		i++
	}
}

// starts the break timer
func (tf *Timeframe) startBreak() {
	t := 0
	fmt.Println("its break time!")
	for t < tf.breakTime {
		fmt.Printf("%v minutes into break\n", t)
		time.Sleep(time.Second)
		t++
	}
	fmt.Println("break time over, make it happen captain.")
}

// starts the timer
func (tf *Timeframe) BeginPomo() {
	i := 0
	// for the loop amount, do these things:
	for i < tf.loopAmount {
		tf.startWork()
		tf.startBreak()
		i++
	}
	fmt.Println("good work, hope you worked hard but not too much!")
}
