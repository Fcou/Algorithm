package bitmap

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	b := newBitMap(50)
	for i := 20; i <= 80; i++ {
		b.set(i)
	}
	for i := 0; i < 200; i++ {
		if b.get(i) {
			fmt.Println("存在", i)
		} else {
			fmt.Println("不存在", i)
		}
	}
}
