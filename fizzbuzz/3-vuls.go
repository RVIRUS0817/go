package main

import "fmt"

func main() {
    vuls := "バルス!"
    vulssan := "バルスさん!"

    for i := 1; i <= 40 ; i++ {
        if i%3 == 0 {
            fmt.Println(vuls)
        } else if i % 3 == 0 {
            fmt.Println(vulssan)
        } else {
            fmt.Println(i)
          }
    }
}
