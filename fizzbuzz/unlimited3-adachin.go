package main

import (
    "fmt"
    "strconv"
)

func main() {

    fmt.Print("いくつまで言わせるんだね!?: ")
    var input int
    fmt.Scanf("%d", &input)

    for i := 1; i <= input; i++ {
        fizzbuzz(i)
    }
}

func ThreeIncluded(num int) bool {
    var strInt = strconv.Itoa(num)
    for k := 0; k < len([]rune(strInt)); k++ {
        if '3' == strInt[k] {
            return true
        }
    }
    return false
}

func fizzbuzz(i int) {
    adachin := "あだちん！"
    adachinsan := "あだちん3！"

    if ThreeIncluded(i) && i%3 == 0 {
        fmt.Println(i, adachinsan)
//        fmt.Println(i, adachin+""+adachinsan)
    } else if ThreeIncluded(i) {
        fmt.Println(i, adachinsan)
    } else if i%3 == 0 {
        fmt.Println(i, adachin)
    } else {
        fmt.Println(i)
    }
}
