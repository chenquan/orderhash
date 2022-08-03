# orderhash

[![Godoc](https://img.shields.io/badge/godoc-reference-brightgreen)](https://pkg.go.dev/github.com/chenquan/orderhash)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenquan/orderhash)](https://goreportcard.com/report/github.com/chenquan/orderhash)
[![Release](https://img.shields.io/github/v/release/chenquan/orderhash.svg?style=flat-square)](https://github.com/chenquan/orderhash)
[![codecov](https://codecov.io/gh/chenquan/orderhash/branch/master/graph/badge.svg?token=74phc5KVI7)](https://codecov.io/gh/chenquan/orderhash)
[![Download](https://goproxy.cn/stats/github.com/chenquan/orderhash/badges/download-count.svg)](https://github.com/chenquan/orderhash)
[![GitHub](https://img.shields.io/github/license/chenquan/orderhash)](https://github.com/chenquan/orderhash/blob/master/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/chenquan/orderhash)](https://github.com/chenquan/orderhash/stargazers)
[![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/chenquan/orderhash)](https://github.com/chenquan/orderhash/pulls)
[![GitHub issues](https://img.shields.io/github/issues/chenquan/orderhash)](https://github.com/chenquan/orderhash/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/chenquan/orderhash?color=red)](https://github.com/chenquan/orderhash/issues?q=is%3Aissue+is%3Aclosed)

> A local stateful orderly hash algorithm.

## installation


```shell
go get -u github.com/chenquan/orderhash
```

## purpose

It is used to guarantee the order of concurrent consumer and producer messages and to provide concurrent throughput capability.
