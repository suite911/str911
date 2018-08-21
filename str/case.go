package str

import (
	"strings"
)

func CaseEqual(a, b string) bool {
	return len(a) == len(b) && strings.ToLower(a) == strings.ToLower(b)
}

func CaseHasPrefix(s, prefix string) bool {
	match, _ := CaseTrimPrefix(s, prefix)
	return match
}

func CaseHasSuffix(s, suffix string) bool {
	match, _ := CaseTrimSuffix(s, suffix)
	return match
}

func CaseTrimPrefix(s, prefix string) (match bool, tail string) {
	l := len(prefix)
	if len(s) < l {
		return false, s
	}
	var head string
	head, tail = s[:l], s[l:]
	if match = strings.ToLower(head) == strings.ToLower(prefix); !match {
		tail = s
	}
	return
}

func CaseTrimPrefixInPlace(s *string, prefix string) bool {
	match, tail := CaseTrimPrefix(*s, prefix)
	if match {
		*s = tail
	}
	return match
}

func CaseTrimSuffix(s, suffix string) (match bool, head string) {
	l := len(suffix)
	if len(s) < l {
		return false, s
	}
	var tail string
	head, tail = s[:len(s)-l], s[len(s)-l:]
	if match = strings.ToLower(tail) == strings.ToLower(suffix); !match {
		head = s
	}
	return
}

func CaseTrimSuffixInPlace(s *string, suffix string) bool {
	match, tail := CaseTrimSuffix(*s, suffix)
	if match {
		*s = tail
	}
	return match
}
