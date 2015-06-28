## 검색 엔진 어플리케이션

Go가 동작하는 것을 보기 위해서 완전한 Go 프로그램을 만들어 보겠습니다. 오늘 개발할 프로그램은 다양한 Go프로그램에서 찾아볼 수 있는 기능 구현입니다. 이 프로그램은 단순한 검색엔진을 생성하기 위해서 html 패키지에 대한 예제를 제공합니다. 엔진은 Google, Bing과 Blekko 검색을 지원합니다. 이 3개 검색 엔진 모두에게 겨로가를 요청할 수 있습니다. 혹은 첫번째 결과만 요청할 수도 있습니다. 검색은 동시성을 지원합니다. GOMAXPROCS 환경 변수를 사용해서 병렬로 검색을 실행합니다.  

## 프로그램 구조

![Image of App.](client_image.png)

이 프로그램은 다른 goroutine들 사이에서 실행되는 여러 독립적인 단계로 나눠집니다. main goroutine에서 검색 goroutines로 그리고 다시 main goroutine 흐름으로 코드를 살펴볼 것입니다. 시작에 앞서 프로젝트의 구조를 알아봅시다.

*cd $GOPATH/src/github.com/ArdanStudios/web_app/sample*

* **sample**
	* **search**
		* [bing.go](sample/search/bing.go) -- Performs searches against Bing
		* [blekko.go](sample/search/blekko.go) -- Performs searches against Blekko
		* [google.go](sample/search/google.go)-- Performs searches against Google
		* [rss.go](sample/search/rss.go)-- Boilerplate code for handling RSS feeds
		* [search.go](sample/search/search.go)-- Searching framework code
	* **service**
		* [index.go](sample/service/index.go)-- Handles the rendering of the index page
		* [service.go](sample/service/service.go)-- Initializes and runs the web app
		* [templates.go](sample/service/templates.go)-- Support for handling html templates
	* **static/css**
		* [main.css](sample/static/css/main.css)-- Stylesheet for web app
	* **tests**
		* [endpoint_test.go](sample/tests/endpoint_test.go)-- Tests for endpoint testing
	* **views**
		* [basic-layout.html](sample/views/basic-layout.html)-- Layout HTML for the web app
		* [index.html](sample/views/index.html)-- HTML for the index page
		* [results.html](sample/views/resuls.html)-- HTML for rendering the search results
	* [main.go](sample/main.go) -- Programs entry point

이 코드는 2개의 패키지로 구성되어 있습니다. service 패키지는 HTTP request과 response의 처리를 다룹니다. HTML template는 view를 렌더링하는데 사용합니다. search 패키지는 각기 다른 검색 엔진의 검색 과정을 다룹니다. Searcher interface는 새 Searcher의 구현하도록 선언합니다.

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).