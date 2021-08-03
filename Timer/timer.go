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

	for tf.workTime > 0 {
		fmt.Printf("%v minutes remaining\n", tf.workTime)
		time.Sleep(time.Second)
		tf.workTime--
	}
}

// starts the break timer
func (tf *Timeframe) startBreak() {
	fmt.Println("its break time!")
	retime := &tf.breakTime
	for *retime > 0 {
		fmt.Printf("%v minutes remaining in break.\n", tf.breakTime)
		time.Sleep(time.Second)
		tf.breakTime--
	}
	fmt.Println("break time over, make it happen captain.")
}

// starts the timer
func (tf *Timeframe) BeginPomo() {

	// for the loop amount, do these things:
	for tf.loopAmount > 0 {
		tf.startWork()
		tf.startBreak()
		tf.loopAmount--
	}
	fmt.Println("good work, hope you worked hard but not too much!")
}
