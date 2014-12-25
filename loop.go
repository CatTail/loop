package loop

import (
	"reflect"
	"sync"
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
	mutex    sync.Mutex
}

// Create default loop struct
func DefaultLoop() (loop *Loop) {
	loop = new(Loop)
	initialize(loop)
	return
}

// Start event loop
func Run(loop *Loop) {
	if len(loop.dq) != 0 || len(loop.wq) != 0 {
		for !loop.stopFlag {
			var work *Work

			for len(loop.dq) == 0 {
				time.Sleep(1000 * time.Microsecond)
			}

			// pop up last element
			work, loop.dq = loop.dq[len(loop.dq)-1], loop.dq[:len(loop.dq)-1]
			// execute the associated callback
			work.done.Call(work.results)

			if len(loop.dq) == 0 && len(loop.wq) == 0 {
				loop.stopFlag = true
			}
		}
	}
}

// Initialize loop struct
func initialize(loop *Loop) {
	loop.wq = make([]*Work, 0)
	loop.dq = make([]*Work, 0)
	loop.stopFlag = false
}
