package loop

import (
	"log"
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

	work := reflect.ValueOf(func(filename string, options map[string]string) (error, []byte) {
		file, err := os.Open(filename)
		if err != nil {
			return err, nil
		}

		fileinfo, err := file.Stat()
		if err != nil {
			return err, nil
		}

		data := make([]byte, fileinfo.Size())
		_, err = file.Read(data)
		if err != nil {
			return err, nil
		}

		return nil, data
	})

	log.Println("Submit work")
	submit(loop, &Work{
		work: work,
		args: args,
		done: done,
	})
}
