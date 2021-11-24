package snow

import (
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	es := &Snow{
		last: time.Now().Add(time.Second).Unix(),
	}
	if es.Next() != 0 {
		t.Fatal(es)
	}
	ti := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-ti.C:
			return
		default:
		}
		for i := 0; i < 100; i++ {
			es.Next()
			Next()
		}
	}
}
