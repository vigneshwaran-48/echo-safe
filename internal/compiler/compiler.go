package compiler

type Token struct {
	tokenType string
	text      string
	level     int
}

type ASTNode struct {
	nodeType string
	value    string
	children []ASTNode
	level    int
}

func Compile(markdown string) (string, error) {
	tokens, err := Tokenize(markdown)
	if err != nil {
		return "", err
	}
	astNode := Parse(tokens)
	return GenerateHTML(astNode), nil
}
