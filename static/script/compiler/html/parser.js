class Parser {
  constructor(tokens) {
    this.tokens = tokens;
    this.position = 0;
  }

  parse() {
    let children = [];
    while (this.position < this.tokens.length) {
      children.push(this.parseNextToken());
    }
    return new DocumentNode(children);
  }

  parseNextToken() {
    const token = this.tokens[this.position++];
    return this.parseToken(token);
  }

  parseToken(token) {
    switch (token.type) {
      case 'HEADING': return new HeadingNode(token.level, token.text);
      case 'BOLD': return new BoldNode(token.text);
      case 'ITALIC': return new ItalicNode(token.text);
      case 'BLOCKQUOTE': return new BlockquoteNode(token.text);
      case 'LISTITEM': return new ListItemNode(token.text);
      case 'CODEBLOCK': return new CodeBlockNode(token.language, token.content);
      case 'INLINECODE': return new InlineCodeNode(token.text);
      case 'TEXT': return new TextNode(token.text, token.inParagraph);
      case 'PARAGRAPH_START': return new ParagraphStartNode();
      case 'PARAGRAPH_CONTENT':
        const childNodes = [];
        token.children.forEach(child => {
          childNodes.push(this.parseToken(child));
        })
        return new ParagraphContentNode(childNodes);
    }
  }
}
