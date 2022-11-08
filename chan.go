package iter

// Chan implements an iterator for a channel.
type Chan[T any] <-chan T

func (c Chan[T]) ForEach(f func(t T) (stop bool)) {
	for t := range c {
		if f(t) {
			return
		}
	}
}
