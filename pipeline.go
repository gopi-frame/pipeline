package pipeline

import (
	"github.com/gopi-frame/contract/pipeline"
)

// Pipeline pipeline
type Pipeline[P, R any] struct {
	passable    P
	stops       []pipeline.Pipe[P, R]
	destination func(passable P) R
}

// Send send
func (p *Pipeline[P, R]) Send(passable P) pipeline.Pipeline[P, R] {
	p.passable = passable
	return p
}

// Through through stops
func (p *Pipeline[P, R]) Through(stops ...pipeline.Pipe[P, R]) pipeline.Pipeline[P, R] {
	p.stops = stops
	return p
}

// Then then
func (p *Pipeline[P, R]) Then(destination func(P) R) R {
	p.destination = destination
	stack := p.destination
	length := len(p.stops)
	for i := length; i > 0; i-- {
		pipe := p.stops[i-1]
		stack = func(s func(passable P) R, pipe pipeline.Pipe[P, R]) func(passable P) R {
			return func(passable P) R {
				return pipe.Handle(passable, s)
			}
		}(stack, pipe)
	}
	return stack(p.passable)
}
