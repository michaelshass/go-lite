package main

const (
	idSize         = 4
	idOffset       = 0
	usernameSize   = 32
	userNameoffset = idOffset + idSize
	mailSize       = 255
	mailOffset     = userNameoffset + usernameSize
	rowSize        = idSize + usernameSize + mailSize
)

// Row in a Table
// To make things easier, the row only supports fixed attributes.
type Row struct {
	ID       [idSize]byte
	Username [usernameSize]byte
	Mail     [mailSize]byte
}

// NewRow creates a row with empty attributes.
func NewRow() *Row {
	row := &Row{}
	row.ID = [idSize]byte{}
	row.Username = [usernameSize]byte{}
	row.Mail = [mailSize]byte{}
	return row
}

func serializeRow(src *Row) (dst [rowSize]byte) {
	copy(dst[idOffset:], src.ID[:])
	copy(dst[userNameoffset:], src.Username[:])
	copy(dst[mailOffset:], src.Mail[:])
	return
}

func deserializeRow(src [rowSize]byte) *Row {
	row := &Row{}
	copy(row.ID[:], src[:])
	copy(row.Username[:], src[userNameoffset:mailOffset])
	copy(row.Mail[:], src[mailOffset:rowSize])
	return row
}
