# 기본 커맨드 라인 프로그램

Go를 이용하면 커맨드 라인 프로그램을 빠르고 쉽게 만들 수 있다.

[http://golang.org/pkg/flag/](Flags Package)을 이용해서 기본 커맨드 라인 컨셉을 설명한다.

## 단순한 CLI

첫번째 프로그램에서 별다른 어려움 없이 box를 꺼낼 수 있다.

`cli.go` 파일을 열어서 다음 프로그램을 생성하고 다음 내용을 추가:

```go
// https://play.golang.org/p/DPigLqZ5Co

package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "hello" or "bye"`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Hello!")
	case "bye":
		fmt.Println("Bye!")
	default:
		flag.Usage()
	}
}
```

아래 명령으로 프로그램을 실행한다:

```sh
go run cli.go
```

아래와 같이 :

```sh
Usage of /var/folders/l7/3s7z7s1s4n72lvj4w6g_fdmm0000gn/T/go-build844850686/command-line-arguments/_obj/exe/basic:
  -cmd="": cmd can be either "hello" or "bye"
```

## 분해하기

인자를 넣지 않으면 이 프로그램의 `Usage`를 출력한다.

인자를 주는 경우:

```sh
go run cli.go -cmd=hello
```

이제 아래와 같은 출력을 볼 수 있다. 

```sh
Hello!
```


## flag.StringVar

이 메소드는 특정 인자 이름을 찾는다고 flag 패키지에게 전달할 수 있다. 이 경우는 `cmd`이다.

추가 정보는 [http://golang.org/pkg/flag/#FlagSet.StringVar](StringVar)에서 찾을 수 있다.


## 바이너리(binary) 컴파일하기

왜 binary로 컴파일이 필요한지 생각해보자. 다음 커맨드로 쉽게 해결:

```sh
go build cli.go
```

You will now have a file called `cli` that is an executable.  To run that, issue this command:

```sh
./cli -cmd=hello
```

## Summary

Congratulations, you just wrote your first command line program!  We only scratched the surface,
but I hope you enjoyed the quick start.





