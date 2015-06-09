## 슬라이스(Slices) - 배열, 슬라이스, Maps(Arrays, Slices and Maps)

Go에서 아주 중요한 자료구조이다. 유연하고 성능이 좋게 그리고 동적으로 data를 관리하고 처리하느냐에 따라 달라진다. 슬라이스를 사용하는 방법을 배우는 것은 Go 개발자에게 아주 중요하다.

## Notes

* 슬라이스는 특별한 빌트인 기능을 가지는 동적 배열과 같다.
* 슬라이스 길이(length)와 용량(capacity)에는 차이가 있고 각각 특정 목적이 있다.
* 슬라이스는 기반이 되는 배열을 다양한 관점에 표현하는 것이다.
* 슬라이스는 빌트인 함수인 append를 통해 증가시킬 수 있다.

## Links

http://blog.golang.org/go-slices-usage-and-internals

http://blog.golang.org/slices

http://www.goinggo.net/2013/08/understanding-slices-in-go-programming.html

http://www.goinggo.net/2013/08/collections-of-unknown-length-in-go.html

http://www.goinggo.net/2013/09/iterating-over-slices-in-go.html

http://www.goinggo.net/2013/09/slices-of-slices-of-slices-in-go.html

http://www.goinggo.net/2013/12/three-index-slices-in-go-12.html

## 코드 리뷰

[Declare and Length](example1/example1.go) ([Go Playground](http://play.golang.org/p/fWJR3Kln4Y))

[참조 타입(Reference Types)](example2/example2.go) ([Go Playground](https://play.golang.org/p/d1kRkbZ-iV))

[슬라이스의 슬라이스(Taking slices of slices)](example3/example3.go) ([Go Playground](https://play.golang.org/p/aizhjTO-br))

[Appending slices](example4/example4.go) ([Go Playground](http://play.golang.org/p/UzmwiMWDwd))

[문자열과 슬라이스(Strings and slices)](example5/example5.go) ([Go Playground](http://play.golang.org/p/6CAkumo0HI))

[Variadic functions](example6/example6.go) ([Go Playground](http://play.golang.org/p/cK3y_qYUgd))

## 고급 코드 리뷰

[슬라이스 실무에서 사용(Practical use of slices)](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/-qQgO7NbLm))

[Three index slicing](advanced/example2/example2.go) ([Go Playground](http://play.golang.org/p/dJk2eycWhH))

## 연습문제

### 문제 1

**Part A** int의 nil 슬라이스를 선언하라. 10개 값을 슬라이스에 append하는 루프를 생성하라. 슬라이스를 이터레이터를 통해 각 값을 출력한다.

**Part B** 5개 string의 슬라이스를 선언하고 literal 값으로 슬라이스를 초기화하라. 모든 element를 출력하라. 인덱스 1, 2를 가지는 슬라이스로 인덱스 위치와 새로운 슬라이스에서 각 element의 값을 출력하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/mPKmyGNR2L)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/BSNAUj2pd-))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).