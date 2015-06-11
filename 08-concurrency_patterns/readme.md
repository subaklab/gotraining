## 동시성 패턴 (Concurrency Patterns)
goroutine과 채널을 만드는데 다양한 패턴이 있다. 2가지 재밌는 패턴으로 리소스 풀링(resource pooling)과 동시성 검색(concurrent searching)이 있다.

## Notes

* 작업할 코드는 안전하게 goroutine의 수를 설정하는 패턴을 제공한다.
* 리소스 풀링 코드는 goroutine이 리소스를 얻거나 해제하는 것과 같이 관리하는 패턴을 제공한다.
* 검색 코드는 다양한 goroutine을 사용해서 동시성을 가지고 동작하는 패턴을 제공한다.

## Links

http://blog.golang.org/pipelines

https://talks.golang.org/2012/concurrency.slide#1

https://blog.golang.org/context

http://blog.golang.org/advanced-go-concurrency-patterns

http://talks.golang.org/2012/chat.slide

## 코드 리뷰

[Task](task)

[Pooling](pool)

[Chat](chat)

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
