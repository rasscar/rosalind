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

	var protein string
	fmt.Fscan(r, &protein)

	var mod int
	mod = 1000000

	var mp [256]int
	mp['A'] = 4
	mp['C'] = 2
	mp['D'] = 2
	mp['E'] = 2
	mp['F'] = 2
	mp['G'] = 4
	mp['H'] = 2
	mp['I'] = 3
	mp['K'] = 2
	mp['L'] = 6
	mp['M'] = 1
	mp['N'] = 2
	mp['P'] = 4
	mp['Q'] = 2
	mp['R'] = 6
	mp['S'] = 6
	mp['T'] = 4
	mp['V'] = 4
	mp['W'] = 1
	mp['Y'] = 2

	var hasil, i, n int
	hasil = 1
	n = len(protein)

	for i = 0; i < n; i++ {
		hasil = (hasil * mp[protein[i]]) % mod
	}

	hasil = (hasil * 3) % mod

	fmt.Fprintln(w, hasil)
}
