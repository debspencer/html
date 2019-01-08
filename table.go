package html

// Table is the contaner for a table
type TableElement struct {
	Attributes
	rows []*RowElement
}

// RowElement is the contaner for a table row
type RowElement struct {
	Attributes
	rowType HtmlTag
	cells   []*CellElement
}

// CellElement is the contaner for a table row
type CellElement struct {
	Attributes
	tagType HtmlTag
	data    Element
}

// Table returns a TableElement object
func Table() *TableElement {
	return &TableElement{}
}

// Write writes the HTML table tag and table data
func (table *TableElement) Write(tw *TagWriter) {
	if len(table.rows) == 0 {
		return
	}
	tw.WriteTag(TagTable, table)
}

// WriteContent writes the HTML table data
func (table *TableElement) WriteContent(tw *TagWriter) {
	for _, row := range table.rows {
		row.Write(tw)
	}
}

// NewTable returns a Table RowElement object
func (table *TableElement) Row() *RowElement {
	return table.addRow(TagTd)
}

// NewTable returns a Table RowElement object
func (table *TableElement) Header() *RowElement {
	return table.addRow(TagTh)
}

func (table *TableElement) addRow(tagType HtmlTag) *RowElement {
	row := &RowElement{
		rowType: tagType,
	}
	table.rows = append(table.rows, row)
	return row
}

// Write writes the HTML table row tag and row and column
func (row *RowElement) Write(tw *TagWriter) {
	if len(row.cells) == 0 {
		return
	}
	tw.WriteTag(TagTr, row)
}

// WriteContent writes the HTML table row and column data
func (row *RowElement) WriteContent(tw *TagWriter) {
	for _, cell := range row.cells {
		cell.tagType = row.rowType
		cell.Write(tw)
	}
}

// Cell adds an element to the table
func (row *RowElement) Cell(e Element) *CellElement {
	cell := &CellElement{
		tagType: row.rowType,
		data:    e,
	}
	row.cells = append(row.cells, cell)
	return cell
}

// CellString adds an string element to the table
func (row *RowElement) CellString(s string) {
	row.Cell(Text(s))
}

// Write writes the HTML table row tag and row and column
func (cell *CellElement) Write(tw *TagWriter) {
	tw.WriteTag(cell.tagType, cell)
}

// WriteContent writes the HTML table row and column data
func (cell *CellElement) WriteContent(tw *TagWriter) {
	if cell.data != nil {
		cell.data.Write(tw)
	}
}
