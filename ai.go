package main

func (g *Game) Minimax(isMaximizing bool, depth int) int {
	winner := g.CheckWinner()
	
	if winner == "O" {
		return 1
	} else if winner == "X" {
		return -1
	} else if g.IsDraw() {
		return 0
	}
	
	if isMaximizing {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if g.board[i][j] == " " {
					g.board[i][j] = "O"
					score := g.Minimax(false, depth+1)
					g.board[i][j] = " "
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if g.board[i][j] == " " {
					g.board[i][j] = "X"
					score := g.Minimax(true, depth+1)
					g.board[i][j] = " "
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
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