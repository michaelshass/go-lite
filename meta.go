package main

import (
	"os"
	"strings"
)

// MetaCommand type.
type MetaCommand string

const (
	// MetaCommandExit tells the program to stop.
	MetaCommandExit MetaCommand = ".exit"
)

// MetaCommandResultCode type.
type MetaCommandResultCode int

const (
	// MetaCommandResultSuccess indicates that the meta command was successfully executed.
	MetaCommandResultSuccess MetaCommandResultCode = iota
	// MetaCommandResultUnknwon indicates that the meta command was unkown.
	MetaCommandResultUnknwon
)

// IsMetaCommand returns true if the given strings starts with a '.'
func IsMetaCommand(text string) bool {
	return strings.HasPrefix(text, ".")
}

// HandleMetaCommand executes known commands. Returns a MetaCommandResultCode.
func HandleMetaCommand(command MetaCommand) (result MetaCommandResultCode) {

	switch command {
	case MetaCommandExit:
		os.Exit(int(ExitSuccess))
	default:
		result = MetaCommandResultUnknwon
	}

	return
}
