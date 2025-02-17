package compiler

func Parse(tokens []Token) ASTNode {
	root := ASTNode{
		nodeType: "root",
		value:    "",
		children: []ASTNode{},
	}
	current := root

	for _, token := range tokens {
		switch token.tokenType {
		case "heading":
			root.children = append(root.children, ASTNode{
				nodeType: "heading",
				value:    token.text,
				level:    token.level,
			})
		case "bold":
			root.children = append(root.children, ASTNode{
				nodeType: "bold",
				value:    token.text,
			})
		case "list_item":
			if current.nodeType != "list" {
				listNode := ASTNode{nodeType: "list", children: []ASTNode{}}
				root.children = append(root.children, listNode)
				current = listNode
			}
			current.children = append(current.children, ASTNode{
				nodeType: "list_item",
				value:    token.text,
			})
		default:
			root.children = append(root.children, ASTNode{nodeType: "paragraph", value: token.text})
		}
	}

	return root
}
