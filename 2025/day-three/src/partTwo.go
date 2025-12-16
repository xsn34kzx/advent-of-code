package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const digitAmount = 12

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

		var digits [digitAmount]int

		joltageAmount := len(joltages)

		for i := 0; i < joltageAmount; i++ {
			joltage := int(joltages[i] - '0')
			zeroRest := false

			for j := 0; j < digitAmount; j++ {
				if zeroRest {
					digits[j] = 0
				} else if j >= i - joltageAmount + digitAmount && joltage > digits[j] {
					digits[j] = joltage
					zeroRest = true
				}
			}
		}

		maximumJoltage := 0

		for i, value := range digits {
			maximumJoltage += int(math.Pow10(digitAmount - 1 - i)) * value
		}

		totalOutputJoltage += maximumJoltage
	}

	fmt.Println(totalOutputJoltage)
}
