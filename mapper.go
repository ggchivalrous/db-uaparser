package dbuaparser

import "strings"

type strMapper struct{}

func (strMapper) handle(str []byte, maps handleMaps) string {
	str1 := string(str)

	for k, v := range maps {
		if has(str1, k) {
			if v == _UNKNOWN {
				return ""
			}
			return v
		}
	}

	return str1
}

type toLower struct{}

func (toLower) handle(str []byte, maps handleMaps) string {
	return lowerize(string(str))
}

type trimMapper struct{}

func (trimMapper) handle(str []byte, maps handleMaps) string {
	return trim(string(str), 0)
}

type firstUpperMapper struct{}

func (firstUpperMapper) handle(str []byte, maps handleMaps) string {
	s := string(str)
	if s == "" {
		return ""
	}
	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}
