package html

import "sort"

// Attributes is a contaner for element attributes, implements BaseElement
type Attributes struct {
	// attrs is a map of key/value attributes
	attrs map[string]string
}

// AddAttr will all a key/value attribute to an element
func (a *Attributes) AddAttr(key string, value string) {
	if a.attrs == nil {
		a.attrs = make(map[string]string)
	}
	a.attrs[key] = value
}

func (a *Attributes) GetAttr(key string) string {
	return a.attrs[key]
}

// StyleAttr will all a style key/value attribute to an element
func (a *Attributes) Style(key string, value string) {
	if a.attrs == nil {
		a.attrs = make(map[string]string)
	}
	style := a.attrs["style"]
	if len(style) > 0 {
		style += ";"
	}
	style += key + ":" + value

	a.attrs["style"] = style
}

// GetAttr will return a serialized list of attrs in the form of ` attr1="attr" attr2="attr"`
func (a *Attributes) GetAttrs() string {
	if len(a.attrs) == 0 {
		return ""
	}
	var ret string
	keys := make([]string, 0, len(a.attrs))
	for k := range a.attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := a.attrs[k]
		switch v {
		case "false":
		case "true":
			ret += " " + k
		default:
			ret += " " + k + `="` + v + `"`
		}
	}
	return ret
}

func (a *Attributes) AddClass(c *Class) {
	a.AddAttr("class", c.Name)
}

func (a *Attributes) AddClassName(className string) {
	a.AddAttr("class", className)
}
