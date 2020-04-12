package html

import (
	"io"
	"net/http"
)

// BaseElement is an interface that all HTML Element tags inherit from that allows for attributes
// This is done so that all element implement attributes in a consistent way
type BaseElement interface {
	// AddAttr adds a key/value attribute to an Element
	AddAttr(key string, value string)
	GetAttr(key string) string
	Style(key string, value string)

	GetAttrs() string

	AddClass(c *Class)
	AddClassName(className string)
}

// Element is an interface the implements an HTML element
type Element interface {
	BaseElement

	// Write the element using the TagWriter.
	Write(tw *TagWriter)

	// WriteContnt is called to write the conent of the (between the open and close tags) using the TagWriter
	WriteContent(tw *TagWriter)
}

type Version int

const (
	HTML4 = Version(4)
	HTML5 = Version(5)
)

// Document is the top Level HTML document
type Document struct {
	Attributes
	version Version
	head    *HeadElement
	body    *BodyElement
}

type bufferWriter struct {
	w io.Writer
}

func (b *bufferWriter) Header() http.Header {
	return http.Header{}
}
func (b *bufferWriter) Write(d []byte) (int, error) {
	return b.w.Write(d)
}
func (b *bufferWriter) WriteHeader(int) {
}

// NewDocument creates a new HTML Document container
// This is the top level method
func NewDocument( /* ver Version */ ) *Document {
	return &Document{
		//		version: ver,
		head: Head(),
		body: &BodyElement{},
	}
}

// Head returns the head object
func (doc *Document) Head() *HeadElement {
	return doc.head
}

// Body returns the body object
func (doc *Document) Body() *BodyElement {
	return doc.body
}

// Render will write the HTML document to the supplied io.Writer
func (doc *Document) Render(w http.ResponseWriter) {
	tw := NewTagWriter(w)
	//	switch doc.version {
	//	case HTML4:
	tw.WriteString("<!DOCTYPE html>")
	//	case HTML5:
	//		tw.WriteString("<!DOCTYPE html>")
	//	}
	doc.Write(tw)
}

func (doc *Document) IoRender(w io.Writer) {
	bw := bufferWriter{
		w: w,
	}

	doc.Render(&bw)
}

// Write writes the HTML tag and html data
func (doc *Document) Write(tw *TagWriter) {
	tw.WriteTag(TagHtml, doc)
}

// Write writes the HTML head/styles/body
func (doc *Document) WriteContent(tw *TagWriter) {
	doc.head.Write(tw)
	doc.body.Write(tw)
}

// AddStyle will add a style into the document
func (doc *Document) AddStyle(style *Style) {
	doc.head.styles.Add(style)
}

// AddCSS will add a CSS into the document
func (doc *Document) AddCSS(css *CSSData) {
	doc.head.css.Add(css)
}
