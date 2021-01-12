package main

import "fmt"

func main() {
	// 열거형 = iota
	const (
		A = iota // 0 부터 시작하고 1 씩 증가 iota * 10 하면 0 10 20 출력됨. iota +1 하면 1 부터 시작
		B
		C
	)
	fmt.Println(A, B, C)
}
