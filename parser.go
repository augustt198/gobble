package gobble

import "strings"

type Parser struct {
    reader *strings.Reader
    prev   *Matcher
}

func New(str string) *Parser {
    return &Parser{
        reader: strings.NewReader(str),
        prev:   nil}
}

func (p *Parser) Parse() (Matcher, error) {
    ch, _, err := p.reader.ReadRune()

    if err != nil {
        return nil, err
    }

    switch ch {
    case '\\':
        return p.ParseEscape()
    }

    return nil, nil
}

type UnknownEscapeSequence struct{}
func (e UnknownEscapeSequence) Error() string {
    return "Uknown escape sequence"
}

func (p *Parser) ParseEscape() (Matcher, error) {
    ch, _, err := p.reader.ReadRune()
    
    if err != nil {
        return nil, err
    }

    switch ch {
    case 'd':
        return &DigitMetaMatcher{}, nil
    case 'w':
        return &WordMetaMatcher{}, nil
    default:
        return nil, UnknownEscapeSequence{}
    }
}
