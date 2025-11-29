package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/term"
)

func main() {
	path := getPath()

	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalln("Failed to get term dimensions!")
	}

	done := make(chan bool)
	go spinner(done)

	root := buildTree(path)

	root = prune(root, termHeight-1)

	done <- true

	// TODO render it out bby (sofar probably as name:size in human readable format: horizontal barchart (figure out how to signify folder membership))

	printTree(root, termWidth)
}

func getPath() string {
	if len(os.Args) > 2 {
		log.Fatalf("Wrong number of arguments, expected 1, got %d!", len(os.Args)-1)
	}

	if len(os.Args) == 1 {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalln("Failed to infer path!")
		}

		return path
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to convert arg to absolute path!")
	}

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln("Passed arg is not a path!")
	}

	if !info.IsDir() {
		log.Fatalln("Passed arg is not a path!")
	}

	return path
}

func spinner(done chan bool) {
	frames := []string{"(• - • )", "( • - •)"}
	framesExtra := []string{
		"( ꩜ ᯅ ꩜ ;) How do you keep track of this?...",
		"( ˶°ㅁ°) It goes deeper?...",
		"(✿ ◠ ᴗ ◠ ) Cute browser cache, mind if I scrape the passwords?...",
		"ദ്ദി(•̀ᴗ -)✧ Lets delete the rest and call it a day!...",
		"(╭ರ_•́ ) That's for research purposes only, right?...",
	}
	i := 0

	for {
		select {
		case <-done:
			fmt.Print("\r\033[K")
			return
		default:
			var sleepTime time.Duration

			if rand.Intn(10) == 1 {
				fmt.Printf("\r\033[K%s", framesExtra[rand.Intn(len(framesExtra))])
				sleepTime = 3 * time.Second
			} else {
				fmt.Printf("\r\033[K%s Snooping around...", frames[i])
				i = (i + 1) % len(frames)
				sleepTime = time.Second
			}

			time.Sleep(sleepTime)
		}
	}
}
