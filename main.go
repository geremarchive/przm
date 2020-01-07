package main

import (
	"fmt"
	"strings"
	"strconv"
	"math/rand"
	"time"
	flag "github.com/spf13/pflag"
	"os"
	fu "przm/funcs"
)

const help = `Usage: przm [OPTION] [COLOR]
A simple, yet feature rich color picker and manipulator

┏━┓┏━┓╺━┓┏┳┓
┣━┛┣┳┛┏━┛┃┃┃
╹  ╹┗╸┗━╸╹ ╹

--help, -h: Display this information
--rgb, -r: Return the color in the RGB format
--hex, -x: Return the color in the hexadecimal format
--output, -o: Don't clean up the output
--foreground, -f: Color the text foreground
--background, -b: Color the text background

h: Increment the 'R' value
j: Increment the 'G' value
k: Increment the 'B' value
l: Increment all values (brightens the color)

b: Decrement the 'R' value
n: Decrement the 'G' value
m: Decrement the 'B' value
,: Decrement all values (dims the color)

H: Randomly set the 'R' value
J: Randomly set the 'G' value
K: Randomly set the 'B' value
L: Randomly set all values 

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

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(help)
			os.Exit(0)
		}
	}

	printRGB := flag.BoolP("rgb", "r", false, "Return the color in the RGB format")
	printHex := flag.BoolP("hex", "x", false, "Return the color in the hexadecimal format")
	printOutput := flag.BoolP("output", "o", false, "Don't clean up the output")
	ColorForeground := flag.BoolP("foreground", "f", false, "Color the text foreground")
	ColorBackground := flag.BoolP("background", "b", false, "Color the text background")

	flag.Parse()

	fu.HideCursor()

	args := flag.Args()

	rand.Seed(time.Now().UnixNano())

	if len(args) == 1 {
		// hsl, rgb, etc later.
		if rune(args[0][0]) == '#' {
			r, g, b = fu.GetRGB(args[0][1:])
		} else {
			r, g, b = fu.GetRGB(args[0])
		}
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
		} else if ch == 'H' {
			r = rand.Intn(256)
		} else if ch == 'J' {
			g = rand.Intn(256)
		} else if ch == 'K' {
			b = rand.Intn(256)
		} else if ch == 'L' {
			r = rand.Intn(256)
			g = rand.Intn(256)
			b = rand.Intn(256)
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
