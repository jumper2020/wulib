package md5

import (
	"bytes"
)

//1. 首先进行填充， 补充0x10000... 补足  %512 余数 为448, 之后添加补充前长度 (bit长度) 此处注意cpu 使用为大端还是小端， 比如长度 0x28,
//	对于大端cpu应该append 0x0000000000000028, 对于小端cpu应该append 0x2800000000000000, 否则cpu读取到的长度与正确的长度不一致
//	注意：填充是必须的，就算当前 %512 余数为 448 也需要填充 512bit
//2. 以512bit 为一个大组，对于每个大组进行计算，不断更新 a/b/c/d, 在一个大组中，划分16个小组，来进行计算。这里划分16个小组的时候，事实上
//	将之前以字符串方式存入的数据，以数值的方式读出了。
//	注意： FF GG HH II 函数中 << 表示的是循环左移， 非左移运算
//3. 将step 2中 a/b/c/d 组合成为 128bit 结果即可, 这里需要按照内存中的顺序来链接 a/b/c/d
//	比如 a = 0x12345678 cpu为小端， 内存中应该是 0x78 0x56 0x34 0x12， 这样构造出结果。具体参见 transNumtoByte函数

//参考资料：
//https://zhuanlan.zhihu.com/p/37257569
//https://www.ietf.org/rfc/rfc1321.txt

type Md5Hash struct {
	a,b,c,d uint32
}

func (self *Md5Hash) init(){
	self.a,self.b,self.c,self.d = 0x67452301,0xEFCDAB89,0x98BADCFE,0x10325476
}

func pad(src []byte) []byte{

	var padedSrc []byte
	var padLen uint64
	left := uint64(len(src)*8%512)

	if left < 448{
		padLen = (448-left)/8
	}else {
		padLen = (512-(left-448))/8
	}

	srcLen := uint64(len(src))
	totalLen := srcLen + padLen + 8
	padedSrc = make([]byte, totalLen)
	copy(padedSrc, src)
	padedSrc[srcLen] = 0x80
	copy(padedSrc[srcLen+1 : srcLen+padLen], bytes.Repeat([]byte{0x00}, int(padLen-1)))

	//为了尽可能减少依赖， 这里没有使用 binary.LittleEndian 相关函数
	lenByte := [8]byte{}
	bitLen := srcLen * 8
	lenByte[0] = byte((bitLen) & 0x00000000000000ff)
	lenByte[1] = byte(((bitLen) & 0x000000000000ff00) >> 8)
	lenByte[2] = byte(((bitLen) & 0x0000000000ff0000) >> 16)
	lenByte[3] = byte(((bitLen) & 0x00000000ff000000) >> 24)
	lenByte[4] = byte(((bitLen) & 0x000000ff00000000) >> 32)
	lenByte[5] = byte(((bitLen) & 0x0000ff0000000000) >> 40)
	lenByte[6] = byte(((bitLen) & 0x00ff000000000000) >> 48)
	lenByte[7] = byte(((bitLen) & 0xff00000000000000) >> 56)

	copy(padedSrc[srcLen+padLen:], lenByte[:])

	return padedSrc
}

//F( X ,Y ,Z ) = ( X & Y ) | ( (~X) & Z )
//G( X ,Y ,Z ) = ( X & Z ) | ( Y & (~Z) )
//H( X ,Y ,Z ) =X ^ Y ^ Z
//I( X ,Y ,Z ) =Y ^ ( X | (~Z) )

func F(x,y,z uint32) uint32{
	return (x & y) | ((^x) & z)
}

func G(x,y,z uint32) uint32{
	return (x & z) | (y & (^z))
}

func H(x,y,z uint32) uint32{
	return x ^ y ^ z
}

func I(x,y,z uint32) uint32{
	return y ^ (x | (^z))
}



//FF(a ,b ,c ,d ,Mj ,s ,ti ) 操作为 a = b + ( (a + F(b,c,d) + Mj + ti) << s)
//GG(a ,b ,c ,d ,Mj ,s ,ti ) 操作为 a = b + ( (a + G(b,c,d) + Mj + ti) << s)
//HH(a ,b ,c ,d ,Mj ,s ,ti) 操作为 a = b + ( (a + H(b,c,d) + Mj + ti) << s)
//II(a ,b ,c ,d ,Mj ,s ,ti) 操作为 a = b + ( (a + I(b,c,d) + Mj + ti) << s)

func circularLeftShift(src uint32, shift uint32) uint32{
	return (src << shift) | (src >> (32-shift))
}

func FF(a uint32, b uint32,c uint32, d uint32,Mj uint32,s uint32,ti uint32) uint32{
	return b + circularLeftShift((a+F(b,c,d) + Mj +ti), s)
}

func GG(a uint32, b uint32,c uint32, d uint32,Mj uint32,s uint32,ti uint32) uint32{
	return b + circularLeftShift((a+G(b,c,d) + Mj +ti), s)
}

func HH(a uint32, b uint32,c uint32, d uint32,Mj uint32,s uint32,ti uint32) uint32{
	return b + circularLeftShift((a+H(b,c,d) + Mj +ti), s)
}

func II(a uint32, b uint32,c uint32, d uint32,Mj uint32,s uint32,ti uint32) uint32{
	return b + circularLeftShift((a+I(b,c,d) + Mj +ti), s)
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
		tmp := uint32(src[index]) | uint32(src[index+1]) << 8 | uint32(src[index+2]) << 16 | uint32(src[index+3]) << 24
		rst[index/4] = tmp
	}

	return rst
}

func transNumtoByte(src ...uint32) []byte{
	//fmt.Printf("a: %x, b: %x, c: %x, d: %x\n", src[0], src[1], src[2], src[3])
	rst := make([]byte, 0, len(src)*4)
	for _,v := range src{
		rst = append(rst, byte((v & 0x000000ff)))
		rst = append(rst, byte((v & 0x0000ff00) >> 8))
		rst = append(rst, byte((v & 0x00ff0000) >> 16))
		rst = append(rst, byte((v & 0xff000000) >> 24))
	}
	return rst
}

func (self *Md5Hash)trans(src []byte) {

	if len(src) != 64{
		return
	}

	g := groupUint32(src)

	a:=self.a
	b:=self.b
	c:=self.c
	d:=self.d

	a = FF(a ,b ,c ,d ,g[0] ,7 ,0xd76aa478 )
	d = FF(d ,a ,b ,c ,g[1] ,12 ,0xe8c7b756 )
	c = FF(c ,d ,a ,b ,g[2] ,17 ,0x242070db )
	b = FF(b ,c ,d ,a ,g[3] ,22 ,0xc1bdceee )
	a = FF(a ,b ,c ,d ,g[4] ,7 ,0xf57c0faf )
	d = FF(d ,a ,b ,c ,g[5],12 ,0x4787c62a )
	c = FF(c ,d ,a ,b ,g[6],17 ,0xa8304613 )
	b = FF(b ,c ,d ,a ,g[7],22 ,0xfd469501)
	a = FF(a ,b ,c ,d ,g[8],7 ,0x698098d8 )
	d = FF(d ,a ,b ,c ,g[9],12 ,0x8b44f7af )
	c = FF(c ,d ,a ,b ,g[10],17 ,0xffff5bb1 )
	b = FF(b ,c ,d ,a ,g[11],22 ,0x895cd7be )
	a = FF(a ,b ,c ,d ,g[12],7 ,0x6b901122 )
	d = FF(d ,a ,b ,c ,g[13],12 ,0xfd987193 )
	c = FF(c ,d ,a ,b ,g[14],17 ,0xa679438e )
	b = FF(b ,c ,d ,a ,g[15] ,22 ,0x49b40821 )

	a = GG(a ,b ,c ,d ,g[1] ,5 ,0xf61e2562 )
	d = GG(d ,a ,b ,c ,g[6],9 ,0xc040b340 )
	c = GG(c ,d ,a ,b ,g[11],14 ,0x265e5a51 )
	b = GG(b ,c ,d ,a ,g[0],20 ,0xe9b6c7aa )
	a = GG(a ,b ,c ,d ,g[5],5 ,0xd62f105d )
	d = GG(d ,a ,b ,c ,g[10],9 ,0x02441453 )
	c = GG(c ,d ,a ,b ,g[15],14 ,0xd8a1e681 )
	b = GG(b ,c ,d ,a ,g[4],20 ,0xe7d3fbc8 )
	a = GG(a ,b ,c ,d ,g[9],5 ,0x21e1cde6 )
	d = GG(d ,a ,b ,c ,g[14],9 ,0xc33707d6 )
	c = GG(c ,d ,a ,b ,g[3],14 ,0xf4d50d87 )
	b = GG(b ,c ,d ,a ,g[8],20 ,0x455a14ed )
	a = GG(a ,b ,c ,d ,g[13],5 ,0xa9e3e905 )
	d = GG(d ,a ,b ,c ,g[2],9 ,0xfcefa3f8 )
	c = GG(c ,d ,a ,b ,g[7],14 ,0x676f02d9 )
	b = GG(b ,c ,d ,a ,g[12],20 ,0x8d2a4c8a )

	a = HH(a ,b ,c ,d ,g[5],4 ,0xfffa3942 )
	d = HH(d ,a ,b ,c ,g[8],11 ,0x8771f681 )
	c = HH(c ,d ,a ,b ,g[11],16 ,0x6d9d6122 )
	b = HH(b ,c ,d ,a ,g[14],23 ,0xfde5380c )
	a = HH(a ,b ,c ,d ,g[1],4 ,0xa4beea44 )
	d = HH(d ,a ,b ,c ,g[4],11 ,0x4bdecfa9 )
	c = HH(c ,d ,a ,b ,g[7],16 ,0xf6bb4b60 )
	b = HH(b ,c ,d ,a ,g[10],23 ,0xbebfbc70 )
	a = HH(a ,b ,c ,d ,g[13],4 ,0x289b7ec6 )
	d = HH(d ,a ,b ,c ,g[0],11 ,0xeaa127fa )
	c = HH(c ,d ,a ,b ,g[3],16 ,0xd4ef3085 )
	b = HH(b ,c ,d ,a ,g[6],23 ,0x04881d05 )
	a = HH(a ,b ,c ,d ,g[9],4 ,0xd9d4d039 )
	d = HH(d ,a ,b ,c ,g[12],11 ,0xe6db99e5 )
	c = HH(c ,d ,a ,b ,g[15],16 ,0x1fa27cf8 )
	b = HH(b ,c ,d ,a ,g[2],23 ,0xc4ac5665 )

	a = II(a ,b ,c ,d ,g[0],6 ,0xf4292244 )
	d = II(d ,a ,b ,c ,g[7],10 ,0x432aff97 )
	c = II(c ,d ,a ,b ,g[14],15 ,0xab9423a7 )
	b = II(b ,c ,d ,a ,g[5],21 ,0xfc93a039 )
	a = II(a ,b ,c ,d ,g[12],6 ,0x655b59c3 )
	d = II(d ,a ,b ,c ,g[3],10 ,0x8f0ccc92 )
	c = II(c ,d ,a ,b ,g[10],15 ,0xffeff47d )
	b = II(b ,c ,d ,a ,g[1],21 ,0x85845dd1 )
	a = II(a ,b ,c ,d ,g[8],6 ,0x6fa87e4f )
	d = II(d ,a ,b ,c ,g[15],10 ,0xfe2ce6e0 )
	c = II(c ,d ,a ,b ,g[6],15 ,0xa3014314 )
	b = II(b ,c ,d ,a ,g[13],21 ,0x4e0811a1 )
	a = II(a ,b ,c ,d ,g[4],6 ,0xf7537e82 )
	d = II(d ,a ,b ,c ,g[11],10 ,0xbd3af235 )
	c = II(c ,d ,a ,b ,g[2],15 ,0x2ad7d2bb )
	b = II(b ,c ,d ,a ,g[9],21 ,0xeb86d391 )

	self.a += a
	self.b += b
	self.c += c
	self.d += d
}

func (self *Md5Hash) Hash(src []byte) []byte {

	self.init()
	padedSrc := pad(src)
	//fmt.Printf("padedSrc: %x\n", padedSrc)
	gs := group(padedSrc)

	for index := 0; index < len(gs); index++ {
		self.trans(gs[index])
	}

	rst := transNumtoByte(self.a, self.b, self.c, self.d)
	return rst
}


