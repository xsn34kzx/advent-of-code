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

	inputBuf := bufio.NewScanner(input)
	amountNumberLines := 0

	for inputBuf.Scan() {
		line := inputBuf.Text()
		maybeOp := line[0]

		if maybeOp == '+' || maybeOp == '*' {
			operations := strings.Fields(line)
			cols := len(operations)
			finalValues := make([]int, cols)

			for i := range cols {
				if operations[i] == "*" {
					finalValues[i] = 1
				}
			}

			input.Seek(0, io.SeekStart)

			for i := 0; inputBuf.Scan() && i < amountNumberLines; i++ {
				values := strings.Fields(inputBuf.Text()) 

				for i, v := range values {
					op := operations[i]

					value, err := strconv.Atoi(v)

					if err != nil {
						log.Fatal(err)
					}

					switch op {
					case "+":
						finalValues[i] += value
					case "*":
						finalValues[i] *= value
					default:
					}
				}
			}

			finalValue := 0
			for _, v := range finalValues {
				finalValue += v
			}

			fmt.Println(finalValue)
		} else {
			amountNumberLines++
		}
	}
}
