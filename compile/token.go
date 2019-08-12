package compileT

type tokenType uint

const (
	_        tokenType = iota
	Number    // 数字 0-9
	Operator  // 操作符
	Lparen    // (
	Rparen    // )
)

type Operation uint

const (
	_        Operation = iota
	Add       // +
	Minus     // -
	Mutiply   // *
	Division  // /
)

type Token struct {
	Typ tokenType
	Val interface{}
}


type String string

func (str String) GetTokenType() tokenType {
	if len(str) == 1 {
		switch str {
		case "+","-","*","\\":
			return Operator
		case "(":
			return Lparen
		case ")":
			return Rparen
		}
	}
	return Number
}