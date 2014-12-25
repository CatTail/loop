package loop

// Submit job to another corroutine
func submit(loop *Loop, work *Work) {
	work.id = uuid()
	work.loop = loop
	loop.wq = append(loop.wq, work)

	go func() {
		work.results = work.work.Call(work.args)
		// pop up work into done queue
		//loop.mutex.Lock()
		work.loop.wq = remove(work.loop.wq, work)
		work.loop.dq = append(work.loop.dq, work)
		//loop.mutex.Unlock()
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
