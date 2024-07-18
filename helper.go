package pipeline

import "github.com/gopi-frame/contract/pipeline"

// New new pipeline
func New[P, R any]() pipeline.Pipeline[P, R] {
	return &Pipeline[P, R]{}
}

// Pipe create pipe from an anonymous function
func Pipe[P, R any](callback func(P, func(P) R) R) pipeline.Pipe[P, R] {
	return &pipe[P, R]{callback}
}
