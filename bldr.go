package mtgsearch

import (
	"strings"
)

func buildQuery(expr []string) cardsQuery {
	stack := make([]cardsQuery, 0)
	var q1 cardsQuery
	var q2 cardsQuery
	for _, token := range expr {
		switch token {
		case and:
			q1, stack = pop(stack)
			q2, stack = pop(stack)
			stack = append(stack, q1.combine(q2))
		case or:
			q1, stack = pop(stack)
			q2, stack = pop(stack)
			stack = append(stack, q1.pipe(q2))
		default:
			k, v := parseKeyValue(token)
			q := simpleQuery(k, v)
			stack = append(stack, q)
		}
	}
	return stack[0]
}

func parseKeyValue(kv string) (string, string) {
	k, v, _ := strings.Cut(kv, ":")
	return k, v
}
