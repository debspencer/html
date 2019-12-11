package html

// Head defines the HTML body element
type ScriptElement struct {
	Container
}

func Script(js string) *ScriptElement {
	s := &ScriptElement{}
	s.Add(Raw(js))
	return s
}

// Write writes the HTML body tag and body data
func (s *ScriptElement) Write(tw *TagWriter) {
	tw.WriteTag(TagScript, s)
}
