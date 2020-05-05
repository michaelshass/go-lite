package main

import "fmt"

const (
	pageSize      = 4096
	rowsPerPage   = pageSize / rowSize
	tableMaxPages = 100
	tableMaxRows  = rowsPerPage * tableMaxPages
)

// Table holds data in pages 4096 in size.
// Number of total pages is limited to const 'tableMaxPages'
type Table struct {
	numRows int
	pages   [tableMaxPages][]byte
}

func (table *Table) rowSlot(rowIndex int) []byte {

	pageIndex := rowIndex / rowsPerPage
	page := table.pages[pageIndex]
	if page == nil {
		table.pages[pageIndex] = make([]byte, pageSize)
		page = table.pages[pageIndex]
	}

	rowOffset := rowIndex % rowsPerPage
	byteOffset := rowOffset * rowSize

	return page[byteOffset : byteOffset+rowSize]
}

// TableCapacityError indicates that the table is not able to store more data.
type TableCapacityError struct {
	capacity int
}

func (err *TableCapacityError) Error() string {
	return fmt.Sprintf("Operation will exceed max capacity of %d", err.capacity)
}

// NewTableCapacityError creates a new error for exceeding the given capacity.
func NewTableCapacityError(capacity int) *TableCapacityError {
	return &TableCapacityError{capacity}
}
