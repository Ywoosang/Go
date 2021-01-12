package main

import "fmt"

func main() {
	// 짧은 선언 (go 에만 있음 )
	// 일반 변수는 함수 밖에서 사용할 수 있지만 반드시 함수 안에서만 사용 (전역으로 사용이 불가능합니다.)
	// 선언 후 재할당시 예외(에러) 발생
	// var 변수처럼 ex)  var a int32 = 4   a = 6 처럼 다시 할당할 수 없음.
	// 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있기 때문에 사용함.
	shortVar1 := 3      // 짧은선언
	shortVar2 := "Test" // 짧은 선언은 내부적으로 int string 다 할당이 됨.
	shortVar3 := false
	// shortVar3 := true 처럼 다시 선언하면 에러가 납니다.
	// Println 에서 같은 변수는 한번만 들어가야함. 그렇지 않으면 오류남.
	fmt.Println("shortVar1 :", shortVar1, "shortVar2 :", shortVar2, "shortVar3 :", shortVar3)
	// if 문에서만 사용할 것이기 때문에 짧은 선언으로 go 에만 있는 if 문을 사용할 수 있음
	// go 에서는 if 안에서 i 를 할당하고 i 가 조건에 맞는지 체크함.
	// 타 언어와 구별되는 go 만의 특징
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}

}
