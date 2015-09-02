## 에러 처리(Error Handling)

에러 처리는 프로그램을 사용 사람이 안정적이며 신뢰할 수 있게 만드는 것이 핵심이다. 적절한 에러 값은 구체적이고 정보를 제공해야한다. 호출하는 쪽이 발생한 에러에 대해서 정보를 기반으로 의사결정을 할 수 있게 해줘야 한다. 에러 값을 만드는데 여러가지 방법이 있다. 이는 제공할 필요가 있는 콘텍스트의 양에 달려있다.

## Notes

* static이고 단순한 형태의 에러 메시지일 경우에는 기존 베공하는 error 값을 이용한다.
* 호출자가 특정 에러를 구별하도록 하기 위해서 에러 값을 생성해서 반환하도록 한다.
* 에러의 컨텍스트가 더 복잡해지면 custom 에러 타입을 생성한다.
* 에러 값은 특별한 것이 아니다. 다른 것과 마찬가지로 그냥 값이다. 따라서 원하는대로 처리할 수 있다.

## Links

http://blog.golang.org/error-handling-and-go

http://www.goinggo.net/2014/10/error-handling-in-go-part-i.html

http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html

http://clipperhouse.com/2015/02/07/bugs-are-a-failure-of-prediction/

http://dave.cheney.net/2014/12/24/inspecting-errors

## 코드 리뷰

[Default Error Values](example1/example1.go) ([Go Playground](http://play.golang.org/p/8x6kDZxPWK))

[Error Variables](example2/example2.go) ([Go Playground](https://play.golang.org/p/4YHAbpynl3))

[Type As Context](example3/example3.go) ([Go Playground](http://play.golang.org/p/Eu3X54PnWm))

[Behavior As Context](example4/example4.go) ([Go Playground](http://play.golang.org/p/6GYqwSxHjI))

[Find The Bug](example5/example5.go) ([Go Playground](http://play.golang.org/p/czXpjvWWTT))

## 연습문제

### 문제 1
2개 에러 변수를 생성한다. ErrInvalidValue와 ErrAmountTooLarge이다. 각 변수에 대해서 static message를 제공한다. 다음으로 checkAmount라는 함수를 작성한다. 이 함수는 float64 타입 값을 받아서 에러 값을 반환한다. zero 값을 검사해서 만약 zero면 ErrInvalidValue를 반환한다. $1,000보다 큰 값을 검사해서 ErrAmountTooLarge값을 반환한다. checkAmount 함수를 호출하는 main 함수를 작성하고 에러 값을 반환한다. 적절한 메시지를 화면에 출력한다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/Rt3O-7ndtJ)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/8KETdvYk17))

### 문제 2

appError라는 custom 에러 타입을 생성하자. 이 타입은 Err error, Message string and Code int와 같이 3개 필드를 가진다. 이 3개 필드를 이용해서 우리 메시지를 제공하는 error 인터페이스를 구현하라. bool값을 받는 checkFlag라는 함수를 작성하라. 만약 값이 false라면 초기화된 custom error type의 포인터를 반환하라. 값이 true라면 default error를 반환하라. checkFlag함수를 호출하는 main 함수를 작성하고 구체적인 타입에 대해서 에러를 검사하라.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/x6UimVQMMQ)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/-v-sxBl_ER))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
