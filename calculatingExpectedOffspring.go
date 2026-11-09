package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var ar [6]int
	var p [6]float64
	var i int
	var hasil float64

	for i = 0; i < 6; i++ {
		fmt.Fscan(r, &ar[i])
	}

	p[0] = 1.0
	p[1] = 1.0
	p[2] = 1.0
	p[3] = 0.75
	p[4] = 0.5
	p[5] = 0.0

	hasil = 0
	for i = 0; i < 6; i++ {
		hasil += float64(ar[i]) * p[i]
	}

	hasil *= 2.0

	fmt.Fprintln(w, hasil)
}
