## 인터페이스(Interfaces) - 메소드, 인터페이스, 임베딩(Methods, Interfaces and Embedding)

인터페이스는 behavior만 정의하는 타입을 선언하는데 사용된다. 이 behavior는 메소드를 통해 구체적인 타입(구조체나 커스텀 타입)으로 구현될 수 있다. 구체적인 타입은 인터페이스에 대해서 메소드의 집합을 구현한다. 구체적인 타입의 값은 인터페이스 타입의 변수에 할당할 수 있다. 인터페이스 값에서 메소드 호출은 실제로 실제 값의 메소드를 호출하는 것과 같다. 구체적 타입은 어떤 인터페이스도 구현할 수 있기 때문에 인터페이스 값에 대한 메소드 호출은 자연스럽게 다형성(polymorphic) 특징을 가진다.

## Notes

* 어떤 값에 대한 메소드의 집합은 receiver로 구현되는 메소드만 포함한다.
* 포인터에 대한 메소드 집합은 포인터와 value reciever를 구현하는 메소드를 포함한다.
* 포인터 receiver로 선언한 메소드는 포인터 값을 가지는 인터페이스만 구현한다.
* value receiver로 선언한 메소드는 값과 포인터 receiver 둘다 가지는 인터페이스를 구현한다.
* 메소드 집합의 원칙은 인터페이스 타입에도 적용된다.
* 인터페이스는 레퍼런스 타입이며 포인터로 공유하지 않는다.
* 이것이 다형성 behavior를 생성하는 방법이다.

## Links

https://golang.org/doc/effective_go.html#interfaces

http://blog.golang.org/laws-of-reflection

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## 코드 리뷰

[Method Receivers](example1/example1.go) ([Go Playground](https://play.golang.org/p/eY2Ms-UF-t))

[다형성](example2/example2.go) ([Go Playground](https://play.golang.org/p/Lo1ucf1e9d))

[Address Of Value](example3/example3.go) ([Go Playground](https://play.golang.org/p/yaUWLZjidB))

## 연습문제

### 문제 1

**Part A** 메소드 이름 SayHello를 가지는 Speaker 인터페이스를 선언한다. 영어를 말할 수 있는 사람을 대표하는 구조체 English를 선언하고 중국말을 하는 사람을 대표하는 Chinese 구조체를 선언한다. value receiver를 사용해서 각 구조체에 대해서 Speaker 인터페이스를 구현하고 이 literal string은 "Hello World"과 "你好世界"이다. Speaker 타입의 변수를 선언하고 English 타입의 값의 주소를 할당한다. Chinese 타입의 값에 대해서도 동일하게 해준다. 

**Part B** Speaker 타입의 값을 만족시키는 SayHello라는 새로운 함수를 추가한다. 인터페이스 값에서 SayHello 메소드를 호출하기 위해서 해당 함수를 구현한다. 그런 다음 각 타입의 새로운 값을 생성하고 해당 함수를 사용한다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/oijJdRW3cD)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/iHOEV4YEKr))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).