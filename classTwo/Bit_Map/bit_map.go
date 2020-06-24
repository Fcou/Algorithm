package bitmap

type BitMap struct {
	bytes []byte
	nbits int
}

func (b *BitMap) newBitMap(nbits int) {
	b.nbits = nbits
	b.bytes = make([]byte, nbits/8+1) // Go中byte类型占8bit，也即是1个字节
}

// 存在数字k
func (b *BitMap) set(k int) {
	if k > b.nbits {
		return
	}
	byteIndex := k / 8
	bitIndex := k % 8
	b.bytes[byteIndex] |= (1 << bitIndex)
}

// 判断是否存在数字k
func (b *BitMap) get(k int) bool {
	if k > b.nbits {
		return false
	}
	byteIndex := k / 8
	bitIndex := k % 8
	return (b.bytes[byteIndex] & (1 << bitIndex)) != 0
}
