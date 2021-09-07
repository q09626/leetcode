package main

import "fmt"

type BitMap struct {
	bits []byte
	len  int // 表示数据量的个数
}

func NewBitMap(len int) *BitMap {
	return &BitMap{
		make([]byte, len>>3+1), // >>3表示除8，余数不足8也要用1个byte表示
		len,
	}
}

func (b *BitMap) Set(n uint) {
	index := n >> 3           // n>>3等价于n/8，8=2^3
	pos := n & 7              // n&7等价于n%8，7=8-1
	b.bits[index] |= 1 << pos // 1<<pos等价于 2^pos
}

func (b *BitMap) Reset(n uint) {
	index := n >> 3
	pos := n & 7
	b.bits[index] &= ^(1 << pos)
}

func (b *BitMap) Test(n uint) bool {
	index := n >> 3
	pos := n & 7
	return b.bits[index]&(1<<pos) != 0
}

func (b *BitMap) String() string {
	return fmt.Sprint(b.bits)
}

func main() {
	b := NewBitMap(10)
	b.Set(8)
	b.Set(0)
	fmt.Println(b.String())
}
