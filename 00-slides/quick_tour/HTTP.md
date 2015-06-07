# 첫번째 웹 서버(Web Server)

아래와 같은 내용을 다룹니다:

- 웹서버(webserver) 돌리기
- 기본 라우팅(Routing)
- 첫번째 템플릿(template)

## 웹서버 실행

처음으로 해야하는 일은 웹서버를 실행하는 것이다. 몇 줄의 코드로 이 일을 할 수 있다.

`webserver.go` 파일을 열고 다음 내용을 추가하자:

```go
// https://play.golang.org/p/x4iJhctz8e

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
```

실행하기 위해 다음과 같은 명령을 실행하자:

```sh
go run website.go
```

다음으로  [http://localhost:8080](http://localhost:8080) 주로소 웹브라우저를 연다.

보는 바와같이 기본 "Hello world"를 출력한다.

이제 다음 url로 이동하자: [http://localhost:8080/foo](http://localhost:8080/foo)

이 웹서버는 입력 라우트에 따라 동작한다. 다음과 같은 동작을 하는 것이 최종 목표는 아니다:

```go
// https://play.golang.org/p/CUnPy2CKqI

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)
	
	// Add the "index" or root path
	rt.HandleFunc("/", Index)
	
	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
}
```

`github.com/gorilla/mux`라는 `external` 패키지를 사용했다는 것을 알수 있다.

이 패키지에 여러분 컴퓨터에 있는지와 해당 프로그램이 컴파일 되는지 확인하기 위해서 `get`을 이용한 다음 명령을 실행한다:

```sh
go get github.com/gorilla/mux
```

이제 서버를 다시 실행한다:

```sh
go run website.go
```

이제 [http://localhost:8080](http://localhost:8080) 주소로 웹브라우저를 연다. 보는 바와 같이 새로운 것은 없다.

하지만 [http://localhost:8080/foo](http://localhost:8080/foo)로 가면 페이지를 찾을 수 없다는 메시지를 보게 된다.


## 템플릿(Templates)

기본 템플릿을 만든다. 이렇게 하기 위해서 [http://golang.org/pkg/html/template/](html/template) 패키지를 이용한다.

프로그램을 다음과 같이 변경한다:

```go
// https://play.golang.org/p/N5c1LMZWe_

package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// This is a basic struct to hold basic page data variables
type PageData struct {
	Title string
	Body  string
}

func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out page data for index
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page.",
	}

	// Render a template with our page data
	tmpl, err := htmlTemplate(pd)

	// If we got an error, write it out and exit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// All went well, so write out the template
	w.Write([]byte(tmpl))
}

func htmlTemplate(pd PageData) (string, error) {
	// Define a basic HTML template
	html := `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	{{.Body}}
	</body>
	</HTML>`

	// Parse the template
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	// We need somewhere to write the executed template to
	var out bytes.Buffer

	// Render the template with the data we passed in
	if err := tmpl.Execute(&out, pd); err != nil {
		// If we couldn't render, return a error
		return "", err
	}

	// Return the template
	return out.String(), nil
}
```

### 요약

여러분의 첫번째 웹서버를 축하한다.


