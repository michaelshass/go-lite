package main

import (
	"fmt"
	"os"
	"strings"
)

// MetaCommand type.
type MetaCommand string

const (
	// MetaCommandExit tells the program to stop.
	MetaCommandExit MetaCommand = ".exit"
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

// IsMetaCommand returns true if the given strings starts with a '.'
func IsMetaCommand(text string) bool {
	return strings.HasPrefix(text, ".")
}

// ExecuteMetaCommand executes known commands.
// Returns an error if command could not be handled.
func ExecuteMetaCommand(command MetaCommand) (err error) {

	switch command {
	case MetaCommandExit:
		os.Exit(int(ExitSuccess))
	default:
		err = NewUnknownMetaCommandError(command)
	}
	return err
}
