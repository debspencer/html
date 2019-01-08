package html

type styleWriter func(tw *TagWriter)

// Style describes an individual CSS style, which contians a set of
// associations (element, class or id) and set of key/values which
// describe individual styles
type CSSData struct {
	Attributes
	css string
}

// NewStyle will create a Style object, identified by name.
// An optional list of StyleDef can be passed to add individual styles, or Add() can be called later
// The Style will need to be added to the Styles list to be rendered
func CSS(css string) *CSSData {
	return &CSSData{
		css: css,
	}
}

func (css *CSSData) WriteContent(tw *TagWriter) {
	tw.WriteString(css.css)
}

type CSSElement struct {
	Container
	css []*CSSData
}

// NewCSS will create a container to contain CSS objects
func NewCSS() *CSSElement {
	css := &CSSElement{}
	css.AddAttr("type", "text/css")
	return css
}

// Write all styles
func (s *CSSElement) Write(tw *TagWriter) {
	// nothing to do
	if len(s.css) == 0 {
		return
	}
	tw.WriteTag(TagStyle, s)
}

// Write each style
func (s *CSSElement) WriteContent(tw *TagWriter) {
	for _, v := range s.css {
		v.WriteContent(tw)
	}
}

// Add CSS to the CSS container
func (s *CSSElement) Add(css *CSSData) {
	s.css = append(s.css, css)
}

var (
	styleSep      = []byte(", ")
	styleIndent   = []byte("    ")
	styleOpen     = []byte(" {\n")
	styleClose    = []byte("}\n")
	styleBreak    = []byte(": ")
	styleComplete = []byte(";\n")
)

// StyleDef describes an individual style key: value;
type StyleDef struct {
	// Key is the key name
	Key string

	// Value is value for the key
	Value string
}

// Style describes an individual CSS style, which contians a set of
// associations (element, class or id) and set of key/values which
// describe individual styles
type Style struct {
	name string

	// associations is an slice of associations which links a
	// style to one or more HTML elements, classes or styles
	// element:   table
	// class:     .center  or table.center
	// id:        #table1
	associations []string

	// styles key: value
	styles map[string]string
}

// NewStyle will create a Style object, identified by name.
// An optional list of StyleDef can be passed to add individual styles, or Add() can be called later
// The Style will need to be added to the Styles list to be rendered
func NewStyle(name string, defs ...StyleDef) *Style {
	style := &Style{
		name:   name,
		styles: make(map[string]string),
	}
	style.Add(defs...)
	return style
}

// Add will add individual styles
func (style *Style) Add(defs ...StyleDef) {
	for _, def := range defs {
		style.styles[def.Key] = def.Value
	}
}

// Write renders the style to the io.Writer
func (s *Style) Write(tw *TagWriter) {

	// If there are no associations, then can not write the style
	if len(s.associations) == 0 {
		return
	}

	// If there are no styles, then nothing to do
	if len(s.styles) == 0 {
		return
	}

	multiAssoc := false
	for _, a := range s.associations {
		if multiAssoc {
			tw.Write(styleSep)
		}
		tw.WriteString(a)
		multiAssoc = true
	}
	s.writeStyle(tw, func(tw *TagWriter) {
		for k, v := range s.styles {
			tw.Write(styleIndent)
			tw.WriteString(k)
			tw.Write(styleBreak) // :
			tw.WriteString(v)
			tw.Write(styleComplete) // ;\n
		}
	})
}

// Add Class to this style
func (s *Style) AddClass(c *Class) {
	if c != nil {
		s.associations = append(s.associations, "."+c.Name)
	}
}

func (s *Style) writeStyle(tw *TagWriter, sw styleWriter) {
	tw.Write(styleOpen)
	sw(tw)
	tw.Write(styleClose)
}

// StyleElement contains a map of styles
type StyleElement struct {
	Attributes
	styles map[string]*Style
}

// NewStyles will create a container to contain Style objects
func NewStyles() *StyleElement {
	styles := &StyleElement{
		styles: make(map[string]*Style),
	}
	styles.AddAttr("type", "text/css")
	return styles
}

// Write all styles
func (s *StyleElement) Write(tw *TagWriter) {
	// nothing to do
	if len(s.styles) == 0 {
		return
	}
	tw.WriteTag(TagStyle, s)
}

// Write each style
func (s *StyleElement) WriteContent(tw *TagWriter) {
	for k, v := range s.styles {
		tw.Comment("Style", k)
		v.Write(tw)
	}
}

// Add a Style to the Styles container
func (styles *StyleElement) Add(style *Style) {
	styles.styles[style.name] = style
}

// Class is an association between Styles and elements
type Class struct {
	// Name is the name of the Class
	Name string
}

// NewClass creates a new CSS Class with name as the class name
func NewClass(name string) *Class {
	return &Class{
		Name: name,
	}
}

// AddStyle is a convenience function for style.AddClass() which is awkward
func (c *Class) AddStyle(s *Style) {
	s.AddClass(c)
}

// StyleBackgroundColor defines a background color
func StyleBackgroundColor(color string) StyleDef {
	return StyleDef{
		Key:   "background",
		Value: color,
	}
}
