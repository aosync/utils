package perms

import (
	"strconv"
)

var text string

func Eat(r rune) bool {
	if len(text) == 0 || rune(text[0]) != r {
		return false
	}
	text = text[1:]
	return true
}

func Who(r rune) bool {
	switch r {
	case 'u':
	case 'g':
	case 'o':
	case 'a':
		return true
	}
	return false
}

func EatWho() (rune, bool) {
	if len(text) == 0 || !Who(rune(text[0])) {
		return '0', false
	}
	r := rune(text[0])
	text = text[1:]
	return r, true
}

func EatBreak(r rune) (string, string) {
	for i := len(text) - 2; i >= 0; i-- {
		if rune(text[i]) == r {
			return text[:i], text[i+1:]
		}
	}
	return "", ""
}

func ParseWhoList() ([]rune, bool) {
	var r []rune
	for {
		w, e := EatWho()
		if !e {
			break
		}
		r = append(r, w)
	}
	if len(r) == 0 {
		return nil, false
	}
	return r, true
}

func ParseActionList() ([]rune, bool) {

}

func ParseClause() (uint64, bool) {

}

func ParseSymbolicMode() (uint64, bool) {
	if p, b := ParseClause(); b {
		return p, true
	}
	ltext := text

	oa, ob := EatBreak(',')
	if oa == "" && ob == "" {
		return 0, false
	}

	text = oa
	ca, va := ParseSymbolicMode()
	text = ob
	cb, vb := ParseClause()
	text = ltext

	if va && vb {
		return ca | cb, true
	}

	return 0, false
}

func ParsePerm(perm string) (uint64, bool) {
	if d, err := strconv.ParseUint(perm, 8, 64); err == nil {
		return d, true
	}
	text = perm
	return ParseSymbolicMode()
}
