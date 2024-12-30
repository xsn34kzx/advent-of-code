package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Must provide two arguments to program!")
    }

    input, err := os.Open(os.Args[1])

    if err != nil {
        log.Fatal(err)
    }

    inputBuf := bufio.NewScanner(input)

    safeReports := 0

    for inputBuf.Scan() {
        inputLine := inputBuf.Text()
        nums := strings.Fields(inputLine)
        var isIncreasing bool
        var isUnsafe bool
    
        for i := 0; i < len(nums) - 1; i++ {
            firstVal, firstErr := strconv.Atoi(nums[i])
            secondVal, secondErr := strconv.Atoi(nums[i + 1])

            if firstErr != nil || secondErr != nil {
                log.Fatal(err)
            }

            diff := secondVal - firstVal
            absDiff := math.Abs(float64(diff))
            isCurrentlyIncreasing := diff >= 0

            if i == 0 {
                isIncreasing = isCurrentlyIncreasing
            }

            isUnsafe = isCurrentlyIncreasing != isIncreasing || absDiff < 1 || absDiff > 3

            if isUnsafe {
                break
            } 
        }

        if !isUnsafe {
            safeReports += 1
        }
    }

    fmt.Println(safeReports)
}
