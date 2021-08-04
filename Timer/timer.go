// TIMER:
// NEED to take in the user work time, breaktime, loops.

//  NOTE: THESE ARE THE BASIC FUNCTIONS AND WILL BE MOVED TO COBRA EVENTUALLY.

/* NOTES: The crew helped me figure that the events in the beginPomo()
 might be all triggering at once. Will investigate!
THANKS CARTER, LINH, WADE! */

package timer

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/schollz/progressbar/v3"
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
	fmt.Println("*Work session starting*")

	// progress bar for work time
	bar := progressbar.Default(int64(tf.workTime))
	i := 0
	for i <= tf.workTime {
		bar.Add(1)
		time.Sleep(time.Second)
		fmt.Printf("%v minutes into work\n", i)
		i++
	}
}

// starts the break timer
func (tf *Timeframe) startBreak() {
	t := 0
	fmt.Println("its break time!")
	for t <= tf.breakTime {
		fmt.Printf("%v minutes into break\n", t)
		time.Sleep(time.Second)
		t++
	}
}

// starts the timer
func (tf *Timeframe) BeginPomo() {
	i := 0
	// for the loop amount, do these things:
	for i < tf.loopAmount {
		Clear()
		fmt.Printf("this is loop #%v out of %v\n", i+1, tf.loopAmount)
		tf.startWork()
		tf.startBreak()
		i++
	}
	fmt.Println("good work, hope you worked hard but not too much!")
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
