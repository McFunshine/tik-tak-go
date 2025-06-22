package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Tic-Tac-Go")
	fmt.Println("==========")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Play against computer? (y/n): ")
	scanner.Scan()
	vsComputer := strings.ToLower(scanner.Text()) == "y"
	
	var game *Game
	if vsComputer {
		game = NewGameVsComputer()
		fmt.Println("\nYou are X, computer is O")
	} else {
		game = NewGame()
	}

	for {
		game.board.Display()
		
		if game.vsComputer && game.currentPlayer == "O" {
			fmt.Println("\nComputer's turn...")
			game.MakeComputerMove()
		} else {
			fmt.Printf("\nPlayer %s's turn\n", game.currentPlayer)
			fmt.Print("Enter row and column (e.g., 1 2): ")
			
			scanner.Scan()
			input := scanner.Text()
			parts := strings.Fields(input)
			
			if len(parts) != 2 {
				fmt.Println("Invalid input! Please enter row and column separated by space.")
				continue
			}
			
			row, err1 := strconv.Atoi(parts[0])
			col, err2 := strconv.Atoi(parts[1])
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input! Please enter numbers.")
				continue
			}
			
			if !game.MakeMove(row, col) {
				fmt.Println("Invalid move! Try again.")
				continue
			}
		}
		
		if game.IsGameOver() {
			game.board.Display()
			winner := game.CheckWinner()
			if winner != "" {
				fmt.Printf("\nPlayer %s wins!\n", winner)
			} else {
				fmt.Println("\nIt's a draw!")
			}
			break
		}
	}
}