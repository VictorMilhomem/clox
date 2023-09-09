package compiler

import (
	"testing"
)

func TestIdentifierType(t *testing.T) {
	testCases := []struct {
		source        string
		expectedToken TokenType
	}{
		{"and", TOKEN_AND},
		{"class", TOKEN_CLASS},
		{"else", TOKEN_ELSE},
		{"false", TOKEN_FALSE},
		{"for", TOKEN_FOR},
		{"fun", TOKEN_FUN},
		{"if", TOKEN_IF},
		{"nil", TOKEN_NIL},
		{"or", TOKEN_OR},
		{"print", TOKEN_PRINT},
		{"return", TOKEN_RETURN},
		{"super", TOKEN_SUPER},
		{"this", TOKEN_THIS},
		{"true", TOKEN_TRUE},
		{"var", TOKEN_VAR},
		{"while", TOKEN_WHILE},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		// Create a Scanner instance with the source set to the current test case
		scanner := NewScanner([]rune(testCase.source + "\x00"))

		// Call the scanToken function
		tokenType := scanner.scanToken().kind

		// Assert the expected token type
		if tokenType != testCase.expectedToken {
			t.Errorf("For source '%s', expected token type %d, but got %d", testCase.source, testCase.expectedToken, tokenType)
		}
	}
}
