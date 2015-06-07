## 변수 - 언어 문법

변수는 언어의 핵심중에 하나로 메모리에 읽고 쓰는 방법을 제공한다. Go에서 메모리 접근은 타입 안전성(type safe)을 지킨다. 이 뜻은 컴파일러가 타입이 아주 중요하다는 뜻이고 변수를 사용하는데 있어 선언한 타입의 범위를 벗어나는 것을 허용하지 않는다.

## Notes

* 변수를 zero value로 선언하고 싶을 때, 키워드 var를 이용하라. 
* 변수가 선언하고 초기화할 때, 짧은 변수 선언 연선자를 사용하라.
* 값이 힙(heap)을 빠져나올 때를 결정하는데 Escape analysis를 이용한다.

## Links

[Built-In 타입들](http://golang.org/ref/spec#Boolean_types)

https://golang.org/doc/effective_go.html#variables

http://www.goinggo.net/2013/08/gustavos-ieee-754-brain-teaser.html

## 코드 리뷰

[변수 선언 및 초기화](example1/example1.go) ([Go Playground](http://play.golang.org/p/6w6hBNE75a))

## 연습문제

### 문제 1 

**Part A:** 3개 변수를 선언하는데 zero value로 초기화하고 literal 값으로 선언한다. string, int 그리고 bool 이렇게 3개 타입 변수 선언한다. 이 변수의 값을 표시하라.

**Part B:** float32 타입의 새로운 변수를 선언하고 Pi(3.14) literal 값을 변환으로 변수를 초기화하라.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/1xUWjHMB3I)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/d2M0Q3mRnd))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
