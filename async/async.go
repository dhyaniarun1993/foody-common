// Copy of sync errgroup with custom error

package async

import (
	"context"
	"sync"

	"github.com/dhyaniarun1993/foody-common/errors"
)

// Async provides goroutines working on subtasks of a single task
type Async struct {
	cancel  func()
	wg      sync.WaitGroup
	errOnce sync.Once
	err     errors.AppError
}

// WithContext returns a new Async and associated Context
func WithContext(ctx context.Context) (*Async, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Async{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Async) Wait() errors.AppError {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Async) Go(f func() errors.AppError) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
