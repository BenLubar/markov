package main

type MarkovCorpus map[[MarkovLength - 1]string]map[string]uint

func Reduce(in <-chan [MarkovLength]string) <-chan MarkovCorpus {
	out := make(chan MarkovCorpus)

	go func() {
		corpus := make(MarkovCorpus)

		var prefix [MarkovLength - 1]string
		for chain := range in {
			copy(prefix[:], chain[:])
			suffix := chain[len(chain)-1]

			if m, ok := corpus[prefix]; ok {
				m[suffix]++
			} else {
				corpus[prefix] = map[string]uint{suffix: 1}
			}
		}

		out <- corpus
	}()

	return out
}
