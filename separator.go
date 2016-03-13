package main

import (
	"strings"
	"unicode/utf8"
)

func toDelimiters(list string) []string {
	var isEscaping bool
	var a []string
	for _, ch := range list {
		if isEscaping {
			switch ch {
			case '\\':
				a = append(a, "\\")
			case 't':
				a = append(a, "\t")
			case 'n':
				a = append(a, "\n")
			case '0':
				a = append(a, "")
			default:
				a = append(a, string(ch))
			}
			isEscaping = false
			continue
		}
		if ch == '\\' {
			isEscaping = true
			continue
		}
		a = append(a, string(ch))
	}
	if len(a) == 0 {
		return []string{""}
	}
	return a
}

func sizeOfHeadRune(s string) int {
	_, size := utf8.DecodeRuneInString(s)
	return size
}

type Separator struct {
	isSerial       bool
	delimiters     []string
	delimiterIndex int
}

func NewSeparator(isSerial bool, list string) *Separator {
	return &Separator{
		isSerial:       isSerial,
		delimiters:     toDelimiters(list),
		delimiterIndex: 0,
	}
}

func (s *Separator) Separate(t string) []string {
	if t == "" {
		return []string{}
	}

	var beg, end int
	var a []string
	for {
		d := s.delimiters[s.delimiterIndex]
		if d == "" {
			n := sizeOfHeadRune(t)
			if n == 0 {
				break
			}
			end = beg + n
		} else {
			i := strings.Index(t[beg:], d)
			if i == -1 {
				a = append(a, t[end+len(d):])
				break
			}
			end = beg + i
		}

		a = append(a, t[beg:end])
		if end >= len(t) {
			break
		}
		beg = end + len(d)

		if s.delimiterIndex < len(s.delimiters)-1 {
			s.delimiterIndex++
		} else {
			s.delimiterIndex = 0
		}
	}
	return a
}
