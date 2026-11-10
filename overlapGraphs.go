package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var ids []string
	ids = make([]string, 0)

	var seqs map[string]string
	seqs = make(map[string]string)

	var curId string
	var curSeq strings.Builder

	var line string
	var err error

	for {
		line, err = r.ReadString('\n')
		if err != nil && len(line) == 0 {
			break
		}
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			if line[0] == '>' {
				if curId != "" {
					seqs[curId] = curSeq.String()
				}
				var parts []string
				parts = strings.Split(line[1:], " ")
				curId = parts[0]
				ids = append(ids, curId)
				curSeq.Reset()
			} else {
				curSeq.WriteString(line)
			}
		}
	}

	if curId != "" {
		seqs[curId] = curSeq.String()
	}

	var k int
	k = 3

	var prefToIds map[string][]string
	prefToIds = make(map[string][]string)

	var i, j int
	for i = 0; i < len(ids); i++ {
		var id, seq string
		id = ids[i]
		seq = seqs[id]
		if len(seq) >= k {
			var pref string
			pref = seq[:k]
			prefToIds[pref] = append(prefToIds[pref], id)
		}
	}

	for i = 0; i < len(ids); i++ {
		var sId, sSeq string
		sId = ids[i]
		sSeq = seqs[sId]
		if len(sSeq) < k {
			continue
		}

		var suff string
		suff = sSeq[len(sSeq)-k:]

		var matches []string
		var ada bool
		matches, ada = prefToIds[suff]
		if ada {
			for j = 0; j < len(matches); j++ {
				var tId string
				tId = matches[j]
				if sId != tId {
					fmt.Fprintln(w, sId, tId)
				}
			}
		}
	}
}
