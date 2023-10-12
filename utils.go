package mtgsearch

import (
	"strings"
)

const (
	and = "and"
	or  = "or"
)

func pop[T any](arr []T) (T, []T) {
	li := len(arr) - 1
	v := arr[li]
	return v, arr[:li]
}

func replaceRelativeCmc(cmc string) string {
	if strings.HasPrefix(cmc, ">=") {
		return strings.Replace(cmc, ">=", "gte", 1)
	}

	if strings.HasPrefix(cmc, "<=") {
		return strings.Replace(cmc, "<=", "lte", 1)
	}

	if strings.HasPrefix(cmc, ">") {
		return strings.Replace(cmc, ">", "gt", 1)
	}

	if strings.HasPrefix(cmc, "<") {
		return strings.Replace(cmc, "<", "lt", 1)
	}

	return cmc
}

func appendAll[T any](t, s []T) []T {
	r := make([]T, len(t)+len(s))
	for i := range t {
		r[i] = t[i]
	}

	for i := range s {
		r[len(t)+i] = s[i]
	}
	return r
}
