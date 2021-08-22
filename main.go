package main

import (
"WordChain/Chain"
"WordChain/Dictionary"
"log"
)

func main() {

	dictionary, err := Dictionary.New("wordlist.txt")
	if err != nil {
		log.Fatalf("fatal error %e", err)
	}
	start := "ruby"
	end := "code"



	chain := Chain.New(dictionary, start, end)
	log.Println(chain.Solve())
}

