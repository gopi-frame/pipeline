# Pipeline

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
	result := p.Send(number).Through(
		pipeline.Pipe(func(value int, f func(int) int) int {
            value *= value			
            return f(value)    
	    },
		Pow{},
    ).Then(func(value int) int {
		return value
	}
	fmt.Println(result) // 16
}
```