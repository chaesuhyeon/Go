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

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // for index , number : = range numbers -> index출력하고 싶지 않을 시 _ 써서 ignore variable 샤용
		total = total + number
	}
	return total
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
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

	result := superAdd(1, 2, 3, 4, 5, 6) // for문 사용
	fmt.Println(result)

	fmt.Println(canIDrink(18))

	// a := 2
	// b := 2
	//a = 10
	//fmt.Println(a, b)   // 10, 2
	//fmt.Println(&a, &b) // 0x140000a4018 0x140000a4030 메모리 주소

	a := 2
	b := &a
	fmt.Println(b)  // a의 메모리 주소
	fmt.Println(*b) // 2

	*b = 20        // a의 주소와 연결된 b로 a의 값을 변경
	fmt.Println(a) // 20

	//names := [5]string{"nico", "lynn", "dal"} // 길이가 5이고 타입이 string인 array
	//names[3] = "alala"
	//names[4] = "alalala"

	names := []string{"nico", "lynn", "dal"} // length없이 사용하는 array
	names = append(names, "flynn")           // array에 요소 추가
	fmt.Println(names)

	nico := map[string]string{"name": "nico", "age": "12"} // key string타입 value string타입
	fmt.Println(nico)                                      // map[age:12 name:nico]

	for key, value := range nico {
		fmt.Println(key, value)
	}

	favFood := []string{"kimchi", "ramen"}
	// nicoStruct := person{"nico", 18, favFood}
	nicoStruct := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nicoStruct)      // {nico 18 [kimchi ramen]}
	fmt.Println(nicoStruct.name) //nico
}
