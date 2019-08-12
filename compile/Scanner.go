package compileT


type Scanner struct {
	strBuf  []byte
	nextIndex int
	strLen  int // 字符的总长度
}

func NewScanner(str string) *Scanner {
	scanner := &Scanner{
		strBuf:[]byte(str),
	}
	scanner.strLen = len(scanner.strBuf)
	return scanner
}

func(scan *Scanner) NextChar() (byte) {

	for {
		if scan.nextIndex >= scan.strLen {
			return 0
		}
		c := scan.strBuf[scan.nextIndex]
		scan.nextIndex++

		if c == ' ' || c == '\n' {
			continue
		}
		return c
	}

}

func(scan *Scanner) RollBack() {
	scan.nextIndex--
	if scan.nextIndex < 0 {
		scan.nextIndex = 0
	}
}
