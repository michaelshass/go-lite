package main

import (
	"bytes"
	"fmt"
)

// StatementType describing the type of a Statement.
type StatementType string

const (
	// InsertStatement type for inserting elements
	InsertStatement = StatementType("insert ")
	// SelectStatement type for selecting elements
	SelectStatement = StatementType("select ")
)

// Statement represents an instruction to the database.
type Statement struct {
	Type StatementType
	data []byte
}

// Data returns a slice without the StatementType
func (stmt *Statement) Data() []byte {
	if bytes.HasPrefix(stmt.data, []byte(stmt.Type)) {
		return stmt.data[len(stmt.Type):]
	}
	return stmt.data
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
		return &Statement{InsertStatement, data}

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
func ExecuteStatement(stmt Statement) {
	switch stmt.Type {
	case InsertStatement:
		break
	case SelectStatement:
		break
	default:
		panic(NewStatementExecutionError(stmt))
	}
}
