package conditions

import "unicode/utf8"

type State int

const (
	Start State = iota
	Identifier
	OpenParen
	CloseParen
	String
	Compare
	Operator
	End
)

type lexer struct {
	input  string
	pos    int
	state  State
	tokens []token
}

type token struct {
	Value string
	Type  State
	Start int
}

func Lex(c string) []token {
	lex := lexer{state: Start, input: c, pos: 0}
	lex.getTokens()
	return lex.tokens
}

func (lex *lexer) getTokens() {
	for i := 0; i < len(lex.input); i++ {
		c := rune(lex.input[i])

		switch c {
		case '(':
			lex.tokens = append(lex.tokens, token{Value: string(c), Start: i, Type: OpenParen})
			break
		case ')':
			lex.tokens = append(lex.tokens, token{Value: string(c), Start: i, Type: CloseParen})
			break
		case '"', '\'':
			tk := token{Value: lex.handleQuote(c, i), Start: i, Type: String}
			ln := utf8.RuneCountInString(tk.Value)

			lex.tokens = append(lex.tokens, tk)
			if ln > 0 {
				i += ln + 1
			} else {
				i += 1
			}
			break
		case ' ', '\n', '\t', '\r':
			break
		default:
			tk := token{Value: lex.handleIdentifier(i), Start: i, Type: Identifier}
			lex.tokens = append(lex.tokens, tk)
			ln := utf8.RuneCountInString(tk.Value)
			if ln > 0 {
				i += ln - 1
			}
			break
		}
	}
}

func (l *lexer) handleIdentifier(pos int) string {
	ret := ""
	done := false
	for i := pos; i < len(l.input) && !done; i++ {
		c := rune(l.input[i])

		switch c {
		case ' ', '\n', '\r', '\t', '(', ')', '"', '\\', '\'', '=', ',', '.', ']', '[':
			done = true
			break
		default:
			ret = ret + string(c)
		}
	}
	return ret
}

func (l *lexer) handleQuote(q rune, pos int) string {
	ret := ""
	done := false
	isEscape := false
	for i := pos + 1; i < len(l.input) && !done; i++ {
		isClose := rune(l.input[i]) == q
		isEscape = rune(l.input[i-1]) == '\\' && isClose
		if (rune(l.input[i]) != q && !isEscape) || isEscape {
			ret += string(l.input[i])
		} else if isClose && !isEscape {
			done = true
		}
	}
	return ret
}
