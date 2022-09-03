package main

import (
	"fmt"
	"strings"
)

func mutiply(a, b int) int {
	return a * b

}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done ") // defer : func이 실행되고(return ) 난 후에 실행됨
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(name ...int) int {
	 
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// main은 컴파일을 위한 것
func main() {
	// fmt.Println("Hello World!")
	// something.SayHello()
	// const name string = "hyunha"
	// var name string = "hyunha"
	name := "nico" // const name string = "hyunha" 과 똑같음
	fmt.Println(name)

	fmt.Println(mutiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	// totalLength, _ := lenAndUpper("nico")
	// fmt.Println(totalLength) //  return 값이 2개인데 하나만 출력하고 싶을 때 _ 사용해서 ignore var로 사용

	repeatMe("nico", "lynn", "dal", "marl", "flynn ")

	// total := superAdd(1,2,3,4,5,6)

}
