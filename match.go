package gobble

type Matcher interface {
    Match(*StringReader) bool
}

type WordMetaMatcher struct {}
func (_ *WordMetaMatcher) Match(e *StringReader) bool {
    ch, eos := e.Read()
    if eos {
        return false
    }
    return  ch >= 'a' && ch <= 'z' ||
            ch >= 'A' && ch <= 'Z' ||
            ch >= '0' && ch <= '9' ||
            ch == '_'
}

type DigitMetaMatcher struct {}
func (_ *DigitMetaMatcher) Match(e *StringReader) bool {
    ch, eos := e.Read()
    if eos {
        return false
    }
    return ch >= '0' && ch <= '9'
}

type LiteralMatcher struct {
    Ch rune
}
func (l *LiteralMatcher) Match(e *StringReader) bool {
    pos := e.Position()

    ch, eos := e.Read()
    if eos {
        return false
    }
    if ch == l.Ch {
        return true
    } else {
        e.Reset(pos)
        return false
    }
}

type StarMatcher struct {
    Token Matcher
}
func (s *StarMatcher) Match(e *StringReader) bool {
    pos := e.Position()

    for {
        match := s.Token.Match(e)
        if match {
            pos = e.Position()
        } else {
            e.Reset(pos)
            break
        }
    }

    return true
}

type PlusMatcher struct {
    Token Matcher
}
func (p *PlusMatcher) Match(e *StringReader) bool {
    pos := e.Position()
    any := false

    for {
        match := p.Token.Match(e)
        if match {
            pos = e.Position()
            any = true
        } else {
            e.Reset(pos)
            break
        }
    }

    return any
}

type AlternationMatcher struct {
    Left Matcher
    Right Matcher
}
func (a *AlternationMatcher) Match(e *StringReader) bool {
 
    pos := e.Position()
    if a.Left.Match(e) {
        return true
    }
    e.Reset(pos)
    if a.Right.Match(e) {
        return true
    }
    e.Reset(pos)
    return false
}

type ChainMatcher struct {
    Tokens []Matcher
}
func (c *ChainMatcher) Match(e *StringReader) bool {
    for _, Token := range c.Tokens {
        if !Token.Match(e) {
            return false
        }
    }
    return true
}
