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