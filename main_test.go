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

func rmPre(buf *string) {
	// bufから、<pre>と</pre>を1つずつ削除
	*buf = strings.Replace(*buf, "<pre>", "", 1)
	*buf = strings.Replace(*buf, "</pre>", "", 1)
}

func createSampleFile(url, filename string) {
	// urlのhtmlをstringで取得
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

	var samples []sample
	previousOfPre := func() int { return strings.Index(html, "<pre>") + 5 }
	endOfPre := func() int { return strings.Index(html, "</pre>") - 2 }
	rmPre(&html)

	// 日本語のhtmlの範囲の間、構造体samplesに入出力例を代入
	for i := 0; strings.Index(html, "<pre>") < strings.Index(html, "Problem Statement"); i++ {
		samples = append(samples, sample{})
		samples[i].Input = html[previousOfPre():endOfPre()]
		rmPre(&html)
		samples[i].Output = html[previousOfPre():endOfPre()]
		rmPre(&html)
	}

	// jsonファイルを作成、samplesをjsonエンコードしてそのファイルに書き込む
	f, _ := os.Create(filename)
	defer f.Close()
	fmt.Println(filename + "を作成しました")
	data, _ := json.Marshal(samples)
	f.Write(data)
}

// ファイル、フォルダが存在するかどうか。
// ある => true
// ない => false
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readSampleFile(filename string) []sample {
	// jsonファイル読み込み
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// jsonデコードして、構造体に変換してret
	var samples []sample
	if err := json.Unmarshal(bytes, &samples); err != nil {
		log.Fatal(err)
	}
	return samples
}

func TestSolve(t *testing.T) {
	url := "https://atcoder.jp/contests/abc148/tasks/abc148_b"
	filename := "pages/" + url[strings.LastIndex(url, "/")+1:] + ".json" // filename = pages/abc0_a.json

	// pagesフォルダがない場合、作成
	if !isExist("pages") {
		err := os.Mkdir("pages", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// filenameのファイルがない場合、作成・書き込み
	if !isExist(filename) {
		createSampleFile(url, filename)
	}

	// sampleファイルから読み込み、[]sampleで返ってくる
	samples := readSampleFile(filename)

	// 答え合わせ
	for count, sample := range samples {
		fmt.Printf("Q%v answer: %s\treply : ", count+1, sample.Output)

		solve(strings.Fields(sample.Input)) // main.goの関数。reply = 自分の出力
		if sample.Output != reply {
			t.Errorf("\x1b[1;31mQ%v: %v != %v\x1b[0m", count+1, sample.Output, reply)
		}
	}
}
