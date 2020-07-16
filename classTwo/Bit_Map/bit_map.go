// 将数字 A 的第 k 位设置为1：A = A | (1 << (k - 1))
// 将数字 A 的第 k 位设置为0：A = A & ~(1 << (k - 1))
// 检测数字 A 的第 k 位：A & (1 << (k - 1)) != 0
package bitmap

type BitMap struct {
	bytes []byte
	nbits int //可存储的最大值
}

func newBitMap(nbits int) *BitMap {
	bytes := make([]byte, nbits/8+1) // Go中byte类型占8bit，也即是1个字节
	return &BitMap{nbits: nbits, bytes: bytes}
}

// 存在数字k
func (b *BitMap) set(k int) {
	if k > b.nbits {
		return
	}
	byteIndex := k / 8
	bitIndex := k % 8
	b.bytes[byteIndex] |= (1 << bitIndex) // 从0到无穷大，每8个数对应一个byte存储
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
