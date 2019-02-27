package bitmap

func setBit(n int, pos uint) int {
	n |= 1 << pos
	return n
}

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}

type BitMap struct {
	words []uint64
}


func New() *BitMap {
	return &BitMap{}
}

func (bmp *BitMap) HasBit(pos int) bool {
	idx, bit := pos/64, uint(pos%64)
	return idx < len(bmp.words) && bmp.words[idx] & 1 << bit != 0
}

func (bmp *BitMap) SetBit(pos int)  {
	idx, bit := pos/64, uint(pos%64)
	for idx > len(bmp.words) {
		bmp.words = append(bmp.words, 0)
	}
	if !bmp.HasBit(pos) {
		bmp.words[idx] |= 1 << bit
	}

}

func (bmp *BitMap) ClearBit(pos int) {
	idx, bit := pos/64, uint(pos%64)
	mask := ^(1 << bit)
	bmp.words[idx] &= uint64(mask)
}

func main() {

}

