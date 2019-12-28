package main

func tictactoe(moves [][]int) string {

	gameMap := [3][3]string{}

	for index, step := range moves {

		if index%2 == 0 {
			gameMap[step[0]][step[1]] = "X"
		} else {
			gameMap[step[0]][step[1]] = "O"
		}
	}

	return checkGameStatus(gameMap)
}

func checkGameStatus(gameMap [3][3]string) string {

	for _, line := range LineList {
		full, win, winner := checkLine(gameMap, line)
		if win {
			return winner
		}

		if full {
			return "Draw"
		}
	}
	return "Pending"
}

func checkLine(gameMap [3][3]string, line [3][2]int) (full bool, win bool, winer string) {

	s := gameMap[line[0][0]][line[0][1]]

	if s == "" {
		return false, false, ""
	}
	win = true
	for i := 1; i < 3; i++ {
		c := gameMap[line[i][0]][line[i][1]]
		if c == "" {
			return false, false, ""
		}
		if s != c {
			win = false
		}
	}
	if win {
		winer = getOperoter(s)
	}
	return
}

func getOperoter(c string) string {
	if c == "X" {
		return "A"
	} else if c == "O" {
		return "B"
	}
	return ""
}

var LineList = [8][3][2]int {
{{0, 0},{0, 1},{0, 2}},
{{1, 0},{1, 1},{1, 2}},
{{2, 0},{2, 1},{2, 2}},

{{0, 0},{1, 0},{2, 0}},
{{0, 1},{1, 1},{2, 1}},
{{0, 2},{1, 2},{2, 2}},

{{0, 0},{1, 1},{2, 2}},
{{0, 2},{1, 1},{2, 0}},
}
