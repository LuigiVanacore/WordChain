package Chain

import (
	"WordChain/Dictionary"
	"log"
	"strings"
)

// Chain struct that contains list
type Chain struct {
	words     []string
	startWord string
	lastWord  string
	result    *[]string
}

func New(dictionary *Dictionary.Dictionary, startWord, lastWord string) *Chain {
	if len(startWord) != len(lastWord) {
		log.Fatal("start and end words must have the same size")
	}
	if !dictionary.Contain(startWord) {
		log.Fatal("startWord not present in the Dictionary")
	}
	if !dictionary.Contain(lastWord) {
		log.Fatal("lastWord not present in the Dictionary")
	}
	chain := &Chain{startWord: startWord, lastWord: lastWord}
	chain.filteringWords(dictionary)
	var slice []string
	chain.result = &slice
	return chain
}

func (c *Chain) filteringWords(dictionary *Dictionary.Dictionary) {
	c.words = []string{}
	l := len(c.startWord)
	for _, w := range dictionary.GetWords() {
		if len(w) == l {
			c.words = append(c.words, w)
		}
	}
}

type node struct {
	parents []node
	word    string
}

func (c *Chain) Solve() []string {
	toVisit := []node{
		node{
			word:    c.startWord,
			parents: nil,
		},
	}
	var visited []string
	var neighbourhoods []string

	for i := 0; i < len(toVisit); i++ {
		if strings.Compare(toVisit[i].word, c.lastWord) == 0 {
			visit(toVisit[i].parents, c.result)
			break
		}
		getNeighborhoods(toVisit[i].word, &c.words, &neighbourhoods)
		for _, n := range neighbourhoods {
			if !contains(n, &visited) {
				visited = append(visited, n)
				no := node{word: n, parents: []node{toVisit[i]}}
				toVisit = append(toVisit, no)
			}
		}
		neighbourhoods = nil
	}
	*c.result = append([]string{c.lastWord}, *c.result...)
	return *c.result
}

func visit(chain []node, res *[]string) {
	if chain == nil {
		return
	}
	*res = append(*res, chain[0].word)
	visit(chain[0].parents, res)
}

func contains(word string, words *[]string) bool {
	for _, w := range *words {
		if strings.Compare(word, w) == 0 {
			return true
		}
	}
	return false
}

// getNeighborhoods return  the  word
// with 1 character of difference from the given word
func getNeighborhoods(word string, chain, neighborhood *[]string) {
	for _, w := range *chain {
		if isNeighborhood(w, word) {
			*neighborhood = append(*neighborhood, w)
		}
	}
}

func isNeighborhood(w, word string) bool {
	var diff int
	for i := 0; i < len(w); i++ {
		if w[i] != word[i] {
			diff++
		}
	}
	return diff == 1
}
