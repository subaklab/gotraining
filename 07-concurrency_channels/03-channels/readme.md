## 채널(Channels) - 동시성과 채널(Concurrency and Channels)

goroutine 사이에서 data를 안전하게 공유하기 위한 방법을 제공하느 참조 타입이다. unbuffered 채널은 goroutine에서 다른 goroutine으로 data를 안전하게 전달하는 것을 보장한다. buffered 채널은 이런 보장 없이도 채널을 통해서 data를 전달할 수 있다. send/receive 동작이 완료되기 전에 unbuffered 채널은 sending/receiving goroutine이 필요하다. buffered 채널은 goroutine이 send/receive를 수행하기 위해 동일 인스턴스에 준비하고 있을 필요가 없다.

## Notes

* unbuffered 채널은 어떤 인스턴스에서 data가 교환이 이뤄지는 것을 보장한다.
* buffered 채널은 goroutine과 리소스를 관리를 돕는다.
* 닫힌 채널은 시스템에서 통보하는 역할을 한다.
* unbuffered 채널에서 send는 해당 채널의 receive 전에 일어난다.
* unbuffered 채널에서 receive는 채널이 완료되는 send 전에 발생한다.
* 채널을 closing은 채널이 닫혀서 zero value를 반환하는 receive 전에 발생한다.

## 문서

[Channel Diagrams](documentation/channels.md)

## Links

http://blog.golang.org/share-memory-by-communicating

http://www.goinggo.net/2014/02/the-nature-of-channels-in-go.html

## 코드 리뷰

[Unbuffered channels - Tennis game](example1/example1.go) ([Go Playground](http://play.golang.org/p/7WO_eOJx_G))

[Unbuffered channels - Relay race](example2/example2.go) ([Go Playground](http://play.golang.org/p/0qJeU4GFnF))

[Buffered channels - Retrieving results](example3/example3.go) ([Go Playground](http://play.golang.org/p/jT4-vZBpMm))

[Timer channels and Select](example4/example4.go) ([Go Playground](http://play.golang.org/p/KuMG3o_7-C))

## 연습문제

### 문제 1
2개 goroutine이 10번에 걸쳐 정수를 주고받는 프로그램을 작성하라. 각 goroutine이 정수를 받는 것을 출력하라. 각 전달에서 정수를 증가시켜라. 일단 정수가 10이 되면 프로그램을 종료시켜라.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/G7O-DnJrEA)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/v7fEyd86i3))

### 문제 2
4개 string의 버퍼를 유지하기 위해 buffered 채널을 사용하는 프로그램을 작성하라. main에서 string 'A', 'B', 'C' 그리고 'D'를 채널로 보내라. 다음으로 20개 goroutine을 생성해서 채널로부터 하나의 string을 받고 값을 출력하라. 그리고 string을 채널로 돌려보내라. 일단 각 goroutine이 이 일을 완료하면 goroutine이 종료되도록 한다.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/vc6c1-M2EB)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/K9gNyTRA0s))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
