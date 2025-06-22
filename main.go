package main

import "fmt"

func main() {
	fmt.Println("Tic-Tac-Go")
	fmt.Println("==========")
	fmt.Println()

	board := NewBoard()
	board.Display()
}