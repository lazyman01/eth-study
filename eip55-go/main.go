package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

const AddressLength = 20

type Address [AddressLength]byte

func (a Address) hex() []byte {
	var buf [len(a)*2 + 2]byte
	copy(buf[:2], "0x")
	hex.Encode(buf[2:], a[:])
	return buf[:]
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

func FromHex(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}
func (a *Address) checksumHex() []byte {
	buf := a.hex()
	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[2:])
	hash := sha.Sum(nil)
	fmt.Println("hash: ",hash)
	for i := 2; i < len(buf); i++ {
		//hash里每个字节取两次
		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			//取四位比特
			hashByte = hashByte >> 4
		} else {
			//只取i*4位比特
			hashByte &= 0xf
		}
		//如果是字母，且第i*4位比特是1（即值为8）
		if buf[i] > '9' && hashByte > 7 {
			//小写转大写
			buf[i] -= 32
		}
	}
	return buf[:]
}

func main()  {
	var addr Address
	addr = BytesToAddress(FromHex("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(addr)
	fmt.Println("checkhex:",addr.checksumHex())
	fmt.Println(string(addr.checksumHex()))
}
