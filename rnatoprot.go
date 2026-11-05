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

	fmt.Fprintln(w, "=== PROGRAM TRANSLATING RNA INTO PROTEIN ===")

	var s string
	fmt.Fprint(w, "Masukkan string mRNA: ")
	w.Flush()
	fmt.Fscan(r, &s)

	var mp map[string]byte
	mp = map[string]byte{
		"UUU": 'F', "CUU": 'L', "AUU": 'I', "GUU": 'V',
		"UUC": 'F', "CUC": 'L', "AUC": 'I', "GUC": 'V',
		"UUA": 'L', "CUA": 'L', "AUA": 'I', "GUA": 'V',
		"UUG": 'L', "CUG": 'L', "AUG": 'M', "GUG": 'V',
		"UCU": 'S', "CCU": 'P', "ACU": 'T', "GCU": 'A',
		"UCC": 'S', "CCC": 'P', "ACC": 'T', "GCC": 'A',
		"UCA": 'S', "CCA": 'P', "ACA": 'T', "GCA": 'A',
		"UCG": 'S', "CCG": 'P', "ACG": 'T', "GCG": 'A',
		"UAU": 'Y', "CAU": 'H', "AAU": 'N', "GAU": 'D',
		"UAC": 'Y', "CAC": 'H', "AAC": 'N', "GAC": 'D',
		"UAA": 0, "CAA": 'Q', "AAA": 'K', "GAA": 'E',
		"UAG": 0, "CAG": 'Q', "AAG": 'K', "GAG": 'E',
		"UGU": 'C', "CGU": 'R', "AGU": 'S', "GGU": 'G',
		"UGC": 'C', "CGC": 'R', "AGC": 'S', "GGC": 'G',
		"UGA": 0, "CGA": 'R', "AGA": 'R', "GGA": 'G',
		"UGG": 'W', "CGG": 'R', "AGG": 'R', "GGG": 'G',
	}

	var i, panjang int
	panjang = len(s)

	var hasil []byte
	hasil = make([]byte, 0)

	for i = 0; i < panjang; i += 3 {
		if i+3 <= panjang {
			var kodon string
			kodon = s[i : i+3]

			var asam byte
			var ada bool
			asam, ada = mp[kodon]

			if ada {
				if asam == 0 {
					break
				}
				hasil = append(hasil, asam)
			}
		}
	}

	fmt.Fprintln(w, "\n--- HASIL ---")
	fmt.Fprintf(w, "Protein String: %s\n", string(hasil))
}
