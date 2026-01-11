package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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
	allAvailableFreshIngredients := 0

	for inputBuf.Scan() {
		line := inputBuf.Text()

		if len(line) != 0 {
			ids := strings.Split(line, "-")

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
		} else {
			slices.SortFunc(freshIngredientRanges, func(a Range, b Range) int {
				return cmp.Compare(a.Start, b.Start)
			})

			currentEnd := 0

			for _, idRange := range freshIngredientRanges {
				if idRange.Start > currentEnd {
					allAvailableFreshIngredients += idRange.End - idRange.Start + 1
					currentEnd = idRange.End
				} else if idRange.End >= currentEnd {
					allAvailableFreshIngredients += idRange.End - currentEnd
					currentEnd = idRange.End
				}
			}
			
			break
		}
	}

	fmt.Println(allAvailableFreshIngredients)
}
