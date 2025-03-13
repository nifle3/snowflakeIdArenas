package snowflakeid

import "errors"

var (
	ErrLocalCounterOverflow = errors.New("local counter overflow")
	ErrMachineIdOverflow    = errors.New("machine id overflow")
)
