package sha256

//主要步骤：
//1. 添加pad 和 length 补足长度为 512bit 整数倍
//2. 对于每个分组 512bit 进行迭代计算，每次输入 h0~h8， 输出 h0~h8, 每组中 512bit分为 16个uint32(大端序),
//	然后扩展成64个uint32, 之后经过计算获得 a~h，最后 h0+=a h1+=b...
//3. 通过每组512bit 修改一次 h0~h8, 最后一组h0~h8 使用大端序转换为[]byte， 然后链接起来即可
//与md5很相似，不过需要小心md5中多使用小端序，sha256中使用大端序
//参考：https://blog.csdn.net/u011583927/article/details/80905740
type Sha256Hash struct {
	h0, h1, h2, h3, h4, h5, h6, h7 uint32
	ks [64]uint32
}

func (self *Sha256Hash) init(){
	self.h0 = 0x6a09e667
	self.h1 = 0xbb67ae85
	self.h2 = 0x3c6ef372
	self.h3 = 0xa54ff53a
	self.h4 = 0x510e527f
	self.h5 = 0x9b05688c
	self.h6 = 0x1f83d9ab
	self.h7 = 0x5be0cd19

	self.ks = [64]uint32{
		0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
	}
}

func pad(src []byte) []byte{
	dataLen := len(src)

	var padLen int
	left := dataLen%64
	if 56 > left{
		padLen = 56 - left
	}else{
		padLen = 64 - left + 56
	}

	rst := make([]byte, dataLen + padLen + 8)
	copy(rst, src)
	rst[dataLen] = 0x80

	//不同于md5, sha256中长度字段使用大端序
	bitLen := dataLen*8
	length := make([]byte, 8)
	length[0] = byte(uint64(bitLen)>>56)
	length[1] = byte(uint64(bitLen)>>48)
	length[2] = byte(uint64(bitLen)>>40)
	length[3] = byte(uint64(bitLen)>>32)
	length[4] = byte(uint64(bitLen)>>24)
	length[5] = byte(uint64(bitLen)>>16)
	length[6] = byte(uint64(bitLen)>>8)
	length[7] = byte(uint64(bitLen))


	copy(rst[dataLen+padLen:], length)

	return rst
}


func group(src []byte) [][]byte{

	if len(src) % 64 != 0{
		return nil
	}

	count := len(src)/64
	rst := make([][]byte, 0, count)

	for index:=0; index<len(src); index+=64{
		rst = append(rst, src[index:index+64])
	}

	return rst
}

func groupUint32(src []byte) []uint32{

	if len(src) != 64 {
		return nil
	}

	rst := make([]uint32, 16)
	for index := 0; index < 64; index += 4 {
		tmp := uint32(src[index]) << 24 | uint32(src[index+1]) << 16 | uint32(src[index+2]) << 8 | uint32(src[index+3])
		rst[index/4] = tmp
	}

	return rst
}

func circularRightShift(src uint32, shift uint32) uint32{
	return (src >> shift) | (src << (32-shift))
}

func (self *Sha256Hash)trans(src []byte) {

	if len(src) != 64{
		return
	}

	w := make([]uint32, 64)
	tmp := groupUint32(src)
	copy(w, tmp)
	for index := 16; index < 64; index++ {
		s0 := circularRightShift(w[index-15], 7) ^ circularRightShift(w[index-15], 18) ^ (w[index-15] >> 3)
		s1 := circularRightShift(w[index-2], 17) ^ circularRightShift(w[index-2], 19) ^ (w[index-2] >> 10)
		w[index] = w[index-16] + s0 + w[index-7] + s1
	}

	a := self.h0
	b := self.h1
	c := self.h2
	d := self.h3
	e := self.h4
	f := self.h5
	g := self.h6
	h := self.h7

	for i := 0; i< 64; i++ {

		s0 := circularRightShift(a, 2) ^ circularRightShift(a, 13) ^ circularRightShift(a, 22)
		maj := (a & b) ^ (a & c) ^ (b & c)
		t2 := s0 + maj
		s1 := circularRightShift(e, 6) ^ circularRightShift(e, 11) ^ circularRightShift(e, 25)
		ch := (e & f) ^ ((^e) & g)
		t1 := h + s1 +ch + self.ks[i] + w[i]
		h = g
		g = f
		f = e
		e = d + t1
		d= c
		c = b
		b = a
		a = t1 + t2
	}

	self.h0 = self.h0 + a
	self.h1 = self.h1 + b
	self.h2 = self.h2 + c
	self.h3 = self.h3 + d
	self.h4 = self.h4 + e
	self.h5 = self.h5 + f
	self.h6 = self.h6 + g
	self.h7 = self.h7 + h

	return
}

func transNumtoByte(src ...uint32) []byte{

	rst := make([]byte, 0, len(src)*4)
	for _,v := range src{
		rst = append(rst, byte(v >> 24))
		rst = append(rst, byte(v >> 16))
		rst = append(rst, byte(v >> 8))
		rst = append(rst, byte(v))
	}
	return rst

}

func (self *Sha256Hash) Hash(src []byte) []byte{

	self.init()
	padedSrc := pad(src)
	gs := group(padedSrc)

	for _,v := range gs{
		self.trans(v)
	}

	return transNumtoByte(self.h0, self.h1, self.h2, self.h3, self.h4, self.h5, self.h6, self.h7)
}
