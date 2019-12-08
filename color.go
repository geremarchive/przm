package przm

import (
	"fmt"
)

func GetHex(r int, g int, b int) (hex string) {
	sr := fmt.Sprintf("%x", r)
	sg := fmt.Sprintf("%x", g)
	sb := fmt.Sprintf("%x", b)

	hex = "#"

	if len(sr) == 1 {
		hex += fmt.Sprintf("0%x", r)
	} else {
		hex += sr
	}

	if len(sg) == 1 {
		hex += fmt.Sprintf("0%x", g)
	} else {
		hex += sg
	}

	if len(sb) == 1 {
		hex += fmt.Sprintf("0%x", b)
	} else {
		hex += sb
	}

	return
}

func IncVal(v int, inc int) (out int) {
	if v+inc < 255 {
		out = v + inc
	} else {
		out = 255
	}

	return
}

func DecVal(v int, inc int) (out int) {
	if v-inc >= 0 {
		out = v - inc
	} else {
		out = 0
	}

	return
}
