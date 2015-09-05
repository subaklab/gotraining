## 인코딩 - 표준 라이브러리

인코딩은 다른 형태로 바꾸는 처리를 말한다.(marshaling or unmarshaling) JSON 문자열 문서와 이를 미리 정의한 타입의 값으로 변경하는 것은 오늘날 Go프로그램에서 매우 일반적인 사례이다. 인코딩에 대해서 Go의 지원은 놀랍고 매번 릴리즈마다 향상되어 속도가 빨라지고 있다. 

## Notes

* JSON과 XML에 대해서 인코딩과 디코딩을 지원을 표준 라이브러리로 제공한다.
* 이 패키지는 매번 릴리즈마다 향상되고 있다.

## Links

http://www.goinggo.net/2014/01/decode-json-documents-in-go.html

## 코드 리뷰

[Unmarshal JSON documents](example1/example1.go) ([Go Playground](http://play.golang.org/p/ocxFH62yaw))

[Unmarshal JSON files](example2/example2.go) ([Go Playground](http://play.golang.org/p/IWfOJbmMdL))

[Marshal a user defined type](example3/example3.go) ([Go Playground](http://play.golang.org/p/rLDpqYbnGR))

[Custom Marshaler and Unmarshler](example4/example4.go) ([Go Playground](http://play.golang.org/p/TOYrZJoLei))

## 연습문제

### 문제 1

**Part A** 사용자 이름과 이메일 주소를 포함하는 JSON 문서의 배열을 가지는 파일을 생성한다. JSON 문서로 매핑되는 구조체 타입을 선언한다. json 패키지를 사용해서 파일을 읽고 이 구초제 타입의 슬라이스를 생성하라. 이 슬라이스를 출력해보자.

**Part B** 이 슬라이스를 바로 프린트 문자열로 Marshal하고 각 요소들을 출력해보자.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/OkIHsVwMQ7)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/Huf8jEDUJO))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
