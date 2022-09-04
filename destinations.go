package salt

import (
	"strings"
	"unicode"
)

// Dst - destination that has id, where bot will send logs
//
// Use Chat, Channel, Private
type Dst interface {
	ID() string
	Valid() bool
}

func onlyDigits(s string) bool {
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

type (
	Chat string
)

func (c Chat) ID() string {
	return string(c)
}

func (c Chat) Valid() bool {
	if len(c) == 0 {
		return false
	}
	if !onlyDigits(string(c[1:])) {
		return false
	}
	if !strings.HasPrefix(string(c), "-100") {
		return false
	}
	return true
}

type (
	Channel string
)

func (c Channel) Valid() bool {
	if len(c) == 0 {
		return false
	}
	if !onlyDigits(string(c[1:])) {
		return false
	}
	return true
}

func (c Channel) ID() string {
	return string(c)
}

type (
	Private string
)

func (p Private) Valid() bool {
	if len(p) == 0 {
		return false
	}
	if !onlyDigits(string(p[1:])) {
		return false
	}
	return true
}

func (p Private) ID() string {
	return string(p)
}
