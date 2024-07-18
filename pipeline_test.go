package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	var pipeline = New[string, string]()
	var str = "pipeline"
	str = pipeline.Send(str).Through(
		Pipe(func(s string, next func(string) string) string {
			s += "pipe1"
			return next(s)
		}),
		Pipe(func(s string, next func(string) string) string {
			s += "pipe2"
			return next(s)
		}),
	).Then(func(s string) string {
		s += "then"
		return s
	})
	assert.Equal(t, "pipelinepipe1pipe2then", str)
}
