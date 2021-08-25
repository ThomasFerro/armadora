package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type RunesInCell = []rune

func fold(unfolded [][]RunesInCell) string {
	if len(unfolded) == 1 {
		return string(unfolded[0][0])
	}

	// TODO: While needed to fold

	foldRightToLeft := make([][]RunesInCell, len(unfolded))
	for i := 0; i < len(unfolded); i++ {
		foldRightToLeft[i] = make([]RunesInCell, len(unfolded[0]))
		for j := 0; j < len(unfolded[0]); j++ {
			// TODO: Check that the use of "i" is correct instead of "j" (cf foldBottomToTop)
			foldRightToLeft[i][j] = append(unfolded[i][len(unfolded[i])-1-i], unfolded[i][i]...)
		}
	}
	if len(foldRightToLeft) == 1 {
		return string(foldRightToLeft[0][0])
	}

	foldBottomToTop := make([][]RunesInCell, len(foldRightToLeft)/2)
	for i := 0; i < len(foldRightToLeft)/2; i++ {
		foldBottomToTop[i] = make([]RunesInCell, len(foldRightToLeft[0])/2)
		for j := 0; j < len(foldRightToLeft[0])/2; j++ {
			foldBottomToTop[i][j] = append(foldRightToLeft[len(unfolded[i])-1-i][j], foldRightToLeft[i][j]...)
		}
	}
	if len(foldBottomToTop) == 1 {
		return string(foldBottomToTop[0][0])
	}
	// Third fold (foldLeftToRight)
	// Fourth fold (foldTopToBottom)
	return "??"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	unfolded := make([][]RunesInCell, N)

	for i := 0; i < N; i++ {
		unfolded[i] = make([]RunesInCell, N)
		scanner.Scan()
		L := scanner.Text()
		for j, element := range L {
			unfolded[i][j] = RunesInCell{element}
		}
	}
	// fmt.Fprintf(os.Stderr, "unfolded:%v\n", unfolded)

	// TODO: Fold while there are still cases ?
	decodedMessage := fold(unfolded)

	fmt.Println(decodedMessage) // Write answer to stdout
}
