package main

import (
	"fmt"
	"strconv"
)

func main(){
	end := false

	fmt.Println("계산기 기능이 작동됩니다.")
	for end == false{
		var err error

		num1 := 0
		num2 := 0

		v1, oper, v2 := "", "", ""

		fmt.Println("연산값을 입력하시오(종료하려면 end를 입력하시오)")
		fmt.Scanln(&v1, &oper, &v2)

		num1, err = strconv.Atoi(v1)
		num2, _ = strconv.Atoi(v2)

		if oper == "+" {
			fmt.Println(add(num1, num2))
		}else if oper == "-"{
			fmt.Println(substract(num1, num2))
		}else if oper == "*"{
			fmt.Println(multiply(num1, num2))
		}else if oper == "/"{
			fmt.Println(divide(num1, num2))
		}
		if err != nil {
			fmt.Println("종료합니다.")
			end = true
		}
	}
}

func add(num1, num2 int) int {
	return num1+num2
}

func substract(num1, num2 int) int{
	return num1-num2
}

func multiply(num1, num2 int) int{
	return num1*num2
}

func divide(num1, num2 int) int{
	return num1/num2
}
