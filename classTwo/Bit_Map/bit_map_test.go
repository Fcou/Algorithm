package bitmap

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	b := newBitMap(201326592)

	b.set(178288549)

	if b.get(178288549) {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

}
