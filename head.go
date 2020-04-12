package html

import "strconv"

// Head defines the HTML head element
type HeadElement struct {
	Container
	title  string
	css    *CSSElement
	styles *StyleElement
}

func Head() *HeadElement {
	head := &HeadElement{
		css:    NewCSS(),
		styles: NewStyles(),
	}
	head.Add(head.css)
	head.Add(head.styles)
	return head
}

// Write writes the HTML head tag and head data
func (head *HeadElement) Write(tw *TagWriter) {
	tw.WriteTag(TagHead, head)
}

// AddTitle Adds a title to the Header
func (head *HeadElement) AddTitle(title string) {
	head.title = title
	head.Add(NewTitle(title))
}

// GetTitle will return the last title set
func (head *HeadElement) GetTitle() string {
	return head.title
}

// Head defines the HTML head title element
type Title struct {
	Attributes
	Title string
}

// NewTitle will create a new HTML head title with the given title string
func NewTitle(title string) *Title {
	return &Title{
		Title: title,
	}
}

// Write writes the HTML head title tag and title
func (t *Title) Write(tw *TagWriter) {
	tw.WriteTag(TagTitle, t)
}

// WriteContent writes the HTML title
func (t *Title) WriteContent(tw *TagWriter) {
	tw.WriteString(t.Title)
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
