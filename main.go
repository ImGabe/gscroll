package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	text, prefix string
	length       int
	speed        int64
	marquee      bool
)

func init() {
	flag.IntVar(&length, "length", 20, "length of characters to scroll.")
	flag.Int64Var(&speed, "speed", 5, "Speed of text scrolling.")
	flag.StringVar(&text, "text", "", "Text to scroll.")
	flag.BoolVar(&marquee, "marquee", true, "Text scroll left.")
}

func debugPrint(s []rune) {
	escape := "\r%s"

	if !marquee {
		escape = "%s\n"
	}

	fmt.Printf(escape, string(s))
}

func main() {
	flag.Parse()

	if text == "" {
		log.Fatal("Missing text input.")
	}

	runes := []rune(text)
	rulesLen := len(runes)

	sleepDuration := (1 * time.Second) - (time.Duration(speed) * 100 * time.Millisecond)

	if length > rulesLen {
		empty := []rune(strings.Repeat(" ", length-rulesLen))
		runes = append(runes, empty...)
		rulesLen = len(runes)
	}

	for {
		split := rulesLen - length

		for i := 0; i < split; i++ {
			s := runes[i : i+length]
			debugPrint(s)

			time.Sleep(sleepDuration)
		}

		for i := 0; i < length; i++ {
			a := runes[0:i]
			b := runes[split+i:]

			b = append(b, a...)
			debugPrint(b)

			time.Sleep(sleepDuration)
		}
	}
}
