# Pipeline
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/pipeline.svg)](https://pkg.go.dev/github.com/gopi-frame/pipeline)
[![Go](https://github.com/gopi-frame/pipeline/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/gopi-frame/pipeline/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gopi-frame/pipeline/graph/badge.svg?token=UGVGP6QF5O)](https://codecov.io/gh/gopi-frame/pipeline)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/pipeline)](https://goreportcard.com/report/github.com/gopi-frame/pipeline)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
## Installation

```shell
go get -u -v github.com/gopi-frame/pipeline
```

## Import

```go
import "github.com/gopi-frame/pipeline"
```

## Usage

```go
package main

import "github.com/gopi-frame/pipeline"

type Pow struct {}

func (p Pow) Handle(value int, f func(int) int) int {
	value *= value
    return f(value)
}

func (p *Pow) Calculate(i int) int {

func main() {
    var number = 2
	p := pipeline.New[int, int]()
	pipe1 := pipeline.Pipe(func(value int, f func(int) int) int {
		value *= value
		return f(value)
	})
	pipe2 := Pow{}
	dest := func(value int) int {
		return value *= value
	}
	result := p.Send(number).Through(pipe1, pipe2).Then(dest)
	fmt.Println(result) // 256
}
```