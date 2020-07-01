package conditions

type parser struct {
	tokens       []token
	index        int
	currentToken token
	nextToken    token
}

func NewParser(tokens []token) *parser {
	p := &parser{tokens: tokens, index: -1}
	p.moveFirst()
	return p
}
func (p *parser) Parse() ([]Expression, error) {
	statements := []Expression{}
	done := false
	for !done {
		statement, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if statement != nil {
			statements = append(statements, statement)
		}
		done = p.currentToken.Type == End
	}
	return statements, nil
}

func (p *parser) parseExpression() (Expression, error) {
	var st Expression
	var err error
	switch p.currentToken.Type {
	case Identifier:
		switch p.peekType() {
		case OpenParen:
			st, err = p.parseExpression()
		default:
			st, err = p.parseIdentifier()
		}
		break
	case String:
		st, err = p.parseString()
	case lex.End:
		break
	default:
		p.moveNext()
	}
	return st, err
}

func (p *parser) parseIdentifier() (*IdentifierExpression, error) {
	ident := &IdentifierExpression{token: p.currentToken, lit: p.currentToken.Value, Value: p.currentToken.Value}
	p.moveNext()
	return ident, nil
}

func (p *parser) moveFirst() {
	p.index = -1
	p.moveNext()
}

func (p *parser) moveNext() {
	p.index++
	p.currentToken = p.tokens[p.index]
	if len(p.tokens) > p.index+1 {
		p.nextToken = p.tokens[p.index+1]
	}
}

func (p *parser) peekType() State {
	return p.nextToken.Type
}

func (p *parser) curType() State {
	return p.currentToken.Type
}

func getComp(s string) Comparison {
	c := Undefined
	switch s {
	case "Equals":
		c = Equals
	case "NotEquals":
		c = NotEquals
	case "LessThan":
		c = LessThan
	case "LessThanEqual":
		c = LessThanEqual
	case "GreaterThan":
		c = GreaterThan
	case "GreaterThanEqual":
		c = GreaterThanEqual
	case "Contains":
		c = Contains
	}
	return c
}
