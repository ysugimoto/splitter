package splitter

import "bytes"

const (
	SQ = 39
	DQ = 34
)

func inByte(haystack []byte, needle byte) (has bool) {
	for _, b := range haystack {
		if b == needle {
			has = true
			break
		}
	}

	return
}

func SplitBytes(bt, sep []byte) (result [][]byte) {
	stack := []byte{}
	quote := false
	squote := false
	dquote := false

	for _, v := range bt {
		if inByte(sep, v) && !quote {
			result = append(result, stack)
			stack = []byte{}
			continue
		}
		if v == SQ {
			if squote {
				quote = false
				squote = false
			} else {
				quote = true
				squote = true
			}
		} else if v == DQ {
			if dquote {
				quote = false
				dquote = false
			} else {
				quote = true
				dquote = true
			}
		}
		stack = append(stack, v)
	}

	if len(stack) > 0 {
		if squote || dquote {
			result = append(result, bytes.Split(stack, sep)...)
		} else {
			result = append(result, stack)
		}
	}

	return
}

func SplitString(str, sep string) (result []string) {
	parsed := SplitBytes([]byte(str), []byte(sep))
	for _, v := range parsed {
		result = append(result, string(v))
	}

	return
}
