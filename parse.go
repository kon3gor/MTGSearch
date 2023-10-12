package mtgsearch

import (
	"strings"
)

const (
	delim = 32 // ASCII whitespace
	quote = 34 // ASCII quote
	lepa  = 40 // ASCII (
	repa  = 41 // ASCII )
	col   = 58 // ASCII :
	newl  = 10 // ASCII \n
	cor   = 13 // Ascii \r
)

func parseExpression(expr string) []string {
	ostack := make([]string, 0)
	rstack := make([]string, 0)
	tokens := tokenize(expr)

	for _, token := range tokens {
		switch token {
		case ")":
			var top string
			for len(ostack) > 0 {
				top, ostack = pop(ostack)
				if top != "(" {
					rstack = append(rstack, top)
				}
			}
		default:
			if isOperator(token) {
				var top string
				for shouldPopTopOperator(ostack, token) {
					top, ostack = pop(ostack)
					rstack = append(rstack, top)
				}
				ostack = append(ostack, token)
			} else if token == "(" {
				ostack = append(ostack, token)
			} else {
				rstack = append(rstack, token)
			}
		}
	}

	var v string
	for len(ostack) > 0 {
		v, ostack = pop(ostack)
		rstack = append(rstack, v)
	}

	return rstack
}

func tokenize(expr string) []string {
	buf := strings.Builder{}
	hasOpenQuote := false
	metColon := false
	tokens := make([]string, 0, len(expr)/2)
	runes := []rune(expr)
	var i int
	for i < len(runes) {
		c := runes[i]
		switch c {
		case lepa:
			tokens = append(tokens, "(")
		case repa:
			s := buf.String()
			if s != "" { // quick fix for "... <key>:<value>)"
				tokens = append(tokens, s)
				buf.Reset()
			}
			tokens = append(tokens, ")")
		case delim:
			if hasOpenQuote {
				buf.WriteRune(delim)
				break
			}
			s := buf.String()
			if s != "" {
				if s == and || s == or {
					tokens = append(tokens, s)
					buf.Reset()
				} else {
					if s[len(s)-1] == 58 { // string ends with colon, meaning there is smth like "type: elf"
						for runes[i+1] == delim { // just skip all spaces
							i++
						}
					} else if metColon { // Normal string type:elf
						tokens = append(tokens, s)
						buf.Reset()
						metColon = false
					} else { // smth like type : elf
						for runes[i+1] == delim { // just skip all spaces
							i++
						}
					}
				}
			}
		case col:
			metColon = true
			buf.WriteRune(c)
		case quote:
			hasOpenQuote = !hasOpenQuote
		case cor, newl:
			break
		default:
			buf.WriteRune(c)
		}
		i++
	}

	l := buf.String()
	if l != "" {
		tokens = append(tokens, l)
	}

	return tokens
}

func isOperator(s string) bool {
	return s == and || s == or
}

func shouldPopTopOperator(s []string, cur string) bool {
	if len(s) == 0 {
		return false
	} else {
		top := s[len(s)-1]
		return (top == and && cur == or) || top == cur
	}
}
