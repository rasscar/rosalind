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

	fmt.Fprintln(w, "=== PROGRAM HITUNG HAMMING DISTANCE ===")

	var s, t string

	fmt.Fprint(w, "Masukkan string DNA pertama (s): ")
	w.Flush()
	fmt.Fscan(r, &s)

	fmt.Fprint(w, "Masukkan string DNA kedua   (t): ")
	w.Flush()
	fmt.Fscan(r, &t)

	var panjangS, panjangT int
	panjangS = len(s)
	panjangT = len(t)

	if panjangS != panjangT {
		fmt.Fprintln(w, "\n[ERROR] Panjang kedua string DNA tidak sama!")
		fmt.Fprintf(w, "Panjang string s: %d karakter\n", panjangS)
		fmt.Fprintf(w, "Panjang string t: %d karakter\n", panjangT)
		return
	}

	var beda, i int
	beda = 0

	for i = 0; i < panjangS; i++ {
		if s[i] != t[i] {
			beda++
		}
	}

	fmt.Fprintln(w, "\n--- HASIL ---")
	fmt.Fprintf(w, "Hamming Distance: %d\n", beda)
}
