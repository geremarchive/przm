package funcs

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
	"github.com/geremachek/escape"
	"golang.org/x/crypto/ssh/terminal"
)

func Getru() rune {
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	in := bufio.NewReader(os.Stdin)
	r, _, err := in.ReadRune()

	if err != nil {
		log.Println("STDIN:", err)
	}

	return r
}

func HideCursor() {
	fmt.Print("[?25l")
}

func ShowCursor() {
	fmt.Print("[?121[?25h")
}

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

func GetRGB(hex string) (r int, g int, b int) {
	iSlice := []int{}
	var iconv int64

	for i := 0; i < len(hex); i += 2 {
		iconv, _ = strconv.ParseInt(string(hex[i:i+2]), 16, 64)
		iSlice = append(iSlice, int(iconv))
	}

	return iSlice[0], iSlice[1], iSlice[2]
}

func PrintInfo(col string, r int, g int, b int) int {
	var beg string

	if col == "normal" {
		beg = escape.Vint(38, 2, r, g, b) + "\u2588\u2588" + escape.Vint(0) + " "
	} else if col == "fore" {
		beg = escape.Vint(38, 2, r, g, b)
	} else if col == "back" {
		beg = escape.Vint(38, 2, r, g, b, 7)
	}

	output := beg + "rgb(" + strconv.Itoa(r) + ", " + strconv.Itoa(g) + ", " + strconv.Itoa(b) + ") " +
		GetHex(r, g, b) + escape.Vint(0)
	fmt.Print(output)

	return len(output)
}

func IncVal(v int, inc int) (out int) {
	if v + inc < 255 {
		out = v + inc
	} else {
		out = 255
	}

	return
}

func DecVal(v int, inc int) (out int) {
	if v - inc >= 0 {
		out = v - inc
	} else {
		out = 0
	}

	return
}
