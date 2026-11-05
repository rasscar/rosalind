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

	fmt.Fprintln(w, "=== PROGRAM INTRODUCTION TO MENDELIAN INHERITANCE ===")

	var k, m, n, tot int
	fmt.Fprint(w, "Masukkan nilai k m n (pisahkan dengan spasi): ")
	w.Flush()
	
	_, err := fmt.Fscan(r, &k, &m, &n)
	if err != nil {
		fmt.Fprintln(w, "[ERROR] Masukkan harus berupa 3 angka integer yang dipisahkan spasi!")
		return
	}

	tot = k + m + n

	var tp, p1, p2, p3, res, ans float64
	tp = float64(tot) * float64(tot-1)

	if tp == 0 {
		fmt.Fprintln(w, "[ERROR] Total populasi tidak boleh 0!")
		return
	}

	p1 = (float64(m) * float64(m-1)) / tp * 0.25
	p2 = (float64(m) * float64(n) * 2.0) / tp * 0.5
	p3 = (float64(n) * float64(n-1)) / tp * 1.0

	res = p1 + p2 + p3
	ans = 1.0 - res

	fmt.Fprintln(w, "\n--- HASIL ---")
	fmt.Fprintf(w, "Peluang fenotipe dominan: %.5f\n", ans)
}
