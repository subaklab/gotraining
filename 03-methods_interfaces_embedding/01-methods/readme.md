## 메소드(Methods) - 메소드, 인터페이스, 임베딩(Methods, Interfaces and Embedding)

receiver를 가지고 함수를 메소드라고 한다. 따라서 이 메소드는 특정 타입과 연결된다. 이 메소드는 해당 타입의 포인터나 값에서 대해서 동작할 수 있다.

## Notes

* 메소드는 receiver 값을 포함하는 함수이다.
* receiver는 메소드를 타입에 연결하며 값이거나 포인터다.
* 메소드는 패키지가 아니라 값이나 포인터에 대해 호출된다.
* Go는 함수와 메소드 변수를 지원한다.

## Links

https://golang.org/doc/effective_go.html#methods

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## 코드 리뷰

[Declare and receiver behavior](example1/example1.go) ([Go Playground](https://play.golang.org/p/AYsB78Dlxb))

[Named typed methods](example2/example2.go) ([Go Playground](https://play.golang.org/p/zHePe-yTUw))

## 고급 코드 리뷰

[함수/메소드 변수](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/MNI1jR8Ets))

## 연습문제

### 문제 1

야구 선수를 표현하는 구조체(이름, atBat, hits)를 선언한다. 선수의 베팅 평균을 계산하는 메소드를 선언한다. 공식은 Hits/AtBats다. 이 타입의 슬라이스를 선언하고 여러 선수들로 슬라이스를 초기화한다. 슬라이스를 이터레이팅해서 야구선수와 타율을 출력한다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/Rj0QfwVPhX)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/C8Z_MiYKbc))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).