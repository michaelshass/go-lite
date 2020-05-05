package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func craeteStatements(count int) []Statement {
	statements := []Statement{}
	for i := 0; i < count; i++ {
		data := []byte(fmt.Sprintf("insert %d name mail@mail.com", i))
		statements = append(statements, *NewStatementFromInput(data))
	}
	return statements
}

func TestMaxCapacity(t *testing.T) {

	statements := craeteStatements(tableMaxRows + 1)
	table := &Table{}
	for index, stmt := range statements {
		if index < len(statements)-1 {
			ExecuteStatement(&stmt, table)
		}
	}

	func() {
		defer func() {
			err, ok := recover().(error)
			var capErr *TableCapacityError
			if !ok || !errors.As(err, &capErr) {
				t.Errorf("Did not throw capacity error '%t' '%+v'", ok, err)
			}
		}()

		ExecuteStatement(&statements[len(statements)-1], table)
	}()

}

func TestInsertStatement(t *testing.T) {

	statements := craeteStatements(22)
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
}
