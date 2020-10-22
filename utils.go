package main

import "utils/commands"
import "os"
import "path"

func main() {
	exec := path.Base(os.Args[0])
	os.Exit(commands.Execute(exec, os.Args[1:]))
}
