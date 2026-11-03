package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hitungGc(dna string) float64 {
	var panjang, jum, i int
	panjang = len(dna)
	if panjang == 0 {
		return 0.0
	}
	jum = 0
	for i = 0; i < panjang; i++ {
		if dna[i] == 'C' || dna[i] == 'G' {
			jum++
		}
	}
	return float64(jum) / float64(panjang)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	scanner := bufio.NewScanner(r)

	var idTerbaik, idSkr string
	var gcTerbaik, skr float64
	gcTerbaik = -1.0

	var dnaSkr strings.Builder
	var baris string

	for scanner.Scan() {
		baris = strings.TrimSpace(scanner.Text())
		if len(baris) == 0 {
			continue
		}

		if baris[0] == '>' {
			if idSkr != "" {
				skr = hitungGc(dnaSkr.String())
				if skr > gcTerbaik {
					gcTerbaik = skr
					idTerbaik = idSkr
				}
			}
			idSkr = baris[1:]
			dnaSkr.Reset()
		} else {
			dnaSkr.WriteString(baris)
		}
	}

	if idSkr != "" {
		skr = hitungGc(dnaSkr.String())
		if skr > gcTerbaik {
			gcTerbaik = skr
			idTerbaik = idSkr
		}
	}

	fmt.Fprintln(w, idTerbaik)
	fmt.Fprintf(w, "%.6f\n", gcTerbaik*100)
}
