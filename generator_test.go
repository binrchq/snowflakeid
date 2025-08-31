package snowflakeid

import (
	"testing"
)

func TestGen(t *testing.T) {
	g := NewGenerator(0b001, 0b00001)
	g.Prefix = HostTen
	for i := 0; i < 100; i++ {
		// time.Sleep(5 * time.Millisecond)
		t.Log(g.NextID())
	}
}
