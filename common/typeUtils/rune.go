package typeUtils

import "unicode/utf8"

func UTF8RunesToBytes(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}
	bs := make([]byte, size)
	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}
	return bs
}

func UTF8BytesToRunes(bs []byte) []rune {
	size := utf8.RuneCount(bs)
	rs := make([]rune, size)
	count := 0
	for i := 0; i < len(bs); i++ {
		charRune, charLen := utf8.DecodeRune(bs[count:])
		count += charLen
		rs[i] = charRune
	}
	return rs
}
