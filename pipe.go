package pipeline

import pc "github.com/gopi-frame/contract/pipeline"

type pipe[P, R any] struct {
	callback func(P, func(P) R) R
}

func (s *pipe[P, R]) Handle(passable P, next func(P) R) R {
	return s.callback(passable, next)
}

func Pipe[P, R any](callback func(P, func(P) R) R) pc.Pipe[P, R] {
	return &pipe[P, R]{callback}
}
