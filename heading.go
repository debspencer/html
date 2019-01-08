package html

// Table is the contaner for a table
type HeadingElement struct {
	Attributes
	level int
	data  Element
}

func Heading(level int, e Element) *HeadingElement {
	if level < 1 {
		level = 1
	} else if level > 6 {
		level = 6
	}
	return &HeadingElement{
		level: level,
		data:  e,
	}
}

func (h *HeadingElement) Write(tw *TagWriter) {
	// go:nofmt
	switch h.level {
	case 1:	tw.WriteTag(TagH1, h)
	case 2:	tw.WriteTag(TagH2, h)
	case 3:	tw.WriteTag(TagH3, h)
	case 4:	tw.WriteTag(TagH4, h)
	case 5:	tw.WriteTag(TagH5, h)
	case 6:	tw.WriteTag(TagH6, h)
	}
	// go:fmt
}

// WriteContent writes the HTML table row and column data
func (h *HeadingElement) WriteContent(tw *TagWriter) {
	if h.data != nil {
		h.data.Write(tw)
	}
}
