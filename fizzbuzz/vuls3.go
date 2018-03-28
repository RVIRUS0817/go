package main

import (
    "fmt"
    "strings"
)

func main() {
    vuls := "vuls!"
    vulssan := "vulssan!"

    for i := 1; i <= 40; i++ {
        if strings.Contains(fmt.Sprint(i), "3") {
            fmt.Println(vulssan)
        } else if i%3 == 0 {
            fmt.Println(vuls)
        } else {
            fmt.Println(i)
        }
    }
}
