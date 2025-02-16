package compiler

func Tokenize(input string) ([]Token, error) {
	tokens := []Token{}
	markdown := []rune(input)
	index := 0

	for index < len(markdown) {

		character := markdown[index]

		if character == '#' && markdown[index+1] == '#' {
			index++
			level := 1
			for markdown[index+level] == '#' {
				level++
			}
			index += level
			start := index
			for index < len(markdown) && markdown[index] != '\n' {
				index++
			}
			tokens = append(tokens, Token{tokenType: "heading", text: string(markdown[start:index])})
		}
	}
	return tokens, nil
}
