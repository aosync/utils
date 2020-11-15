package nice

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

var adjustment int = 10
var opts []string

/* Only parsable argument should be in first position to avoid `nice X -n 4` */
func ParseArgs() int {
	N := opts[0]
	if strings.HasPrefix(N, "-n") {
		if len(N[2:]) == 0 {
			if len(opts) == 1 {
				fmt.Fprintf(os.Stderr, "nice:-n: missing operand\n")
				return 1
			}
			d, err := strconv.ParseUint(opts[1], 10, 8)
			if err != nil {
				fmt.Fprintf(os.Stderr, "nice:-n: invalid operand\n")
				return 1
			}
			adjustment = int(d)
			opts = opts[2:]
		} else {
			d, err := strconv.ParseUint(opts[0][2:], 10, 8)
			if err != nil {
				fmt.Fprintf(os.Stderr, "nice:-n: invalid operand\n")
				return 1
			}
			adjustment = int(d)
			opts = opts[1:]
		}
	}
	return 0
}

func Nice(o []string) int {
	opts = o
	if len(opts) > 0 {
		if d := ParseArgs(); d != 0 {
			return d
		}
	}
	pid := syscall.Getpid()
	prio, err := syscall.Getpriority(syscall.PRIO_PROCESS, pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "nice: could not get niceness\n")
		return 1
	}
	if len(opts) == 0 { /* nonstandard */
		fmt.Println(prio)
		return 0
	}
	syscall.Setpriority(syscall.PRIO_PROCESS, pid, prio+adjustment)
	cmd := exec.Command(opts[0], opts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Start()
	e := cmd.Wait()
	if exiterr, ok := e.(*exec.ExitError); ok {
		fmt.Println(exiterr)
	}
	return 0
}
