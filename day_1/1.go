package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)

func main() {
    inputBytes, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := string(inputBytes)

    lastVal := 0
    largerCount := 0
    for _, line := range strings.Split(input, "\n") {
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }
        lineInt, err := strconv.Atoi(line)
        if err != nil {
            panic(err)
        }
        //fmt.Println(line)
        if lastVal != 0 && lineInt > lastVal {
            largerCount++
        }
        lastVal = lineInt
    }

    fmt.Printf("%d increases\n", largerCount)
}
