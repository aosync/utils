package xbd

import "strings"

type OptRule struct {
	R    rune
	Oper bool
}

type Enc struct {
	R    rune
	Oper string
}

func Encountered(enc []Enc, r rune) (bool, string) {
	for _, e := range enc {
		if e.R == r {
			return true, e.Oper
		}
	}
	return false, ""
}

func Lookup(rules []OptRule, r rune) *OptRule {
	for _, rule := range rules {
		if rule.R == r {
			return &rule
		}
	}
	return nil
}

func GetOpts(work []string, rules []OptRule) ([]Enc, []string) {
	var operands []string
	var enc []Enc
Outer:
	for len(work) > 0 {
		if !strings.HasPrefix(work[0], "-") || work[0] == "-" {
			operands = append(operands, work[0])
			work = work[1:]
			continue
		}
		work[0] = work[0][1:]
		for len(work[0]) > 0 {
			opt := rune(work[0][0])
			behav := Lookup(rules, opt)
			if behav == nil {
				// Error
				work[0] = work[0][1:]
				continue
			}
			if !(*behav).Oper {
				enc = append(enc, Enc{R: opt, Oper: ""})
				work[0] = work[0][1:]
				continue
			}
			if len(work[0][1:]) > 0 {
				enc = append(enc, Enc{R: opt, Oper: work[0][1:]})
				work = work[1:]
				continue Outer
			}
			if len(work) >= 2 {
				enc = append(enc, Enc{R: opt, Oper: work[1]})
				work = work[2:]
				continue Outer
			}
			enc = append(enc, Enc{R: opt, Oper: ""})
			work = work[1:]
			continue Outer
		}
		work = work[1:]
	}

	return enc, operands
}
