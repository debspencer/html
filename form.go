package html

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Form is the contaner for a form
type FormElement struct {
	Container
}

// List returns a TableElement object`
func Form(action *URL) *FormElement {
	f := &FormElement{}
	f.AddAttr("action", action.Link())
	f.AddAttr("method", "GET")
	f.AddAttr("name", "html_form") // This can be overriden by SetName
	return f
}

func (f *FormElement) SetName(name string) *FormElement {
	f.AddAttr("name", name)
	return f
}

func (f *FormElement) formName() string {
	return f.GetAttr("name")
}
func (f *FormElement) validateFunc() string {
	name := f.GetAttr("name")
	return "webForm_" + name + "_Validate"
}

func (f *FormElement) MethodPOST(tw *TagWriter) *FormElement {
	f.AddAttr("method", "POST")
	return f
}

func (f *FormElement) Validate(name string, msg string) *FormElement {
	docname := "document." + f.formName() + "." + name

	script := " if (" + docname + ".value.length < 1) {\n"
	script += "  alert(\"" + msg + "\");\n"
	script += "  " + docname + ".focus();\n"
	script += "  return false;\n"
	script += " }\n"

	f.AddJavaScript(f.validateFunc(), script)
	return f
}

// Write writes the HTML form tag and container data
func (f *FormElement) Write(tw *TagWriter) {
	f.AddAttr("onsubmit", "return "+f.validateFunc()+"()")
	f.AddJavaScript(f.validateFunc(), " return true;\n")
	tw.WriteTag(TagForm, f)
}

type InputElement struct {
	Attributes

	Name string
}

func (i *InputElement) Write(tw *TagWriter) {
	tw.WriteTag(TagInput, i)
}
func (i *InputElement) WriteContent(tw *TagWriter) {
}

func Hidden(name string, value interface{}) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "hidden")
	i.AddAttr("name", name)
	i.AddAttr("value", fmt.Sprintf("%v", value))
	return i
}

func Submit(label string) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "submit")
	i.AddAttr("name", "submit")
	i.AddAttr("value", label)
	return i
}

func (i *InputElement) SetName(name string) *InputElement {
	i.AddAttr("name", name)
	return i
}

func TextInput(name string, size int) *InputElement {
	i := &InputElement{}
	i.AddAttr("type", "text")
	i.AddAttr("name", name)
	i.AddAttr("size", strconv.Itoa(size))
	i.Name = name
	return i
}

func (i *InputElement) SetDefault(value string) *InputElement {
	i.AddAttr("value", value)
	return i
}

func (i *InputElement) SetDefaultInt(value interface{}) *InputElement {
	switch value.(type) {
	case sql.NullInt64:
		sv := value.(sql.NullInt64)
		if sv.Valid {
			i.SetDefault(strconv.FormatInt(sv.Int64, 10))
		}
	default:
		i.SetDefault(fmt.Sprintf("%d", value))
	}
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

type TextAreaElement struct {
	Attributes
	text string
}

func TextArea(name string, rows int, cols int) *TextAreaElement {
	ta := &TextAreaElement{}
	ta.AddAttr("name", name)
	ta.AddAttr("rows", strconv.Itoa(rows))
	ta.AddAttr("cols", strconv.Itoa(cols))
	return ta
}
func (ta *TextAreaElement) SetDefault(value string) *TextAreaElement {
	ta.text = value
	return ta
}
func (ta *TextAreaElement) Write(tw *TagWriter) {
	tw.WriteTag(TagTextArea, ta)
}
func (ta *TextAreaElement) WriteContent(tw *TagWriter) {
	tw.WriteString(ta.text)
}

type FormSelectElement struct {
	Attributes
	options []*OptionElement
}

func FormSelect(name string) *FormSelectElement {
	s := &FormSelectElement{}
	s.AddAttr("name", name)
	return s
}

func (e *FormSelectElement) Option(display, value string) *OptionElement {
	opt := &OptionElement{
		display: display,
	}
	opt.AddAttr("value", value)
	e.options = append(e.options, opt)
	return opt
}

func (e *FormSelectElement) OptionInt(display string, value interface{}) *OptionElement {
	var s string
	switch value.(type) {
	case sql.NullInt64:
		sv := value.(sql.NullInt64)
		if sv.Valid {
			s = strconv.FormatInt(sv.Int64, 10)
		}
	default:
		s = fmt.Sprintf("%d", value)
	}
	return e.Option(display, s)
}

func (e *FormSelectElement) Write(tw *TagWriter) {
	tw.WriteTag(TagSelect, e)
}
func (e *FormSelectElement) WriteContent(tw *TagWriter) {
	for _, opt := range e.options {
		opt.Write(tw)
	}
}

type OptionElement struct {
	Attributes
	display string
}

func (e *OptionElement) Selected(b ...bool) *OptionElement {
	selected := true
	if len(b) > 0 {
		selected = b[0]
	}
	e.AddAttr("selected", strconv.FormatBool(selected))
	return e
}

func (e *OptionElement) Write(tw *TagWriter) {
	tw.WriteTag(TagOption, e)
}
func (e *OptionElement) WriteContent(tw *TagWriter) {
	tw.WriteString(e.display)
}
