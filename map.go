package iter

type IterMap[K comparable, V any] map[K]V

func (i IterMap[K, V]) ForEach(f func(KeyValue[K, V]) (stop bool)) {
	for k, v := range i {
		f(KeyValue[K, V]{k, v})
	}
}

func (i IterMap[K, V]) Count() int {
	return len(i)
}

type KeyValue[K comparable, V any] struct {
	K K
	V V
}
