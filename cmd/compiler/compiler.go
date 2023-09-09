package compiler

import "fmt"

func Compile(source string) {
	scanner := NewScanner([]rune(source))
	line := -1
	for {
		token := scanner.scanToken()
		if token.line != line {
			fmt.Printf("%4d ", token.line)
			line = token.line
		} else {
			fmt.Printf("   | ")
		}

		fmt.Printf("%2d '%.*s'\n", token.kind, token.length, token.start)

		if token.kind == TOKEN_EOF {
			break
		}
	}
}
