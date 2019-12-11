package html

import (
	"html"
	"io"
	"sort"
	"strconv"
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

type MetaElement struct {
	Attributes
	text string
}

// MetaRefresh will create a meta tag forcing a refresh in number of seconds.  Link is optional
func MetaRefresh(seconds int, link string) *MetaElement {
	m := &MetaElement{}
	m.AddAttr("http-equiv", "refresh")
	content := strconv.Itoa(seconds)
	if len(link) > 0 {
		content += "; URL=" + link
	}
	m.AddAttr("content", content)
	return m
}
func (m *MetaElement) Write(tw *TagWriter) {
	tw.WriteTag(TagMeta, m)
}

func (m *MetaElement) WriteContent(tw *TagWriter) {
}

type ImageElement struct {
	Attributes
}

func Image(src string) *ImageElement {
	img := &ImageElement{}
	img.AddAttr("src", src)
	return img
}
func (e *ImageElement) Height(h int) *ImageElement {
	e.AddAttr("heigth", strconv.Itoa(h))
	return e
}
func (e *ImageElement) Width(w int) *ImageElement {
	e.AddAttr("width", strconv.Itoa(w))
	return e
}

func (e *ImageElement) AddMap(m *MapElement) *ImageElement {
	name := m.GetAttr("name")
	e.AddAttr("usemap", "#"+name)
	return e
}

func (e *ImageElement) Write(tw *TagWriter) {
	tw.WriteTag(TagImg, e)
}

func (e *ImageElement) WriteContent(tw *TagWriter) {
}

// Attributes is a contaner for element attributes, implements BaseElement
type Attributes struct {
	// attrs is a map of key/value attributes
	attrs map[string]string
}

// AddAttr will all a key/value attribute to an element
func (a *Attributes) AddAttr(key string, value string) {
	if a.attrs == nil {
		a.attrs = make(map[string]string)
	}
	a.attrs[key] = value
}

func (a *Attributes) GetAttr(key string) string {
	return a.attrs[key]
}

// StyleAttr will all a style key/value attribute to an element
func (a *Attributes) Style(key string, value string) {
	if a.attrs == nil {
		a.attrs = make(map[string]string)
	}
	style := a.attrs["style"]
	if len(style) > 0 {
		style += ";"
	}
	style += key + ":" + value

	a.attrs["style"] = style
}

// GetAttr will return a serialized list of attrs in the form of ` attr1="attr" attr2="attr"`
func (a *Attributes) GetAttrs() string {
	if len(a.attrs) == 0 {
		return ""
	}
	var ret string
	keys := make([]string, 0, len(a.attrs))
	for k := range a.attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := a.attrs[k]
		switch v {
		case "false":
		case "true":
			ret += " " + k
		default:
			ret += " " + k + `="` + v + `"`
		}
	}
	return ret
}

func (a *Attributes) AddClass(c *Class) {
	a.AddAttr("class", c.Name)
}

func (a *Attributes) AddClassName(className string) {
	a.AddAttr("class", className)
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

// NewDocument creates a new HTML Document container
// This is the top level method
func NewDocument( /* ver Version */ ) *Document {
	return &Document{
		//		version: ver,
		head: NewHead(),
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
func (doc *Document) Render(w io.Writer) {
	tw := NewTagWriter(w)
	//	switch doc.version {
	//	case HTML4:
	tw.WriteString("<!DOCTYPE html>")
	//	case HTML5:
	//		tw.WriteString("<!DOCTYPE html>")
	//	}
	doc.Write(tw)
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

type BreakElement struct {
	Attributes
	count int
}

func Br(n ...int) *BreakElement {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	return &BreakElement{count: count}
}
func (br *BreakElement) Write(tw *TagWriter) {
	for i := 0; i < br.count; i++ {
		tw.WriteTag(TagBr, br)
	}
}

func (br *BreakElement) WriteContent(tw *TagWriter) {
}

type NonBreakingSpace struct {
	Attributes // does not implement
	count      int
}

func Nbsp(n ...int) *NonBreakingSpace {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	return &NonBreakingSpace{count: count}
}
func (nbsp *NonBreakingSpace) Write(tw *TagWriter) {
	for i := 0; i < nbsp.count; i++ {
		tw.WriteString("&nbsp;")
	}
}
func (nbsp *NonBreakingSpace) WriteContent(tw *TagWriter) {
}
