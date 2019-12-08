package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	fu "./funcs"
	"github.com/geremachek/escape"
)

const help = `Usage: przm [OPTION]
A simple, yet feature rich color picker and manipulator

--help, -h: Display this information
-r: Return the color in the RGB format
-x: Return the color in the hexadecimal format
-o: Don't clean up the output
-f: Color the text foreground
-b: Color the text background

h: Increment the 'R' value
j: Increment the 'G' value
k: Increment the 'B' value
l: Increment all values (brightens the color)
b: Decrement the 'R' value
n: Decrement the 'G' value
m: Decrement the 'B' value
,: Decrement all values (dims the color)
[space]: Sets the color to black

q: Exit the program
[number]: Set the increment to [number] (0 = 10)`

func main() {
	var (
		args []string = os.Args[1:]
		optR bool
		optH bool
		optO bool
		optF bool
		optB bool

		ch rune
		olen int

		r int
		g int
		b int

		inc int = 1
	)

	if len(args) > 0 {
		if args[0] == "-h" || args[0] == "--help" {
			fmt.Println(help)
			os.Exit(0)
		} else if rune(string(args[0])[0]) == '-' {
			for _, elem := range string(args[0])[1:] {
				if elem == 'r' && !(optR) && !(optH) && !(optO) {
					optR = true
				} else if elem == 'x' && !(optR) && !(optH) && !(optO) {
					optH = true
				} else if elem == 'o' && !(optR) && !(optH) && !(optO) {
					optO = true
				} else if elem == 'f' && !(optF) && !(optB) {
					optF = true
				} else if elem == 'b' && !(optF) && !(optB) {
					optB = true
				} else {
					fmt.Println(escape.Vint(31, 1) + "Error: invalid options" + escape.Vint(0))
					os.Exit(0)
				}
			}
		}
	}

	fu.HideCursor()

	for {
		if optF {
			olen = fu.PrintInfo("fore", r, g, b)
		} else if optB {
			olen = fu.PrintInfo("back", r, g, b)
		} else {
			olen = fu.PrintInfo("normal", r, g, b)
		}

		ch = fu.Getru()

		if ch == 'q' {
			if optO {
			} else {
				fmt.Print("\r" + strings.Repeat(" ", olen) + "\r")
				if optR {
					fmt.Print("rgb(" + strconv.Itoa(r) + ", " +
						strconv.Itoa(g) + ", " + strconv.Itoa(b) + ")")
				} else if optH {
					fmt.Print(fu.GetHex(r, g, b))
				} else {
					fmt.Print("\033[1A")
				}
			}

			break
		} else if ch == 'h' {
			r = fu.IncVal(r, inc)
		} else if ch == 'j' {
			g = fu.IncVal(g, inc)
		} else if ch ==  'k' {
			b = fu.IncVal(b, inc)
		} else if ch == 'l' {
			r = fu.IncVal(r, inc)
			g = fu.IncVal(g, inc)
			b = fu.IncVal(b, inc)
		} else if ch == 'b' {
			r = fu.DecVal(r, inc)
		} else if ch == 'n' {
			g = fu.DecVal(g, inc)
		} else if ch == 'm' {
			b = fu.DecVal(b, inc)
		} else if ch == ',' {
			r = fu.DecVal(r, inc)
			g = fu.DecVal(g, inc)
			b = fu.DecVal(b, inc)
		} else if ch == ' ' {
			r, g, b = 0, 0, 0
		} else if ch > 47 && ch < 58 {
			if ch == '0' {
				inc = 10
			} else {
				conv, _ := strconv.Atoi(string(ch))
				inc = conv
			}
		}

		fmt.Print("\r" + strings.Repeat(" ", olen) + "\r")
	}

	fmt.Println()
	fu.ShowCursor()
}
