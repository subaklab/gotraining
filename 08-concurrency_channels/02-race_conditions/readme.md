## 경쟁 조건(Race Conditions) - 동시성과 채널(Concurrency and Channels)

레이스 조건은 2개 이상 goroutine이 동시에 동일한 리소스에 읽고 쓰기를 할 때를 말한다. 통제가 불가능하게 발생하는 버그를 발생시키거나 data를 망가트리는 것과 같이 드러나지 않는다. atomic 함수와 뮤텍스(mutexes)는 goroutine 사이에서 공유 자원의 접근을 동기화 시키는 방식이다.

## Notes

* goroutine은 조정하고 동기화가 필요하다.
* 2개 이상 goroutine이 동일한 자원에 접근하려고 할 때, 레이스 조건이 된다.
* atomic 함수와 뮤텍스 지원이 필요하다.

## Links

http://blog.golang.org/race-detector

http://www.goinggo.net/2013/09/detecting-race-conditions-with-go.html

https://golang.org/doc/articles/race_detector.html

## 문서

[Race Condition Diagram](documentation/race_condition.md)

## 코드 리뷰

[Race Condition](example1/example1.go) ([Go Playground](https://play.golang.org/p/dMHhzuM-TM))

[Atomic Increments](example2/example2.go) ([Go Playground](https://play.golang.org/p/LJETaLkVl0))

[Atomic Store/Load](example3/example3.go) ([Go Playground](https://play.golang.org/p/qifiyxrX1R))

[Mutex](example4/example4.go) ([Go Playground](https://play.golang.org/p/nr8BM7lvNA))

[Read/Write Mutex](example5/example5.go) ([Go Playground](https://play.golang.org/p/p9V1R-_1T2))
## 연습문제

### 문제 1
다음 문제가 주어졌을 때, 레이스 상태를 찾고 바로잡기 위한 레이스 디텍터(race detector)를 사용하라.

	// https://play.golang.org/p/lNXhQJ8gz4

	// Program for an exercise to fix a race condition.
	package main

	import (
		"fmt"
		"math/rand"
		"sync"
		"time"
	)

	// numbers maintains a set of random numbers.
	var numbers []int

	// wg is used to wait for the program to finish.
	var wg sync.WaitGroup

	// init is called prior to main.
	func init() {
		rand.Seed(time.Now().UnixNano())
	}

	// main is the entry point for all Go programs.
	func main() {
		// Add a count for each goroutine we will create.
		wg.Add(3)

		// Create three goroutines to generate random numbers.
		go random(10)
		go random(10)
		go random(10)

		// Wait for all the goroutines to finish.
		wg.Wait()

		// Display the set of random numbers.
		for i, number := range numbers {
			fmt.Println(i, number)
		}
	}

	// random generates random numbers and stores them into a slice.
	func random(amount int) {
		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}

		// Tell main we are done.
	    wg.Done()
	}


[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/yBFA-MDcMw)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/wFTNvVoBpz))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
