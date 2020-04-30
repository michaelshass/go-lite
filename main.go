package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ExitCode type of the code indicating the exit reason.
type ExitCode int

const (
	// ExitSuccess code 0 - program excited successfully without an error.
	ExitSuccess ExitCode = iota
	// ExitReadInputError code 1 - an error occured while reading the user input.
	ExitReadInputError
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

func readInput(reader *bufio.Reader) (string, error) {

	delimiter := byte('\n')
	text, err := reader.ReadString(delimiter)

	if err != nil {
		return text, err
	}

	// Clean string before evaluation
	if index := strings.IndexByte(text, delimiter); index >= 0 {
		text = text[:index]
	}

	return text, nil
}

func isMetaCommand(text string) bool {
	return strings.HasPrefix(text, ".")
}

func handleMetaCommand(command MetaCommand) (result MetaCommandResultCode) {

	switch command {
	case MetaCommandExit:
		os.Exit(int(ExitSuccess))
	default:
		result = MetaCommandResultUnknwon
	}

	return
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		text, err := readInput(reader)

		if err != nil {
			fmt.Println(err)
			os.Exit(int(ExitReadInputError))
		}

		if isMetaCommand(text) {
			result := handleMetaCommand(MetaCommand(text))
			if result == MetaCommandResultUnknwon {
				fmt.Printf("Unrecognized Command: %s\n", text)
			}
			continue
		}
	}
}
