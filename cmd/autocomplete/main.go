package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cnnrznn/ds/trie"
)

// main reads from stdin the tags for each order note
// it computes a count for each unique tag
// it sorts this list of counts by count greates to smallest
// it inserts each tag into the trie, appending it's count to each prefix node
// All nodes now have a sorted list of tags they prefix from greatest to smallest
func main() {
	counts := map[string]int{}
	t := trie.New()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		ls := strings.Split(s, ",")
		for _, tag := range ls {
			key := strings.TrimSpace(tag)
			counts[key]++
		}
	}

	datas := []Data{}
	for k, v := range counts {
		datas = append(datas, Data{
			Key:   k,
			Count: v,
		})
	}
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Count > datas[j].Count
	})

	for _, data := range datas {
		t.Insert(data.Key, func(node *trie.Node) {
			if node.Data == nil {
				node.Data = []Data{}
			}
			if len(node.Data.([]Data)) > 5 {
				return
			}
			node.Data = append(node.Data.([]Data), data)
		})
	}

	fmt.Println(getPrefixes(t, "r"))
}

func getPrefixes(t *trie.Trie, prefix string) []string {
	result := []string{}
	node := t.Find(prefix)
	for _, tag := range node.Data.([]Data) {
		result = append(result, tag.Key)
	}
	return result
}

type Data struct {
	Key   string
	Count int
}

type ByCount []Data

func (b ByCount) Less(i, j int) bool {
	return b[i].Count < b[j].Count
}
