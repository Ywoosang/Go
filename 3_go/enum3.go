package main

import "fmt"

func main() {
	const (
		_ = iota // 증가하는 도중 생략하고 싶은 게 있다면 밑줄로 대체할 수 있습니다.
		A
		_
		C
		_
		D
	)
	const (
		// 쇼핑몰에서 많이 구매한 회원 등급
		_ = iota + 0.75*2 // 열거형이 포인트를계산하는 공식이라고 가정
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)
	fmt.Println(A, C, D) // 1,3,5
	fmt.Println(DEFAULT, SILVER, GOLD, PLATINUM)
}
