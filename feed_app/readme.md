## Feed 검색 어플리케이션

Go가 동작하는 것을 보기 위해서 완전한 Go 프로그램을 만들어 보겠습니다. 오늘 개발할 프로그램은 다양한 Go프로그램에서 찾아볼 수 있는 기능 구현입니다. 이 프로그램은 web에서 다른 data를 가져와서 내용과 검색어를 비교합니다. 매치되는 내용이 있으면 터미널 윈도우에 표시합니다. 프로그램이 text 파일을 읽고, web call을 만들고, XML과 JSON을 구조체 타입 값으로 디코딩하며 마지막으로 이런 동작을 빠르게 수행하기 위해 Go의 동시성을 이용할 것입니다.

## 프로그램 구조

![프로그램 구조의 흐름](architecture.png)

이 프로그램은 다른 goroutine들 사이에서 실행되는 여러 독립적인 단계로 나눠집니다. main goroutine에서 시작해서 검색과 트랙킹 goroutines으로 그리고 다시 main goroutine 흐름으로 코드를 살펴볼 것입니다. 시작에 앞서 프로젝트의 구조를 알아봅시다.

*cd $GOPATH/src/github.com/ArdanStudios/go_in_action/sample*

* **sample**
	* **data**
		* [data.json](sample/data/data.json) -- Contains a list of data feeds
	* **matchers**
		* [rss.go](sample/matchers/rss.go) -- Matcher for searching rss feeds
	* **search**
		* [default.go](sample/search/default.go) -- Default matcher for searching data
		* [feed.go](sample/search/feed.go) -- Support for reading the json data file
		* [match.go](sample/search/match.go)-- Interface support for using different matchers
		* [search.go](sample/search/search.go)-- Main program logic for performing search
	* [main.go](sample/main.go) -- Programs entry point

이 코드는 알파벳 순서의 4개 폴더로 구성되어 있다. 이 폴더에는 JSON 문서가 있고 이 문서에는 프로그램이 꺼내서 검색 단어와 매치할 데이터가 있다. matcher 폴더는 프로그램이 지원하는 피드의 다른 타입에 대해서 코드를 포한한다. 현재 프로그램은 단지 하나의 matcher만 지원하고 이 matcher는 RSS 타입 피드를 처리한다. 검색 폴더는 다른 matcher를 사용해서 내용을 검색하는 비지니스 로직을 포함하고 있다. 마지막으로 sample이라는 상위 폴더를 가지고 이 폴더는 main.go 코드 파일을 포함한다. 이 파일은 프로그램의 entry 포인트이다.

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).