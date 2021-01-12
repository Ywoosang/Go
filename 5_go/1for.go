// 반복문이 for 밖에 없음
package main

import "fmt"

func main() {
	// 반복문 for
	// Go 에서 유일하게 제공되는 반복문/
	// 다양한 사용법 숙지

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 에러 발생  => 괄호 없이 썼을 경우, 블록 내렸을 경우

	// while 문에 쓰는 패턴 (무한 루프) for 만 쓰면 무한적으로 실행  => 초기화 조건식 증감 모두 뺀 경우
	// for {
	// 	fmt.Println("Go")
	// 	fmt.Println("lang")
	// }

	// 예제3 (Range 용법)
	// 리스트 순회 
	loc := []string{"First", "Second", "Third"}
	// index 와 name 둘 다 출력할 경우 
	for index, name := range loc {
		fmt.Println(index, name)
	}
	// 인덱스만 출력할 경우 
	for index := range loc {
		fmt.Println(index)
	}
	// 이름만 출력할 경우 
	for _, name := range loc {
		fmt.Println(name)
	}
}
