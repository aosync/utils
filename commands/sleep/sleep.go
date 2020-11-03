package sleep

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
	To consider: add nonstandard unit handling
*/

func Sleep(opts []string) int {
	if len(opts) == 0 {
		fmt.Fprintf(os.Stderr, "sleep: missing operand\n")
		return 1
	}
	d, err := strconv.ParseUint(opts[0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sleep: invalid operand\n")
		return 1
	}
	time.Sleep(time.Second * time.Duration(d))
	return 0
}
