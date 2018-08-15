package str

func Simp(complex string) (simple string) {
	clen := len(complex)
	simp := make([]byte, clen)
	slen := 0
	for cidx := 0; cidx < clen; cidx++ {
		codeunit := complex[cidx]
		switch {
		case codeunit >= 'A' && codeunit <= 'Z':
			codeunit += 'a' - 'A'
			fallthrough
		case codeunit >= 'a' && codeunit <= 'z':
			fallthrough
		case codeunit >= '0' && codeunit <= '9':
			fallthrough
		case codeunit == '-' || codeunit == '.' || codeunit == '/' || codeunit == '_':
			simp[slen] = codeunit
			slen++
		}
	}
	return string(simp[:slen])
}
