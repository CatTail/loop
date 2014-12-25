package loop

import (
	"io/ioutil"
	"net/http"
	"reflect"
)

func HTTPGet(loop *Loop, options map[string]string, callback func(err error, data []byte)) {
	args := []reflect.Value{
		reflect.ValueOf(options),
	}
	done := reflect.ValueOf(callback)

	work := reflect.ValueOf(HTTPGetSync)

	submit(loop, &Work{
		work: work,
		args: args,
		done: done,
	})
}

func HTTPGetSync(options map[string]string) (err error, data []byte) {
	resp, err := http.Get(options["url"])
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	return nil, body
}
