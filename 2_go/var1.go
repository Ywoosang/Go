// 변수
package main

import "fmt"

func main() {
	// 기본 초기화
	// 정수타입 : 0, 실수(소수점) : 0.0, 문자열 : "", Boolean :true,false
	// 변수명 : 숫자로 첫글자X, 대문자로 된 변수와 소문자로 된 변수는 다르다. 문자,숫자,밑줄,특수기호 사용 가능
	// 변수 및 상수 : 함수 내외 사용 가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3 // 선언과 동시 초기화
	var i float32 = 11.4      // 실수형 지정
	var j string = "Go"
	var k = 5.55      // 자료형을 지정하지 않았지만 자동으로 Go 가 실수의 크기에 맞게 자료형을 선언 동시 초기화
	var l = "Go lang" // 문자도 마찬가지로 자동으로 자료형
	var m = true

	fmt.Println("a : ", a) // int 형은 선언만 했는데 0 으로 초기화 됨
	fmt.Println("a : ", b) // string 형은 아무것도 출력되지 않음 null 공백이 들어감
	fmt.Println("a : ", c)
	fmt.Println("a : ", d)
	fmt.Println("a : ", e)
	fmt.Println("a : ", f)
	fmt.Println("a : ", g)
	fmt.Println("a : ", h)
	fmt.Println("a : ", i)
	fmt.Println("a : ", j)
	fmt.Println("a : ", k)
	fmt.Println("a : ", l)
	fmt.Println("a : ", m)

}
