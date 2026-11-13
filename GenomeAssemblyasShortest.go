package main

import (
	"fmt"
	"os"
	"strings"
)

func parseFasta(text string) []string {
	var seqs []string
	seqs = make([]string, 0)
	var cur strings.Builder
	var lines []string
	lines = strings.Split(text, "\n")

	var i int
	for i = 0; i < len(lines); i++ {
		var line string
		line = strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		if line[0] == '>' {
			if cur.Len() > 0 {
				seqs = append(seqs, cur.String())
				cur.Reset()
			}
		} else {
			cur.WriteString(line)
		}
	}
	if cur.Len() > 0 {
		seqs = append(seqs, cur.String())
	}
	return seqs
}

func cariOverlap(s1 string, s2 string) int {
	var l1, l2, m, l int
	l1 = len(s1)
	l2 = len(s2)

	m = l1
	if l2 < m {
		m = l2
	}

	for l = m; l > 0; l-- {
		if strings.HasSuffix(s1, s2[:l]) {
			return l
		}
	}
	return 0
}

func shortestSuperstring(reads []string) string {
	var i, j int
	for len(reads) > 1 {
		var bestO, bi, bj int
		bestO = -1
		bi = 0
		bj = 1
		var merged string

		for i = 0; i < len(reads); i++ {
			for j = 0; j < len(reads); j++ {
				if i == j {
					continue
				}
				var o int
				o = cariOverlap(reads[i], reads[j])
				if o > bestO {
					bestO = o
					bi = i
					bj = j
					merged = reads[i] + reads[j][o:]
				}
			}
		}

		var nextReads []string
		nextReads = make([]string, 0, len(reads)-1)
		for i = 0; i < len(reads); i++ {
			if i != bi && i != bj {
				nextReads = append(nextReads, reads[i])
			}
		}
		nextReads = append(nextReads, merged)
		reads = nextReads
	}
	return reads[0]
}

func main() {
	var data []byte
	var err error
	data, err = os.ReadFile("input.txt")
	if err != nil {
		return
	}

	var reads []string
	reads = parseFasta(string(data))

	var hasil string
	hasil = shortestSuperstring(reads)

	os.WriteFile("output.txt", []byte(hasil), 0644)

	fmt.Println("Selesai! Jawaban telah disimpan di file output.txt")
}
