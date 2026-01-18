package unicode

import (
	"unicode/utf8"
)

const RunesCount = utf8.MaxRune + 1

const (
	N0, N9   uint8 = 48, 57
	AUp, ZUp uint8 = 65, 90
	ALo, ZLo uint8 = 97, 122
)
