package iter

import (
	"fmt"
)

type Result[T any] struct {
	T T
	Err error
}

func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(fmt.Sprintf("Attempted unwrapping failed result: %s", r.Err.Error()))
	}
	
	return r.T
}

type Responder[Request, ResponseMedium any] struct {
	Request Request
	ResponseMedium ResponseMedium
}
