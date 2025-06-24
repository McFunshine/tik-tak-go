package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Board represents the 3x3 game board.
type Board [3][3]string

// Game represents the game state including the board, current player, and game mode.
type Game struct {
	board       Board
	currentPlayer string
	vsComputer   bool
}

// NewGame creates a new two-player game instance.
func NewGame() *Game {
	return &Game{
		board:       NewBoard(),
		currentPlayer: "X",
		vsComputer:   false,
	}
}

// NewGameVsComputer creates a new game instance with computer opponent.
func NewGameVsComputer() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		board:       NewBoard(),
		currentPlayer: "X",
		vsComputer:   true,
	}
}

// NewBoard creates a new empty game board.
func NewBoard() Board {
	return Board{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

// Display prints the current board state to the console.
func (b Board) Display() {
	fmt.Println("\n    1   2   3")
	fmt.Println("  +---+---+---+")
	fmt.Printf("1 | %s | %s | %s |\n", b[0][0], b[0][1], b[0][2])
	fmt.Println("  +---+---+---+")
	fmt.Printf("2 | %s | %s | %s |\n", b[1][0], b[1][1], b[1][2])
	fmt.Println("  +---+---+---+")
	fmt.Printf("3 | %s | %s | %s |\n", b[2][0], b[2][1], b[2][2])
	fmt.Println("  +---+---+---+")
}

// MakeMove attempts to place the current player's mark at the specified position.
// Returns true if the move was valid and successful, false otherwise.
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

// IsValidMove checks if a move is valid at the specified position.
func (g *Game) IsValidMove(row, col int) bool {
	if row < 1 || row > 3 || col < 1 || col > 3 {
		return false
	}
	return g.board[row-1][col-1] == " "
}

// CheckWinner determines if there is a winner.
// Returns "X" or "O" if there is a winner, empty string otherwise.
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

// IsDraw checks if the game is a draw (board full with no winner).
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

// IsGameOver checks if the game has ended (either by win or draw).
func (g *Game) IsGameOver() bool {
	return g.CheckWinner() != "" || g.IsDraw()
}

// GetAvailableMoves returns a slice of all empty board positions.
func (g *Game) GetAvailableMoves() [][2]int {
	moves := make([][2]int, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == " " {
				moves = append(moves, [2]int{i, j})
			}
		}
	}
	return moves
}

// MakeComputerMove executes the computer's turn using the minimax algorithm.
func (g *Game) MakeComputerMove() {
	moves := g.GetAvailableMoves()
	if len(moves) == 0 {
		return
	}
	
	// Use minimax algorithm for optimal play
	bestMove := g.GetBestMove()
	g.board[bestMove[0]][bestMove[1]] = g.currentPlayer
	
	fmt.Printf("Computer plays: %d %d\n", bestMove[0]+1, bestMove[1]+1)
	
	if g.currentPlayer == "X" {
		g.currentPlayer = "O"
	} else {
		g.currentPlayer = "X"
	}
}