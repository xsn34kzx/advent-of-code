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

	inputBuf := bufio.NewScanner(input)

	totalOutputJoltage := 0

	for inputBuf.Scan() {
		joltages := inputBuf.Text()

		tensPlace := 0
		onesPlace := 0

		for i := 0; i < len(joltages); i++ {
			joltage := int(joltages[i] - '0')

			if joltage > tensPlace && i != len(joltages) - 1 {
				tensPlace = joltage
				onesPlace = 0
			} else if joltage > onesPlace {
				onesPlace = joltage
			}
		}

		maximumJoltage := tensPlace * 10 + onesPlace
		totalOutputJoltage += maximumJoltage
	}

	fmt.Println(totalOutputJoltage)
}
