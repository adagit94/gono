package unicode

import (
	"unicode/utf8"
)

const RunesCount = utf8.MaxRune + 1

const (
	ASCIIFirst, ASCIILast rune = 33, 126
	N0, N9                rune = 48, 57
	AUp, ZUp              rune = 65, 90
	ALo, ZLo              rune = 97, 122
)
