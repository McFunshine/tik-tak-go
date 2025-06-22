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

	game := NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		game.board.Display()
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
}