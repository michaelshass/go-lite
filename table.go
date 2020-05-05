package main

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
