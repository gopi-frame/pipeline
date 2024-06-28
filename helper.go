package pipeline

import "github.com/gopi-frame/pipeline/contract"

// New new pipeline
func New[P, R any]() contract.Pipeline[P, R] {
	return &Pipeline[P, R]{}
}

// Stop create stop from an anonymous function
func Stop[P, R any](callback func(P, func(P) R) R) contract.Pipe[P, R] {
	return &stop[P, R]{callback}
}
