package main

import (
	"fmt"
	"io"
	"os"
)

func FromFile(path string, buf []byte) error {
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

func main() {
	code := 0
	var buf []byte

	opts := os.Args[1:]

	/* -u: unbuffered mode */
	if len(opts) > 0 {
		if opts[0] == "-u" {
			buf = make([]byte, 1)
			opts = opts[1:]
		} else {
			buf = make([]byte, 4096)
		}
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
			err := FromFile(a, buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: %s not found\n", a)
				code = 1
			}
		}
	}

	os.Exit(code)
}
