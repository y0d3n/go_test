package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type body struct{ str string }

func (b *body) index(pre string) int {
	return strings.Index(b.str, pre)
}

func (b *body) rmPre() {
	// <pre>と</pre>を先頭から1つずつ削除
	b.str = strings.Replace(b.str, "<pre>", "", 1)
	b.str = strings.Replace(b.str, "</pre>", "", 1)
}

func TestSolve(t *testing.T) {
	res, err := http.Get("https://atcoder.jp/contests/abc148/tasks/abc148_a")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	b := &body{string(byteBody)} // b.str = 問題ページのHTML
	b.rmPre()                    // 使わないpreタグがあるので、1組消して調節

	var input, answer string
	for count := 1; b.index("<pre>") < b.index("Problem Statement"); count++ {
		input = b.str[b.index("<pre>")+5 : b.index("</pre>")-2] // input = 入力例
		b.rmPre()
		answer = b.str[b.index("<pre>")+5 : b.index("</pre>")-2] // answer = 出力例
		b.rmPre()

		fmt.Printf("Q%v answer: %v\treply : ", count, answer)
		solve(strings.Fields(input)) // reply = 出力

		if answer != reply {
			t.Errorf("\x1b[1;31mQ%v: %v != %v\x1b[0m", count, answer, reply)
		}
	}
}
