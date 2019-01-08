package html

import "strconv"

// Form is the contaner for a form
type FormElement struct {
	Container
}

// List returns a TableElement object
func Form(action *URL) *FormElement {
	f := &FormElement{}
	f.AddAttr("action", action.Link())
	f.AddAttr("method", "GET")
	return f
}

func (f *FormElement) Method(tw *TagWriter) *FormElement {
	f.AddAttr("method", "POST")
	return f
}

// Write writes the HTML form tag and container data
func (f *FormElement) Write(tw *TagWriter) {
	tw.WriteTag(TagForm, f)
}

type InputElement struct {
	Attributes
}

func (i *InputElement) Write(tw *TagWriter) {
	tw.WriteTag(TagInput, i)
}
func (i *InputElement) WriteContent(tw *TagWriter) {
}

func Hidden(name string, value string) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "hidden")
	i.AddAttr("name", name)
	i.AddAttr("value", value)
	return i
}

func Submit(name string, label string) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "submit")
	i.AddAttr("name", name)
	i.AddAttr("value", label)
	return i
}

func TextInput(name string, size int) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "text")
	i.AddAttr("name", name)
	i.AddAttr("size", strconv.Itoa(size))
	return i
}

func (i *InputElement) SetDefault(value string) *InputElement {
	i.AddAttr("value", value)
	return i
}

type LabelElement struct {
	Attributes
	label string
}

func Label(label string) *LabelElement {
	l := &LabelElement{
		label: label,
	}
	return l
}
func (l *LabelElement) Write(tw *TagWriter) {
	tw.WriteTag(TagLabel, l)
}
func (l *LabelElement) WriteContent(tw *TagWriter) {
	tw.WriteString(l.label)
}
