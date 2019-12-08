package przm

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func HideCursor() {
	fmt.Print("[?25l")
}

func ShowCursor() {
	fmt.Print("[?121[?25h")
}
