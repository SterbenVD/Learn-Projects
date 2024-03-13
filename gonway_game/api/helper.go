package api

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func string_To_Uint8(str string) uint8 {
	var val uint8
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			val = val*16 + uint8(str[i]-'0')
		} else if str[i] >= 'A' && str[i] <= 'F' {
			val = val*16 + uint8(str[i]-'A'+10)
		} else if str[i] >= 'a' && str[i] <= 'f' {
			val = val*16 + uint8(str[i]-'a'+10)
		}
	}
	return val
}

func String_To_Color(str string) rl.Color {
	var r, g, b uint8
	r = string_To_Uint8(str[0:2])
	g = string_To_Uint8(str[2:4])
	b = string_To_Uint8(str[4:6])
	return rl.NewColor(r, g, b, 255)
}
