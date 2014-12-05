package loop

import (
	"log"
)

// Submit job to another corroutine
func submit(loop *Loop, work *Work) {
	work.id = uuid()
	work.loop = loop
	loop.wq = append(loop.wq, work)

	go func() {
		log.Println(work.work)
		work.results = work.work.Call(work.args)
		// pop up work into done queue
		work.loop.wq = remove(work.loop.wq, work)
		work.loop.dq = append(work.loop.dq, work)
		log.Printf("Work %s finished", work.id)
	}()

}

func remove(array []*Work, item *Work) []*Work {
	index := -1
	for idx, value := range array {
		if value.id == item.id {
			index = idx
			break
		}
	}

	if index < 0 || index > len(array) {
		return array
	}
	return append(array[0:index], array[index+1:]...)
}

var uuid = func() func() int {
	id := 0
	return func() int {
		id = id + 1
		return id
	}
}()
