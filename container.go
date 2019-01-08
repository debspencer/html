package html

// Container is a generic container of HTML elements, like div, body, etc...
type Container struct {
	Attributes
	elements []Element
}

// WriteContent write all elements in the container
func (c *Container) WriteContent(tw *TagWriter) {
	for _, e := range c.elements {
		e.Write(tw)
	}
}

// Add an element to the container
func (c *Container) Add(e Element) {
	c.elements = append(c.elements, e)
}
