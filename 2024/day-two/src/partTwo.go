package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "slices"
    "strconv"
)

func analyzeReport(nums []string) int {
    levels := len(nums) - 1
    badIndex := -1

    var isCurrentlyIncreasing bool
    var wasIncreasing bool

    for i := 0; i < levels; i++ {
        firstVal, firstErr := strconv.Atoi(nums[i])
        secondVal, secondErr := strconv.Atoi(nums[i + 1])

        if firstErr != nil {
            log.Fatal(firstErr)
        } else if secondErr != nil {
            log.Fatal(secondErr)
        }

        diff := secondVal - firstVal
        absDiff := diff
        isCurrentlyIncreasing = diff >= 0

        if !isCurrentlyIncreasing {
            absDiff *= -1
        }

        if i == 0 {
            wasIncreasing = isCurrentlyIncreasing
        }

        if wasIncreasing != isCurrentlyIncreasing {
            if i < levels - 1 {
                badIndex = i - 1
            } else {
                badIndex = i + 1
            }
        } 

        if absDiff < 1 {
            badIndex = i
        } else if absDiff > 3 {
            if i < levels - 1 {
                badIndex = i
            } else {
                if wasIncreasing != isCurrentlyIncreasing {
                    badIndex = i
                } else {
                    badIndex = i + 1
                }
            }
        } else {
            wasIncreasing = isCurrentlyIncreasing
        }
    }

    return badIndex
}

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
        badIndex := analyzeReport(nums)

        if badIndex == -1 {
            safeReports++
        } else {
            nums := slices.Delete(nums, badIndex, badIndex + 1)
            rerunIndex := analyzeReport(nums)
            if rerunIndex == -1 {
                safeReports++
            }
        }
    }

    fmt.Println(safeReports)
}
