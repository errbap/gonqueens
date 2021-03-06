package main

import (
	"math/rand"
	"reflect"
)

// HillClimbingSolver is an interface that must be met to solve a problem using
// the Hill-Climbing algorithm. Sucessor is responsable for the method used to
// generate the successor, first choice or best choice.
type HillClimbingSolver interface {
	// New creates a new object of the problem to be solved.
	New(size int) HillClimbingSolver

	// Successor generates a new successor state from a problem object.
	Successor() HillClimbingSolver

	// Objective checks if the current state is an objective state.
	Objective() bool
}

// SolveWithHillClimbing is an implementation of the Hill Climbing algorithm
// for local search problems.
func SolveWithHillClimbing(size int, h HillClimbingSolver) HillClimbingSolver {
	current := h.New(size)

	for {
		for i := 0; i < size*3; i++ {
			successor := current.Successor()
			// Found better successor
			if !reflect.DeepEqual(successor, current) {
				current = successor
			} else {
				break
			}
		}

		if current.Objective() {
			return current
		}
	}
}

// N-Queen structure is a board of integers where the index represents the
// column of the queen in the board and the content represents the line of the
// queen.
type Queen struct {
	board []int
}

// MakeQueen creates a new N-Queen object with a board of mixed values after a
// sequential initialization.
func MakeQueen(size int) Queen {
	qq := Queen{board: make([]int, size)}

	for i := 0; i < size; i++ {
		qq.board[i] = i
	}

	qq.mixBoard()

	return qq
}

// New implements the HillClimbingSolver interface function to create a new
// N-Queen object.
func (q Queen) New(size int) HillClimbingSolver {
	return MakeQueen(size)
}

// BoardSize returns the size of the current board.
func (q *Queen) BoardSize() int {
	return len(q.board)
}

// duplicates replicates a N-Queen object by creating a new on and copying its
// contents to the new one.
func (q *Queen) duplicate() Queen {
	newBoard := make([]int, len(q.board))

	for i := 0; i < len(q.board); i++ {
		newBoard[i] = q.board[i]
	}

	newQueen := Queen{board: newBoard}
	return newQueen
}

// randInt generates a random integer from 0 to the size of the board.
func (q Queen) randInt() int {
	return rand.Intn(len(q.board))
}

// swapTwo swaps two random queens from the N-Queens board.
func (q *Queen) swapTwo() {
	first := q.randInt()
	second := q.randInt()

	q.board[first], q.board[second] = q.board[second], q.board[first]
}

// mixBoard mixes the N-Queens board by swapping two random queens as many times.
// as the size of the board.
func (q *Queen) mixBoard() {
	for i := 0; i < len(q.board); i++ {
		q.swapTwo()
	}
}

// areThreats checks if two given queens are a threat to each other.
func (q *Queen) areThreats(first int, second int) bool {
	return q.board[first]-first == q.board[second]-second ||
		q.board[first]+first == q.board[second]+second ||
		q.board[first] == q.board[second]
}

// Heuristic function returns the number of threats in a board of N-Queens.
func (q *Queen) Heuristic() int {
	threats := 0

	for i := 0; i < len(q.board); i++ {
		for j := i + 1; j < len(q.board); j++ {
			if q.areThreats(i, j) {
				threats++
			}
		}
	}
	return threats
}

// Objective function checks if a given board is a solution to the problem, that
// is, if its heuristic is 0.
func (q Queen) Objective() bool {
	return q.Heuristic() == 0
}

// Sucessor generates a possible list of successors and selects the first one
// found where its heuristic is smaller or equal than the current one.
func (q Queen) successor() Queen {
	listSize := len(q.board) * 2
	currentHeuristic := q.Heuristic()

	for i := 0; i < listSize; i++ {
		new := q.duplicate()
		new.swapTwo()

		if new.Heuristic() <= currentHeuristic {
			return new
		}
	}
	return q
}

// Successor implements the HillClimbingSolver interface function in order to
// generate a new successor from a internal board.
func (q Queen) Successor() HillClimbingSolver {
	return q.successor()
}
