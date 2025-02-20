
class Lexer {
  constructor(input) {
    this.input = input;
    this.position = 0;
    this.inParagraph = false;
  }

  nextChar() {
    return this.position < this.input.length ? this.input[this.position++] : null;
  }

  peekChar() {
    return this.position < this.input.length ? this.input[this.position] : null;
  }

  peekNextChar() {
    return this.position + 1 < this.input.length ? this.input[this.position + 1] : null;
  }

  isInlineMarker(char) {
    return char === '*' || char === '`';
  }

  tokenizeParagraphContent() {
    const tokens = [];
    let textBuffer = '';

    while (this.position < this.input.length) {
      const char = this.peekChar();

      // End of paragraph on double newline
      if (char === '\n' && this.peekNextChar() === '\n') {
        if (textBuffer) {
          tokens.push({
            type: 'TEXT',
            text: textBuffer,
            inParagrap: true
          });
        }
        this.inParagraph = false;
        this.position += 2; // Skip both newlines
        break;
      }

      // Single newline continues paragraph
      if (char === '\n') {
        textBuffer += ' '; // Convert newline to space
        this.nextChar();
        continue;
      }

      // Bold
      if (char === '*' && this.peekNextChar() === '*') {
        if (textBuffer) {
          tokens.push({
            type: 'TEXT',
            text: textBuffer,
            inParagraph: true
          });
          textBuffer = '';
        }
        this.nextChar(); // Skip first *
        this.nextChar(); // Skip second *
        let boldText = '';
        while (this.position < this.input.length) {
          if (this.input[this.position] === '*' && this.input[this.position + 1] === '*') {
            this.position += 2;
            break;
          }
          boldText += this.input[this.position++];
        }
        tokens.push({
          type: 'BOLD',
          text: boldText,
          inParagraph: true
        });
        continue;
      }

      // Italic
      if (char === '*' && this.peekNextChar() !== '*') {
        if (textBuffer) {
          tokens.push({
            type: 'TEXT',
            text: textBuffer,
            inParagraph: true
          });
          textBuffer = '';
        }
        this.nextChar(); // Skip *
        let italicText = '';
        while (this.position < this.input.length) {
          if (this.input[this.position] === '*') {
            this.position++;
            break;
          }
          italicText += this.input[this.position++];
        }
        tokens.push({
          type: 'ITALIC',
          text: italicText,
          inParagraph: true
        });
        continue;
      }

      // Inline Code
      if (char === '`') {
        if (textBuffer) {
          tokens.push({
            type: 'TEXT',
            text: textBuffer,
            inParagraph: true
          });
          textBuffer = '';
        }
        this.nextChar(); // Skip `
        let codeText = '';
        while (this.position < this.input.length) {
          if (this.input[this.position] === '`') {
            this.position++;
            break;
          }
          codeText += this.input[this.position++];
        }
        tokens.push({
          type: 'INLINECODE',
          text: codeText,
          inParagraph: true
        });
        continue;
      }

      textBuffer += this.nextChar();
    }

    // Handle any remaining text
    if (textBuffer) {
      tokens.push({
        type: 'TEXT',
        text: textBuffer,
        inParagraph: true
      });
    }

    return tokens;
  }

  nextToken() {

    if (this.position >= this.input.length) return null;

    // If we're in a paragraph, continue parsing paragraph content
    if (this.inParagraph) {
      const paragraphTokens = this.tokenizeParagraphContent();
      if (paragraphTokens.length > 0) {
        return {
          type: 'PARAGRAPH_CONTENT',
          children: paragraphTokens
        };
      }
    }

    let char = this.nextChar();

    // Skip whitespace
    while (char === ' ' || char === '\t' || char === '\n' || char === '\r') {
      char = this.nextChar();
      if (!char) return null;
    }

    // Heading
    if (char === '#') {
      let level = 1;
      while (this.peekChar() === '#') {
        this.nextChar();
        level++;
      }
      if (this.peekChar() === ' ') {
        this.nextChar();
        let text = this.readUntilNewline();
        return {
          type: 'HEADING',
          level,
          text
        };
      }
    }

    // Blockquote
    if (char === '>') {
      if (this.peekChar() === ' ') this.nextChar();
      let text = this.readUntilNewline();
      return {
        type: 'BLOCKQUOTE',
        text
      };
    }

    // List Item
    if (char === '-' && this.peekChar() === ' ') {
      this.nextChar();
      let text = this.readUntilNewline();
      return {
        type: 'LISTITEM',
        text
      };
    }

    // Code Block
    if (char === '`' && this.peekChar() === '`' && this.input[this.position + 1] === '`') {
      this.nextChar();
      this.nextChar();
      let language = this.readUntilNewline();
      let content = this.readUntil('```');
      return {
        type: 'CODEBLOCK',
        language,
        content
      };
    }

    // Start of paragraph
    this.position--;
    this.inParagraph = true;
    return {
      type: 'PARAGRAPH_START'
    };
  }

  readUntil(stop) {
    let start = this.position;
    while (this.position < this.input.length) {
      if (this.input.startsWith(stop, this.position)) {
        let result = this.input.slice(start, this.position);
        this.position += stop.length;
        return result;
      }
      this.position++;
    }
    return this.input.slice(start);
  }

  readUntilNewline() {
    let start = this.position;
    while (this.position < this.input.length && this.input[this.position] !== '\n') {
      this.position++;
    }
    return this.input.slice(start, this.position);
  }
}
