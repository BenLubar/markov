#!/bin/bash

go build

# the second sed expression is from http://stackoverflow.com/a/5751555/2664560
curl http://www.gutenberg.org/cache/epub/10/pg10.txt | sed -e 's/\r//g' | sed -n -e '1{${p;b};h;b};/^$/!{H;$!b};x;s/\(.\)\n/\1 /g;p' | grep '^[0-9]\+:[0-9]\+' | sed -e 's/^[0-9]\+:[0-9]\+ //' | ./markov -c > bible.markov

# the second sed expression is from http://stackoverflow.com/a/5751555/2664560
curl http://www.gutenberg.org/cache/epub/11/pg11.txt | sed -e 's/\r//g' | sed -n -e '1{${p;b};h;b};/^$/!{H;$!b};x;s/\(.\)\n/\1 /g;p' | tail -n +35 | head -n +1668 | grep . | ./markov -c > alice.markov
