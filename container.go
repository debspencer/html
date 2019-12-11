package html

import "fmt"

// Container is a generic container of HTML elements, like div, body, etc...
type Container struct {
	Attributes
	elements   []Element
	javaScript map[string]string
}

// WriteContent write all elements in the container
func (c *Container) WriteContent(tw *TagWriter) {
	if len(c.javaScript) > 0 {
		for k, v := range c.javaScript {
			Script(fmt.Sprintf("function %s() {\n%s\n}\n", k, v)).Write(tw)
		}
	}

	for _, e := range c.elements {
		e.Write(tw)
	}
}

// Add an element to the container
func (c *Container) Add(e ...Element) {
	c.elements = append(c.elements, e...)
}

func (c *Container) AddJavaScript(scriptName string, script string) {
	if c.javaScript == nil {
		c.javaScript = make(map[string]string)
	}
	c.javaScript[scriptName] += script
}
