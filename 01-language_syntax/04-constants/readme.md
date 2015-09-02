## 상수(Constants) - 언어 문법

변경되지 않는 값에 대해서 고유 이름을 만드는 방법입니다. 이를 이용하여 언어의 유연성을 높여줍니다. 다른 언어에 비해서 Go에서는 특별한 방법으로 사용됩니다.

## Notes

* 상수는 변수가 아니다.
* 컴파일 시에만 존재한다.
* 타입을 가지지 않는 상수(untyped constant)는 묵시형 변환이 가능하다.(변수와 타입을 가지는 상수는 불가!)
* 타입을 가지지 않는 상수는 타입(Type)이 아니라 종류(Kind)를 가진다고 생각하면 된다.

## Links

https://golang.org/ref/spec#Constants

http://blog.golang.org/constants

http://www.goinggo.net/2014/04/introduction-to-numeric-constants-in-go.html

## 코드 리뷰

[상수 선언 및 초기화](example1/example1.go) ([Go Playground](http://play.golang.org/p/OLuzwK1oHT))

[패러럴 타입 시스템(Parallel type system (Kind))](example2/example2.go) ([Go Playground](http://play.golang.org/p/ExxRWe6jUz))

## 연습문제

### 문제 1

**Part A:** 타입이 없는 그리고 타입을 가지는 상수를 선언하고 그 값을 출력하라.

**Part B:** 2개 literal 상수를 곱해서 타입을 가지는 변수에 넣고 그 값을 출력하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/kzZZ24O23g)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/d2gkKxEftw))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).