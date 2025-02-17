package compiler

import (
	"fmt"
	"strings"
)

func GenerateHTML(astNode ASTNode) string {
	var htmlBuilder strings.Builder
	for _, child := range astNode.children {
		htmlBuilder.WriteString(visit(child))
	}
	return htmlBuilder.String()
}

func visit(astNode ASTNode) string {
	switch astNode.nodeType {
	case "heading":
		return fmt.Sprintf("<h%d>%s</h%d>", astNode.level, astNode.value, astNode.level)
	case "bold":
		return fmt.Sprintf("<b>%s</b>", astNode.value)
	case "list":
		listBuilder := strings.Builder{}
		listBuilder.WriteString("<ul>")
		for _, child := range astNode.children {
			listBuilder.WriteString(fmt.Sprintf("<li>%s</li>", child.value))
		}
		listBuilder.WriteString("</ul>")
		return listBuilder.String()
	default:
		return fmt.Sprintf("<p>%s</p>", astNode.value)
	}
}
