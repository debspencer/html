package html

import (
	"fmt"
	"net/http"
	"strings"
)

// TagWrite is the base struct for writing output
// it contains an io write which the HTML document is rendered into
type TagWriter struct {
	w http.ResponseWriter
}

// HtmlTag defines the open/close structure for the tag
// Attributes are automatically inserted before the > in open tag
type HtmlTag struct {
	Open  string
	Close string
}

var (
	TagNone = HtmlTag{}

	// go:nofmt
	TagA        = HtmlTag{Open: "<a>",        Close: "</a>"}
	TagArea     = HtmlTag{Open: "<area>",     Close: ""}
	TagBody     = HtmlTag{Open: "<body>",     Close: "</body>"}
	TagBr       = HtmlTag{Open: "<br>",       Close: ""}
	TagComment  = HtmlTag{Open: "<!--",       Close: "-->"}
	TagDd       = HtmlTag{Open: "<dd>",       Close: "</dd>"}
	TagDiv      = HtmlTag{Open: "<div>",      Close: "</div>"}
	TagDl       = HtmlTag{Open: "<dl>",       Close: "</dl>"}
	TagDt       = HtmlTag{Open: "<dt>",       Close: "</dt>"}
	TagForm     = HtmlTag{Open: "<form>",     Close: "</form>"}
	TagH1       = HtmlTag{Open: "<h1>",       Close: "</h1>"}
	TagH2       = HtmlTag{Open: "<h2>",       Close: "</h2>"}
	TagH3       = HtmlTag{Open: "<h3>",       Close: "</h3>"}
	TagH4       = HtmlTag{Open: "<h4>",       Close: "</h4>"}
	TagH5       = HtmlTag{Open: "<h5>",       Close: "</h5>"}
	TagH6       = HtmlTag{Open: "<h6>",       Close: "</h6>"}
	TagHead     = HtmlTag{Open: "<head>",     Close: "</head>"}
	TagHtml     = HtmlTag{Open: "<html>",     Close: "</html>"}
	TagImg      = HtmlTag{Open: "<img>",      Close: ""}
	TagInput    = HtmlTag{Open: "<input>",    Close: ""}
	TagLabel    = HtmlTag{Open: "<label>",    Close: "</label>"}
	TagLi       = HtmlTag{Open: "<li>",       Close: "</li>"}
	TagMap      = HtmlTag{Open: "<map>",      Close: "</map>"}
	TagMeta     = HtmlTag{Open: "<meta>",     Close: ""}
	TagOl       = HtmlTag{Open: "<ol>",       Close: "</ol>"}
	TagOption   = HtmlTag{Open: "<option>",   Close: "</option>"}
	TagPre      = HtmlTag{Open: "<pre>",      Close: "</pre>"}
	TagScript   = HtmlTag{Open: "<script>",   Close: "</script>"}
	TagSelect   = HtmlTag{Open: "<select>",   Close: "</select>"}
	TagStyle    = HtmlTag{Open: "<style>",    Close: "</style>"}
	TagTable    = HtmlTag{Open: "<table>",    Close: "</table>"}
	TagTd       = HtmlTag{Open: "<td>",       Close: "</td>"}
	TagTextArea = HtmlTag{Open: "<textarea>", Close: "</textarea>"}
	TagTitle    = HtmlTag{Open: "<title>",    Close: "</title>"}
	TagTh       = HtmlTag{Open: "<th>",       Close: "</th>"}
	TagTr       = HtmlTag{Open: "<tr>",       Close: "</tr>"}
	TagUl       = HtmlTag{Open: "<ul>",       Close: "</ul>"}
	Newline     = []byte("\n")
	// go:fmt
)

// NewTagWriter creates a TagWrite to render into an io.Writer
func NewTagWriter(w http.ResponseWriter) *TagWriter {
	return &TagWriter{
		w: w,
	}
}

// WriteTag emits the tag and any attributes, and calls WriteContent to emit the payload followed by the end tag
// The playload may in turn call this method to render sub elements
func (tw *TagWriter) WriteTag(tag HtmlTag, e Element) {
	open := tag.Open
	if e != nil {
		attrs := e.GetAttrs()
		if len(attrs) > 0 {
			open = strings.Replace(open, ">", attrs+">", 1)
		}
	}
	tw.WriteString(open)
	e.WriteContent(tw)
	tw.WriteString(tag.Close)
	tw.Nl()
}

// WriteString writes a string to the io.Writer
func (tw *TagWriter) WriteString(s string) {
	tw.Write([]byte(s))
}

// Write writes a byteslice to the io.Writer
func (tw *TagWriter) Write(b []byte) {
	tw.w.Write(b)
}

// Comment will insert an HTML Comment into the stream
func (tw *TagWriter) Comment(data ...interface{}) {
	comment := " "
	for _, v := range data {
		comment += fmt.Sprintf("%v ", v)
	}
	tw.WriteTag(TagComment, &htmlComment{comment: comment})
}

// Nl will insert a newline into the stream
func (tw *TagWriter) Nl() {
	tw.w.Write(Newline)
}

// htmlComment is a temporary struct for rendering comments
type htmlComment struct {
	Attributes
	comment string
}

// Write writes the HTML tag and comment
func (c *htmlComment) Write(tw *TagWriter) {
	tw.WriteTag(TagComment, c)
}

// WriteContent writes the HTML comment string
func (c *htmlComment) WriteContent(tw *TagWriter) {
	tw.WriteString(c.comment)
}
