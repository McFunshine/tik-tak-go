package main

import "fmt"

type Board [3][3]string

type Game struct {
	board       Board
	currentPlayer string
}

func NewGame() *Game {
	return &Game{
		board:       NewBoard(),
		currentPlayer: "X",
	}
}

func NewBoard() Board {
	return Board{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func (b Board) Display() {
	fmt.Println("\n  1   2   3")
	fmt.Println("1   |   |   ")
	fmt.Printf("  %s | %s | %s \n", b[0][0], b[0][1], b[0][2])
	fmt.Println(" ___|___|___")
	fmt.Println("2   |   |   ")
	fmt.Printf("  %s | %s | %s \n", b[1][0], b[1][1], b[1][2])
	fmt.Println(" ___|___|___")
	fmt.Println("3   |   |   ")
	fmt.Printf("  %s | %s | %s \n", b[2][0], b[2][1], b[2][2])
	fmt.Println("    |   |   ")
}

func (g *Game) MakeMove(row, col int) bool {
	if row < 1 || row > 3 || col < 1 || col > 3 {
		return false
	}
	
	if g.board[row-1][col-1] != " " {
		return false
	}
	
	g.board[row-1][col-1] = g.currentPlayer
	
	if g.currentPlayer == "X" {
		g.currentPlayer = "O"
	} else {
		g.currentPlayer = "X"
	}
	
	return true
}

func (g *Game) IsValidMove(row, col int) bool {
	if row < 1 || row > 3 || col < 1 || col > 3 {
		return false
	}
	return g.board[row-1][col-1] == " "
}

func (g *Game) CheckWinner() string {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] != " " && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			return g.board[i][0]
		}
	}
	
	// Check columns
	for j := 0; j < 3; j++ {
		if g.board[0][j] != " " && g.board[0][j] == g.board[1][j] && g.board[1][j] == g.board[2][j] {
			return g.board[0][j]
		}
	}
	
	// Check diagonals
	if g.board[0][0] != " " && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		return g.board[0][0]
	}
	if g.board[0][2] != " " && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		return g.board[0][2]
	}
	
	return ""
}

func (g *Game) IsDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func (g *Game) IsGameOver() bool {
	return g.CheckWinner() != "" || g.IsDraw()
}