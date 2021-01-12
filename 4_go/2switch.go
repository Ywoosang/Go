package main

import ( // 패키지 여러개 import 할 때 ( ) 사용
	"fmt"
	"math/rand" // 난수 생성할 때
	"time"
)

// 패키지 가져온 상태에서 사용하지 않고 저장하면 자동으로 지워버림.
// if 문에서 한 형식을 switch 문에서 모두 할 수 있음. 
func main() {
	// 예제 1
	// Seed : 난수 나옴
	// 현재 시간 타임스탬프로 반환
	// 다양한 형태로 switch 사용 가능 && and 연산자로 범위에 대한 매칭 
	// if 와 switch 문 두개로 조건에 대한 모든 로직을 처리 
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); { // 0~ 99 까지 랜덤
	case i >= 50 && i < 100:
		fmt.Println("i :", i, "50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i :", i, "25 이상 50 미만")
	// 
	default:
		fmt.Println("i :", i, "기본 값")
	}
	// 예제 2
	// 여러가지 값 나열해서 조건 실행에 맞는지 (걸리는지) 확인하는 경우
	a := 30 / 15 // 2
	switch a {
	case 2, 4, 6: // a 가 2,4,6 이 나올 경우
		fmt.Println("a :", a, "는 짝수")
	case 1, 3, 5: //
		fmt.Println("a :", a, "는 홀수")
	}
	// 예제 2
	// js 같은 다른 언어에는 break 문이 있어야 빠져나가는데, golang 은 케이스 하나 걸리면 자동으로 빠져나감
	// fallthrough 를 써서 걸리더라도 다음으로 넘어가게 할 수 있다. 다음 케이스문으로 진입
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
	case "go":
		fmt.Println("golang")
		fallthrough
	case "python":
		fmt.Println("python")
	case "C":
		fmt.Println("C")
	case "C++":
		fmt.Println("C++")
		// fallthrough 에러 
	}

}

//switch 문은 다양한 사용법이 있기 때문에 숙지 필요
// case 안에 if 문 등등 필요에 따라 사용 가능
