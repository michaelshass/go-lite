package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type testError struct{}

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

func main() {

	reader := bufio.NewReader(os.Stdin)

	task := func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				fmt.Println(recovered)
			}
		}()

		fmt.Print("db > ")
		text, err := readInput(reader)

		if err != nil {
			fmt.Println(err)
			os.Exit(int(ExitReadInputError))
		}

		if IsMetaCommand(text) {
			if err := HandleMetaCommand(MetaCommand(text)); err != nil {
				panic(err)
			} else {
				return
			}
		}
	}

	for {
		task()
	}
}
