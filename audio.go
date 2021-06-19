package html

type AudioPreload string

const (
	PreloadNone     = AudioPreload("none")
	PreloadMetaData = AudioPreload("metadata")
	PreloadAuto     = AudioPreload("auto")
)

type AudioElement struct {
	Container
}

func Audio(src string) *AudioElement {
	audio := &AudioElement{}
	if len(src) > 0 {
		audio.AddAttr("src", src)
	}
	return audio
}

func (e *AudioElement) Unimplemented(text string) *AudioElement {
	e.Add(Text(text))
	return e
}

func (e *AudioElement) Preload(preload AudioPreload) *AudioElement {
	e.AddAttr("preload", string(preload))
	return e
}

func (e *AudioElement) Controls() *AudioElement {
	e.AddAttr("controls", "true")
	return e
}

func (e *AudioElement) Source(s *SourceElement) *AudioElement {
	e.Add(s)
	return e
}

func (e *AudioElement) Write(tw *TagWriter) {
	tw.WriteTag(TagAudio, e)
}

type SourceElement struct {
	Attributes
}

func Source(src string) *SourceElement {
	e := &SourceElement{}
	e.AddAttr("src", src)
	return e
}

func (e *SourceElement) Type(t string) *SourceElement {
	e.AddAttr("type", t)
	return e
}
func (e *SourceElement) Write(tw *TagWriter) {
	tw.WriteTag(TagSource, e)
}

func (e *SourceElement) WriteContent(tw *TagWriter) {
}
