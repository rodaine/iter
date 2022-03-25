package iter

import "golang.org/x/exp/constraints"

// CountUpBy returns an unbounded Iterator beginning (inclusively) at the
// provided start value, and incrementing by the provided step. For an iterator
// that terminates, use RangeBy instead.
func CountUpBy[I constraints.Integer](start, step I) Iterator[I] {
	return FromCore[I](&countUpByCore[I]{
		next: start,
		step: step,
	})
}

// CountDownBy returns an unbounded Iterator beginning (inclusively) at the
// provided start value, and decrements by the provided step. For an iterator
// that terminates, use RangeBy instead.
func CountDownBy[I constraints.Integer](start, step I) Iterator[I] {
	return FromCore[I](&countDownByCore[I]{
		next: start,
		step: step,
	})
}

// RangeBy returns an Iterator that begins (inclusively) at the provided start
// value, approaches the end value by the provided step, and terminates when
// the end value is reached or passed (exclusive).
func RangeBy[I constraints.Integer](start, end, step I) Iterator[I] {
	if start <= end {
		return FromCore[I](rangeUpCore[I]{
			ct:  &countUpByCore[I]{next: start, step: step},
			max: end,
		})
	}

	return FromCore[I](rangeDownCore[I]{
		ct:  &countDownByCore[I]{next: start, step: step},
		min: end,
	})
}

type countUpByCore[I constraints.Integer] struct {
	next I
	step I
}

func (cubc *countUpByCore[I]) Next() (next I, ok bool) {
	next = cubc.next
	cubc.next += cubc.step
	return next, true
}

type countDownByCore[I constraints.Integer] struct {
	next I
	step I
}

func (cdbc *countDownByCore[I]) Next() (next I, ok bool) {
	next = cdbc.next
	cdbc.next -= cdbc.step
	return next, true
}

type rangeUpCore[I constraints.Integer] struct {
	ct  *countUpByCore[I]
	max I
}

func (rc rangeUpCore[I]) Next() (I, bool) {
	if rc.ct.next >= rc.max {
		return empty[I]()
	}

	return rc.ct.Next()
}

type rangeDownCore[I constraints.Integer] struct {
	ct  *countDownByCore[I]
	min I
}

func (rc rangeDownCore[I]) Next() (I, bool) {
	if rc.ct.next <= rc.min {
		return empty[I]()
	}

	return rc.ct.Next()
}
