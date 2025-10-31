package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the width of the board: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)

	width, err := strconv.ParseInt(strings.TrimSpace(widthInput), 10, 32)
	if err != nil {
		fmt.Println("You did not enter a numeric value!")
		return
	}

	fmt.Print("Enter the height of the board: ")
	heightInput, _ := reader.ReadString('\n')
	heightInput = strings.TrimSpace(heightInput)

	height, err := strconv.ParseInt(strings.TrimSpace(heightInput), 10, 32)
	if err != nil {
		fmt.Println("You did not enter a numeric value!")
		return
	}

	result, err := CreateBoard(width, height)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s", result)
}

func CreateBoard(width, height int64) (string, error) {
	if width <= 0 || height <= 0 {
		return "", errors.New("Width and height must be positive")
	}

	if width == 1 && height == 1 {
		return "#", nil
	}

	boardSize := width*height + (height - 1)
	board := make([]byte, boardSize)
	pos := 0
	for h := range height {
		for w := range width {
			if (w+h)%2 == 0 {
				board[pos] = '#'
			} else {
				board[pos] = ' '
			}
			pos++
		}
		if h < height-1 {
			board[pos] = '\n'
			pos++
		}
	}
	return string(board), nil
}
