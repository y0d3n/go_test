# go_test

[AtCoder ABC 148 A](https://atcoder.jp/contests/abc148/tasks/abc148_a)

## ファイル構成
```
$ tree
```
>.
├── main.go
├── main_test.go
└── pages
    ├── abc148_a.json
    ├── abc148_b.json
    └── abc148_c.json....

## 使用法
```
$ go test
```
### ok
```
Q1 answer: 2    reply : 2
Q2 answer: 3    reply : 3
PASS
ok      [CUR PATH]  0.015s
```

### fail
```
Q1 answer: 2    reply : 4
Q2 answer: 3    reply : 3
--- FAIL: TestSolve (0.00s)
        main_test.go:109: Q1: 2 != 4
FAIL
exit status 1
FAIL    [CUR PATH]  0.013s
```