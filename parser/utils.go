package parser

import "regexp"

func AppendManyNew[T any](v []T, as ...any) []T {
	l := len(v)
	for _, a := range as {
		switch a.(type) {
		case T:
			l += 1
		case []T:
			l += len(a.([]T))
		default:
			E("type error:", vt(a))
		}
	}
	n := make([]T, l)
	n = n[:0]
	n = append(n, v...)
	for _, a := range as {
		switch a.(type) {
		case T:
			v = append(v, a.(T))
		case []T:
			v = append(v, a.([]T)...)
		default:
			E("type error:", vt(a))
		}
	}
	return n
}

func regexSplit(s0, frx s) []s {
	return regexp.MustCompile(frx).FindAllString(s0, -1)
}
