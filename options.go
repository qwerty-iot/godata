package godata

import (
	"strings"
)

type ListFilter struct {
	Raw   string
	Field string
	Op    string
	Value string
}

type ListOptions struct {
	Top       int
	Skip      int
	SkipToken string
	Expand    string
	Filter    *ListFilter
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func (o *ListOptions) ParseFilter(filter string) {
	parts := strings.Split(filter, " ")
	if len(parts) == 3 {
		o.Filter = &ListFilter{Raw: filter, Field: parts[0], Op: parts[1], Value: trimQuotes(parts[2])}
	}
	return
}
