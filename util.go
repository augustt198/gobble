package gobble

type StringReader struct {
    runes []rune
    pos int
}

func NewReader(str string) *StringReader {
    return &StringReader{[]rune(str), 0 }
}

func (s *StringReader) IsEOS() bool {
    return s.pos >= len(s.runes)
}

func (s *StringReader) Read() (rune, bool) {
    if s.IsEOS() {
        return 0, true
    }
    ch := s.runes[s.pos]
    s.pos++
    return ch, false
}

func (s *StringReader) Position() int {
    return s.pos
}

func (s *StringReader) Reset(pos int) {
    s.pos = pos
}
