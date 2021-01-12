package main

import "fmt"

func main() {
	// 상수
	// const 로 초기화하고 한번 선언 후 값 변경이 불가합니다. 고정된 값 관리용 입니다.
	const a string = "Test1"
	const b = "Test2" // 내부적으로 자료형 선언해 줌
	const c int32 = 10 * 10
	// const d = getHeight() 함수의 반환값을 const 로 받으면 에러가 발생합니다.
	// 함수가 항상 똑같은 값을 반환한다는 보장이 없기 때문에 내부적으로 막은 것입니다.
	const e = 35.6
	const f = false
	/*
		에러가 발생할 때
		const a = 3
		a = 4

	*/
	fmt.Println("a:", a, "b:", b, "c:", c, "e:", e, "f:", f)
}
