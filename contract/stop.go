package contract

// Pipe stop interface
type Pipe[P, R any] interface {
	Handle(P, func(P) R) R
}
