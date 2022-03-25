package iter

import "context"

// FromChan converts the provided channel into an Iterator. Most methods on the
// Iterator will block until the channel can be received from, and possibly
// will not terminate until the channel is empty and closed.
func FromChan[E any](c <-chan E) Iterator[E] {
	return FromCoreFunc[E](func() (E, bool) {
		out, more := <-c
		return out, more
	})
}

// ToChan converts the current Iterator into a channel. Note that the returned
// context.CancelFunc should always be called once use of the channel is
// complete to ensure background goroutines are cleaned up.
func (iter Iterator[E]) ToChan() (ch <-chan E, cancel context.CancelFunc) {
	return iter.ToChanWithContext(context.Background())
}

// ToChanWithContext converts the current Iterator into a channel. If ctx is
// canceled, the returned channel will be closed, and the iterator may not be
// fully exhausted. Note that the returned context.CancelFunc should always be
// called once use of the channel is complete to ensure background goroutines
// are cleaned up.
func (iter Iterator[E]) ToChanWithContext(ctx context.Context) (ch <-chan E, cancel context.CancelFunc) {
	c := make(chan E)
	ctx, cancel = context.WithCancel(ctx)

	go func() {
		defer close(c)
		defer cancel()

		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			select {
			case c <- next:
				// continue
			case <-ctx.Done():
				return
			}
		}
	}()

	return c, cancel
}
