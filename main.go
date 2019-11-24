package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	filePath := "./moby.txt"

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	words := 0
	inWord := false

	b := bufio.NewReader(f)
	for {
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Failed to read byte from file", err)
		}

		if unicode.IsSpace(r) && inWord {
			words++
			inWord = false
		}

		inWord = unicode.IsLetter(r)
	}

	log.Printf("Word count %d", words)
}
