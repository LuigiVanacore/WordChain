package Dictionary

import (
	"bufio"
	"os"
)

type Dictionary struct {
	words []string
}

// Load dictionary from path
func New(path string) (*Dictionary, error) {
	d := Dictionary{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		d.words = append(d.words,  text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &d, nil
}


func (d *Dictionary) GetWords() []string {
	return d.words
}

func (d *Dictionary) Contain(word string) bool {
	for _, w :=  range d.words {
		if w == word {
			return true
		}
	}
	return false
}