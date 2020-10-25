package cat

import (
	"fmt"
	"io"
	"os"
	"utils/xbd"
)

var code int = 0
var buf []byte

func FromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.CopyBuffer(os.Stdout, f, buf)
	if err != nil {
		return err
	}

	return nil
}

func Cat(opts []string) int {
	rs := []xbd.OptRule{xbd.OptRule{'u', false}}
	e, o := xbd.GetOpts(opts, rs)
	opts = o

	/* -u: unbuffered mode */
	if b, _ := xbd.Encountered(e, 'u'); b {
		buf = make([]byte, 1)
	} else {
		buf = make([]byte, 4096)
	}

	/* read stdin if no arg */
	if len(opts) == 0 {
		if _, err := io.CopyBuffer(os.Stdout, os.Stdin, buf); err != nil {
			fmt.Fprintf(os.Stderr, "cat: cannot read from stdin")
			code = 1
		}
	}

	for _, a := range opts {
		if a == "-" {
			if _, err := io.CopyBuffer(os.Stdout, os.Stdin, buf); err != nil {
				fmt.Fprintf(os.Stderr, "cat:-: cannot read from stdin")
				code = 1
			}
		} else {
			err := FromFile(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: %s not found\n", a)
				code = 1
			}
		}
	}

	return code
}
