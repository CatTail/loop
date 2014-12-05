package loop

import (
	"log"
	"reflect"
	"time"
)

type Work struct {
	id      int
	work    reflect.Value
	args    []reflect.Value
	done    reflect.Value
	results []reflect.Value
	loop    *Loop
}

type Loop struct {
	wq       []*Work // work queue
	dq       []*Work // done queue
	stopFlag bool
}

// Create default loop struct
func DefaultLoop() (loop *Loop) {
	log.SetFlags(log.Llongfile)

	loop = new(Loop)
	initialize(loop)
	return
}

// Start event loop
func Run(loop *Loop) {
	for !loop.stopFlag {
		var work *Work

		// WHY????
		for len(loop.dq) == 0 {
			time.Sleep(100 * time.Millisecond)
		}

		// pop up last element
		work, loop.dq = loop.dq[len(loop.dq)-1], loop.dq[:len(loop.dq)-1]
		log.Printf("Available %d, Remaining %d, Current %d",
			len(loop.dq), len(loop.wq), work.id)
		// execute the associated callback
		work.done.Call(work.results)

		if len(loop.dq) == 0 && len(loop.wq) == 0 {
			loop.stopFlag = true
		}
	}
}

// Initialize loop struct
func initialize(loop *Loop) {
	loop.wq = make([]*Work, 0)
	loop.dq = make([]*Work, 0)
	loop.stopFlag = false
}
