package gogger

import (
	"fmt"
	"runtime/debug"
)

func Trace(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w\n%s", err, debug.Stack())
}
