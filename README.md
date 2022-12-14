# orderhash

[![GitHub](https://img.shields.io/github/license/chenquan/orderhash)](https://github.com/chenquan/orderhash/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/chenquan/orderhash/branch/master/graph/badge.svg?token=74phc5KVI7)](https://codecov.io/gh/chenquan/orderhash)
[![Release](https://img.shields.io/github/v/release/chenquan/orderhash.svg?style=flat-square)](https://github.com/chenquan/orderhash)
[![Godoc](https://img.shields.io/badge/godoc-reference-brightgreen)](https://pkg.go.dev/github.com/chenquan/orderhash)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenquan/orderhash)](https://goreportcard.com/report/github.com/chenquan/orderhash)
[![Download](https://goproxy.cn/stats/github.com/chenquan/orderhash/badges/download-count.svg)](https://github.com/chenquan/orderhash)


> A local stateful orderly hash algorithm.

## installation


```shell
go get -u github.com/chenquan/orderhash
```

## purpose

It is used to ensure the order of consumer and producer message consumption and sending during concurrency, and improve the concurrent throughput.

example: https://github.com/chenquan/go-queue/blob/master/kafka
