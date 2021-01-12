package main

import "fmt"

func main() {
	// 조건문
	// if 문 : 반드시 Bollean 검사 ->자바스크립트처럼 1은 True , 0 False 그런것 없음. 무조건 Boolean 으로 입력해야함. 자동 형 변환 불가
	// 소괄호 사용 X

	var a int = 20
	b := 20
	// 예제 1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}
	// 에러 발생
	//if b >=25  자동으로 ; 붙여주는데 여기서 문장 끝나므로, {} 인식 못해서 에러
	//{
	//
	//}
	// 에러 발생2
	//if b >= 25
	//fmt.Println("25 이상") 괄호 생략 시 에러
	// 에러 발생 3
	// if c:=1; c  {     Go 는  1을 Boolean 으로 인식하지 못함.
	// 	fmt.Println("True")
	// }
	//
	// c 를 25 로 초기화 하고 ; 뒤 내용 은 20 보다 큰지  여부 Boolean
	if c := 25; c > 20 {
		fmt.Println("20보다 큼")
	}

}
