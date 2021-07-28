// TIMER:
// NEED to take in the user work time, breaktime, loops.

//  NOTE: THESE ARE THE BASIC FUNCTIONS AND WILL BE MOVED TO COBRA EVENTUALLY.

package timer

import (
	"fmt"
	"time"
)

type Timeframe struct {
	workTime, breakTime, loopAmount int
}

func NewTimeframe(workTime, breakTime, loopAmount int) Timeframe {
	return Timeframe{
		workTime:   workTime,
		breakTime:  breakTime,
		loopAmount: loopAmount,
	}
}

// start func: tf Timeframe is how the func inherits this.
func (tf *Timeframe) startWork() {
	// currentTime := time.Now().Minute()
	fmt.Println("*Pomodoro session starting*")
	for tf.workTime > 0 {
		fmt.Printf("%v minutes remaining\n", tf.workTime)
		time.Sleep(time.Minute)
		tf.workTime--
	}
}

func (tf *Timeframe) startBreak() {
	fmt.Println("its break time!")

	for tf.breakTime > 0 {
		fmt.Printf("%v minutes remaining in break.\n", tf.breakTime)
		time.Sleep(time.Minute)
		tf.breakTime--
	}
	fmt.Println("break time over, make it happen captain.")
}

func (tf *Timeframe) BeginPomo() {

	// for the loop amount, do these things:

	for tf.loopAmount > 0 {
		tf.startWork()
		tf.startBreak()
		tf.loopAmount--
	}
	fmt.Println("good work, hope you worked hard but not too much!")
}
