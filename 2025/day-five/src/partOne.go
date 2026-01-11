package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide two arguments!")
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	inputBuf := bufio.NewScanner(input)
	
	freshIngredientRanges := make([]Range, 0)
	availableFreshIngredients := 0
	finishedRanges := false

	for inputBuf.Scan() {
		curLine := inputBuf.Text()

		if len(curLine) != 0 {
			if finishedRanges {
				id, err := strconv.Atoi(curLine)

				if err != nil {
					log.Fatal(err)
				}

				for _, idRange := range freshIngredientRanges {
					if id >= idRange.Start && id <= idRange.End {
						availableFreshIngredients++
						break;
					}
				}
			} else {
				ids := strings.Split(curLine, "-")

				if len(ids) != 2 {
					log.Fatal("Did not get the ID range correctly!")
				}

				firstId, err := strconv.Atoi(ids[0])

				if err != nil {
					log.Fatal(err)
				}

				secondId, err := strconv.Atoi(ids[1])

				if err != nil {
					log.Fatal(err)
				}

				idRange := Range{firstId, secondId}

				freshIngredientRanges = append(freshIngredientRanges, idRange)
			}
		} else {
			finishedRanges = true
		}
	}

	fmt.Println(availableFreshIngredients)
}

