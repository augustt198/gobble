package gobble

const(
    // quantifiers
    KLEENE_STAR = iota
    QUESTION
  
    // matching    
    META_CHAR
    LITERAL

    // grouping
    L_BRACKET
    R_BRACKET
    L_PAREN
    R_PAREN
)

type Token struct {
    Content string
    Type int
    Start int
    End int
}
