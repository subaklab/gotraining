## 패키지와 Exporting(Package and Exporting)

패키지는 컴파일되는 코드의 기본 단위이다. 내부에 선언된 식별자에 대한 범위를 정의한다. 코드를 작성할 때 어떻게 패키징해야하는지를 배우는 것이 패키지 API의 일부를 외부에서 접근가능한 식별자가 되기 때문에 중요하다. 안정적이고 유용한 API는 아주 중요하다.

## Notes

* Go에 코드는 패키지로 컴파일되고 이후에 함께 링크된다.
* 패키지에서 식별자를 외부로 노출시키거나 노출시키지 않을 수 있다.
* 노출된 식별자에 접근하기 위해서 패키지를 import한다.
* 어떤 패키지든 노출된 타입의 값을 사용할 수 있다.

## Links

http://blog.golang.org/organizing-go-code

http://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html

http://www.goinggo.net/2013/08/organizing-code-to-support-go-get.html

## 코드 리뷰

[Declare and access exported identifiers](example1/example1.go)

[Declare unexported identifiers and restrictions](example2/example2.go)

[Access values of unexported identifiers](example3/example3.go)

[Unexported struct type fields](example4/example4.go)

[Unexported embedded types](example5/example5.go)

## 연습문제

### 문제 1
**Part A** toy라는 이름의 패키지를 생성한다. 이 패키지는 Toy라는 구조체를 노출한다. 노출된 필드는 Name과 Weight이다. 다음으로 노출하지 않는 2개의 필드 onHand와 sold를 추가한다. New라는 팩토리 함수를 선언해서 toy 타입의 값을 생성하고 노출할 필드를 위해 인자를 받아들인다. 그리고 노출하지 않는 필드에 대해서 값을 수정하고 반환하는 메소드를 선언한다.

**Part B** toy 패키지를 import하는 프로그램을 생성한다. New 함수를 사용해서 toy 입의 값을 생성한다. 그리고 counts를 설정하기 위해 메소드를 사용하고 해당 toy 값의 필드값을 출력한다.

[Template](exercises/template1) | 
[Answer](exercises/exercise1)

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
