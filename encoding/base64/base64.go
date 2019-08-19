package base64

//base64 编码
//将字节流表示成为可读字符串
//2^6 = 64 即 6bit可以表示64种字符，3个字节24bit 因此可以转换为4个字符。 3bytes -> 4char
//字节数 % 3, 差几个字节补几个'='
//在编码过程中将3个字节转为一个uint32, uint32(src[0])<<16 | uint32(src[1])<<8 | uint32(src[2])
//之后 << >> & 获取对应6bit 查找转换表获得字符即可。
//解码过程，将4个字符转换为3字节即可。
//注意：解码过程中参考标准库，对于存放对应关系表，发现使用 array 比 map 性能上快太多，因此这里使用 array.

type Base64Coding struct {
}

var tableSlice = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
//var tableMap map[byte]uint
var tableReverseSlice [256]uint8

func init() {
	for k, v := range tableSlice {
		tableReverseSlice[uint8(v)] = uint8(k)
	}
}

func (self *Base64Coding) createSlice(srcLen int) []byte{
	dlen := (srcLen+2)/3*4
	rst := make([]byte, dlen)
	return rst
}

func (self *Base64Coding) createOriginSlice(src []byte) ([]byte, int){
	var rstLen int
	srcLen := len(src)
	count := self.checkEqualCount(src)
	switch count{
	case 0:
		rstLen = srcLen/4*3
	case 1:
		rstLen = srcLen/4*3-1
	case 2:
		rstLen = srcLen/4*3-2
	}
	rst := make([]byte, rstLen)
	return rst, count
}

func (self *Base64Coding) checkEqualCount(src []byte) int{

	srcLen := len(src)
	if src[srcLen-1] == '=' {
		if  src[srcLen-2] == '='{
			return 2
		}
		return 1
	}
	return 0
}




func (self *Base64Coding) Encode(src []byte) []byte {


	si, di := 0, 0
	n := len(src)/3*3
	rst := self.createSlice(len(src))

	for si < n{
		value := uint(src[si]) << 16 | uint(src[si+1]) << 8 | uint(src[si+2])
		rst[di] = tableSlice[((value >> 18) & 0x3F)]
		rst[di+1] = tableSlice[((value >> 12) & 0x3F)]
		rst[di+2] = tableSlice[((value >> 6) & 0x3F)]
		rst[di+3] = tableSlice[((value) & 0x3F)]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0{
		return rst
	}

	value := uint(src[si]) << 16
	if remain == 2{
		value |= uint(src[si+1]) << 8
	}

	rst[di] = tableSlice[((value >> 18) & 0x3F)]
	rst[di+1] = tableSlice[((value >> 12) & 0x3F)]

	switch remain{
	case 2:
		rst[di+2] = tableSlice[((value >> 6) & 0x3F)]
		rst[di+3] = '='
	case 1:
		rst[di+2] = '='
		rst[di+3] = '='
	}

	return rst
}

func (self *Base64Coding) Decode(src []byte) []byte {

	n := len(src)
	si := 0
	di := 0
	rst, count := self.createOriginSlice(src)

	for si < n-4{

		value := uint(tableReverseSlice[src[si]]&0x3F) << 18 | uint(tableReverseSlice[src[si+1]]&0x3F) << 12 | uint(tableReverseSlice[src[si+2]]&0x3F) << 6 | uint(tableReverseSlice[src[si+3]]&0x3F)
		rst[di] = byte(value >> 16)
		rst[di+1] = byte(value >> 8)
		rst[di+2] = byte(value)

		si += 4
		di += 3
	}

	value := uint(tableReverseSlice[src[si]]&0x3F) << 18
	value = value | uint(tableReverseSlice[src[si+1]]&0x3F) << 12

	switch count{
	case 0:
		value = value | uint(tableReverseSlice[src[si+2]]&0x3F) << 6
		value = value | uint(tableReverseSlice[src[si+3]]&0x3F)
		rst[di] = byte(value >> 16)
		rst[di+1] = byte(value >> 8)
		rst[di+2] = byte(value)
	case 1:
		value = value | uint(tableReverseSlice[src[si+2]]&0x3F) << 6
		rst[di] = byte(value >> 16)
		rst[di+1] = byte(value >> 8)
	case 2:
		rst[di] = byte(value >> 16)
	}

	return rst
}



