# 동시성(Concurrency)

여기서는 동시성의 기본에 대해서 알아보자.

## 웹사이트 검색

얼마나 시간이 걸리지 모르는 많은 연산이 있다. 가능한한 빠르고 효율적으로 처리하기를 원하므로 Go에서 제공하는 동시성을 이용하면 된다.

## 프로그램

단순화시켜 응답시간을 모니터링하는 웹사이트의 리스트가 있다고 가정하자.

## 기본 CLI

인자로 웹사이트를 받아 검색하는 기본 커맨드 라인 프로그램으로 시작한다.

```go
// https://play.golang.org/p/GDaMvunNMZ

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	// flag.Args contains all non-flag arguments
	fmt.Println(flag.Args())
}
```

아래와 같은 인자로 이 프로그램을 실행한다.:

```sh
go run monitor.go http://google.com http://yahoo.com
```

다음과 같이 출력된다:

```sh
[http://google.com http://yahoo.com]
```

## 검색과 응답시간 기록

웹사이트를 검색하고 응답시간을 기록하고자 한다.

아래와 같이 프로그램을 변경하자:

```go
// https://play.golang.org/p/I-gUNt3biw

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Lets keep a reference to when we started
	start := time.Now()

	for _, site := range sites {
		// start a timer for this request
		begin := time.Now()
		
		// Retreive the site
		if _, err := http.Get(site); err != nil {
			fmt.Println(site, err)
			continue
		}
		
		fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
	}

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
```

이제 프로그램을 실행하자:

```sh
go run monitor.go http://google.com http://yahoo.com
```

시간은 다를 수 있지만 아래와 같은 출력값을 볼 수 있다:

```sh
Site "http://google.com" took 119.178774ms to retrieve.
Site "http://yahoo.com" took 581.01092ms to retrieve.
Entire process took 700.245713ms
```

## 동시성을 갖도록 만들기

속도를 내기 위해서 동시에 실행되도록 하자.

다음과 같이 변경하자:

```go
// https://play.golang.org/p/Lc3tS8kprX

package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Use a waitgroup to make sure all go routines finish
	var wg sync.WaitGroup

	// Lets keep a reference to when we started
	start := time.Now()
	
	// Set the value for the waitgroup
	wg.Add(len(sites))

	for _, site := range sites {
		// Launch each retrieval in a go routine.  This makes each request concurrent
		go func(site string) {
			defer wg.Done()
			// start a timer for this request
			
			begin := time.Now()
			
			// Retreive the site
			if _, err := http.Get(site); err != nil {
				fmt.Println(site, err)
				continue
			}
			
			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}(site)
	}

	// Block until all routines finish
	wg.Wait()

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
```

다시 실행해 보자:

```sh
go run monitor.go http://google.com http://yahoo.com
```

이제 실행할 때 아래와 같은 출력을 볼 수 있다:

```sh
Site "http://google.com" took 131.805645ms to retrieve.
Site "http://yahoo.com" took 541.624279ms to retrieve.
Entire process took 541.743794ms
```

가장 오래걸리는 request보다 약간더 길어지는 것이 아니라는 것을 알 수 있다. 모든 것이 동시에 실행되기 때문이다.

마지막 프로그램과 가장 큰 차이점은 `WaitGroup`를 사용하고 `go func` 사용해서 go routine으로 모든 것을 실행시킨다는 것이다.

[Sync Package](http://golang.org/pkg/sync)에서 wait group에 대한 자료를 참고하자 .

## 요약

처음으로 동시성 프로그램을 작성했다!
