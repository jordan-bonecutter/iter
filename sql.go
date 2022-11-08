package iter

import (
	"database/sql"
)

// Scanner scans an sql.Rows into some type T.
// It doesn't pass scanning errors up to the next layer although
// this can easily be implemented by setting T to be a type wrapping
// some underlying type in a struct with an error a la:
// type Result[T any] struct {
//	 T T
//	 Err error
// }
type Scanner[T any] interface {
	Scan(*sql.Rows) T
}

// SQL implements iter for sql.Rows. It doesn't close the rows when it's done.
type SQL[T any] struct {
	*sql.Rows
	Scanner[T]
}

// Implement Iter[T].
func (r SQL[T]) ForEach(f func(t T) (stop bool)) {
	for r.Next() {
		t := r.Scanner.Scan(r.Rows)
		if f(t) {
			return
		}
	}
}
