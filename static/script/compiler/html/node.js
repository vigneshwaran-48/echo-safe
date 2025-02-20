class ASTNode {
  accept(visitor) {
    throw new Error("Accept method not implemented");
  }
}

class DocumentNode extends ASTNode {
  constructor(children) {
    super();
    this.children = children;
  }
  accept(visitor) {
    return visitor.visitDocument(this);
  }
}

class HeadingNode extends ASTNode {
  constructor(level, text) {
    super();
    this.level = level;
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitHeading(this);
  }
}

class BoldNode extends ASTNode {
  constructor(text) {
    super();
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitBold(this);
  }
}

class ItalicNode extends ASTNode {
  constructor(text) {
    super();
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitItalic(this);
  }
}

class BlockquoteNode extends ASTNode {
  constructor(text) {
    super();
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitBlockquote(this);
  }
}

class ListItemNode extends ASTNode {
  constructor(text) {
    super();
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitListItem(this);
  }
}

class CodeBlockNode extends ASTNode {
  constructor(language, content) {
    super();
    this.language = language;
    this.content = content;
  }
  accept(visitor) {
    return visitor.visitCodeBlock(this);
  }
}

class InlineCodeNode extends ASTNode {
  constructor(text) {
    super();
    this.text = text;
  }
  accept(visitor) {
    return visitor.visitInlineCode(this);
  }
}

class TextNode extends ASTNode {
  constructor(text, inParagraph) {
    super();
    this.text = text;
    this.inParagraph = inParagraph;
  }
  accept(visitor) {
    return visitor.visitText(this);
  }
}

class ParagraphStartNode extends ASTNode {
  constructor() {
    super();
  }
  accept(visitor) {
    return visitor.visitParagraphStart();
  }
}

class ParagraphContentNode extends ASTNode {
  constructor(children) {
    super();
    this.children = children;
  }
  accept(visitor) {
    return visitor.visitParagraphContent(this);
  }
}
