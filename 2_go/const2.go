package main

import "fmt"

func main() {
	const a, b int = 10, 99            // 선언과 동시에 int 형으로  초기화
	const c, d, e = true, 0.84, "test" // 각각 다른 데이터의 자료형 형태로 선언해도 자동적으로 자료형 맞춰주므로 에러 안남
	// 괄호 안에 있는 것들은 영향을 주고받는 하나의 상태를 설명하는 변수임을 나타낼 때 용이
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7778
	)
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)

}
