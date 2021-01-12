package main

import "fmt"

func main() {
	// 여러개 선언을 한번에
	var (
		name      string = "Ywoosang"
		age       int
		doing     string
		isRunning = false
	)

	age = 250
	doing = "Go"
	isRunning = true
	name = "woosang"

	fmt.Println("name :", name, "age :", age, "doing :", doing, "isRunning :", isRunning)

}
