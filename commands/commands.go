package commands

import (
	"utils/commands/cat"
	"utils/commands/fmt"
	"utils/commands/mkdir"
	"utils/commands/rmdir"
	"utils/commands/yes"
)

func Execute(name string, opts []string) int {
	switch name {
	case "cat":
		return cat.Cat(opts)
	case "fmt":
		return fmt.Fmt(opts)
	case "mkdir":
		return mkdir.Mkdir(opts)
	case "rmdir":
		return rmdir.Rmdir(opts)
	case "yes":
		return yes.Yes(opts)
	case "utils":
		if len(opts) == 0 {
			return 1
		}
		return Execute(opts[0], opts[1:])
	}
	return 1
}
