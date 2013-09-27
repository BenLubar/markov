package main

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	flagCompile = flag.Bool("c", false, "read stdin as lines of text and output a preprocessed markov file")
	flagCount   = flag.Uint("n", 10, "number of lines to produce (incompatible with -c)")
)

func main() {
	flag.Parse()

	if *flagCompile {
		compile()
	} else {
		output()
	}
}

func compile() {
	out := Reduce(emitCh)

	var mapWg sync.WaitGroup

	startMap := func(text string) {
		mapWg.Add(1)
		go func() {
			Map(text)
			mapWg.Done()
		}()
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		startMap(in.Text())
	}

	if err := in.Err(); err != nil {
		panic(err)
	}

	mapWg.Wait()
	close(emitCh)

	if err := gob.NewEncoder(os.Stdout).Encode(<-out); err != nil {
		panic(err)
	}
}

func output() {
	var corpus MarkovCorpus
	if err := gob.NewDecoder(os.Stdin).Decode(&corpus); err != nil {
		panic(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := func(prefix [MarkovLength - 1]string) string {
		m := corpus[prefix]
		var n int
		for _, i := range m {
			n += int(i)
		}
		n = r.Intn(n)
		for s, i := range m {
			n -= int(i)
			if n < 0 {
				return s
			}
		}
		panic("unreachable")
	}

	for i := *flagCount; i > 0; i-- {
		var chain []interface{}

		var prefix [MarkovLength - 1]string
		suffix := random(prefix)

		for suffix != "" {
			chain = append(chain, suffix)
			copy(prefix[:], prefix[1:])
			prefix[len(prefix)-1] = suffix
			suffix = random(prefix)
		}

		fmt.Println(chain...)
	}
}
