## Goroutines - 동시성과 채널(Concurrency and Channels)

Goroutine은 독립적으로 실행되도록 만든 함수이다. 공유 쓰레드에서 멀티플렉스로 동작한다. 스케쥴러는 goroutine의 관리와 실행을 책임진다.

## Notes

* goroutine은 독립적으로 실행되도록 스케줄된 함수이다.
* 스케쥴러는 OS 쓰레드를 가지는 컨텍스를 사용하고 goroutine은 큐를 실행한다.
* goroutine을 account를 유지해야하고 깔끔하게 종료한다.
* 동시성은 병렬성과 다르다.
	* 동시성은 한 번에 여러가지를 처리하는 것에 관한 것이다.Concurrency is about dealing with lots of things at once.
	* 병렬성은 한 번에 여러가지를 하는 것에 관한 것이다.

## 문서

[Scheduler Diagrams](documentation/scheduler.md)

## Links

http://blog.golang.org/advanced-go-concurrency-patterns

http://blog.golang.org/context

http://blog.golang.org/concurrency-is-not-parallelism

http://talks.golang.org/2013/distsys.slide

http://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html

## 코드 리뷰

[Goroutines and concurrency](example1/example1.go) ([Go Playground](http://play.golang.org/p/RFwEqHPt7P))

[Goroutine time slicing](example2/example2.go) ([Go Playground](http://play.golang.org/p/mEcWhL14ha))

[Goroutines and parallelism](example3/example3.go) ([Go Playground](http://play.golang.org/p/vT23eEJxJ1))

## 연습문제

### 문제 1

**Part A** 2개의 익명 함수를 선언하는 프로그램을 생성하라. 하나는 0에서 100까지 카운트하고 다른 하나는 100에서 0까지 카운트한다. 각 goroutine에 대해서 고유식별자를 가지고 각 숫자를 출력한다. 다음으로 이 함수로부터 goroutine을 생성하고 main이 goroutine을 완성할 때까지 반환하지 않게 한다.

**Part B** 병렬로 프로그램을 실행하라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/H-h1cbBW3B)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/mB4QslSNoA))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
