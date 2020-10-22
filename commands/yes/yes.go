package yes

import (
	"fmt"
	"strings"
)

func Yes(opts []string) int {
	y := "y"

	if len(opts) > 0 {
		y = strings.Join(opts, " ")
	}

	for {
		fmt.Println(y)
	}
	return 0
}
