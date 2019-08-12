package compileT

import (
	"github.com/pkg/errors"
	"fmt"
)


type Lexer struct {
	TokenArr []Token // 最终解析完成的token数组
	DfaManager  *DfaManager
	Scanner *Scanner
}


func NewLexer(scanner *Scanner, DfaManager *DfaManager) *Lexer {
	return &Lexer{
		Scanner:scanner,
		DfaManager:DfaManager,
		TokenArr:make([]Token,0),
	}
}


func (lexer *Lexer) Exec() error {

	for {
		c := lexer.Scanner.NextChar()
		if c == 0 {
			break
		}
		liters := ""
		curDfa := lexer.DfaManager.GetStartDfa(c)
		if curDfa == nil {
			return errors.New(fmt.Sprintf("index:%d char:%c, 不符合DFA。", lexer.Scanner.nextIndex-1, c))
		}
		liters += fmt.Sprintf("%c", c)

		for {
			nc := lexer.Scanner.NextChar()
			if nc == 0 {
				if curDfa.IsTerminal {
					lexer.AddToken(liters)
					return nil
				} else {
					return errors.New("不满足1")
				}
			}
			// 还能匹配
			nextDfa := curDfa.MachChar(nc)
			if nextDfa != nil {
				liters += fmt.Sprintf("%c", nc)
				curDfa = nextDfa
			} else { // 下一个就不匹配了
				if curDfa.IsTerminal {
					lexer.AddToken(liters)
					lexer.Scanner.RollBack()
					break
				} else {
					return errors.New("不满足2")
				}
			}
		}
	}
	return nil
}


func(lexer *Lexer) AddToken(expStr string) {
	token := Token{
		Val:expStr,
		Typ: String(expStr).GetTokenType(),
	}
	lexer.TokenArr = append(lexer.TokenArr, token)
}