package exception

import (
	"errors"
	"fmt"
)

func PanicError(panicError any) error {
	if panicError == nil {
		return nil
	}

	if err, ok := panicError.(error); ok {
		return err
	}

	if errText, ok := panicError.(string); ok {
		return errors.New(errText)
	}

	return fmt.Errorf("%s", panicError)
}
