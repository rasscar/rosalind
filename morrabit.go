package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)

	var kel []*big.Int
	kel = make([]*big.Int, m)

	var i, j, k int
	for i = 0; i < m; i++ {
		kel[i] = big.NewInt(0)
	}
	kel[0] = big.NewInt(1)

	var maju, temp *big.Int
	maju = big.NewInt(0)
	temp = big.NewInt(0)

	for i = 2; i <= n; i++ {
		for j = 0; j < m; j++ {
			if j == 0 {
				maju.SetInt64(0)
				for k = 1; k < m; k++ {
					maju.Add(maju, kel[k])
				}
				temp.Set(kel[0])
				kel[0].Set(maju)
			} else {
				maju.Set(kel[j])
				kel[j].Set(temp)
				temp.Set(maju)
			}
		}
	}

	var total *big.Int
	total = big.NewInt(0)
	for j = 0; j < m; j++ {
		total.Add(total, kel[j])
	}

	fmt.Fprintln(w, total.String())
}
