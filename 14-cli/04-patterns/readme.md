## 패턴 매칭

쉘에서 제공하는 편리함 중에 하나가 바로 파일의 패턴 매칭이다. 몇 개 파일을 가지는 디렉토리를 초기화 해보자 :

    touch a b1 c1 01.txt 02.txt

패턴을 이용해서 다양한 방식으로 이 4개 파일에 대한 매칭을 할 수 있다.

    # '.txt'로 끝나는 파일을 출력한다 : 01.txt. 02.txt
    # "start"는 0개 이상의 문자와 매치된다 :
    ls -l *.txt

    # 소문자로 시작하는 모든 파일을 출력한다 : a b1 c1
    # 문자 클래스는 대소문자를 따지고 정확히 1개 문자와 매치한다
    echo [a-z]*

    # 숫자로 끝나지 않는 모든 파일을 출력한다 : a 01.txt 02.txt
    echo *[!0-9]

    # 파일 이름이 2개 문자로 구성된 모든 파일을 삭제한다 : b1 c1
    # 각 물음표는 정확히 하나의 문자와 매치된다
    rm ??

패턴 매칭은 기본 동작이다. 와일드 문자나 공백을 포함하는 인자나 파일이름을 지정하기 위해서는 따옴표나 쌍따옴표를 이용할 수 있다. 특수문자의 경우 백슬래쉬를 사용할 수 있다. 예제는 아래와 같다 :

	# 새로운 2개 파일 생성
	touch filename\ with\ spaces "more spaces"

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
