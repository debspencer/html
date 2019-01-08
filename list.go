package html

import "strconv"

type ListType int
type ListStyle string
type ListNumber string

const (
	Description ListType = iota
	Ordered
	Unordered

	// list-style-type
	LIDisc   = ListStyle("disc")
	LICircle = ListStyle("circle")
	LISquare = ListStyle("square")
	LINone   = ListStyle("none")

	// type="i"
	LINumber = ListNumber("1")
	LIALPHA  = ListNumber("A")
	LIalpha  = ListNumber("a")
	LIROMAN  = ListNumber("I")
	LIroman  = ListNumber("i")
)

// Table is the contaner for a table
type ListElement struct {
	Attributes
	listType ListType
	items    []*ListItemElement
}

// ListItemElement is the contaner for a list item <li>
type ListItemElement struct {
	Attributes
	listType ListType
	data     Element
	desc     Element
	di       int
}

// List returns a TableElement object
func List(t ListType) *ListElement {
	return &ListElement{
		listType: t,
	}
}

// Write writes the HTML list tag and table data
func (l *ListElement) Write(tw *TagWriter) {
	if len(l.items) == 0 {
		return
	}
	switch l.listType {
	case Description:
		tw.WriteTag(TagDl, l)
	case Ordered:
		tw.WriteTag(TagOl, l)
	case Unordered:
		tw.WriteTag(TagUl, l)
	}
}

// WriteContent writes the list elements
func (l *ListElement) WriteContent(tw *TagWriter) {
	for _, item := range l.items {
		item.Write(tw)
	}
}

// ListElement adds an element to the List
func (l *ListElement) AddItem(e Element) *ListItemElement {
	li := &ListItemElement{
		listType: l.listType,
		data:     e,
	}
	l.items = append(l.items, li)
	return li
}

func (li *ListElement) SetListSytle(style ListStyle) {
	li.AddAttr("list-style-type", string(style))
}

func (li *ListElement) SetStart(start int, numbering ListNumber) {
	li.AddAttr("start", strconv.Itoa(start))
	if len(numbering) > 0 {
		li.AddAttr("type", string(numbering))
	}
}

func (li *ListItemElement) AddDescription(e Element) {
	if li.listType == Description {
		li.desc = e
	}
}

// Write writes the HTML table row tag and row and column
func (li *ListItemElement) Write(tw *TagWriter) {
	switch li.listType {
	case Ordered, Unordered:
		tw.WriteTag(TagLi, li)
	case Description:
		tw.WriteTag(TagDt, li)
		tw.WriteTag(TagDd, li)
	}
	li.di = 0
}

// WriteContent writes the HTML table row and column data
func (li *ListItemElement) WriteContent(tw *TagWriter) {
	if li.di%2 == 0 {
		li.data.Write(tw)
	} else if li.desc != nil {
		li.desc.Write(tw)
	}
	li.di++
}
