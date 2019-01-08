package html

import (
	"bytes"

	"io/ioutil"
	"testing"
)

func TestDocumentWrite(t *testing.T) {
	css := "supercss { font: bold; }"

	style := NewStyle("mystyle", StyleBackgroundColor(ColorRed))
	class := NewClass("myclass")

	// golang 1.10 strings buffer
	var b bytes.Buffer
	doc := NewDocument()
	doc.Head().AddTitle("My Document")
	doc.AddStyle(style)
	doc.AddCSS(CSS("loaded", css))

	// tbl := Table(Row1(), Row2())

	tbl := Table()
	row := tbl.Header()
	row.CellString("hello there")
	row.Cell(Table())

	row = tbl.Row()
	row.CellString("hello there")
	row.Cell(Table())

	doc.Body().Add(tbl)
	doc.AddStyle(style)

	class.AddStyle(style)
	//	class.Add(tbl)
	tbl.AddClass(class)
	tbl.AddClassName("foo")
	doc.Render(&b)
	ioutil.WriteFile("/tmp/goodie.html", b.Bytes(), 0644)
}
