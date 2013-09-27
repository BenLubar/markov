package main

import (
	"strings"
)

const MarkovLength = 2

func Map(text string) {
	var chain [MarkovLength]string
	words := strings.Fields(text)

	for _, word := range words {
		copy(chain[:], chain[1:])
		chain[len(chain)-1] = word
		Emit(chain)
	}

	copy(chain[:], chain[1:])
	chain[len(chain)-1] = ""
	Emit(chain)
}

var emitCh = make(chan [MarkovLength]string)

func Emit(chain [MarkovLength]string) {
	// copy so we don't hold the entire corpus
	for i, s := range chain {
		chain[i] = string([]byte(s))
	}

	emitCh <- chain
}
