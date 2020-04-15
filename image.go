package html

import "strconv"

type ImageElement struct {
	Attributes
}

func Image(src string) *ImageElement {
	img := &ImageElement{}
	img.AddAttr("src", src)
	return img
}
func (e *ImageElement) Height(h int) *ImageElement {
	e.AddAttr("heigth", strconv.Itoa(h))
	return e
}
func (e *ImageElement) Width(w int) *ImageElement {
	e.AddAttr("width", strconv.Itoa(w))
	return e
}

func (e *ImageElement) AddMap(m *MapElement) *ImageElement {
	name := m.GetAttr("name")
	e.AddAttr("usemap", "#"+name)
	return e
}

func (e *ImageElement) Write(tw *TagWriter) {
	tw.WriteTag(TagImg, e)
}

func (e *ImageElement) WriteContent(tw *TagWriter) {
}
