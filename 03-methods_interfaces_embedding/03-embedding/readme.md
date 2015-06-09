## 임베딩(Embedding) - 메소드, 인터페이스, 임베딩(Methods, Interfaces and Embedding)

임베딩 타입은 타입들 사이에서 상태(state)와 행위(behavior)를 공유하고 재사용하는 최후의 수단이다. inner type 사용을 통해 outer type의 참조로 inner type의 필드와 메소드에 바로 접근이 가능하다. 

## Notes

* 임베딩 타입은 타입들 사이에서 상태(state)와 행위(behavior)를 공유할 수 있다.
* inner type은 자신의 특성을 잃어버리지 않는다.
* 이것은 상속이 아니다.
* prototion을 통해서 inner type 필드와 메소드는 outer 타입을 통해 접근 가능하다. 
* outer type은 inner 타입의 행위를 override할 수 있다.

## Links

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## 코드 리뷰

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/e5O_Dx5VpM))

[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/UkrDXkk-Ch))

[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/BgEoThS7u9))

[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/jfOfrRMPZR))

## 연습문제

### 문제 1

**Part A** 2개 필드 name, age를 가지는 animal 구조체 타입을 선언한다. bark 필드를 가지는 dog 구조체 타입을 선언한다. animal 타입을 dog 타입에 임베딩한다. dog 타입의 값을 선언하고 초기화한다. 변수의 값을 출력한다.

**Part B** yelp라는 포인터 receiver를 이용해서 animal 타입에 메소드를 추가한다. yelp는 "Not Implemented" literal string을 출력한다. dog 타입의 값으로 메소드를 호출해보자.

**Part C** yelp라는 단일 메소드를 가지는 yelper 인터페이스를 추가하라. yelper 타입의 값을 선언하고 dog 타입의 값의 주소를 할당한다. yelp 메소드를 호출하자.

**Part D** dog 타입에 대해서 yelper 인터페이스를 구현한다. 창의적으로 bark 필드를 만들어보자. yelper 타입의 값에서 다시 yelp 메소드를 호출하자.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/a-Nzng_E6Z)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/hvVA4zB9Bf))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
