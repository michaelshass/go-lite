package main

const (
	pageSize      = 4096 * 1000
	rowsPerPage   = pageSize / rowSize
	tableMaxPages = 100
	tableMaxRows  = rowsPerPage * tableMaxPages
)

// Table holds data in pages each 4096kb in size.
// Number of total pages in limited to const 'tableMaxPages'
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
