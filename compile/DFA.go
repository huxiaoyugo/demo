package compileT

type charType uint

const (
	_ charType  = iota

	Eps     //
	Char
	None
)

type Dfa struct {
	index int // DFA编号
	charSet  []byte // 字符集合
	cTyp  charType
	NextDfa  *Dfa
	IsTerminal bool // 是否是完成节点
	IsStart  bool // 是否是头结点
}

func(dfa *Dfa)MachChar(char byte) *Dfa {
	for _, item := range dfa.charSet {
		if item == char {
			return dfa.NextDfa
		}
	}
	return nil
}





