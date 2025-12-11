package main

import (
	"bufio"
	"fmt"
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

	inputBuf := bufio.NewScanner(input)
	if inputBuf.Scan() {
		inputIdRanges := strings.Split(inputBuf.Text(), ",")

		invalidIdSum := 0

		for _, idRange := range inputIdRanges {
			ids := strings.Split(idRange, "-")

			if len(ids) != 2 {
				log.Fatal("Did not get ID range correctly!")
			}

			firstId, err := strconv.Atoi(ids[0])

			if err != nil {
				log.Fatal(err)
			}

			secondId, err := strconv.Atoi(ids[1])

			if err != nil {
				log.Fatal(err)
			}

			for id := firstId; id <= secondId; id++ {
				idString := strconv.Itoa(id)

				half := len(idString) / 2

				firstHalf := idString[:half]
				secondHalf := idString[half:]

				if firstHalf == secondHalf {
					invalidIdSum += id
				}

			}

		}

		fmt.Println(invalidIdSum)
	}
}
