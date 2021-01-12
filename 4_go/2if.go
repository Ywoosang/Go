package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// 예제 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 이하")
	}

	if b > 70 {
		fmt.Println("70 초과")
	} else {
		fmt.Println("70 이하")
	}

	// 에러
	// if a >= 65 {
	// 	fmt.Println("65 이상")
	// }   else 도 이런식으로 내려쓰면 에러 발생
	// else {
	// 	fmt.Println("65 이하")
	// }

	i := 110
	// if - else if 예제(1)
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50이상 100 미만 ")

	} else {
		fmt.Println("50 미만")
	}
}
