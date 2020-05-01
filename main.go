package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	linePrefix string = "db > "
	delimiter  byte   = '\n'
)

func readInput(reader *bufio.Reader) (text string, err error) {

	text, err = reader.ReadString(delimiter)

	if err != nil {
		return
	}

	// Clean string before evaluation
	if index := strings.IndexByte(text, delimiter); index >= 0 {
		text = text[:index]
	}

	return
}

func main() {

	task := func(reader *bufio.Reader, writer io.Writer) {
		defer func() {
			if recovered := recover(); recovered != nil {
				fmt.Fprintf(writer, "[Error] - %s\n", recovered)
			}
		}()

		fmt.Fprintf(writer, "db > ")

		text, err := readInput(reader)

		if err != nil {
			fmt.Fprintln(writer, err)
			os.Exit(int(ExitReadInputError))
		}

		if IsMetaCommand(text) {
			if err := ExecuteMetaCommand(MetaCommand(text)); err != nil {
				panic(err)
			} else {
				return
			}
		}

		stmt, err := NewStatementFromInput(text)

		if err != nil {
			panic(err)
		}

		if err = ExecuteStatement(*stmt); err != nil {
			panic(err)
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		task(reader, os.Stdout)
	}
}
