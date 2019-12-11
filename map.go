package html

// MapElement is the contaner for area elements
type MapElement struct {
	Attributes
	items []*AreaElement
}

// Area element
type AreaElement struct {
	Attributes
}

// List returns a TableElement object
func Map(name string) *MapElement {
	m := &MapElement{}
	m.AddAttr("name", name)
	return m
}

// Write writes the HTML list tag and table data
func (e *MapElement) Write(tw *TagWriter) {
	tw.WriteTag(TagMap, e)
}

// WriteContent writes the list elements
func (e *MapElement) WriteContent(tw *TagWriter) {
	for _, item := range e.items {
		item.Write(tw)
	}
}

// Rect adds a Cicrle Area element to the map
func (e *MapElement) Circle(href string, coords string) *AreaElement {
	return e.addShape("circle", href, coords)
}

// Rect adds a Polygon Area element to the map
func (e *MapElement) Poly(href string, coords string) *AreaElement {
	return e.addShape("poly", href, coords)
}

// Rect adds a Rectangle Area element to the map
func (e *MapElement) Rect(href string, coords string) *AreaElement {
	return e.addShape("rect", href, coords)
}

func (e *MapElement) addShape(shape string, href string, coords string) *AreaElement {
	a := &AreaElement{}
	a.AddAttr("shape", shape)
	a.AddAttr("href", href)
	a.AddAttr("coords", coords)

	e.items = append(e.items, a)
	return a
}

// Write writes the HTML area tag
func (e *AreaElement) Write(tw *TagWriter) {
	tw.WriteTag(TagArea, e)
}

// WriteContent writes the HTML area data (there is none)
func (e *AreaElement) WriteContent(tw *TagWriter) {
}
