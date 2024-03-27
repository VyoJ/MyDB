package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SELECT string = "SELECT"
	FROM   string = "FROM"
	WHERE  string = "WHERE"
	CREATE string = "CREATE"
	ALTER  string = "ALTER"
	TABLE  string = "TABLE"
	INSERT string = "INSERT"
	INTO   string = "INTO"
	VALUES string = "VALUES"
	AS     string = "AS"
	IDENT  string = "IDENT"
	AND    string = "AND"
	OR     string = "OR"
	NUMBER string = "NUMBER"
)

const (
	SEMICOLON string = "SEMICOLON"
	ASTERISK  string = "ASTERISK"
	COMMA     string = "COMMA"
	LBRACKET  string = "LBRACKET"
	RBRACKET  string = "RBRACKET"
	GT        string = "GT"
	LT        string = "LT"
	GTE       string = "GTE"
	LTE       string = "LTE"
	EQ        string = "EQ"
)

func lex(input string) []string {
	var tokens []string

	words := strings.Fields(input)

	for _, word := range words {
		switch word {
		case "select":
			tokens = append(tokens, SELECT)
		case "from":
			tokens = append(tokens, FROM)
		case "where":
			tokens = append(tokens, WHERE)
		case "create":
			tokens = append(tokens, CREATE)
		case "alter":
			tokens = append(tokens, ALTER)
		case "table":
			tokens = append(tokens, TABLE)
		case "insert":
			tokens = append(tokens, INSERT)
		case "into":
			tokens = append(tokens, INTO)
		case "values":
			tokens = append(tokens, VALUES)
		case "as":
			tokens = append(tokens, AS)
		case "*":
			tokens = append(tokens, ASTERISK)
		case ",":
			tokens = append(tokens, COMMA)
		case ";":
			tokens = append(tokens, SEMICOLON)
		case "(":
			tokens = append(tokens, LBRACKET)
		case ")":
			tokens = append(tokens, RBRACKET)
		case ">":
			tokens = append(tokens, GT)
		case "<":
			tokens = append(tokens, LT)
		case ">=":
			tokens = append(tokens, GTE)
		case "<=":
			tokens = append(tokens, LTE)
		case "=":
			tokens = append(tokens, EQ)
		default:
			_, err := strconv.Atoi(word)
			if err != nil {
				tokens = append(tokens, IDENT)
			} else {
				tokens = append(tokens, NUMBER)
			}

		}
	}
	return tokens
}

func main() {
	fmt.Println("Enter your SQL query: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// query = "SELECT * FROM students WHERE rollno > 3 ;"
	tokens := lex(strings.ToLower(scanner.Text()))

	for _, token := range tokens {
		fmt.Println(token)
	}
}
