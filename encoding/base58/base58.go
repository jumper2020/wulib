package base58

import (
	"bytes"
	"math/big"
	"strings"
)


//首先需要明白这里的 base58　是用于 bitcoin 中的，因此输入是 256bit 的一个大整数，　输出是特定字符集中字符组成的串。
//编码之前是一个十六进制数(使用字符串表示), 因此一定是偶数个字符 (sha256 hash结果为32个字节，可以表示为十六进制超大整数
//编码之前大小写无所谓，因为无论大小写，表示出的整数是一致的
//编码之后大小写需要区分开，因为编码之后为字符串，大小写表示不同字符
//这里使用 math/big 来表示大整数, 该库中与[]byte相互转换时是会忽略最开始的'0'的，也就是说比如在big.Int中存放　0x0123，　
//Text 函数输出为　"123", 因此这里就需要补充一个 '0', 具体参见 Decode 函数
//而github中有些golang　base58的编码还是与bitcoin中的base58不完全相同，　比如bitcoin中需要根据0x00数量来补充'1'
//网上的编解码工具也需要找用于bitcoin中的base58
//编解码测试　参考　http://lenschulwitz.com/base58

type Base58Coding struct {
}

var tableSlice = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
//var tableMap map[byte]uint
var tableReverseSlice [256]uint8

func init() {
	for k, v := range tableSlice {
		tableReverseSlice[uint8(v)] = uint8(k)
	}
}


func getZeroCount(src []byte) int{
	var count int
	for index:=0; index<len(src); index+=2{
		if src[index] == '0' && src[index+1] == '0'{
			count++
		}else{
			break
		}
	}
	return count
}

func getOneCount(src []byte) int{
	var count int
	for index:=0; index<len(src); index++{
		if src[index] == '1'{
			count++
		}else{
			break
		}
	}
	return count
}

func reverse(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}



func (self *Base58Coding) Encode(src []byte) []byte{

	//1. count the 0x00 at the beginning of text.
	//2. divide to 58 and order the remainder.
	//3. add '1' at the beginning of result of step 2.

	//note: src should be a hexadecimal num. so it's length % 2 == 0
	if src == nil || len(src) == 0 || len(src) %2 ==1{
		return nil
	}

	zero := getZeroCount(src)

	var value *big.Int
	var ok bool
	if value,ok = new(big.Int).SetString(string(src), 16); !ok{
		return nil
	}

	result := make([]byte, 0)
	var remainder big.Int
	b := big.NewInt(58)

	for{
		if value.Cmp(big.NewInt(0)) == 0{
			break
		}
		value.DivMod(value, b, &remainder)
		result = append(result, byte(tableSlice[remainder.Uint64()]))
	}

	one := strings.Repeat("1", zero)
	result = append(result, []byte(one)...)
	reverse(result)

	return result
}

func (self *Base58Coding) Decode(src []byte) []byte{


	count := getOneCount(src)
	value := src[count:]
	div := big.NewInt(0)
	b := big.NewInt(58)

	for index:=0; index<len(value); index++{
		remainder := int64(tableReverseSlice[uint8(value[index])])
		div.Add(big.NewInt(remainder), div.Mul(div,b))
	}


	//div 转换成的16进制，是不保证为双数位数的，比如　"123", 这里需要转换为　"0123"
	zero := bytes.Repeat([]byte{'0'}, 2*count)
	if len(div.Text(16))%2 == 1{
		zero= append(zero, '0')
	}
	result := div.Append(zero, 16)

	return result
}


