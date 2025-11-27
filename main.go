package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// TODO extract path from the first arg or infer cwd if no args
	path := getPath()

	// TODO start cli animation (ansi escape chars (kaomoji might be fun to animate)) while the search run
	done := make(chan bool)
	go spinner(done)

	time.Sleep(50 * time.Second) // so I can visualize the cli anim during dev

	// TODO build a tree representation 3 field (name, size, childern), dirs should init nil size

	// TODO reverse populate size on dirs as sum of children

	// TODO get terminal dimentions, figure out line count

	// TODO figure out top x relevant items to display to match line count

	done <- true
	// TODO render it out bby (sofar probably as name:size in human readable format: horizontal barchart (figure out how to signify folder membership))

	fmt.Println(path)
}

func getPath() string {
	if len(os.Args) > 2 {
		log.Fatalln("Lol. Lmao. Pass a single arg at most you moron!")
	}

	if len(os.Args) == 1 {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalln("Where the fuck have you called this from? There is no cwd on here!")
		}

		return path
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalln("Maybe pass in a path, yeah?")
	}

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln("Maybe make sure the path is valid before bothering me...")
	}

	if !info.IsDir() {
		log.Fatalln("Dude... that's a file, not a dir. What do you expect me to do?")
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
			if rand.Int()%10 == 1 {
				fmt.Printf("\r\033[K%s", framesExtra[rand.Int()%len(framesExtra)])
				time.Sleep(2 * time.Second)
			} else {
				fmt.Printf("\r\033[K%s Snooping around...", frames[i])
				i = (i + 1) % len(frames)
			}
			time.Sleep(1 * time.Second)
		}
	}
}
