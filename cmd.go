package main

import (
	"fmt"
	"os"
)

type cmdAction int

const (
	cmdActionConfig cmdAction = iota
	cmdActionFromStdin
	cmdActionFileInput
)

const help = `
Usage:
	omnirun <file>              # run a file
	echo "<file>" | omnirun -   # run a file from stdin
	omnirun config              # print config path
`

func parseArgs() cmdAction {
	if len(os.Args) != 2 {
		errExitWith("Invalid arguments." + help)
	}

	switch os.Args[1] {
	case "config":
		return cmdActionConfig
	case "-":
		return cmdActionFromStdin
	default:
		return cmdActionFileInput
	}
}

func check(err error, extraMessage string) {
	if err != nil {
		errExitWith(extraMessage + ":\n" + err.Error())
	}
}

func errExitWith(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
