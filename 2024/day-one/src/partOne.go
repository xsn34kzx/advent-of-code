package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "sort"
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

    sort.Ints(leftSlice)
    sort.Ints(rightSlice)

    sum := 0

    for i, _ := range leftSlice {
        distance := math.Abs(float64(rightSlice[i] - leftSlice[i]))
        sum = sum + int(distance)
    }

    fmt.Println(sum)
}
