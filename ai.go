package main

func (g *Game) Minimax(isMaximizing bool, depth int) int {
	// Check terminal states
	winner := g.CheckWinner()
	switch winner {
	case "O":
		return 1
	case "X":
		return -1
	}
	
	if g.IsDraw() {
		return 0
	}
	
	// Minimax evaluation
	if isMaximizing {
		return g.evaluateMax()
	}
	return g.evaluateMin()
}

func (g *Game) evaluateMax() int {
	bestScore := -1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] != " " {
				continue
			}
			g.board[i][j] = "O"
			score := g.Minimax(false, 0)
			g.board[i][j] = " "
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func (g *Game) evaluateMin() int {
	bestScore := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] != " " {
				continue
			}
			g.board[i][j] = "X"
			score := g.Minimax(true, 0)
			g.board[i][j] = " "
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func (g *Game) GetBestMove() [2]int {
	bestScore := -1000
	bestMove := [2]int{-1, -1}
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == " " {
				g.board[i][j] = "O"
				score := g.Minimax(false, 0)
				g.board[i][j] = " "
				
				if score > bestScore {
					bestScore = score
					bestMove = [2]int{i, j}
				}
			}
		}
	}
	
	return bestMove
}