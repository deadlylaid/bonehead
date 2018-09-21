package main

import (
    "fmt"
    "unsafe"
)


type Person struct {
    name    string
    age     int
    weight  float64
}

func main() {

    // &를 붙인채 선언할 경우 객체를 생성 후 주소값을 저장한다.
    minsoo := Person{"Han minsoo", 27, 72.4}
    pointer_minsoo := &Person{"Han minsoo", 27, 72.4}

    // new를 통해서 생성하면 모든 데이터가 0인 포인터를 반환한다.
    new_minsoo := new(Person)

    fmt.Println(minsoo)
    fmt.Println(pointer_minsoo)
    fmt.Println(new_minsoo)

    fmt.Println(unsafe.Sizeof(minsoo)) // 32
    fmt.Println(unsafe.Sizeof(pointer_minsoo)) // 8
    fmt.Println(unsafe.Sizeof(new_minsoo)) // 8
}
