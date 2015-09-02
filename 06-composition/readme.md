## Composition

Composition은 단순히 타입 임베딩(type embedding)이 아니다. 디자인 패턴의 일종으로 작은 부품들로 큰 프로그램을 구성하기 위해 중요한 역할을 한다. 작은 부품은 하나의 일에 초점을 맞춘 타입의 선언과 구현에서 시작한다. composition으로 구조화된 프로그램은 변화에 적응하고 커가기 위해 더 나은 기회를 가지고 있다.

## Note

* composition으로 타입과 behavior를 선언한다.
* Composition은 lego 블록으로 소프트웨어를 구성하는 것과 같다.
* 단순히 타입 임베딩 이상의 의미가 있다.

## Links

http://golang.org/doc/effective_go.html#embedding

## 코드 리뷰

[Composition I](example1/example1.go) ([Go Playground](http://play.golang.org/p/W5ya6_LAU6))

[Composition II](example2/example2.go) ([Go Playground](http://play.golang.org/p/xsDJhCYOBA))

## 연습문제

### 문제 1

**Part A** 아래 가이드를 따라 진행한다:

**Part B** administrator interface를 구현하는 sysadmin 타입을 선언한다.

**Part C** developer interface를 구현하는 programmer type을 선언한다.

**Part D** administrator와 developer를 모두를 임베딩하는 company type을 선언한다.

**Part E** sysadmin, programmers, company을 생성한다. 이것들은 미리 정의한 task를 완료하기 위해서 고용하고 사용하는 것이 가능하다.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/b8ww3jd2Xs)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/8t5ns3cqNp))
___
All material is licensed under the [GNU Free Documentation License](https://github.com/gobridge/gotraining/blob/master/LICENSE).
