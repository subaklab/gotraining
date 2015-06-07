# Go 설치하기

Go를 설치하는 과정은 아주 단순하다. 하지만 제공하는 툴 전체를 설치하기 원하는 경우에는 좀 복잡할 수 있다.

여기서는 Quick Tour를 위해 필요한 기본 설치에 대해서만 다룬다.

## 어떤 OS를 사용하는가?

Go는 Windows, Mac, 리눅스와 같이 대부분 OS에서 동작한다.

### Mac

Mac에서 가장 빠르게 설치하는 방법은 `HomeBrew`를 이용하는 방법이다.

`HomeBrew`가 없다면 아래와 같이 쉽게 설치할 수 있다.

```sh
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

일단 설치되면 아래와 같은 명령으로 간단하게 실행한다:

```sh
brew update
brew install go
```

### Windows와 리눅스

인스톨러를 다운받아서 설치하기(Mac도 포함):

[http://golang.org/dl/](http://golang.org/dl/)


## 환경설정하기

컴파일러가 제대로 동작하기 위해서 `GOPATH`를 설정해야 한다.

### Mac과 리눅스에서 GOPATH

```sh
mkdir $HOME/go
export GOPATH=$HOME/go
```

GOPATH에 대해서 좀더 상세한 자료는 [https://code.google.com/p/go-wiki/wiki/GOPATH](here).

## Windows에서 GOPATH

커맨드 프롬프트에서 :

```sh
mkdir "%USERPROFILE%\go"
```

제어판 > 시스템 > Advanced Tab > 환경변수

새로운 *사용자(User)* 변수 추가 (시스템 변수가 아님)

변수 이름 : GOPATH
변수 값 : %USERPROFILE%\go

*NOTE:* 변수 설정이 제대로 동작하기 위해서는 reboot이 필요할 수 있다.

## 첫번째 Go 프로그램

자 이제 처음으로 Go 프로그램을 작성해 볼 차례다. 이렇게 하기 위해서 원하는 편집기로 `hello.go` 파일을 생성하고 다음과 같은 내용을 채운다:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

실행하기 위해서 다음과 같은 명령을 사용한다:

```sh
$ go run hello.go
hello, world
```
"hello, world" 메시지를 보면 설치한 Go가 제대로 동작한다는 것이다.

*NOTE:* [golang.org/doc/install](http://golang.org/doc/install)에서 참고하자.


## 요약

축하! 자 이제 더 많은 Go 코드를 작성해 보자!







