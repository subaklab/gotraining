## 조건 제어 흐름(Conditional Control Flow)

명령이 성공적으로 완료되면 0 상태로 종료된다; 모든 0이 아닌 종료 상태는 실패를 뜻한다. 규칙으로 1인 상태는 일반적인 프로그램 에러 상태고 2는 사용자 에러(개별 입력이나 인자)를 뜻한다.하지만 0-127 범위의 특정값이 사용될 수도 있다. 다양한 시스템에서 128-255 상태는 특별한 의미를(signal reporting) 가지므로 사용하지 않는다. 파이프라인에서 마지막 프로세스의 종료 상태는 전체 파이프라인에 대한 상태로 사용된다.

파이프라인의 성공이나 실패는 추가 파이프라인을 실행하기도 한다:

    rm some-file && echo success || echo failure

"lists"로 알려진 이런 방식으로 파이프라인이 체인으로 엮여서 왼쪽에서 오른쪽으로 우선순위를 가지며 연결된다. 가장 일반적으로 AND 순차는 함께 체인으로 연결되어 태스크의 연속은 모든 태스크가. 
Pipelines chained in this way are known as "lists" and have left-to-right
precedence, such that the above will always either echo success or failure, but
never both. Most commonly, AND sequences are chained together so that a series
of tasks can proceed only if every task before it was successful:

    task1 && task2 && task3

만약 세미콜론을 대신 사용하면(혹은 각 태스크가 자신의 라인을 가지고 있다면) 태스크는 실패하든지 상관없이 실행될 것이다.

## 예제

[Setting exit status](example1/parent.go)

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
