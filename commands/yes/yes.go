package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	y := "y"

	if len(os.Args) > 1 {
		y = strings.Join(os.Args[1:], " ")
	}

	for {
		fmt.Println(y)
	}
}
