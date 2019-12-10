package main

import (
	"fmt"
	"strings"
	"strconv"
	"flag"
	"os"
	fu "przm/funcs"
)

const help = `Usage: przm [OPTION] [COLOR]
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
		ch rune
		olen int

		r, g, b int

		inc int = 1
	)

	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(help)
			os.Exit(0)
		}
	}

	printRGB := flag.Bool("r", false, "Return the color in the RGB format")
	printHex := flag.Bool("x", false, "Return the color in the hexadecimal format")
	printOutput := flag.Bool("o", false, "Don't clean up the output")
	ColorForeground := flag.Bool("f", false, "Color the text foreground")
	ColorBackground := flag.Bool("b", false, "Color the text background")

	flag.Parse()

	fu.HideCursor()

	if len(flag.Args()) == 1 {
		// hsl, rgb, etc later.
		r, g, b = fu.GetRGB(flag.Args()[0])
	}

	for {
		if *ColorForeground {
			olen = fu.PrintInfo("fore", r, g, b)
		} else if *ColorBackground {
			olen = fu.PrintInfo("back", r, g, b)
		} else {
			olen = fu.PrintInfo("normal", r, g, b)
		}

		ch = fu.Getru()

		if ch == 'q' {
			if *printOutput {
			} else {
				fmt.Print("\r" + strings.Repeat(" ", olen) + "\r")
				if *printRGB {
					fmt.Print("rgb(" + strconv.Itoa(r) + ", " +
						strconv.Itoa(g) + ", " + strconv.Itoa(b) + ")")
				} else if *printHex {
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
