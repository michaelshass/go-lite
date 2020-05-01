package main

import (
	"fmt"
	"strings"
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

// NewStatementFromInput creates a statement from the given input string.
// If the StatementType is unknown, the function will return an error.
func NewStatementFromInput(text string) (stmt *Statement, err error) {
	switch {
	case strings.HasPrefix(text, string(InsertStatement)):
		stmt = &Statement{InsertStatement}

	case strings.HasPrefix(text, string(SelectStatement)):
		stmt = &Statement{InsertStatement}

	default:
		err = NewStatementCreateError(text)
	}
	return
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
// Returns an error in case the statement could not be executed
func ExecuteStatement(stmt Statement) (err error) {
	switch stmt.Type {
	case InsertStatement:
		break
	case SelectStatement:
		break
	default:
		err = NewStatementExecutionError(stmt)
	}
	return
}
