package html

// Head defines the HTML body element
type DivElement struct {
	Container
}

// Div creates a new Div.   Optionally can take a list of elements to add to the div
func Div(elements ...Element) *DivElement {
	div := &DivElement{}
	if len(elements) > 0 {
		div.Add(elements...)
	}
	return div
}

// Write writes the HTML body tag and body data
func (div *DivElement) Write(tw *TagWriter) {
	tw.WriteTag(TagDiv, div)
}

// Center creates a new Div that is centered.  Optionally can take a list of elements to add to the div
func Center(elements ...Element) *DivElement {
	div := Div(elements...)
	div.AddAttr("align", "center")
	return div
}

func Error(text string) *DivElement {
	div := Div()
	div.Add(Text(text))
	div.AddAttr("background-color", "red")
	return div
}
