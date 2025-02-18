package compiler

/*
Grammer file for the markdown.

Document ::= (BlockElement | BlankLine)*

BlockElement ::= Heading | Paragraph | Blockquote | List | CodeBlock | Table

Heading ::= ('#' Level) Whitespace Text Newline
Level ::= '#' | '##' | '###' | '####' | '#####' | '######'

Paragraph ::= InlineElement+ Newline

Blockquote ::= '>' Whitespace (BlockElement | InlineElement)+ Newline

List ::= (UnorderedList | OrderedList)
UnorderedList ::= ('-' | '*' | '+') Whitespace ListItem Newline
OrderedList ::= Digit+ '.' Whitespace ListItem Newline
ListItem ::= InlineElement* (Newline List)? (* Nested lists *)

CodeBlock ::= '```' Language? Newline CodeContent '```' Newline
Language ::= [a-zA-Z]+
CodeContent ::= (!'```' AnyChar)* (* Everything inside a code block *)

Table ::= TableRow+ Newline
TableRow ::= '|' (Text '|')+ Newline

InlineElement ::= Bold | Italic | Strikethrough | InlineCode | Link | Image | Text

Bold ::= '**' Text '**' | '__' Text '__'
Italic ::= '*' Text '*' | '_' Text '_'
Strikethrough ::= '~~' Text '~~'
InlineCode ::= '`' Text '`'
Link ::= '[' Text ']' '(' URL ')'
Image ::= '![' Text ']' '(' URL ')'

Text ::= (AnyChar - SpecialChar)+
SpecialChar ::= ('#' | '*' | '_' | '-' | '`' | '[' | ']' | '(' | ')' | '>' | '|')

BlankLine ::= Newline
Whitespace ::= (' ' | '	')+
Newline ::= '\n' | '\r\n'
Digit ::= '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
URL ::= ('http://' | 'https://') AnyChar+

*/

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
