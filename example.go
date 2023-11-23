package golangmeetupkazdream

import (
	"encoding/binary"
	"sort"
)

type Customer struct {
	Level   uint32
	Age     uint32
	Balance uint64
}

func (c Customer) Bytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint32(buf, c.Level)
	binary.BigEndian.PutUint32(buf[4:], c.Age)
	binary.BigEndian.PutUint64(buf[8:], c.Balance)
	return buf
}

type Queue struct {
	customers []*Customer
}

func (q *Queue) Add(c *Customer) {
	q.customers = append(q.customers, c)
	sort.Slice(q.customers, func(i, j int) bool {
		return q.customers[i].Level < q.customers[j].Level
	})
}

func (q *Queue) Next() (*Customer, bool) {
	if len(q.customers) == 0 {
		return nil, false
	}

	c := q.customers[0]
	q.customers = append(q.customers[:0], q.customers[1:]...)
	return c, true
}

func CustomerEncodedLen() int {
	c := Customer{}
	return BytesLen(c)
}

type BytesInterface interface {
	Bytes() []byte
}

func BytesLen(b BytesInterface) int {
	return len(b.Bytes())
}
