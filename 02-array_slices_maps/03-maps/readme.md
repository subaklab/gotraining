## Maps - 배열, 슬라이스, Maps(Arrays, Slices and Maps)

key/value 쌍을 가지는 data 저장 및 관리를 위한 자료구조이다.

## Notes

* key/value 쌍으로 저장 및 추출하는 방법을 제공
* map key는 반드시 할당 구문에서 사용한 value여야 한다.
* map에서 interating시 항상 임의의 순서이다.

## Links

http://blog.golang.org/go-maps-in-action

http://www.goinggo.net/2013/12/macro-view-of-map-internals-in-go.html

## 코드 리뷰

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](http://play.golang.org/p/voXAyiydFf))

[Map literal initialization](example2/example2.go) ([Go Playground](http://play.golang.org/p/ekemu8ZfBu))

[Map key restrictions](example3/example3.go) ([Go Playground](http://play.golang.org/p/0v_VHlYF7f))

## 고급 코드 리뷰

[Composing maps of maps](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/pQsoB02pDl))

## 연습문제

### 문제 1

key로 string을 가지고 value로 정수형을 가지는 map을 선언하고 만들어보자. 5개 값을 가지며 iterating으로 key/value 쌍을 출력하자.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/-JBSUoux-v)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/8K-IZgJFSg))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).