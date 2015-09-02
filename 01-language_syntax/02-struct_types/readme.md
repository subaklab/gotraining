## 구조체 타입(Struct types) - 언어 문법

다양한 data 필드를 모아 복잡한 타입을 만드는 방법입니다. 작성할 프로그램에서 사용할 다양한 데이터를 구성하고 공유할 수 있는 좋은 방법입니다.

## Notes

* 구조체 literal 형태로 값을 초기화할 수 있다.
* '.' 연산자를 이용해서 각 필드 값에 접근할 수 있다.
* 익명 구조체(Anonymous struct)를 생성할 수 있다.

## Links

http://www.goinggo.net/2013/07/understanding-type-in-go.html

http://www.goinggo.net/2013/07/object-oriented-programming-in-go.html

## 코드 리뷰

[구조체 타입의 선언, 생성, 초기화](example1/example1.go) ([Go Playground](http://play.golang.org/p/-oeDmu2et8))

[익명 구조체 타입](example2/example2.go) ([Go Playground](http://play.golang.org/p/_xxuE1Ep6U))

## Advanced Code Review

[Struct type alignments](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/1CL1ACDipG))

## 연습문제

### 문제 1

**Part A:** 사용자(이름, email, 나이)에 대한 정보를 가지는 구조체 타입을 선언하라. 이 타입의 값을 생성하고 초기화하여 각 필드의 값을 출력하라.

**Part B:** 동일한 3개 필드를 가지는 익명 구조체 타입을 선언하고 초기화하라. 값을 출력하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ItPe2EEy9X)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/tnn-8hJPUd))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
