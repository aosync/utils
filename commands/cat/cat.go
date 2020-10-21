package main

import (
	"bufio"
	"fmt"
	"os"
)

func FromFile(path string, buf []byte) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		n, err := r.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	return nil
}

func FromStdin() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
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
		for {
			FromStdin()
		}
	}

	for _, a := range opts {
		if a == "-" {
			FromStdin()
		} else {
			err := FromFile(a, buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: %s not found\n", a)
				code = 1
				continue
			}
		}
	}

	os.Exit(code)
}
