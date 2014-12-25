package loop

import (
	"os"
	"reflect"
)

// Read file
func FSReadFile(loop *Loop, filename string, options map[string]string, callback func(err error, data []byte)) {
	args := []reflect.Value{
		reflect.ValueOf(filename),
		reflect.ValueOf(options),
	}
	done := reflect.ValueOf(callback)

	work := reflect.ValueOf(FSReadFileSync)

	submit(loop, &Work{
		work: work,
		args: args,
		done: done,
	})
}

func FSReadFileSync(filename string, options map[string]string) (err error, data []byte) {
	file, err := os.Open(filename)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return err, nil
	}

	data = make([]byte, fileinfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return err, nil
	}

	return nil, data
}
