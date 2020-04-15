package html

type IFrameElement struct {
	Container
}

// IFrame creates a new IFrame.   Optionally can take a list of elements to add to the iframe
func IFrame(elements ...Element) *IFrameElement {
	e := &IFrameElement{}
	if len(elements) > 0 {
		e.Add(elements...)
	}
	return e
}

// Write writes the IFrame and contents
func (e *IFrameElement) Write(tw *TagWriter) {
	tw.WriteTag(TagIFrame, e)
}
