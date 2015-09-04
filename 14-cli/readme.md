## 커맨드라인 인터페이스 프로그래밍(Command-Line Interface (CLI) Programming)

커맨드라인은 다목적이고 효과적이며 이식성이 좋은 인터페이스다. 따라서 메인프레임 컴퓨터 시대 이래로 계속 사용하고 있다. 최근에는 서버나 다른 제품 시스템을 관리하는 주요 인터페이스로 사용된다. 단순하면서도 복잡한 작업을 쉽게 할 수 있어서 컴퓨터 전문가들 사이에서 계속 사용되고 있다.
커맨드라인 시스템의 능력은 시간이 지나면서 진화되었지만 지금까지도 기반으로 하는 원칙이 있다. *인자*를 받고 *환경*을 상속받아 입출력 텍스트 스트림으로 동작한다.

## 커맨드 쉘(Command Shell)

명령을 해석하는 프로그램과 이에 따라 *shell*이라 프로그램으로 실행된다. *커널(kernel)*이라는 하위 운영체제를 추상하기 때문에 이렇게 이름 붙여졌다. 커널은 직접 시스템 하드웨어와 프로세스 호출과 스케줄링 그리고 다른 중요한 타스크를 관리하는 소프트웨어다. 하지만 쉘과 달리 직접적으로 유저 인터페이스는 없다.

쉘은 종류에 따라서 문법과 제공하는 기능이 다양하다. 하지만 커맨드 파이프라이닝(command pipelining)과 스트림 리디렉션(stream redirection), 조건 제어 흐름(conditional control flow), 패턴 매칭(pattern matching), job control 그리고 상호작용 쉘과 같은 특정은 공통적으로 가지고 있다.

[Bourne Shell][1]은 1997년에 벨 연구소(Bell Labs)에서 개발했으면 훌륭한 프로그래밍 특징을 가지는 첫번째 쉘이다. 쉘 문법의 표준을 만드는 역할을 했고 유사한 쉘들이 따라하였다. 본 쉘 호환"Bourne shell compatible"이라는 말을 생겨났다. 다음 섹션은 본 쉘 호환 문법을 설명한다.


  [1]: http://en.wikipedia.org/wiki/Bourne_shell

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
