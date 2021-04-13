package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		table := make(map[int]int)
		n, _ := strconv.Atoi(input.Text())
		for i := 0; i < n; i++ {
			input.Scan()
			line := strings.Split(input.Text(), " ")
			index, _ := strconv.Atoi(line[0])
			value, _ := strconv.Atoi(line[1])
			if old_value, ok := table[index]; ok {
				table[index] = old_value + value
			} else {
				table[index] = value
			}
		}
		keys := []int{}
		for k := range table {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Printf("%d %d\n", k, table[k])
		}
	}
}
