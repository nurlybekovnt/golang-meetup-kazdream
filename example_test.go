package golangmeetupkazdream

import (
	"bytes"
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	q := &Queue{}
	c := &Customer{Level: 1}
	for i := 0; i < b.N; i++ {
		q.Add(c)
		q.Next()
	}
}

func BenchmarkConversionCustomer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CustomerEncodedLen()
	}
}

func BenchmarkWritingKeyValues(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.Reset()
		WriteKeyValues(buf)
	}
}
