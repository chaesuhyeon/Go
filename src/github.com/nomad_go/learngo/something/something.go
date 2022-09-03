package something

import "fmt"

func sayBye() { // 함수를 소문자로 시작하면  private라서 export안됨
	fmt.Println("Bye")
}

func SayHello() { // 함수를 대문자로 시작해야 export됨
	fmt.Println("Hello")
}
