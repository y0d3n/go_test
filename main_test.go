package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

type data struct{ html string }

func (d *data) index(pre string) int {
	return strings.Index(d.html, pre)
}

func (d *data) rmPre() {
	d.html = strings.Replace(d.html, "<pre>", "", 1)
	d.html = strings.Replace(d.html, "</pre>", "", 1)
}

func createFile(url, filename string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(byteBody)
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readFile(filename string) *data {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1000000)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
	}

	return &data{html: string(buf)}
}

func TestSolve(t *testing.T) {
	url := "https://atcoder.jp/contests/abc148/tasks/abc148_a"
	filename := url[strings.LastIndex(url, "/")+1:]
	if !isExist(filename) { // もし問題のHTMLファイルがない場合、作成する
		createFile(url, filename)
	}
	d := readFile(filename) // b.html = 問題ページのHTML
	d.rmPre()               // 使わないpreタグがあるので、1組消して調節

	for count := 1; d.index("<pre>") < d.index("Problem Statement"); count++ {
		input := d.html[d.index("<pre>")+5 : d.index("</pre>")-2] // input = 入力例
		d.rmPre()
		output := d.html[d.index("<pre>")+5 : d.index("</pre>")-2] // output = 出力例
		d.rmPre()

		fmt.Printf("Q%v answer: %v\treply : ", count, output)
		solve(strings.Fields(input)) // reply = 出力

		if output != reply {
			t.Errorf("\x1b[1;31mQ%v: %v != %v\x1b[0m", count, output, reply)
		}
	}
}
