package main

import (
	"bufio"
	"fmt"
	"strings"
)

type LineWordCounter struct {
	l int
	w int
}

func (c *LineWordCounter) Write(p []byte) (int, error) {
	sw := bufio.NewScanner(strings.NewReader(string(p)))
	sw.Split(bufio.ScanWords)
	cw := 0
	for sw.Scan() {
		cw++
	}

	sl := bufio.NewScanner(strings.NewReader(string(p)))
	sl.Split(bufio.ScanLines)
	cl := 0
	for sl.Scan() {
		cl++
	}

	c.w += cw
	c.l += cl
	return len(p), nil
}

func main() {
	var lwc = LineWordCounter{0, 0}
	fmt.Fprintf(&lwc, "hoge fuga\nddd")
	fmt.Fprintf(&lwc, "hoge")
	fmt.Println(lwc)
}
