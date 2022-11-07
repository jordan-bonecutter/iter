package iter

// Iter is an iterable object.
type Iter[T any] interface {
	ForEach(func(t T) (stop bool))
}

// Counter represents the maximum size of an iterator.
// Not all iterators need to implement this method, we'll check at runtime if they do.
type Counter interface {
	Count() int
}

// Map an Iter[T1] to an Iter[T2] with the mapping function f.
func Map[T1, T2 any](iter Iter[T1], f func(t1 T1) (t2 T2, stop bool)) (out Iter[T2]) {
	if _, ok := iter.(Counter); ok {
		return mapCounter[T1, T2]{
			mapIt: mapIt[T1, T2]{
				in: iter, transformer: f,
			},
		}
	}

	return mapIt[T1, T2]{
		in: iter, transformer: f,
	}
}

// map implements a mapping iterator
type mapIt[T1, T2 any] struct {
	in          Iter[T1]
	transformer func(t1 T1) (t2 T2, stop bool)
}

// ForEach ranges over the map.
func (m mapIt[T1, T2]) ForEach(f func(T2) (stop bool)) {
	m.in.ForEach(func(in T1) (stop bool) {
		t2, stop := m.transformer(in)
		if stop {
			return true
		}

		return f(t2)
	})
}

// mapCounters implement Map and Counter
type mapCounter[T1, T2 any] struct {
	mapIt[T1, T2]
}

// We should only ever create a mapCounter when we know that the underlying
// map's Iter is also a Counter. Thus we don't need to check here.
func (c mapCounter[T1, T2]) Count() int {
	return c.mapIt.in.(Counter).Count()
}

// Filter an Iter[T] with a filtering function.
func Filter[T any](iter Iter[T], f func(t T) bool) Iter[T] {
	if _, ok := iter.(Counter); ok {
		return filterCounter[T]{
			filter: filter[T]{
				in: iter, filter: f,
			},
		}
	}

	return filter[T]{
		in: iter, filter: f,
	}
}

// filter implements the filter operation on Iter[T].
type filter[T any] struct {
	in     Iter[T]
	filter func(T) (include bool)
}

func (filt filter[T]) ForEach(f func(t T) (stop bool)) {
	filt.in.ForEach(func(t T) (stop bool) {
		if filt.filter(t) {
			return f(t)
		}

		return false
	})
}

type filterCounter[T any] struct {
	filter[T]
}

func (f filterCounter[T]) Count() int {
	return f.filter.in.(Counter).Count()
}

// Take at most count items from the given iterable.
func Take[T any](iter Iter[T], count int) Iter[T] {
	return take[T]{
		it: iter, count: count,
	}
}

// take implements the take operator on an iterator.
type take[T any] struct {
	it    Iter[T]
	count int
}

func (taker take[T]) ForEach(f func(t T) (stop bool)) {
	count := 0
	taker.it.ForEach(func(t T) (stop bool) {
		if count == taker.count {
			return true
		}
		count++

		return f(t)
	})
}

func (t take[T]) Count() int {
	if innerCounter, ok := t.it.(Counter); ok {
		innerCount := innerCounter.Count()
		if innerCount < t.count {
			return innerCount
		}
	}

	return t.count
}

// Collect the iterator into a static slice of type T
func Collect[T any](iter Iter[T]) (collection []T) {
	if counter, ok := iter.(Counter); ok {
		collection = make([]T, 0, counter.Count())
	}

	iter.ForEach(func(t T) (stop bool) {
		collection = append(collection, t)
		return false
	})

	return collection
}

// MapCollect collects the iterator into a map using the given key-value extractor.
func MapCollect[T, V any, K comparable](iter Iter[T], kvFunc func(t T) (k K, v V)) (collection map[K]V) {
	if counter, ok := iter.(Counter); ok {
		collection = make(map[K]V, counter.Count())
	}

	iter.ForEach(func(t T) (stop bool) {
		k, v := kvFunc(t)
		collection[k] = v
		return false
	})

	return collection
}
