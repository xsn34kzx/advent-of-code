package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
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

    var leftSlice []int
    var rightSlice []int

    for inputBuf.Scan() {
        inputLine := inputBuf.Text()
        nums := strings.Fields(inputLine)
    
        leftVal, err := strconv.Atoi(nums[0])
        rightVal, err := strconv.Atoi(nums[1])

        if err != nil {
            log.Fatal(err)
        }

        leftSlice = append(leftSlice, leftVal)
        rightSlice = append(rightSlice, rightVal)
    }

    similarityScore := 0

    valueMap := make(map[int]int)

    for _, v := range leftSlice {
        valueMap[v] = 0
    }

    for _, v := range rightSlice {
        count, ok := valueMap[v]
        if ok {
            valueMap[v] = count + 1
        }
    }

    for _, v := range leftSlice {
        similarityScore = similarityScore + (v * valueMap[v])
    }

    fmt.Println(similarityScore)
}
