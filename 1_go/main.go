package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello go")
}

/*
1_go 폴더 내에서

go run main.go  - 테스트 개념
run 은 exe 파일 없이 코드를 유닛테스트 할 때 쓴다.

go build main.go - 테스트 개념
build 는 현재 폴더에 실행 가능한 go 파일을 만든다 파일 이름은 go 파일에 대한 exe 임.
./main.exe  로 파일의 실행이 가능하다.
빌드하면 exe 파일이 생긴다.

go install - 배포 개념
인스톨하면 개발이 완료되었다고 판단 bin 폴더로 가면 1_go.exe 파일이 생성되어 있음
run 과 build 보다는 더 상위임. 하나의 프로젝트를 완성하기 위해 상위폴더 아래 파일들을 묶어서 exe로 배포할 때 씀

bin 폴더 내에서 ./1_go.exe  로 프로젝트 전체의 실행이 가능하다.


참고
go 는 프로젝트 이름을 위치한 폴더 이름으로 생각
1_go  폴더에 있기 때문에
*/
