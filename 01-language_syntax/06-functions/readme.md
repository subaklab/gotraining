## 함수(Functions) - 언어 문법

함수는 언어의 핵심이다. 특정 기능을 담당하는 코드를 구조화하고 모아주는 역할을 한다. 우리가 작성한 패키지에 API를 제공할 수도 있고 이것들은 동시성에 핵심 컴포넌트가 된다.

## Notes

* 함수는 여러 값을 반환하며 대부분 에러 값을 반환한다.
* 에러 값은 프로그래밍 로직의 일부로 검사해야만 한다.
* 공백 식별자(blank identifier) 반환값을 무시한다.

## Links

https://golang.org/doc/effective_go.html#functions

http://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html

http://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html

## 코드 리뷰

[여러 값 반환하기(Return multiple values)](example1/example1.go) ([Go Playground](http://play.golang.org/p/bYY-TRjfH0))

[공백 식별자(Blank identifier)](example2/example2.go) ([Go Playground](http://play.golang.org/p/wPVvgwPlHw))

[재선언(Redeclarations)](example3/example3.go) ([Go Playground](http://play.golang.org/p/XQS-twjrtl))

## 고급 코드 리뷰

[Trapping panics](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/eg14ClW4_y))

## 연습문제

### 문제 1

**Part A** 사용자에 대한 정보를 가지는 구조체 타입을 선언하라. 값을 생성하고 이 타입의 포인터와 에러 값을 반환하는 함수 선언하라. main에서 이 함수를 호출하고 값을 출력하라.

**Part B** 작성한 함수의 2번째 호출을 만든다. 이번에는 에러 값은 무시하고 에러 값을 테스트한다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/p0vlsW5sVL)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/fSjQ3caTy1))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
