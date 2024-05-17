package pipeline

import "github.com/gopi-frame/contract/pipeline"

// New new pipeline
func New[P, R any]() pipeline.Pipeline[P, R] {
	return &Pipeline[P, R]{}
}

// Stop create stop from an anonymous function
func Stop[P, R any](callback func(P, func(P) R) R) pipeline.Stop[P, R] {
	return &stop[P, R]{callback}
}
