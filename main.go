package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	linePrefix string = "db > "
	delimiter  byte   = '\n'
)

func main() {

	task := func(reader *bufio.Reader, writer io.Writer) {
		defer func() {
			if recovered := recover(); recovered != nil {
				fmt.Fprintf(writer, "[error] - %s\n", recovered)
			}
		}()

		fmt.Fprintf(writer, "%s", linePrefix)

		data, err := reader.ReadBytes(delimiter)

		if err != nil {
			panic(err)
		}

		if IsMetaCommand(data) {
			if err = ExecuteMetaCommand(data); err != nil {
				panic(err)
			} else {
				return
			}
		}

		stmt, err := NewStatementFromInput(data)

		if err != nil {
			panic(err)
		}

		if err = ExecuteStatement(*stmt); err != nil {
			panic(err)
		}

		fmt.Printf("%s[info] - executed statement: %+v\n", linePrefix, *stmt)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		task(reader, os.Stdout)
	}
}
