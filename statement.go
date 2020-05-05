package main

import (
	"bytes"
	"errors"
	"fmt"
)

// StatementType describing the type of a Statement.
type StatementType string

const (
	// InsertStatement type for inserting elements
	InsertStatement = StatementType("insert")
	// SelectStatement type for selecting elements
	SelectStatement = StatementType("select")
)

// Statement represents an instruction to the database.
type Statement struct {
	Type StatementType
	data []byte
}

// Data returns a slice without the StatementType
func (stmt *Statement) Data() []byte {
	prefix := fmt.Sprintf("%s ", stmt.Type)
	if bytes.HasPrefix(stmt.data, []byte(prefix)) {
		return stmt.data[len(prefix):]
	}
	return stmt.data
}

// InsertRow returns a row created from the user input data.
// NOTE: Fix format for now (id, username, mail).
func (stmt *Statement) InsertRow() *Row {
	rowData := bytes.SplitAfter(
		stmt.Data(),
		[]byte{' '},
	)

	row := &Row{}
	copy(row.ID[:], rowData[0][:])
	copy(row.Username[:], rowData[1][:])
	copy(row.Mail[:], rowData[2][:])
	return row
}

// StatementCreateError indicating that an error occured
// trying to create a Statement from the given stmtString
type StatementCreateError struct {
	stmtString string
}

// NewStatementCreateError creates a new error with the given stmtString.
func NewStatementCreateError(stmtString string) *StatementCreateError {
	return &StatementCreateError{stmtString}
}

func (err *StatementCreateError) Error() string {
	return fmt.Sprintf("unable to create statement from string: '%s'", err.stmtString)
}

// NewStatementFromInput creates a statement from the given data.
// If the StatementType is unknown, the function will panic.
func NewStatementFromInput(data []byte) *Statement {
	switch {
	case bytes.HasPrefix(data, []byte(InsertStatement)):
		return &Statement{InsertStatement, data}

	case bytes.HasPrefix(data, []byte(SelectStatement)):
		return &Statement{SelectStatement, data}

	default:
		panic(NewStatementCreateError(string(data)))
	}
}

// StatementExecutionError indicating that an error occured
// trying to execute a Statement.
type StatementExecutionError struct {
	stmt Statement
}

// NewStatementExecutionError an error while executing the statement.
func NewStatementExecutionError(stmt Statement) *StatementExecutionError {
	return &StatementExecutionError{stmt}
}

func (err *StatementExecutionError) Error() string {
	return fmt.Sprintf("could not execute statement: %+v", err.stmt)
}

// ExecuteStatement tries to execute the given error.
// Panics in case the statement could not be executed
func ExecuteStatement(stmt *Statement, table *Table) {

	switch stmt.Type {
	case InsertStatement:
		if table.numRows >= tableMaxRows {
			panic(errors.New("Table is full"))
		}

		row := stmt.InsertRow()
		serializedData := serializeRow(row)
		rowSlot := table.rowSlot(table.numRows)
		copy(rowSlot, serializedData[:])
		fmt.Printf(
			"%sinserted row: (%s, %s, %s)\n",
			linePrefix,
			string(row.ID[:]),
			string(row.Username[:]),
			string(row.Mail[:]),
		)
		table.numRows++

	case SelectStatement:
		for i := 0; i < table.numRows; i++ {
			rowData := [rowSize]byte{}
			copy(rowData[:], table.rowSlot(i))
			row := deserializeRow(rowData)
			fmt.Printf(
				"%srow #%d: (%s, %s, %s)\n",
				linePrefix,
				i,
				string(row.ID[:]),
				string(row.Username[:]),
				string(row.Mail[:]),
			)
		}

	default:
		panic(NewStatementExecutionError(*stmt))
	}
}
