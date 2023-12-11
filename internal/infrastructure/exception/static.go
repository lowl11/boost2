package exception

import (
	"fmt"
	"runtime"
	"strings"
)

func GetStackTrace() []string {
	stackTrace := make([]string, 0, 10)

	var file string
	var line int
	var callerCatch bool
	callerIterator := 0
	for {
		_, file, line, callerCatch = runtime.Caller(callerIterator)
		if !callerCatch {
			break
		}

		builder := strings.Builder{}
		builder.Grow(len(file) + 11)
		_, _ = fmt.Fprintf(&builder, "%s: line %d", file, line)
		stackTrace = append(stackTrace, builder.String())
		callerIterator++
	}

	return stackTrace
}
