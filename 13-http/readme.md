## HTTP - API's

Go 표준 라이브러리는 웹사이트와 API를 만드는 중요한 블록을 제공한다.

## Notes

* net/http은 HTTP/1.1 호환 구현 프로토콜을 제공한다.
* SSL/TLS를 지원한다.
* 라우팅과 미들웨어를 어플리케이션에 추가하는데 몇 가지 단순하 패턴만 있으면 된다.

## Links

https://golang.org/pkg/net/http/

https://golang.org/doc/articles/wiki/

https://github.com/bradfitz/http2

https://github.com/interagent/http-api-design/blob/master/README.md

http://www.restapitutorial.com/httpstatuscodes.html

## 코드 리뷰

[Hello World Server](example1/main.go) ([Go Playground](https://play.golang.org/p/leTrm6v1yG))

[1 Line File Server](example2/main.go) ([Go Playground](https://play.golang.org/p/Qmj_C5PEs1))

[Request and Response Basics](example3/main.go) ([Go Playground](https://play.golang.org/p/-dhowrDOO4))

## 고급 코드 리뷰

[Web API](api)  
예제 코드는 Go로 RESTful API를 만드는 최고의 예제다. [httptreemux](https://github.com/dimfeld/httptreemux) 패키지를 사용해서 라우터를 제외하고 표준 라이브러리를 이용해서 제어한다. 이 라우터는 동사 처리나 인자에 접근을 아주 쉽게 할 수 있게 돕는다.

## 연습문제

### 문제 1

TBD

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
