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

func (p *Parser) Parse() (*Matcher, error) {
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


func (p *Parser) ParseEscape() (*Matcher, error) {
    ch, _, err := p.reader.ReadRune()
    
    if err != nil {
        return nil, err
    }

    switch ch {
    }
    return nil, nil
}
