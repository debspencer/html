package html

import (
	"database/sql"
	"fmt"
	"strconv"
)

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
func Table(name ...string) *TableElement {
	table := &TableElement{}
	if len(name) > 0 {
		table.AddAttr("name", name[0])
	}
	return table
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

// Cells adds multiple elements to the table
func (row *RowElement) Cells(elements ...Element) {
	if len(elements) == 0 {
		row.CellString("")
		return
	}
	for _, e := range elements {
		row.Cell(e)
	}
}

// CellString adds an string element to the table
func (row *RowElement) CellString(s string) *CellElement {
	return row.Cell(Text(s))
}

// CellStrings adds an string element to the table
func (row *RowElement) CellStrings(ss ...string) {
	for _, s := range ss {
		row.CellString(s)
	}
}

// CellString adds an string element to the table
func (row *RowElement) CellInt(value interface{}) *CellElement {
	var text string
	switch value.(type) {
	case sql.NullInt64:
		sv := value.(sql.NullInt64)
		if sv.Valid {
			text = strconv.FormatInt(sv.Int64, 10)
		}
	default:
		text = fmt.Sprintf("%d", value)
	}
	return row.Cell(Text(text))
}

// CellStrings adds an string element to the table
func (row *RowElement) CellInts(is ...int) {
	for _, i := range is {
		row.CellInt(i)
	}
}

func (row *RowElement) Rowspan(n int) *RowElement {
	row.AddAttr("rowspan", strconv.Itoa(n))
	return row
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

func (cell *CellElement) Colspan(n int) *CellElement {
	cell.AddAttr("colspan", strconv.Itoa(n))
	return cell
}

func (cell *CellElement) Bg(color string) *CellElement {
	cell.Style("background-color", color)
	return cell
}

func (cell *CellElement) Fg(color string) *CellElement {
	cell.Style("color", color)
	return cell
}

func (cell *CellElement) Center() *CellElement {
	cell.Style("text-align", "center")
	return cell
}
func (cell *CellElement) Right() *CellElement {
	cell.Style("text-align", "right")
	return cell
}
func (cell *CellElement) Left() *CellElement {
	cell.Style("text-align", "left")
	return cell
}
