package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	linePrefix string = "db > "
	delimiter  byte   = '\n'
)

func readInput(reader *bufio.Reader) []byte {
	data, err := reader.ReadBytes(delimiter)

	if err != nil {
		panic(err)
	}

	if bytes.HasSuffix(data, []byte{'\n'}) {
		return data[:len(data)-1]

	}

	return data
}

func main() {

	task := func(reader *bufio.Reader, writer io.Writer, table *Table) {
		defer func() {
			if recovered := recover(); recovered != nil {
				fmt.Fprintf(writer, "[error] - %s\n", recovered)
			}
		}()

		fmt.Fprintf(writer, "%s", linePrefix)

		data := readInput(reader)

		if IsMetaCommand(data) {
			ExecuteMetaCommand(data)
		}

		stmt := NewStatementFromInput(data)
		ExecuteStatement(stmt, table)
	}

	reader := bufio.NewReader(os.Stdin)
	table := &Table{}

	for {
		task(reader, os.Stdout, table)
	}
}
