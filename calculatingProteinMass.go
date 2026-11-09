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

	var s string
	fmt.Fscan(r, &s)

	var mp [256]float64
	mp['A'] = 71.03711
	mp['C'] = 103.00919
	mp['D'] = 115.02694
	mp['E'] = 129.04259
	mp['F'] = 147.06841
	mp['G'] = 57.02146
	mp['H'] = 137.05891
	mp['I'] = 113.08406
	mp['K'] = 128.09496
	mp['L'] = 113.08406
	mp['M'] = 131.04049
	mp['N'] = 114.04293
	mp['P'] = 97.05276
	mp['Q'] = 128.05858
	mp['R'] = 156.10111
	mp['S'] = 87.03203
	mp['T'] = 101.04768
	mp['V'] = 99.06841
	mp['W'] = 186.07931
	mp['Y'] = 163.06333

	var i, n int
	var jum float64
	n = len(s)

	for i = 0; i < n; i++ {
		jum += mp[s[i]]
	}

	fmt.Fprintf(w, "%.3f\n", jum)
}
