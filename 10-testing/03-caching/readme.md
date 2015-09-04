## CPU 캐시(Caches)

HW가 어떻게 동작하는지 이해하는 것이 중요합니다. 왜냐하면 대부분의 성능 좋은 코드를 작성하는 방법을 이해하는 것과 같기 때문입니다. 프로세서 캐싱의 기본을 알고 있다면 관용적 표현 코드를 작성하는 범위내에서 더 나은 결정을 할 수 있기 때문입니다.

## Acknowledgment
이 콘텐츠는 2014년 Dive에서 Scott Meyers가 한 이야기입니다 :

[CPU Caches and Why You Care](https://www.youtube.com/watch?v=WDIkqP4JbkE)

## Notes

* CPU Caches works by caching memory on cache lines.
* On our 64 bit processors, the cache line will be 64k.
* Cache lines are moved and stored in L1, L2 and L3 caches.
* Memory in L1 and L2 caches is also in L3 cache.
* Both data and instructions are stored in these caches.
* Hardware likes to traverse data and instructions linearly along cache lines.
* Access to main memory is incredibly slow, we need the cache.
	* Accessing one byte from main memory will cause an entire cache line to be read.
	* Writes to one byte in a cache line requires the entire cache line to be written.
* Small = Fast
	* Compact, well localized code that fits in cache is fastest.
	* Compact data structures that fit in cache are fastest.
	* Traversals touching only cached data is the fastest.
* Predictable access patterns matter.
	* Provide regular patterns of memory access.
	* Hardware can make better predictions about required memory.

### Cache Hierarchies
This is subject to be different in different processors. For this content, the following is the multi-levels of cache associated with the Intel 4 Core i7-9xx processor:

	L1 - 64KB Cache (Per Core)
		32KB I-Cache
		32KB D-Cache
		2 HW Threads

	L2 - 256KB Cache (Per Core)
		Holds both Instructions and Data
		2 HW Threads

	L3 - 8MB Cache
		Holds both Instructions and Data
		Shared across all 4 cores
		8 HW Threads

This is a diagram of the relationship of the cache hierarchy for each core and main memory:

![figure1](figure1.png)

## Links

https://www.youtube.com/watch?v=WDIkqP4JbkE

http://www.akkadia.org/drepper/cpumemory.pdf

http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips

## Code Review

[Caching](caching.go) ([Go Playground](http://play.golang.org/p/GQQXh3cf15))

[Tests](caching_test.go) ([Go Playground](http://play.golang.org/p/opI__KHj9a))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
