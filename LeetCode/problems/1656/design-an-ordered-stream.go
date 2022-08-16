package main

type OrderedStream struct {
	Stream []string
	Ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.Stream[idKey] = value
	start := s.Ptr
	for s.Ptr < len(s.Stream) && s.Stream[s.Ptr] != "" {
		s.Ptr++
	}
	return s.Stream[start:s.Ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
