package html

import (
	"bufio"
	"html"
	"io"
	"net/http"
)

type TextElement struct {
	Attributes
	text string
	raw  bool
}

func Text(text string) *TextElement {
	return &TextElement{text: text}
}
func (t *TextElement) Write(tw *TagWriter) {
	tw.WriteTag(TagNone, t)
}

func (t *TextElement) WriteContent(tw *TagWriter) {
	s := t.text
	if !t.raw {
		s = html.EscapeString(s)
	}
	tw.WriteString(s)
}

func Raw(text string) *TextElement {
	return &TextElement{
		text: text,
		raw:  true,
	}
}

type IOReaderElement struct {
	Attributes
	reader     io.Reader
	readcloser io.ReadCloser
}

// IoReader will add an io.Reader element
// In most cases, data to be rendered is know ahead of time, but in the case of pre, it might be slow in coming, so alow the reader to fill in as it goes
func IOReader(r io.Reader) *IOReaderElement {
	return &IOReaderElement{
		reader: r,
	}
}

// IoReaderCloser will add an io.ReadCloser element, and then close when done
// In most cases, data to be rendered is know ahead of time, but in the case of pre, it might be slow in coming, so alow the reader to fill in as it goes
func IOReadCloser(r io.ReadCloser) *IOReaderElement {
	return &IOReaderElement{
		readcloser: r,
	}
}

// Write writes the HTML tag and html data
func (e *IOReaderElement) Write(tw *TagWriter) {
	tw.WriteTag(TagNone, e)
}

// WriteContent writes the HTML for the pre
// Will attempt to Flush the data one line at a time
func (e *IOReaderElement) WriteContent(tw *TagWriter) {
	var buf *bufio.Reader
	if e.reader != nil {
		buf = bufio.NewReader(e.reader)
	} else {
		buf = bufio.NewReader(e.readcloser)
	}
	fl, _ := tw.w.(http.Flusher)

	for {
		d, err := buf.ReadBytes('\n')
		if err != nil {
			break
		}
		_, err = tw.w.Write(d)
		if err != nil {
			break
		}
		if fl != nil {
			fl.Flush()
		}
	}
	if e.readcloser != nil {
		e.readcloser.Close()
	}
}

type ParagraphElement struct {
	Container
}

// Paragraph creates a new pargraph (p)
func Paragraph(elements ...Element) *ParagraphElement {
	e := &ParagraphElement{}
	if len(elements) > 0 {
		e.Add(elements...)
	}
	return e
}

// P is an alias for Paragraph
func P(elements ...Element) *ParagraphElement {
	return Paragraph(elements...)
}

// Write writes the Div and Contents
func (e *ParagraphElement) Write(tw *TagWriter) {
	tw.WriteTag(TagP, e)
}

type BoldElement struct {
	Container
}

// Bold creates a new Bold Container (b)
func Bold(elements ...Element) *BoldElement {
	e := &BoldElement{}
	if len(elements) > 0 {
		e.Add(elements...)
	}
	return e
}

// B is an alias for Italic
func B(elements ...Element) *BoldElement {
	return Bold(elements...)
}

// Write writes the Bold contents
func (e *BoldElement) Write(tw *TagWriter) {
	tw.WriteTag(TagB, e)
}

type ItalicElement struct {
	Container
}

// Italic creates a new Italic Container (i)
func Italic(elements ...Element) *ItalicElement {
	e := &ItalicElement{}
	if len(elements) > 0 {
		e.Add(elements...)
	}
	return e
}

// I is an alias for Italic
func I(elements ...Element) *ItalicElement {
	return Italic(elements...)
}

// Write writes the Italic Contents
func (e *ItalicElement) Write(tw *TagWriter) {
	tw.WriteTag(TagI, e)
}
