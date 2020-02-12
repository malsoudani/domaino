package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "site",
	otherWord + "time",
	otherWord + "app",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
	otherWord + "hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

// basic word sprinkler ...
// 	you give it a word and it adds a basic transform picked randomly and prints it out, it keeps running and keeps accepting words and printing sprinkled goodness... at least just for v1
