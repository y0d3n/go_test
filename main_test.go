package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

type sample struct {
	Input  string `json:"in"`
	Output string `json:"out"`
}

func rmPre(buf *string) { // <pre>と</pre>を1つずつ削除
	*buf = strings.Replace(*buf, "<pre>", "", 1)
	*buf = strings.Replace(*buf, "</pre>", "", 1)
}

func createSampleFile(url, filename string) { // urlの問題ページを基に、入出力例のjsonファイルを作成する
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
	fmt.Println("Sampleファイルを作成しました")

	var samples []sample
	for i := 0; strings.Index(html, "<pre>") < strings.Index(html, "Problem Statement"); i++ {
		samples = append(samples, sample{})
		samples[i].Input = html[strings.Index(html, "<pre>")+5 : strings.Index(html, "</pre>")-2]
		rmPre(&html)
		samples[i].Output = html[strings.Index(html, "<pre>")+5 : strings.Index(html, "</pre>")-2]
		rmPre(&html)
	}
	data, _ := json.Marshal(samples)
	f.Write(data)
}

func isExist(filename string) bool { // ファイル、フォルダが存在するかどうか
	_, err := os.Stat(filename)
	return err == nil
}

func readSampleFile(filename string) []sample { // jsonファイルを読み込み、構造体でreturn
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var samples []sample
	if err := json.Unmarshal(bytes, &samples); err != nil {
		log.Fatal(err)
	}
	return samples
}

func TestSolve(t *testing.T) {
	if !isExist("pages") { // pagesフォルダがない場合、作成
		err := os.Mkdir("pages", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	url := "https://atcoder.jp/contests/abc148/tasks/abc148_a"
	filename := "pages/" + url[strings.LastIndex(url, "/")+1:] + ".json"
	if !isExist(filename) { // sampleファイルがない場合、作成
		createSampleFile(url, filename)
	}
	samples := readSampleFile(filename) // sampleファイルから読み込み
	for count, sample := range samples {
		fmt.Printf("Q%v answer: %s\treply : ", count+1, sample.Output)
		solve(strings.Fields(sample.Input)) // reply = 自分の出力

		if sample.Output != reply { // 答え合わせ
			t.Errorf("\x1b[1;31mQ%v: %v != %v\x1b[0m", count+1, sample.Output, reply)
		}
	}
}
