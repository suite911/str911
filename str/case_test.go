package str

import "testing"

type tableCase struct {
	s, affix string
	match    bool
	trimmed  string
}

func TestCaseEqual(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false},
		tableCase{"hello", "hello", true},
		tableCase{"hello", "hell", false},
		tableCase{"hell", "hello", false},
		tableCase{"hello", "HELLO", true},
		tableCase{"hello", "HELL", false},
		tableCase{"hell", "HELLO", false},
		tableCase{"HELLO", "hello", true},
		tableCase{"HELLO", "hell", false},
		tableCase{"HELL", "hello", false},
		tableCase{"HELLO", "HELLO", true},
		tableCase{"HELLO", "HELL", false},
		tableCase{"HELL", "HELLO", false},
		tableCase{"", "hello", false},
		tableCase{"hello", "", false},
		tableCase{"", "", true},
	} {
		if CaseEqual(x.s, x.affix) != x.match {
			t.Errorf("CaseEqual(%q, %q) != %v", x.s, x.affix, x.match)
		}
	}
}

func TestCaseHasPrefix(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "hell", true, "o"},
		tableCase{"hell", "hello", false, "hell"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "HELL", true, "o"},
		tableCase{"hell", "HELLO", false, "hell"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "hell", true, "O"},
		tableCase{"HELL", "hello", false, "HELL"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "HELL", true, "O"},
		tableCase{"HELL", "HELLO", false, "HELL"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		match := CaseHasPrefix(x.s, x.affix)
		if match != x.match {
			t.Errorf("CaseHasPrefix(%q, %q) != %v // %v", x.s, x.affix, x.match, match)
		}
	}
}

func TestCaseHasSuffix(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "ello", true, "h"},
		tableCase{"ello", "hello", false, "ello"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "ELLO", true, "h"},
		tableCase{"ello", "HELLO", false, "ello"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "ello", true, "H"},
		tableCase{"ELLO", "hello", false, "ELLO"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "ELLO", true, "H"},
		tableCase{"ELLO", "HELLO", false, "ELLO"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		match := CaseHasSuffix(x.s, x.affix)
		if match != x.match {
			t.Errorf("CaseHasSuffix(%q, %q) != %v // %v", x.s, x.affix, x.match, match)
		}
	}
}

func TestCaseTrimPrefix(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "hell", true, "o"},
		tableCase{"hell", "hello", false, "hell"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "HELL", true, "o"},
		tableCase{"hell", "HELLO", false, "hell"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "hell", true, "O"},
		tableCase{"HELL", "hello", false, "HELL"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "HELL", true, "O"},
		tableCase{"HELL", "HELLO", false, "HELL"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		match, trimmed := CaseTrimPrefix(x.s, x.affix)
		if match != x.match || trimmed != x.trimmed {
			t.Errorf("CaseTrimPrefix(%q, %q) != %v, %q // %v, %q", x.s, x.affix, x.match, x.trimmed, match, trimmed)
		}
	}
}

func TestCaseTrimPrefixInPlace(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "hell", true, "o"},
		tableCase{"hell", "hello", false, "hell"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "HELL", true, "o"},
		tableCase{"hell", "HELLO", false, "hell"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "hell", true, "O"},
		tableCase{"HELL", "hello", false, "HELL"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "HELL", true, "O"},
		tableCase{"HELL", "HELLO", false, "HELL"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		trimmed := x.s
		match := CaseTrimPrefixInPlace(&trimmed, x.affix)
		if match != x.match || trimmed != x.trimmed {
			t.Errorf("CaseTrimPrefixInPlace(%q, %q) != %v, %q // %v, %q", x.s, x.affix, x.match, x.trimmed, match, trimmed)
		}
	}
}

func TestCaseTrimSuffix(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "ello", true, "h"},
		tableCase{"ello", "hello", false, "ello"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "ELLO", true, "h"},
		tableCase{"ello", "HELLO", false, "ello"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "ello", true, "H"},
		tableCase{"ELLO", "hello", false, "ELLO"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "ELLO", true, "H"},
		tableCase{"ELLO", "HELLO", false, "ELLO"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		match, trimmed := CaseTrimSuffix(x.s, x.affix)
		if match != x.match || trimmed != x.trimmed {
			t.Errorf("CaseTrimSuffix(%q, %q) != %v, %q // %v, %q", x.s, x.affix, x.match, x.trimmed, match, trimmed)
		}
	}
}

func TestCaseTrimSuffixInPlace(t *testing.T) {
	for _, x := range []tableCase{
		tableCase{"hello", "world", false, "hello"},
		tableCase{"hello", "hello", true, ""},
		tableCase{"hello", "ello", true, "h"},
		tableCase{"ello", "hello", false, "ello"},
		tableCase{"hello", "HELLO", true, ""},
		tableCase{"hello", "ELLO", true, "h"},
		tableCase{"ello", "HELLO", false, "ello"},
		tableCase{"HELLO", "hello", true, ""},
		tableCase{"HELLO", "ello", true, "H"},
		tableCase{"ELLO", "hello", false, "ELLO"},
		tableCase{"HELLO", "HELLO", true, ""},
		tableCase{"HELLO", "ELLO", true, "H"},
		tableCase{"ELLO", "HELLO", false, "ELLO"},
		tableCase{"", "hello", false, ""},
		tableCase{"hello", "", true, "hello"},
		tableCase{"", "", true, ""},
	} {
		trimmed := x.s
		match := CaseTrimSuffixInPlace(&trimmed, x.affix)
		if match != x.match || trimmed != x.trimmed {
			t.Errorf("CaseTrimSuffixInPlace(%q, %q) != %v, %q // %v, %q", x.s, x.affix, x.match, x.trimmed, match, trimmed)
		}
	}
}
