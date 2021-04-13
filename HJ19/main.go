package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	lines := []string{}
	table := make(map[string]int)
	for in.Scan() {
		strs := strings.Split(in.Text(), " ")
		index := strings.LastIndex(strs[0], "\\")
		fileName := strs[0][index+1:]
		if len(fileName) > 16 {
			fileName = fileName[len(fileName)-16:]
		}
		if total, ok := table[fileName+" "+strs[1]]; ok {
			table[fileName+" "+strs[1]] = total + 1
		} else {
			table[fileName+" "+strs[1]] = 1
			lines = append(lines, fileName+" "+strs[1])
		}
	}
	for i := max(0, len(lines)-8); i < len(lines); i++ {
		fmt.Printf("%s %d\n", lines[i], table[lines[i]])
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
