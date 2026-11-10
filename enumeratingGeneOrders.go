package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	var total, i int
	total = 1
	for i = 2; i <= n; i++ {
		total *= i
	}

	fmt.Fprintln(w, total)

	var p []int
	p = make([]int, n)
	var use []bool
	use = make([]bool, n+1)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			var j int
			for j = 0; j < n; j++ {
				w.WriteString(strconv.Itoa(p[j]))
				if j < n-1 {
					w.WriteByte(' ')
				}
			}
			w.WriteByte('\n')
			return
		}

		var j int
		for j = 1; j <= n; j++ {
			if !use[j] {
				use[j] = true
				p[idx] = j
				dfs(idx + 1)
				use[j] = false
			}
		}
	}

	dfs(0)
}
