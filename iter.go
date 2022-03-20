package iter

type Iterator[T any] struct {
	core *peekCore[T]
}

func (iter Iterator[T]) Next() (T, bool) {
	return iter.core.Next()
}

func (iter Iterator[T]) Peek() (T, bool) {
	return iter.core.Peek()
}

func (iter Iterator[T]) Count() (n uint) {
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		n++
	}
	return n
}

func (iter Iterator[T]) Last() (last T, ok bool) {
	for out, cont := iter.Next(); cont; out, cont = iter.Next() {
		last = out
		ok = true
	}

	return last, ok
}

func (iter Iterator[T]) AdvanceBy(n uint) (ct uint, ok bool) {
	for i := uint(0); i < n; i++ {
		if _, ok = iter.Next(); !ok {
			return i, false
		}
	}
	return n, true
}

func (iter Iterator[T]) Nth(n uint) (nth T, ok bool) {
	for i := uint(0); i <= n; i++ {
		if nth, ok = iter.Next(); !ok {
			return nth, ok
		}
	}
	return nth, ok
}

func (iter Iterator[T]) ForEach(fn func(T)) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		fn(next)
	}
}

func (iter Iterator[T]) Partition(fn func(T) bool) (trues []T, falses []T) {
	iter.ForEach(func(el T) {
		if fn(el) {
			trues = append(trues, el)
		} else {
			falses = append(falses, el)
		}
	})

	return
}

func (iter Iterator[T]) All(fn func(T) bool) bool {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !fn(next) {
			return false
		}
	}

	return true
}

func (iter Iterator[T]) Any(fn func(T) bool) bool {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if fn(next) {
			return true
		}
	}
	return false
}

func (iter Iterator[T]) Find(fn func(T) bool) (T, bool) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if fn(next) {
			return next, true
		}
	}

	return empty[T]()
}

func (iter Iterator[T]) Position(fn func(T) bool) (uint, bool) {
	var idx uint
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if fn(next) {
			return idx, true
		}
		idx++
	}
	return empty[uint]()
}
