## 리플렉션(Reflection)

리플렉션은 값에 대해서 유도된 타입이나 다른 메타 데이터를 조사하는 것을 뜻한다. 리플렉션은 프로그램에 유연성을 부여해서 다른 타입의 데이터나 값을 생성할 수 있게 한다. 데이터의 인코딩/디코딩에 중요한 역할을 담당하고 있다.

## Notes

* 리플렉션 패키지를 이용해서 타입을 조사할 수 있도록 한다.
* "tags"를 구조체 필드에 추가해서 메타 데이터를 저장하고 사용할 수 있다.
* 인코딩 패키지는 리플렉션을 이용해서 원하는 것을 얻을 수 있다.

## Links

http://blog.golang.org/laws-of-reflection

## 코드 리뷰

[Empty Interface](example1/example1.go) ([Go Playground](http://play.golang.org/p/OSeD9F_P46))

[Reflect struct types with tags](example2/example2.go) ([Go Playground](http://play.golang.org/p/y0WyYezH05))

## 고급 코드 리뷰

[Decoding function for integers](example3/example3.go) ([Go Playground](http://play.golang.org/p/bWQ6hiVECQ))

## 연습문제

### 문제 1
customer invoice를 요청에 대한 구조체 타입을 선언하라. 여기에는 CustomerID와 InvoiceID 필드가 있다. 요청을 유효하게 할 수 있는 테그를 정의하라. ID가 유효한 길이(length)와 범위(range)를 지정하는 테그를 정의하라. 특정 타입의 값을 받고 테그를 처리하는 validate라는 함수를 선언하라. 유효성의 결과를 출력하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/lgMQHWpZul)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/xy-wyPrsjz))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
