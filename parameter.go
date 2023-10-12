package mtgsearch

import (
	"net/url"
	"strings"
)

type parameter interface {
	pipe(o parameter) parameter
	combine(o parameter) parameter
	String() string
	EncodeValues(key string, v *url.Values) error
}

type combination struct {
	values []parameter
}

func (c combination) EncodeValues(key string, v *url.Values) error {
	s := c.String()
	if len(s) > 0 {
		v.Set(key, s)
	}
	return nil
}

func (c combination) String() string {
	s := make([]string, len(c.values))
	for i, v := range c.values {
		s[i] = v.String()
	}
	return strings.Join(s, ",")
}

func (c combination) pipe(o parameter) parameter {
	switch t := o.(type) {
	case combination:
		return pype{[]parameter{c, t}}
	case pype:
		t.values = append(t.values, c)
		return t
	case value:
		return pype{[]parameter{c, t}}
	case empty:
		return c
	}
	return emptyParameter
}

func (c combination) combine(o parameter) parameter {
	switch t := o.(type) {
	case combination:
		c.values = appendAll(c.values, t.values)
		return c
	case pype:
		v := t.split()
		nv := make([]parameter, len(v), cap(v))
		for i, e := range v {
			nv[i] = e.combine(c)
		}
		return pype{nv}
	case value:
		c.values = append(c.values, t)
		return c
	case empty:
		return c
	}
	return emptyParameter
}

type pype struct {
	values []parameter
}

func (p pype) EncodeValues(key string, v *url.Values) error {
	s := p.String()
	if len(s) > 0 {
		v.Set(key, s)
	}
	return nil
}

func (p pype) String() string {
	s := make([]string, len(p.values))
	for i, v := range p.values {
		s[i] = v.String()
	}
	return strings.Join(s, "|")
}

func (p pype) split() []parameter {
	return p.values
}

func (p pype) pipe(o parameter) parameter {
	switch t := o.(type) {
	case combination:
		p.values = append(p.values, t)
		return p
	case pype:
		p.values = appendAll(p.values, t.values)
		return p
	case value:
		p.values = append(p.values, t)
		return p
	case empty:
		return p
	}
	return emptyParameter
}

func (p pype) combine(o parameter) parameter {
	switch t := o.(type) {
	case combination:
		v := p.split()
		nv := make([]parameter, len(v), cap(v))
		for i, e := range v {
			nv[i] = e.combine(t)
		}
		p.values = nv
		return p
	case pype:
		v1 := p.split()
		v2 := t.split()
		nv := make([]parameter, len(v1)*len(v2))
		for i, e1 := range v1 {
			for j, e2 := range v2 {
				nv[i+len(v2)*j] = e1.combine(e2)
			}
		}
		return pype{nv}
	case value:
		v := p.split()
		nv := make([]parameter, len(v), cap(v))
		for i, e := range v {
			nv[i] = e.combine(t)
		}
		p.values = nv
		return p
	case empty:
		return p
	}
	return emptyParameter
}

type empty struct{}

var emptyParameter = empty{}

func (e empty) EncodeValues(key string, v *url.Values) error {
	return nil
}

func (e empty) String() string {
	return ""
}

func (e empty) pipe(o parameter) parameter {
	return o
}

func (e empty) combine(o parameter) parameter {
	return o
}

type value struct {
	value string
}

func (vl value) EncodeValues(key string, v *url.Values) error {
	s := vl.String()
	if len(s) > 0 {
		v.Set(key, s)
	}
	return nil
}

func (v value) String() string {
	return v.value
}

func (v value) pipe(o parameter) parameter {
	switch o.(type) {
	case value:
		return pype{[]parameter{v, o}}
	default:
		return o.pipe(v)
	}
}

func (v value) combine(o parameter) parameter {
	switch o.(type) {
	case value:
		return combination{[]parameter{v, o}}
	default:
		return o.combine(v)
	}
}
