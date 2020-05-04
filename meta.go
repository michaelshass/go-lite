package main

import (
	"bytes"
	"fmt"
	"os"
)

// MetaCommand type.
type MetaCommand string

const (
	// MetaCommandExit tells the program to stop.
	MetaCommandExit MetaCommand = MetaCommand(".exit")
)

// MetaCommandUnknownError type.
type MetaCommandUnknownError struct {
	command MetaCommand
}

func (err *MetaCommandUnknownError) Error() string {
	return fmt.Sprintf("Unkown meta command: '%s'", string(err.command))
}

// NewUnknownMetaCommandError constructs a new unknown meta command error.
func NewUnknownMetaCommandError(command MetaCommand) *MetaCommandUnknownError {
	return &MetaCommandUnknownError{
		command: command,
	}
}

// IsMetaCommand returns true if the first element equals to '.'
func IsMetaCommand(data []byte) bool {
	return bytes.HasPrefix(data, []byte{'.'})
}

// ExecuteMetaCommand executes known commands.
// Returns an error if command could not be handled.
func ExecuteMetaCommand(data []byte) {
	switch {
	case bytes.HasPrefix(data, []byte(MetaCommandExit)):
		os.Exit(int(ExitSuccess))
	default:
		panic(NewUnknownMetaCommandError(MetaCommand(data)))
	}
}
