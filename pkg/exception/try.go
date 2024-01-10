package exception

func Try(action func() error) (err error) {
	defer func() {
		panicError := PanicError(recover())
		if panicError != nil {
			err = panicError
		}
	}()

	if err = action(); err != nil {
		return err
	}

	return nil
}
