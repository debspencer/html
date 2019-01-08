package html

// Head defines the HTML body element
type DivElement struct {
	Container
}

func Div() *DivElement {
	return &DivElement{}
}

// Write writes the HTML body tag and body data
func (div *DivElement) Write(tw *TagWriter) {
	tw.WriteTag(TagDiv, div)
}
