package main

import "fmt"

func main() {
    adachin := "あだちん!"
    adachinsan := "あだちんさん!"

    for i := 1; i <= 40 ; i++ {
        if i%3 == 0 {
            fmt.Println(adachin)
        } else if 3 == 0 {
            fmt.Println(adachinsan)
        } else {
            fmt.Println(i)
          }
    }
}

