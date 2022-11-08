package iter

import (
	zmq "github.com/pebbe/zmq4"
)

// ZMQSocket implements an Iter[T] for a zmq socket.
type ZMQSocket zmq.Socket

func (s *ZMQSocket) ForEach(f func(Result[[][]byte]) (stop bool)) {
	for {
		b, err := ((*zmq.Socket)(s)).RecvMessageBytes(0)
		if f(Result[[][]byte]{b, err}) {
			return
		}
	}
}
