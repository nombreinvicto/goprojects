package chessboard

// Declare a type named File which stores if a square is occupied by a piece
// - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files,
// accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// a file is a column
	//in the chessboard - slice of bool
	fl, exists := cb[file]
	squaresOccupied := 0
	if !exists {
		return squaresOccupied
	}

	// loop over the fl = []bool
	for _, val := range fl {
		if val == true {
			squaresOccupied += 1
		}
	}
	return squaresOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	squaresOccupied := 0
	if rank < 1 || rank > 8 {
		return squaresOccupied
	}

	//loop over chekboard each rank
	// each row = [A B C D E F G H]
	// recall rank varies from 1 to 8
	for _, sliceValue := range cb {
		if sliceValue[rank-1] == true {
			squaresOccupied += 1
		}
	}
	return squaresOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0

	for _, sliceValue := range cb {
		totalSquares += len(sliceValue)
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalSquaresOccupied := 0

	for rank := 1; rank <= 8; rank++ {
		totalSquaresOccupied += CountInRank(cb, rank)
	}
	return totalSquaresOccupied
}
