package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

	var colLengths []int
	inputBuf := bufio.NewScanner(input)
	cols := 0
	amountNumLines := 0

	for inputBuf.Scan() {
		line := inputBuf.Text()
		maybeOp := line[0]

		if maybeOp == '+' || maybeOp == '*' {
			operations := strings.Fields(line)
			values := make([]int, cols)

			for i := range cols {
				if operations[i] == "*" {
					values[i] = 1
				}
			}

			startIndex := 0
			endIndex := 0

			for i, op := range operations {
				endIndex += colLengths[i]

				for j := startIndex; j < endIndex; j++ {
					input.Seek(0, io.SeekStart)

					var valueString strings.Builder

					for n := 0; inputBuf.Scan() && n < amountNumLines; n++ {
						line = inputBuf.Text()
						char := line[j]
						if char != ' ' {
							valueString.WriteByte(line[j])
						}
					}

					currentVal, err := strconv.Atoi(valueString.String())

					if err != nil {
						log.Fatal(err)
					}

					switch op {
					case "+":
						values[i] += currentVal
					case "*":
						values[i] *= currentVal
					default:
					}
				}

				endIndex += 1
				startIndex = endIndex
			}

			finalValue := 0

			for _, v := range values {
				finalValue += v
			}

			fmt.Println(finalValue)
			break
		} else {
			horizontalNumbers := strings.Fields(line)

			if colLengths != nil {
				for i, v := range horizontalNumbers {
					curLength := len(v)

					if colLengths[i] < curLength {
						colLengths[i] = curLength
					}
				}
			} else {
				cols = len(horizontalNumbers)
				colLengths = make([]int, cols)

				for i := range cols {
					colLengths[i] = len(horizontalNumbers[i])
				}
			}

			amountNumLines++
		}
	}

}
