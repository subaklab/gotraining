## 포인터(Pointers) - 언어 문법

함수 경계를 data 공유하는 방법을 제공한다. 포인터로 data를 공유하고 참조함으로써 유연성을 제공한다. 프로그래에서 필요한 메모리의 양을 최소화하고 성능을 개선시킬 수 있다.

## Notes

* data를 공유하는데 포인터를 사용하라.
* Go에서 값은 항상 값에 의한 전달(pass by value)이다.
* "Value of"는 값을 뜻하고, "Address of" ( **&** ), 주소를 뜻한다.
* (*) 연산자는 포인터 변수를 선언하고 "포인터가 가리키는 값"을 뜻한다.

## Links

https://golang.org/doc/effective_go.html#pointers_vs_values

http://www.goinggo.net/2013/07/understanding-pointers-and-memory.html

http://www.goinggo.net/2014/12/using-pointers-in-go.html

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)

## Code Review

[값에 의한 전달(Pass by value)](example1/example1.go) ([Go Playground](https://play.golang.org/p/nNnsK6hWdP))

[데이터 공유 I(Sharing data I)](example2/example2.go) ([Go Playground](https://play.golang.org/p/FWmGnVUDoA))

[데이터 공유 II(Sharing data II)](example3/example3.go) ([Go Playground](http://play.golang.org/p/VYqb11RiWr))

## Advanced Code Review

[스탭 vs 힙(Stack vs Heap)](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/931Cw6uzsn))

## 연습문제

### 문제 1

**Part A** int 타입의 변수를 20 값으로 선언하고 초기화하라. _주소_(address of)와 _값_(value of)을 출력하라.

**Part B** 생성한 마지막 변수를 가리키는 int 타입의 포인터 변수를 선언하고 초기화하라. _주소_(address of)와 _값_(value of) 그리고 _포인터가 가리키는 값_(value that the pointer points to)을 출력하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/asM7GXfJNk)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/pUtHSTN0Ef))

### 문제 2

구조체 타입을 선언하고 이 타입의 값을 생성하라. 이 구조체 타입에서 특정 필드의 값을 변경하는 함수를 선언하라. 만든 함수를 호출하기 전과 후의 값을 출력하라.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/b6-FNFOToO)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/oEtveMoO1s))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).