package contract

// Pipeline pipeline interface
type Pipeline[P, R any] interface {
	Send(P) Pipeline[P, R]
	Through(steps ...Pipe[P, R]) Pipeline[P, R]
	Then(destination func(P) R) R
}
