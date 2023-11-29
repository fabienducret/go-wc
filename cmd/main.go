package main

import (
	"fmt"
	"log"
	"os"
	"wc/wc"
)

func main() {
	fileNames := os.Args[1:]
	var tcc, twc, tlc int

	for _, fName := range fileNames {
		file, err := os.Open(fName)
		if err != nil {
			log.Printf("Error on open file: %s", err)
			continue
		}

		c := wc.CountsFor(file)
		tcc += c.Characters
		twc += c.Words
		tlc += c.Lines

		display(c.Lines, c.Words, c.Characters, fName)
	}
	display(tlc, twc, tcc, "total")
}

func display(lc, wc, cc int, name string) {
	fmt.Printf("%d %2d %2d %s\n", lc, wc, cc, name)
}
