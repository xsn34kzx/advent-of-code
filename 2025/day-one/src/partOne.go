package main 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	timesAtZero := 0
	dialVal := 50

	for inputBuf.Scan() {
		inputLine := inputBuf.Text()

		direction := inputLine[0]
		distance, err := strconv.Atoi(inputLine[1:])

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'L':
			dialVal -= distance % 100

			if dialVal < 0 {
				dialVal += 100
			}
		case 'R':
			dialVal += distance % 100

			if dialVal > 99 {
				dialVal -= 100 
			}
		default:
			log.Fatal("Unknown direction!")
		}

		if dialVal == 0 {
			timesAtZero++
		}
	}

	fmt.Println(timesAtZero)
}
