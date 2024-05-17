package pipeline

type stop[P, R any] struct {
	callback func(P, func(P) R) R
}

func (s *stop[P, R]) Handle(passable P, next func(P) R) R {
	return s.callback(passable, next)
}
