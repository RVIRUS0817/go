package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    countMap := map[string]int{}

    for _, word := range strings.Fields(s) {
        countMap[word]++
    }

    return countMap
}

func main() {
    wc.Test(WordCount)
}
