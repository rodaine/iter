package iter

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

type countUpCore[T Integer] struct {
	next T
}

func (cc *countUpCore[T]) Next() (T, bool) {
	next := cc.next
	cc.next++
	return next, true
}

func CountUp[T Integer](start T) Iterator[T] {
	return FromCore[T](&countUpCore[T]{
		next: start,
	})
}

type countDownCore[T Integer] struct {
	next T
}

func (cc *countDownCore[T]) Next() (T, bool) {
	next := cc.next
	cc.next--
	return next, true
}

func CountDown[T Integer](start T) Iterator[T] {
	return FromCore[T](&countDownCore[T]{
		next: start,
	})
}

type rangeUpCore[T Integer] struct {
	ct  *countUpCore[T]
	max T
}

func (rc rangeUpCore[T]) Next() (T, bool) {
	if rc.ct.next >= rc.max {
		return empty[T]()
	}

	return rc.ct.Next()
}

type rangeDownCore[T Integer] struct {
	ct  *countDownCore[T]
	min T
}

func (rc rangeDownCore[T]) Next() (T, bool) {
	if rc.ct.next <= rc.min {
		return empty[T]()
	}

	return rc.ct.Next()
}

func Range[T Integer](start, end T) Iterator[T] {
	if start <= end {
		return FromCore[T](rangeUpCore[T]{
			ct:  &countUpCore[T]{next: start},
			max: end,
		})
	}

	return FromCore[T](rangeDownCore[T]{
		ct:  &countDownCore[T]{next: start},
		min: end,
	})
}
