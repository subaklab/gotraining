## 테스팅과 디버깅(Testing and Debugging)

Go는 테스팅과 디버깅 모두를 제공한다. Go 프로그램을 프로파일링하고 벤치마킹하는 것을 포함한다.

## 패키지 리뷰

[Testing](../10-testing/01-testing/readme.md)

[Prediction](../10-testing/02-prediction/readme.md)

[Caching](../10-testing/03-caching/readme.md)

[Godebug](../10-testing/04-godebug/readme.md)

[Profiling](../10-testing/05-profiling/readme.md)

## 연습문제

### 문제 1
정수형을 문자열로 변환하는 3가지 벤치마킹 테스트를 작성하라. 먼저 fmt.Sprintf 함수를 사용하고 다음으로 strconv.FormatInt 함수를 마지막으로 strconv.Itoa를 사용하라. 어떤 함수가 성능이 좋은지를 알아보라.

[Template](exercises/template1/bench_test.go) | 
[Answer](exercises/exercise1/bench_test.go)

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).