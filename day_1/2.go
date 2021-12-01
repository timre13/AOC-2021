package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)

func strToInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        panic(err)
    }
    return i
}

func main() {
    inputBytes, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := string(inputBytes)

    lastVal := 0
    largerCount := 0
    lines := strings.Split(input, "\n")
    for i := 0; i < len(lines)-2-1; i++ {
        window := strToInt(strings.TrimSpace(lines[i])) +
                  strToInt(strings.TrimSpace(lines[i+1])) +
                  strToInt(strings.TrimSpace(lines[i+2]))
        if lastVal != 0 && window > lastVal {
            largerCount++
        }
        lastVal = window
    }

    fmt.Printf("%d increases\n", largerCount)
}
