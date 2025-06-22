# Tic-Tac-Go

A straightforward implementation of noughts and crosses in Go, developed as an exercise in learning the language fundamentals.

## Overview

This project demonstrates Go's core features through a classic game implementation. The programme offers both two-player and single-player modes, with the latter featuring an optimal computer opponent.

## Technical Details

### Game Mechanics
The game employs a standard 3×3 grid with coordinate-based input. Players enter their moves using row and column numbers (1-3), with validation to prevent illegal moves. The board state is maintained using a simple two-dimensional array structure.

### Computer Opponent
The AI player utilises the minimax algorithm to evaluate all possible game states and select optimal moves. This recursive approach guarantees the computer will never lose—at best, a skilled human player can achieve a draw. The algorithm assigns scores to terminal game states (win, loss, draw) and propagates these values up the decision tree to determine the best available move.

## Requirements
- Go 1.21 or later

## Running the Game
```bash
go run *.go
```

Alternatively, compile and execute:
```bash
go build -o tictacgo
./tictacgo
```

## Usage
Upon starting, you'll be prompted to select single-player or two-player mode. Moves are entered as space-separated coordinates (e.g., "2 3" for row 2, column 3). In single-player mode, you play as X whilst the computer plays as O.

## Development Notes
This project was developed incrementally to explore Go's syntax, structure, and standard library. The commit history reflects a typical learning progression: establishing basic game logic, implementing win detection, and finally adding the AI component.