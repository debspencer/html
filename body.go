package html

// Head defines the HTML body element
type BodyElement struct {
	Container
}

// Write writes the HTML body tag and body data
func (body *BodyElement) Write(tw *TagWriter) {
	tw.WriteTag(TagBody, body)
}
