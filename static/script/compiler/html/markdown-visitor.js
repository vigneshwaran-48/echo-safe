class MarkdownToHTML {
  visitDocument(node) {
    return node.children.map(child => child.accept(this)).join('');
  }
  visitHeading(node) {
    return `<h${node.level}>${node.text}</h${node.level}>`;
  }
  visitBold(node) {
    return `<b>${node.text}</b>`;
  }
  visitItalic(node) {
    return `<i>${node.text}</i>`;
  }
  visitBlockquote(node) {
    return `<blockquote>${node.text}</blockquote>`;
  }
  visitListItem(node) {
    return `<li>${node.text}</li>`;
  }
  visitCodeBlock(node) {
    return `<pre><code>${node.content}</code></pre>`;
  }
  visitInlineCode(node) {
    return `<code>${node.text}</code>`;
  }
  visitText(node) {
    if (node.inParagraph) {
      return node.text;
    } else {
      return `<p>${node.text}</p>`;
    }
  }
  visitParagraphStart() {
    return `<p>`;
  }
  visitParagraphContent(node) {
    let content = ``;
    node.children.forEach(child => {
      content += child.accept(this);
    })
    content += `</p>` // Will this work in all cases?
    return content;
  }
}
