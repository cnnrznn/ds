package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cnnrznn/ds/trie"
)

const (
	listSize = 5
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
		fmt.Println(k, v)
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
			if len(node.Data.([]Data)) > listSize {
				return
			}
			node.Data = append(node.Data.([]Data), data)
		})
	}

	fmt.Println(getPrefixes(t, ""))
	fmt.Println(getPrefixes(t, "r"))
	fmt.Println(getPrefixes(t, "c"))

	newTag := Data{
		Key:   "created a new tag",
		Count: 10000,
	}

	t.Insert("created a new tag never before seen", func(node *trie.Node) {
		node.Lock()
		defer node.Unlock()

		if node.Data == nil {
			node.Data = []Data{}
		}

		datas := node.Data.([]Data)
		insertAfter := -1
		for i := len(datas) - 1; i >= 0; i-- {
			if datas[i].Count > newTag.Count {
				insertAfter = i
				break
			}
			if datas[i].Key == newTag.Key {
				datas[i].Count = newTag.Count
				bubbleUp(datas, i)
				node.Data = datas
				return
			}
		}

		dst := make([]Data, len(datas)+1)
		copy(dst, datas[:insertAfter+1])
		copy(dst[insertAfter+2:], datas[insertAfter+1:])
		dst[insertAfter+1] = newTag
		datas = dst

		if len(datas) > listSize {
			datas = datas[:listSize]
		}

		node.Data = datas
	})

	fmt.Println(getPrefixes(t, "c"))
	fmt.Println(getPrefixes(t, "ca"))
}

func bubbleUp(datas []Data, index int) {
	for ; index > 0 && datas[index].Count > datas[index-1].Count; index-- {
		datas[index], datas[index-1] = datas[index-1], datas[index]
	}
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
