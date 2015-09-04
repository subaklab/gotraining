## 포인터 사용하기

새로운 타입을 만들 때,메소드를 선언하기 전에 아래 질문에 대해서 생각해보라. 이 타입의 값으로부터 어떤 것을 추가하거나 제거하는 것이 새로운 값을 생성하거나 기존 것을 변형시킬 필요가 있는가? 만약 답변이 새로운 값을 생성해야한다는 것이라면 value receiver를 사용하라. 그렇지 않으면 pointer receiver를 사용하라. 이것은 이 타입의 값이 프로그램의 다른 부분으로 어떻게 전달해야만 하는지에도 적용된다. 항상 값 혹은 포인터를 선택해서 사용해야하고 섞어서 사용하지 말자. 여기에는 몇 가지 예외가 있다.

## Notes

* 타입의 속성이 어떻게 전달하느냐를 결정한다.
* 타입은 primitive와 non-primitive data qualities를 구현할 수 있다.
* 두가지 속성을 가지는 구조체를 선언하지마라.
* 일반적으로 포인터를 가지는 빌트인 타입 값을 전달하지마라.
* 일반적으로 기능의 unmarshal 타입을 구현하지 않는다면 포인터를 가지는 참조 타입 값을 전달하지마라.
* 일반적으로 구조체 타입이 primitive data 값처럼 동작하지 않는다면 포인터를 가지는 구조체 타입 값을 전달해라.

## Links

http://www.goinggo.net/2014/12/using-pointers-in-go.html

http://play.golang.org/p/ki991PuHhk

## 코드 리뷰

[Primitive Types](example1/example1.go) ([Go Playground](https://play.golang.org/p/H5HRoElN6q))

[Reference Types](example3/example3.go) ([Go Playground](https://play.golang.org/p/E-Bb5cRuyz))

[Struct Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/xD6PCx--GG))

## 연습문제

### 문제 1

float64 타입의 X와 Y 필드를 가지는 Point 구조체를 선언한다. 이 타입과 메소드를 위해 facotry 함수를 구현하며 이 메소드는 타입의 값을 받어서 2 point들 사이의 거리를 계산한다. 이 타입의 속성은 무엇일까?

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/9_MSdcdlNQ)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/5KL4HipSJ-))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
