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
    vuls := "vuls！"
    vulssan := "vulssan！"

    if ThreeIncluded(i) && i%3 == 0 {
        fmt.Println(i, vulssan)
//        fmt.Println(i, vuls+""+vulssan)
    } else if ThreeIncluded(i) {
        fmt.Println(i, vulssan)
    } else if i%3 == 0 {
        fmt.Println(i, vuls)
    } else {
        fmt.Println(i)
    }
}
