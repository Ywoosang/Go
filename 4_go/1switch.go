package main

import "fmt"

func main() {
	// 제어문(조건문) - switch
	// Switch 뒤 표현식 생략 가능
	// case 뒤 표현식 사용 가능
	// 자동 break 문 때문에 fallthrough 존재
	// Type 분기 -> 값이 아닌 변수 type 으로 분기 가능

	// 예제 1
	a := 7
	switch {
	case a < 0:
		fmt.Println(a, "a 는 음수")
	case a == 0:
		fmt.Println(a, "a 는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제 2
	// Go 에서 추천하는 방식 switch 영역 안에서 쓸 때는 이런 패턴을 많이 선호
	switch b := 27; {
	case b < 20:
		fmt.Println(b, "20보다 작음")
	case b > 20:
		fmt.Println(b, "20 보다 큼")
	}
	// 예제 3
	// 위에서 한번 선언해 놓고 사용
	switch c := "go"; c {
	case "go":
		fmt.Println("Go")
	case "Java":
		fmt.Println("JAVA")
	default: // Go 도 아니고 Java 도 아닐 때
		fmt.Println("not match")
	}
	// 예제 4
	switch d := "go"; d + "lang" {
	case "golang":
		fmt.Println("golang")
	case "C++":
		fmt.Println("c++")
	default:
		fmt.Println("Others")

		// 예제 5
		// 여러개 변수를 짧은 선언으로 switch 문 안에서 초기화해두고 case 문 안에서 비교할 때
		switch i, j := 20, 30; {
		case i < j:
			fmt.Println("i < j")
		case i == j:
			fmt.Println("i = j")
		case i > j:
			fmt.Println("i > j")
		}
		// if 문보다 다양하게 사용가능 비교,계산 등등..

	}

}
