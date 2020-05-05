package main

// +const uint32_t PAGE_SIZE = 4096;
// +#define TABLE_MAX_PAGES 100
// +const uint32_t ROWS_PER_PAGE = PAGE_SIZE / ROW_SIZE;
// +const uint32_t TABLE_MAX_ROWS = ROWS_PER_PAGE * TABLE_MAX_PAGES;
// +
// +typedef struct {
// +  uint32_t num_rows;
// +  void* pages[TABLE_MAX_PAGES];
// +} Table;

const (
	pageSize      = 1000 * 4096
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
