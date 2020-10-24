package mkdir

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var opts []string
var p bool = false
var mode int = (0400 | 0200 | 0100) | (040 | 020 | 010) | (04 | 02 | 01)

func parseOpts() bool {
	for {
		deferEat := false
		o := opts[0]
		if !strings.HasPrefix(o, "-") {
			break
		}
		o = o[1:]
		olast := len(o) - 1
		for j, _ := range o {
			r := o[olast-j]
			switch r {
			case 'p':
				p = true
				break
			case 'm':
				if len(opts) < 2 {
					fmt.Fprintf(os.Stderr, "mkdir: missing -m operand\n")
					return false
				}
				f, e := strconv.Atoi(opts[1])
				if e != nil {
					return false
				}
				mode = f
				deferEat = true
				break
			default:
				return false
			}
		}
		opts = opts[1:]
		if deferEat {
			opts = opts[1:]
		}
		if len(opts) == 0 {
			break
		}
	}
	return true
}

func Mkdir(optsf []string) int {
	opts = optsf
	if len(opts) == 0 || !parseOpts() {
		return 1
	}

	fmt.Println(mode)
	fmt.Println(p)
	fmt.Println(opts)
	return 0
}
