package main

import (
    "fmt"
    "strconv"
)

//func main() {
//
//    fmt.Print("いくつまでやってほしいのだね？: ")
//    var input int
//    fmt.Scanf("%d", &input)
//
//    for i := 1; i <= input; i++ {
//        fizzbuzz(i)
//    }
//}
//
//func isThreeIncluded(num int) bool {
//    var strInt = strconv.Itoa(num)
//    for k := 0; k < len([]rune(strInt)); k++ {
//        if '3' == strInt[k] {
//            return true
//        }
//    }
//    return false
//}

func main(i int) {
    adachin := "あだちん！"

    if isThreeIncluded(i) && i%3 == 0 {
        fmt.Println(i, adachin)
    } else if isThreeIncluded(i) {
        fmt.Println(i, adachin)
    } else if i%3 == 0 {
        fmt.Println(i, adachin)
    } else {
        fmt.Println(i)
    }
}
