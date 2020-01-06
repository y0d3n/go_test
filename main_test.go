package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

var (
	answer string
	buf    string
	count  int
	input  string
)

func check(t *testing.T) {
	if strings.Index(buf, "<pre>") > strings.Index(buf, "Problem Statement") {
		return
	}
	count++

	input = buf[strings.Index(buf, "<pre>")+5 : strings.Index(buf, "</pre>")-2]
	set(&buf)
	answer = buf[strings.Index(buf, "<pre>")+5 : strings.Index(buf, "</pre>")-2]
	set(&buf)

	fmt.Printf("Q%v answer: %s\treply: ", count, answer)
	solve(strings.Fields(input))

	if answer != reply {
		t.Errorf("%v != %v", answer, reply)
	}
}

func set(buf *string) {
	*buf = strings.Replace(*buf, "<pre>", "", 1)
	*buf = strings.Replace(*buf, "</pre>", "", 1)
}
func TestQ1(t *testing.T) {
	res, err := http.Get("https://atcoder.jp/contests/abc148/tasks/abc148_a")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	buf = string(body)
	set(&buf)

	check(t)
}

func TestQ2(t *testing.T) {
	check(t)
}

func TestQ3(t *testing.T) {
	check(t)
}

func TestQ4(t *testing.T) {
	check(t)
}
