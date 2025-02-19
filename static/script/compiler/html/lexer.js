
class Lexer {
  constructor(input) {
    this.input = input;
    this.position = 0;
  }

  nextToken() {
    if (this.position >= this.input.length) return null;

    let match;
    const text = this.input.slice(this.position);

    // Heading
    if (match = text.match(/^(#{1,6})\s+(.*)\n?/)) {
      this.position += match[0].length;
      return { type: 'HEADING', level: match[1].length, text: match[2] };
    }

    // Bold
    if (match = text.match(/^\*\*(.*?)\*\*/)) {
      this.position += match[0].length;
      return { type: 'BOLD', text: match[1] };
    }

    // Italic
    if (match = text.match(/^\*(.*?)\*/)) {
      this.position += match[0].length;
      return { type: 'ITALIC', text: match[1] };
    }

    // Code Block
    if (match = text.match(/^```(.*?)\n([\s\S]*?)```/)) {
      this.position += match[0].length;
      return { type: 'CODEBLOCK', language: match[1], content: match[2] };
    }

    // Inline Code
    if (match = text.match(/^`(.*?)`/)) {
      this.position += match[0].length;
      return { type: 'INLINECODE', text: match[1] };
    }

    // Blockquote
    if (match = text.match(/^>\s+(.*)\n?/)) {
      this.position += match[0].length;
      return { type: 'BLOCKQUOTE', text: match[1] };
    }

    // List Item
    if (match = text.match(/^-\s+(.*)\n?/)) {
      this.position += match[0].length;
      return { type: 'LISTITEM', text: match[1] };
    }

    // Text
    if (match = text.match(/^(.+?)\n?/)) {
      this.position += match[0].length;
      return { type: 'TEXT', text: match[1] };
    }

    return null;
  }
}
