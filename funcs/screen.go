package funcs

import (
	"fmt"
	"strconv"
	"github.com/geremachek/escape"
)

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

