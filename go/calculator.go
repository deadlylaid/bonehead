package main

import (
	"fmt"
	"strconv"
	"math"
)

func main(){
	end := false

	fmt.Println("계산기 기능이 작동됩니다.")
	for end == false{
		var err error

		num1 := 0
		num2 := 0

		v1, oper, v2 := "", "+", ""

		fmt.Println("")
		fmt.Println("연산값을 입력하시오(종료하려면 exit을 입력하시오)")
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
		}else if oper == "SQR"{
			fmt.Println(sqr(num1))
		}else{
			fmt.Println("허용되지 않는 연산자 입니다.")
		}

		if err != nil{
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

func divide(num1, num2 int) float64{
	var float_num1, float_num2 float64
	float_num1 = float64(num1)
	float_num2 = float64(num2)
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("Warning! 0으로 나누는 것은 불가능 합니다.")
		}
	}()
	return float_num1/float_num2	
}

func sqr(num1 int) float64{
	float_num1 := float64(num1)
	d_value := float64(1)
	pre_value := float64(0)
	is_end := false
	for is_end == false {
		d_value = d_value - ( d_value * d_value - float_num1 )	/ (2 * d_value)
		if math.Abs(pre_value-d_value) > 1e-14{
			pre_value = d_value
		}else{
			is_end = true
		}
	}
	return d_value
}