package html

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type URL struct {
	Attributes

	Name    string  // Name or
	Element Element // Element

	Scheme   string
	UserPass string
	Host     string
	Port     string
	App      string
	Page     string
	RawQuery string
	Query    url.Values
	Anchor   string
}

func NewLink(link string) *URL {
	u, err := url.Parse(link)
	if err != nil {
		return &URL{
			Page: link,
		}
	}
	return NewURL(u, nil)
}

func NewURL(u *url.URL, formValues url.Values) *URL {
	hp := strings.SplitN(u.Host, ":", 2)
	port := ""
	if len(hp) > 1 {
		port = hp[1]
	}
	path := strings.SplitN(strings.TrimLeft(u.Path, "/"), "/", 2)
	app := path[0]
	page := ""
	if len(path) > 1 {
		page = path[1]
	}

	if formValues == nil {
		formValues = make(url.Values)
		rq := strings.Split(u.RawQuery, "&")

		for _, q := range rq {
			if len(q) == 0 {
				continue
			}
			qs := strings.SplitN(q, "=", 2)
			k := qs[0]
			v := ""
			if len(q) > 1 {
				v = qs[1]
			}
			v, _ = url.QueryUnescape(v)
			formValues[k] = append(formValues[k], v)
		}
	}

	r := &URL{
		Scheme:   u.Scheme,
		Host:     hp[0],
		Port:     port,
		App:      app,
		Page:     page,
		RawQuery: u.RawQuery,
		Query:    formValues,
		Anchor:   u.Fragment,
	}
	return r
}

func (u *URL) Clone() *URL {
	// need to deep copy query map
	q := make(url.Values)
	for k, v := range u.Query {
		q[k] = v
	}
	r := *u
	r.Query = q
	return &r
}

func (u *URL) SetName(name string) *URL {
	u.Name = name
	return u
}

func (u *URL) SetApp(app string) *URL {
	u.App = app
	return u
}

func (u *URL) SetPage(page string) *URL {
	u.Page = page
	return u
}

func (u *URL) GetQuery(k string) string {
	s, _ := u.Query[k]
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

func (u *URL) GetQueryInt(k string) (int, bool) {
	s := u.Query[k]
	if len(s) == 0 {
		return 0, false
	}

	i, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, false
	}
	return i, true
}

func (u *URL) HasQuery(k string) bool {
	_, ok := u.Query[k]
	return ok
}

func (u *URL) AddQuery(k string, v interface{}) *URL {
	if u.Query == nil {
		u.Query = make(map[string][]string)
	}
	u.Query[k] = []string{fmt.Sprintf("%v", v)}
	return u
}

func (u *URL) DelQuery(k string) *URL {
	delete(u.Query, k)
	return u
}

func (u *URL) Link() string {
	sb := strings.Builder{}

	if len(u.Scheme) > 0 || len(u.Host) > 0 {
		if len(u.Scheme) > 0 {
			sb.WriteString(u.Scheme)
		} else {
			sb.WriteString("http")
		}
		sb.WriteString("://")
		if len(u.UserPass) > 0 {
			sb.WriteString(u.UserPass)
			sb.WriteString("@")
		}
	}
	if len(u.Host) > 0 {
		sb.WriteString(u.Host)

		if len(u.Port) > 0 {
			sb.WriteString(":")
			sb.WriteString(u.Port)
		}
	}
	if len(u.App) > 0 {
		sb.WriteString("/")
		sb.WriteString(u.App)
	}
	if len(u.Page) > 0 {
		sb.WriteString("/")
		if u.Page != "/" {
			sb.WriteString(u.Page)
		}
	}
	if len(u.Query) > 0 {
		first := true
		for k, vs := range u.Query {
			for _, v := range vs {
				if first {
					sb.WriteString("?")
					first = false
				} else {
					sb.WriteString("&")
				}
				sb.WriteString(k)
				sb.WriteString("=")
				sb.WriteString(v)
			}
		}
	}
	if len(u.Anchor) > 0 {
		sb.WriteString("#")
		sb.WriteString(u.Anchor)
	}
	return sb.String()
}

// Write writes the HTML head title tag and title
func (u *URL) Write(tw *TagWriter) {
	u.AddAttr("href", u.Link())
	tw.WriteTag(TagA, u)
}

// WriteContent writes the HTML title
func (u *URL) WriteContent(tw *TagWriter) {
	if u.Element != nil {
		u.Element.Write(tw)
	} else {
		tw.WriteString(u.Name)
	}
}
