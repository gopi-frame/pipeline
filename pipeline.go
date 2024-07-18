package pipeline

import (
	"github.com/gopi-frame/contract/pipeline"
)

// Pipeline is a struct that implements the [pipeline.Pipeline] interface.
//
// It contains the passable value and the pipes of the pipeline.
// The destination function is the function that is applied to the passable value of the pipeline.
type Pipeline[P, R any] struct {
	passable    P
	pipes       []pipeline.Pipe[P, R]
	destination func(passable P) R
}

// New creates a new instance of the Pipeline[P, R] struct, which implements the pipeline.Pipeline[P, R] interface.
func New[P, R any]() pipeline.Pipeline[P, R] {
	return &Pipeline[P, R]{}
}

// Send sets the passable value of the pipeline and returns the pipeline.
func (p *Pipeline[P, R]) Send(passable P) pipeline.Pipeline[P, R] {
	p.passable = passable
	return p
}

// Through adds the provided pipes to the pipeline. The pipes will be executed in the order they are provided.
// The pipeline will return itself, allowing for method chaining.
func (p *Pipeline[P, R]) Through(pipes ...pipeline.Pipe[P, R]) pipeline.Pipeline[P, R] {
	p.pipes = pipes
	return p
}

// Then applies the provided destination function to the passable value after executing the pipes in the pipeline.
// The pipes are executed in the reverse order they were added to the pipeline.
// The final result of the pipeline is returned.
func (p *Pipeline[P, R]) Then(destination func(P) R) R {
	p.destination = destination
	stack := p.destination
	length := len(p.pipes)
	for i := length; i > 0; i-- {
		pipe := p.pipes[i-1]
		stack = func(s func(passable P) R, pipe pipeline.Pipe[P, R]) func(passable P) R {
			return func(passable P) R {
				return pipe.Handle(passable, s)
			}
		}(stack, pipe)
	}
	return stack(p.passable)
}
