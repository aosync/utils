package mkdir

import (
	"fmt"
	"strconv"
	// "strings"
	"os"
	"utils/xbd"
)

var opts []string
var p bool = false
var mode uint = 0755

func Mkdir(opts []string) int {
	rs := []xbd.OptRule{xbd.OptRule{'p', false}, xbd.OptRule{'m', true}}
	e, o := xbd.GetOpts(opts, rs)
	opts = o

	b, _ := xbd.Encountered(e, 'p')
	p = b
	if b, v := xbd.Encountered(e, 'm'); b {
		d, err := strconv.ParseUint(v, 8, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: invalid mode\n")
			return 1
		}
		mode = uint(d)
	}

	fmt.Println(mode)
	fmt.Println(p)
	return 0
}
