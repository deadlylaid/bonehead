package main

import (
    "fmt"
    "strconv"
)


func main(){
    fmt.Println("100을 2진수로 변환")
    fmt.Println("strconv.FormatInt(8,2))")
    fmt.Println(strconv.FormatInt(8, 2))

    fmt.Println("8을 3진수로 변환")
    fmt.Println("strconv.FormatInt(8,3))")
    fmt.Println(strconv.FormatInt(8,3))
}
