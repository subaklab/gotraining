## 타입 변환(Type Conversions) - 언어 문법

엄격하게 타입을 지켜야하는 언어다. 따라서 타입을 가지는 값은 반드시 타입 변환시 명시적으로 변환을 해야한다. 반면에 타입이 없는 상수는 묵시적으로 컴파일러가 변환을 시켜준다. 이 값들은 종류 시스템(kind system)에 속해서 더 유연하다. 여기 예제에서는 유연하면서 type safe API를 제공하기 위해서 타입과 time 패키지를 어떻게 이용하는지 알아본다.

## Notes

* 빌트인(built-in)과 사용자 정의 타입을 기반으로 타입을 선언하라.
* 명시/묵시적 변환에 대해서 배운다.
* 상수가 가진 특징을 보고 표준 library에서 어떻게 사용했는지 알아보자.
* 타입을 가지는 상수를 사용은 패키지 함수와 메소드를 위한 API의 일부다.

## 코드 리뷰

[새로운 타입 선언, 생성, 초기화(Declare, create and initalize named types)](example1/example1.go) ([Go Playground](http://play.golang.org/p/mhKlxSyuxr))

[표준 라이브러리에서 named 타입(Named types in the standard library)](example2/example2.go) ([Go Playground](http://play.golang.org/p/XJ4Ia1lMWl))

[타입변환 I(Conversions I)](example3/example3.go) ([Go Playground](http://play.golang.org/p/Rgoqvg8dNv))

[타입변환 II(Conversions II)](example4/example4.go) ([Go Playground](http://play.golang.org/p/B75FURdQ7t))

## 연습문제

### 문제 1

**Part A** int 타입을 기반으로 counter라는 타입을 선언한다. 이 타입의 변수를 zero value로 선언하고 초기화한다. 이 변수의 값과 변수의 타입을 출력하라.

**Part B** 위에서 만든 타입의 새로운 변수를 선언하고 10을 할당한다. 값을 출력하라.

**Part C** 위에서 만든 타입을 기반으로 새로운 타입의 변수를 선언하라. 2개 타입의 값을 할당하는 코드를 작성한다. 컴파일러는 아무 문제 없이 동작할까?

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/Eg9m_rYm4V)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/x-a-6J0s-_))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
