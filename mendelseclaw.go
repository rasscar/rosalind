package main

import (
	"bufio"
	"fmt"
	"os"
)

func pangkat(b float64, e int) float64 {
	var r float64
	r = 1.0
	var i int
	for i = 0; i < e; i++ {
		r *= b
	}
	return r
}

func kombinasi(n int, k int) float64 {
	if k < 0 || k > n {
		return 0.0
	}
	if k == 0 || k == n {
		return 1.0
	}
	if k > n-k {
		k = n - k
	}
	var r float64
	r = 1.0
	var i int
	for i = 1; i <= k; i++ {
		r = r * float64(n-k+i) / float64(i)
	}
	return r
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var k, N int
	fmt.Fscan(r, &k, &N)

	var n int
	n = 1 << k

	var p, q, ans float64
	p = 0.25
	q = 0.75
	ans = 0.0

	var i int
	for i = N; i <= n; i++ {
		ans += kombinasi(n, i) * pangkat(p, i) * pangkat(q, n-i)
	}

	fmt.Fprintf(w, "%.10f\n", ans)
}
