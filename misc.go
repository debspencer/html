package html

type BreakElement struct {
	Attributes
	count int
}

func Br(n ...int) *BreakElement {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	return &BreakElement{count: count}
}
func (br *BreakElement) Write(tw *TagWriter) {
	for i := 0; i < br.count; i++ {
		tw.WriteTag(TagBr, br)
	}
}

func (br *BreakElement) WriteContent(tw *TagWriter) {
}

type NonBreakingSpace struct {
	Attributes // does not implement
	count      int
}

func Nbsp(n ...int) *NonBreakingSpace {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	return &NonBreakingSpace{count: count}
}
func (nbsp *NonBreakingSpace) Write(tw *TagWriter) {
	for i := 0; i < nbsp.count; i++ {
		tw.WriteString("&nbsp;")
	}
}
func (nbsp *NonBreakingSpace) WriteContent(tw *TagWriter) {
}

type PreElement struct {
	Container
}

func Pre(elements ...Element) *PreElement {
	pre := &PreElement{}
	if len(elements) > 0 {
		pre.Add(elements...)
	}
	return pre
}

// Write writes the HTML tag and html data for the pre element
func (pre *PreElement) Write(tw *TagWriter) {
	tw.WriteTag(TagPre, pre)
}
