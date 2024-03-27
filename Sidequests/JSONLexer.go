package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	JSON_QUOTE        = '"'
	JSON_COMMA        = ','
	JSON_COLON        = ':'
	JSON_LEFTBRACKET  = '['
	JSON_RIGHTBRACKET = ']'
	JSON_LEFTBRACE    = '{'
	JSON_RIGHTBRACE   = '}'
)

var (
	JSON_WHITESPACE = []rune{' ', '\t', '\b', '\n', '\r'}
	JSON_SYNTAX     = []rune{JSON_COMMA, JSON_COLON, JSON_LEFTBRACKET, JSON_RIGHTBRACKET, JSON_LEFTBRACE, JSON_RIGHTBRACE}

	FALSE_LEN = len("false")
	TRUE_LEN  = len("true")
	NULL_LEN  = len("null")
)

func lexString(s string) (string, string) {
	if s[0] != JSON_QUOTE {
		return "", s
	}
	s = s[1:]
	for i, c := range s {
		if c == JSON_QUOTE {
			return s[:i], s[i+1:]
		}
	}
	panic("Expected end-of-string quote")
}

func lexNumber(s string) (interface{}, string) {
	var jsonNumber strings.Builder
	numberCharacters := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', 'e', '.'}

	for _, c := range s {
		if contains(numberCharacters, c) {
			jsonNumber.WriteRune(c)
		} else {
			break
		}
	}

	rest := s[len(jsonNumber.String()):]
	if jsonNumber.Len() == 0 {
		return nil, s
	}

	if strings.ContainsRune(jsonNumber.String(), '.') {
		f, err := strconv.ParseFloat(jsonNumber.String(), 64)
		if err != nil {
			panic(err)
		}
		return f, rest
	}

	i, err := strconv.ParseInt(jsonNumber.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i, rest
}

func lexBool(s string) (interface{}, string) {
	if len(s) >= TRUE_LEN && s[:TRUE_LEN] == "true" {
		return true, s[TRUE_LEN:]
	} else if len(s) >= FALSE_LEN && s[:FALSE_LEN] == "false" {
		return false, s[FALSE_LEN:]
	}
	return nil, s
}

func lexNull(s string) (interface{}, string) {
	if len(s) >= NULL_LEN && s[:NULL_LEN] == "null" {
		return nil, s[NULL_LEN:]
	}
	return nil, s
}

func contains(slice []rune, item rune) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func lex(s string) []interface{} {
	var tokens []interface{}

	for len(s) > 0 {
		var token interface{}
		var rest string

		token, rest = lexString(s)
		if token != "" {
			tokens = append(tokens, token)
			s = rest
			continue
		}

		token, rest = lexNumber(s)
		if token != nil {
			tokens = append(tokens, token)
			s = rest
			continue
		}

		token, rest = lexBool(s)
		if token != nil {
			tokens = append(tokens, token)
			s = rest
			continue
		}

		token, rest = lexNull(s)
		if token != nil {
			tokens = append(tokens, token)
			s = rest
			continue
		}

		c := rune(s[0])
		if contains(JSON_WHITESPACE, c) {
			s = s[1:]
		} else if contains(JSON_SYNTAX, c) {
			tokens = append(tokens, string(c))
			s = s[1:]
		} else {
			panic(fmt.Sprintf("Unexpected character: %c", c))
		}
	}

	return tokens
}

func main() {
	tokens := lex(`{"name": "John", "age": 30, "city": "New York"}`)
	for _, token  := range tokens {
		fmt.Print("'", token, "', ")
	}
}
