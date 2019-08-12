package compileT

type DfaManager struct {
	StartDfa []*Dfa // 所有的头结点
	index int
}


func NewDfaManager() *DfaManager {
	return &DfaManager{
		StartDfa: make([]*Dfa,0),
		index: 0,
	}
}

func(manager *DfaManager)NewDfa(charSet []byte, cTyp charType,isTerminal, isStart bool) *Dfa {
	manager.index ++
	dfa := &Dfa{
		index: manager.index,
		charSet:charSet,
		cTyp:cTyp,
		IsStart:isStart,
		IsTerminal:isTerminal,
	}
	if isStart {
		manager.StartDfa = append(manager.StartDfa, dfa)
	}
	return dfa
}


func (manager *DfaManager) GetStartDfa(char byte) *Dfa {
	for _, dfa := range manager.StartDfa {
		nt := dfa.MachChar(char)
		if nt != nil {
			return nt
		}
	}
	return nil
}
