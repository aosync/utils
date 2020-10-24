package fmt

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Whitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

func Fmt(opts []string) int {
	var breakLine = 65

	if len(opts) > 0 {
		bl, err := strconv.Atoi(opts[0])
		breakLine = bl
		if err != nil || breakLine < 1 {
			fmt.Fprintf(os.Stderr, "fmt: incorrect argument\n")
			return 1
		}
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fmt: failed to read from stdin\n")
		return 1
	}

	text := string(buf)
	par := strings.Split(text, "\n\n")
	var lines []string

	for _, p := range par {
		p = strings.ReplaceAll(p, "\n", " ")
		for len(p) > breakLine {
			var i int
			for i = breakLine; i >= 0 && !Whitespace(rune(p[i])); i-- {
			}
			if i < 0 {
				i = breakLine
				lines = append(lines, p[:i])
				p = p[i:]
				continue
			}
			lines = append(lines, p[:i])
			p = p[i+1:]
		}
		lines = append(lines, []string{p, ""}...)
	}

	lmax := len(lines) - 1
	fmt.Println(strings.Join(lines[:lmax], "\n"))
	return 0
}
