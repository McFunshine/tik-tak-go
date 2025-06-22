package main

import "fmt"

type Board [3][3]string

func NewBoard() Board {
	return Board{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func (b Board) Display() {
	fmt.Println("   |   |   ")
	fmt.Printf(" %s | %s | %s \n", b[0][0], b[0][1], b[0][2])
	fmt.Println("___|___|___")
	fmt.Println("   |   |   ")
	fmt.Printf(" %s | %s | %s \n", b[1][0], b[1][1], b[1][2])
	fmt.Println("___|___|___")
	fmt.Println("   |   |   ")
	fmt.Printf(" %s | %s | %s \n", b[2][0], b[2][1], b[2][2])
	fmt.Println("   |   |   ")
}