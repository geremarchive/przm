<h1 align="center">przm 🎨</h1>

<p align="center">A simple, yet feature rich color picker and manipulator (w/ vi keybindings!)</p>
<br><br>
<p align="center"><img src="scrot.png"></p>

**Usage:**

```
Usage: przm [OPTION] [COLOR]
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
[number]: Set the increment to [number] (0 = 10)
```

**Dependencies**

```
go get github.com/geremachek/escape 
```

```
go get golang.org/x/crypto/ssh/terminal
```

```
go get github.com/spf13/pflag
```
