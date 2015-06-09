## 배열(Arrays) - 배열, 슬라이스, Maps(Arrays, Slices and Maps)

Go에서 배열은 특수한 자료구조이다. 고정 사이즈의 메모리 블록을 연속으로 할당받는다. 어떻게 선언하고 타입으로 사용되는 특징이 있다.

## Notes

* 변경이 불가능한 고정 사이즈 자료구조이다.
* 사이즈가 다르면 다른 타입으로 간주한다.
* 메모리는 연속된 블록으로 할당된다.

## 코드 리뷰

[Declare, initalize and iterate](example1/example1.go) ([Go Playground](https://play.golang.org/p/DGr8Zp9L_w))

[다른 타입 배열(Different type arrays)](example2/example2.go) ([Go Playground](http://play.golang.org/p/LVD43cYBNS))

[연속된 메모리 할당(Contiguous memory allocations_](example3/example3.go) ([Go Playground](https://play.golang.org/p/s4BSgxz0Y3))

## 연습문제

### 문제 1

5개 string의 배열을 선언하는데 각 요소는 zero value로 초기화한다. 2번째 배열은 5개 string으로 literal string 값으로 초기화한다. 2번째 배열을 첫번째 배열에 할당하고 첫번째 배열의 결과를 출력한다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ggjjRPzhAB)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/Pa3mrTCcpB))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).