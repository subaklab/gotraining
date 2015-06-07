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

Go requires you to set the  `GOPATH` for the compiler to work properly.

### GOPATH on Mac or Linux

```sh
mkdir $HOME/go
export GOPATH=$HOME/go
```

In depth details on `GOPATH can be found [https://code.google.com/p/go-wiki/wiki/GOPATH](here).

## GOPATH on  Windows

From a command prompt:

```sh
mkdir "%USERPROFILE%\go"
```

Go to the Control Panel > System > Advanced Tab > Environment Variables.

Add a new *User* Variable (not a system variable)

Variable name: GOPATH
Variable value: %USERPROFILE%\go

*NOTE:* You may need to reboot for this variable to take effect.

## Your first Go program

Ok, now it's time to create our first go program.  To do so, create a file called `hello.go` with your
preferred text editor, and add the following content:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

To run it, use the following command:

```sh
$ go run hello.go
hello, world
```
If you see the "hello, world" message then your Go installation is working.

*NOTE:* This was taken directly from [golang.org/doc/install](http://golang.org/doc/install)


## Summary

Congratulations, you can now write more Go code!








