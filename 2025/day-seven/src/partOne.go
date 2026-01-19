package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide two arguments!")
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	inputBuf := bufio.NewScanner(input)

	inputBuf.Scan()
	startingLine := inputBuf.Text()
	positions := make(map[int]struct{})

	var emptyStruct struct{}

	for i, v := range startingLine {
		if v == 'S' {
			positions[i] = emptyStruct
		}
	}

	timesSplit := 0

	for inputBuf.Scan() {
		line := inputBuf.Text()
		tempPositions := make(map[int]struct{})

		for i := range positions {
			tempPositions[i] = emptyStruct
		}

		for i := range positions {
			if line[i] == '^' {
				delete(tempPositions, i)

				tempPositions[i - 1] = emptyStruct
				tempPositions[i + 1] = emptyStruct

				timesSplit++
			}
		}

		positions = tempPositions
	}

	fmt.Println(timesSplit)
}
