package commands

import (
	"utils/commands/cat"
	"utils/commands/yes"
)

func Execute(name string, opts []string) int {
	switch name {
	case "cat":
		return cat.Cat(opts)
	case "yes":
		return yes.Yes(opts)
	}
	return 1
}
