package iter

type IterSlice[T any] []T

func (s IterSlice[T]) ForEach(f func(t T) (stop bool)) {
	for _, t := range s {
		f(t)
	}
}

func (s IterSlice[T]) Count() int {
	return len(s)
}
