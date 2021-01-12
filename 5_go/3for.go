package main

import "fmt"

// 반복분 흐름 제어
func main() {

	// 중첩 for 문일 떄
	// 아래와 같이 명시적으로 이름을 붙여주고 break 뒤에 언급하면 Loop1 자리(반복문 밖) 으로 아예 빠져나감.

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				break
			}
			fmt.Println(i, j)
		}

	}
	fmt.Println("The end")

Loop1:
	for k := 0; k < 5; k++ {
		for l := 0; l < 5; l++ {
			if k == 2 && l == 2 {
				break Loop1
			}
			fmt.Println(k, l)
		}
	}
	fmt.Println("The end")
	//break 할 시 두 번째 반복문을 빠져나감. 여전히 첫 번째 반복문 안

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("홀수", i)
	}

	// Loop 레이블 안에는 반드시 for 문이 와야 한다. continue 또는 break 문 사용하는 조건문 또는 반복문와야함

	//Break 와 달리  continue 를 사용하면 밖 for 문으로가 다시 실행하게 된다.
	// 디버깅을 하려면 값을 출력해 봐야하는 문제가 있음.
Loop2:
	// 에러 발생 (Loop 레이블 밑에 관련이 없는 소스코드가 있다면)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("i,j : ", i, j)
		}
	}
}
