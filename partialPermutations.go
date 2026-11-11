package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r *bufio.Reader = bufio.NewReader(os.Stdin)
	var w *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)

	var mod int
	mod = 1000000

	var hasil, i int
	hasil = 1

	for i = 0; i < k; i++ {
		hasil = (hasil * (n - i)) % mod
	}

	fmt.Fprintln(w, hasil)
}
