package base

import (
	"bufio"
	"fmt"
	"os"
)

var stdinR = bufio.NewReader(os.Stdin)
var stdoutW = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(stdoutW, f, a...) }
func print(a ...interface{})            { fmt.Fprint(stdoutW, a...) }
func println(a ...interface{})          { fmt.Fprintln(stdoutW, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(stdinR, f, a...) }
func scan(a ...interface{})             { fmt.Fscan(stdinR, a...) }
func scanln(a ...interface{})           { fmt.Fscanln(stdinR, a...) }

func f() {
	defer stdoutW.Flush()
}
