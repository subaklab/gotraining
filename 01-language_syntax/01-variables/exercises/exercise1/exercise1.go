// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/d2M0Q3mRnd

// 3개 변수를 zero value로 선언하고 3개 변수는 literal 값으로 선언한다.
// 선언 변수는 string, int, bool 타입을 가진다.
// 이 변수 값을 출력해보자.
//
// 새로운 변수를 float32 타입으로 선언하고 Pi(3.14)와 같이
// literal 값으로 변환해서 초기화한다.

package main

import "fmt"

// main is the entry point for the application.
func main() {
	// zero value로 변수 선언
	var age int
	var name string
	var legal bool

	// 변수의 값을 출력
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// 변수 선언 및 초기화
	// 짧은 변수 선언 연산 사용
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// 이 변수 값을 출력
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// 타입 변환 수행
	pi := float32(3.14)

	// 변수의 값을 출력
	fmt.Printf("%T [%v]\n", pi, pi)
}
