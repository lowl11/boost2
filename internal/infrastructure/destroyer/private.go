package destroyer

import "github.com/lowl11/boost2/log"

func runDestroyFunc(destroyFunc func()) {
	// catch panic
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		// print destroyer function error
		log.Error("Run destroyer function error: ", err)
	}()

	// run destroyer function
	destroyFunc()
}
