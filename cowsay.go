package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func buildCloud(lines []string, width int) string {
	var ret []string
	borders := []string{"/", "\\", "|", "<", ">"}
	count := len(lines)

	top := " " + strings.Repeat("_", width+2)
	bottom := " " + strings.Repeat("-", width+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[3], lines[0], borders[4])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s := fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[2])
			ret = append(ret, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[1], lines[i], borders[0])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, c := range lines {
		c = strings.Replace(c, "\t", "    ", -1)
		ret = append(ret, c)
	}
	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}

func normalizeString(lines []string, width int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", width-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		
`
	var rabbit = `		\
		 \
		  (\(\ 
		  ( -.-)
		  o_(")(")
`

	switch name {
	case "cow":
		fmt.Println(cow)
	case "rabbit":
		fmt.Println(rabbit)
	default:
		fmt.Println("We don't got it like that")
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []string

	var figure string
	flag.StringVar(&figure, "f", "cow", "valid values: `cow` and `rabbit`")
	flag.Parse()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo \"Hello, world!\" | cowsay")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, string(input))
	}

	lines := tabsToSpaces(output)
	maxWidth := calculateMaxWidth(lines)
	normalizedString := normalizeString(lines, maxWidth)
	cloud := buildCloud(normalizedString, maxWidth)
	fmt.Println(cloud)
	printFigure(figure)
}
