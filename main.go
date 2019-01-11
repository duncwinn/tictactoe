package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticTacToe struct {
	positionSet [9]bool
	board       [][]string
	player      string
	tryAgain    bool
}

func main() {

	t := ticTacToe{}

	t.resetBoard()

	//Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to TicTacToe! Hit Enter to start:")

	for scanner.Scan() {
		fmt.Print("\nYour board is ready:\n")

		t.printBoard()

		break
	}

	//Loop around 9 times taking turns between 0 and X
	for i := 1; i < 10; i++ {

		if i%2 == 0 {
			t.player = "X"
		} else {
			t.player = "0"
		}

		t.myGo(i)
	}

} // main

func (t *ticTacToe) resetBoard() {

	t.board = [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	t.positionSet = [9]bool{false, false, false, false, false, false, false, false, false}
	t.player = "X"
	t.tryAgain = false
}

func (t *ticTacToe) printBoard() {

	fmt.Println()
	for i := 0; i < cap(t.board); i++ {
		fmt.Printf("%s\n", strings.Join(t.board[i], " "))
	} //for
	fmt.Println()

}

func (t *ticTacToe) myGo(round int) {

	t.tryAgain = true

	fmt.Print("\nROUND ", round, ":\n")

	for t.tryAgain == true {
		fmt.Print("\nPlayer ", t.player, ": Starting top left to bottom right type a number 1 to 9\n:")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		poistion := strings.TrimSuffix(text, "\n")

		i, err := strconv.Atoi(poistion)

		//First check the entry is valid; if true checkPosition
		if err == nil && i-1 >= 0 && i-1 < 9 {
			t.checkPosition(i - 1)
		} else {
			fmt.Println("\nYou messed it up!!!!")
		}
	}

	t.printBoard()
}

func (t *ticTacToe) checkPosition(position int) {

	if t.positionSet[position] == false {
		t.positionSet[position] = t.setPosition(position)
		t.tryAgain = false
	} else {
		fmt.Println("\nposition: ", position, " has already been taken! Please tye again.")
		t.printBoard()
		t.tryAgain = true
	}
}

func (t *ticTacToe) setPosition(position int) bool {

	count := 0

	for i := 0; i < len(t.board); i++ {
		for j := 0; j < len(t.board[i]); j++ {

			if count == position {
				t.board[i][j] = t.player
				return true
			}
			count++
		}
	}
	return false
}
