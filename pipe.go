package pipeline

type pipe[P, R any] struct {
	callback func(P, func(P) R) R
}

func (s *pipe[P, R]) Handle(passable P, next func(P) R) R {
	return s.callback(passable, next)
}
