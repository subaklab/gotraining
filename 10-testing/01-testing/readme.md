## 테스팅과 벤치마킹

테스팅은 go tools과 표준 라이브러리에 기본으로 탑재되어 있습니다. 테스팅은 개발 프로세스에서 중요한 부분입니다. 왜냐하면 프로젝트의 라이브 사이클에서 엄청난 시간을 절약해 줄 수 있기 때문입니다. 벤치마킹은 테스팅 기능에 속하는 아주 강력한 도구입니다. 성능 리뷰를 코드 측면에서 벤치마킹이 가능하다. 프로파일링은 함수 사이에서 상호작용의 관점과 어떤 함수가 가장 많이 사용되는지를 알 수 있다.

## Notes

* Go 도구는 테스팅과 벤치마킹을 지원한다.
* 도구들은 매우 유연하고 다양한 옵션을 제공한다.
* 개발하면서 테스트를 작성할 수 있다.
* 예제 코드는 테스트와 다큐멘트 모두로 가능하다.
* 벤치마킹은 개발, QA, 릴리스 사이클에서 행해진다.
* 성능 문제가 발생하면, 코드를 프로파일해서 어떤 함수에 문제가 있는지를 볼 수 있다.
* 도구는 서로에게 간섭할 수 있다. 예를 들면, 정확한 메모리 프로파일링은 CPU 프로파일에 영향을 미치고, goroutine blocking 프로파일링은 스케줄러 추적에 영향을 미친다. 각 필욯나 프로파일링 모드에 대해서 테스트를 재실행한다.

## Links

http://golang.org/pkg/testing/

http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

http://saml.rilspace.org/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof

http://golang.org/pkg/net/http/pprof/

https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs

## Code Review

[Basic Unit Test](example1/example1_test.go)

[Table Unit Test](example2/example2_test.go)

[Mocking Web Server Response](example3/example3_test.go)

[Testing Internal Endpoints](example4)

[Example Test](example4/handlers/handlers_example_test.go)

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
