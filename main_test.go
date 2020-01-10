package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

func rmPre(buf *string) { // <pre>と</pre>を1つずつ削除
	*buf = strings.Replace(*buf, "<pre>", "", 1)
	*buf = strings.Replace(*buf, "</pre>", "", 1)
}

func createIoFile(url, filename string) { // urlの問題ページを基に、入出力例だけのファイルを作成する
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	html := string(byteBody)
	rmPre(&html)

	f, _ := os.Create(filename)
	defer f.Close()

	for i := 1; strings.Index(html, "<pre>") < strings.Index(html, "Problem Statement"); i++ {
		input := html[strings.Index(html, "<pre>")+5 : strings.Index(html, "</pre>")-2]
		rmPre(&html)
		output := html[strings.Index(html, "<pre>")+5 : strings.Index(html, "</pre>")-2]
		rmPre(&html)

		f.Write([]byte("<input" + strconv.Itoa(i) + "\n" + input + "\n</input" + strconv.Itoa(i) +
			"\n<output" + strconv.Itoa(i) + "\n" + output + "\n</output" + strconv.Itoa(i) + "\n"))
	}
}

func searchIo(buf string, count int) (i, o string) { // ファイルからcountに対応した入力例と出力例をreturn
	i = buf[strings.Index(buf, "<input"+strconv.Itoa(count))+8 : strings.Index(buf, "</input"+strconv.Itoa(count))-1]
	o = buf[strings.Index(buf, "<output"+strconv.Itoa(count))+9 : strings.Index(buf, "</output"+strconv.Itoa(count))-1]
	return
}

func isExist(filename string) bool { // ファイルが存在するかどうか
	_, err := os.Stat(filename)
	return err == nil
}

func readFile(filename string) string { // ファイルを読み込み、中身をreturn
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 10000)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:n])
}

func TestSolve(t *testing.T) {
	if !isExist("pages") { // pagesフォルダがない場合、作成
		err := os.Mkdir("pages", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	url := "https://atcoder.jp/contests/abc148/tasks/abc148_a"
	filename := url[strings.LastIndex(url, "/")+1:]
	if !isExist("pages/" + filename) { // ioファイルがない場合、作成
		createIoFile(url, "pages/"+filename)
	}
	buf := readFile("pages/" + filename) // ioファイルから読み込み
	for count := 1; strings.Index(buf, "<input"+strconv.Itoa(count)) != -1; count++ {
		input, output := searchIo(buf, count) // input = 入力例, output = 出力例

		fmt.Printf("Q%v answer: %v\treply : ", count, output)
		solve(strings.Fields(input)) // reply = 自分の出力

		if output != reply { // 答え合わせ
			t.Errorf("\x1b[1;31mQ%v: %v != %v\x1b[0m", count, output, reply)
		}
	}
}
