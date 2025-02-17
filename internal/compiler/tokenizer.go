package compiler

import "strings"

func Tokenize(input string) ([]Token, error) {
	input = strings.TrimSpace(input)
	tokens := []Token{}
	markdown := []rune(input)
	index := 0

	for index < len(markdown) {

		character := markdown[index]

		if character == '#' {
			level := 1
			for markdown[index+level] == '#' {
				level++
			}
			index += level
			start := index
			for index < len(markdown) && markdown[index] != '\n' {
				index++
			}
			tokens = append(tokens, Token{tokenType: "heading", text: string(markdown[start:index]), level: level})
		} else if character == '*' && (index+1) < len(markdown) && markdown[index+1] == '*' {
			index += 2
			start := index
			for index < len(markdown) && (markdown[index] != '*' || ((index+1) < len(markdown)) && markdown[index+1] != '*') {
				index++
			}
			tokens = append(tokens, Token{tokenType: "bold", text: string(markdown[start:index])})
			index += 2
		} else if character == '-' && (index+1) < len(markdown) && markdown[index+1] == ' ' {
			index += 2
			start := index
			for index < len(markdown) && markdown[index] != '\n' {
				index++
			}
			tokens = append(tokens, Token{tokenType: "list_item", text: string(markdown[start:index])})
		} else if character == '\n' {
			index++
		} else {
			start := index
			for index < len(markdown) && markdown[index] != '\n' {
				index++
			}
			tokens = append(tokens, Token{tokenType: "paragraph", text: string(markdown[start:index])})
		}
		index++
	}
	return tokens, nil
}
