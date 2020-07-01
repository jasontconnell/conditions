package conditions

import "fmt"

type Comparison int

var ParseError error = fmt.Errorf("Parse error")

type Node interface {
	Literal() string
	String() string
}

const (
	Undefined Comparison = iota
	Equals
	NotEquals
	LessThan
	LessThanEqual
	GreaterThan
	GreaterThanEqual
	Contains
)

type Op int

const (
	AND Op = iota
	OR
)

type Condition struct {
	Expression *Expression
}

type Expression interface {
	Node
	expression()
}

type IdentifierExpression struct {
	token token
	lit   string
	Value string
}

type StringExpression struct {
	token token
	lit   string
	Value string
}

type ComparisonExpression struct {
	token      token
	lit        string
	Value      string
	Identifier *IdentifierExpression
	Comparison Comparison
	String     *StringExpression
}

type Clause struct {
	Property   string
	Comparison Comparison
	Value      string
}
