package rmdir

import (
	"fmt"
	"os"
	"path"
	"syscall"
	"utils/xbd"
)

var code = 0
var opts []string
var p = false

func errAt(where string, err string) {
	fmt.Fprintf(os.Stderr, "rmdir:%s: %s\n", where, err)
}

func rmdir(where string) {
	where = path.Clean(where)
	if p {
		rmdirPath(where)
	} else {
		e := syscall.Rmdir(where)
		if e != nil {
			errAt(where, e.Error())
			code = 1
		}
	}
}

func rmdirPath(where string) {
	for where != "." {
		fmt.Println(where)
		e := syscall.Rmdir(where)
		if e != nil {
			errAt(where, e.Error())
			code = 1
		}
		where = path.Dir(where)
	}
}

func Rmdir(opts []string) int {
	rs := []xbd.OptRule{xbd.OptRule{'p', false}}
	e, o := xbd.GetOpts(opts, rs)
	opts = o

	b, _ := xbd.Encountered(e, 'p')
	p = b

	if len(opts) == 0 {
		fmt.Fprintf(os.Stderr, "rmdir: missing dir operand\n")
		return 1
	}

	for _, dir := range opts {
		rmdir(dir)
	}
	return code
}
