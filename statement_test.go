package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInsertStatement(t *testing.T) {

	var statements = func() []Statement {
		statements := []Statement{}
		for i := 0; i < tableMaxRows; i++ {
			data := []byte(fmt.Sprintf("insert %d name mail@mail.com", i))
			statements = append(statements, *NewStatementFromInput(data))
		}
		return statements
	}()

	t.Run("Test insert", func(t *testing.T) {
		table := &Table{}
		for _, stmt := range statements {
			ExecuteStatement(&stmt, table)
		}

		if table.numRows != len(statements) {
			t.Error("Not enough rows were created")
		}

		for index := range statements {
			rowData := [rowSize]byte{}
			copy(rowData[:], table.rowSlot(index))
			row := deserializeRow(rowData)

			testvalue := [idSize]byte{}
			copy(testvalue[:], []byte(fmt.Sprintf("%d", index)))

			if !bytes.Equal(row.ID[:], testvalue[:]) {
				t.Errorf(
					"Unable to create row '%d', '%v' != '%v'",
					index,
					row.ID[:],
					[]byte(fmt.Sprintf("%d", index)),
				)
			}
		}

		fmt.Printf("Succesfully created %d rows\n", table.numRows)

	})
}
