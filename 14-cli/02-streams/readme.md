## 텍스트 스트림, 리다이렉션, 파이프라이닝(Text Streams, Redirection, and Pipelining)

모든 프로세스는 자신의 3개 텍스트 스트림 집합을 가지고 있다. stdio 스트림으로 알려진 : stdin, stdout 그리고 stderr이다. 프로그램은 *pipelining*이라고 알고 있는 하나의 프로세스를 통해 함께 체인으로 연결될 수 있다. 하나의 프로세스의 stdout은 다음 프로세스의 stdin이 되는 것과 같다. 파이프라인에서 각 명령쌍은 "pipe" 문자로 구분되며 아래와 같다 :

    # 현재 디렉토리에 있는 파일의 개수를 카운트한다
    ls | wc -l

    # 압축된 log 파일에서 처음 10개 error 라인을 출력한다
    zcat log.gz | grep -i error | head

Stdio streams can also redirected from or to files. The following two examples
are equivalent: they both lexicographically sort lines from infile, filter out
duplicates, and then write the result to outfile.

    sort < infile | uniq > outfile
    cat infile | sort | uniq | tee outfile

Unless stdin is redirected from a file, the first process in a pipeline will
receive input from the controlling terminal (the keyboard). Likewise, the
stdout of the last process in a pipeline is sent to the controlling terminal,
unless redirected. Even a lone command without explicit redirection is
technically a pipeline, due to its implicit coupling to the terminal.

Stderr is used to communicate information not part of the normal program
output, such as errors; since stderr is not normally redirected, the error
output for all processes in a pipeline are displayed in the terminal.

However, it's often desirable to ignore stderr:

    rm somefile 2>/dev/null

In this case, `2>` means redirect stream #2 (stdin is 0, stdout is 1, and
stderr is 2) to /dev/null, in effect discarding the error output.

## Example

[Running a subprocess](example1/parent.go)

[Replicating a shell pipeline](example2/pipeline.go)

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
