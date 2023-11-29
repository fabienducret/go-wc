package wc

import (
	"bufio"
	"io"
	"strings"
)

type Counts struct {
	Lines      int
	Words      int
	Characters int
}

func CountsFor(reader io.Reader) Counts {
	s := bufio.NewScanner(reader)
	c := Counts{}

	for s.Scan() {
		line := s.Text()
		c.calculateFor(line)
	}

	return c
}

func (c *Counts) calculateFor(line string) {
	c.Lines++
	c.Words += wordsCountFor(line)
	c.Characters += len(line)
}

func wordsCountFor(line string) int {
	return len(strings.Fields(line))
}
