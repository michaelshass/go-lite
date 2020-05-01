package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		text, err := readInput(reader)

		if err != nil {
			fmt.Println(err)
			os.Exit(int(ExitReadInputError))
		}

		if IsMetaCommand(text) {
			result := HandleMetaCommand(MetaCommand(text))
			if result == MetaCommandResultUnknwon {
				fmt.Printf("Unrecognized Command: %s\n", text)
			}
			continue
		}
	}
}
