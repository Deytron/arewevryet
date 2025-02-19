package utils

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenPassPhrase(nb int, separator string) string {
	file, err := os.Open("/app/assets/wordlist-fr.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	caser := cases.Title(language.Und)
	for scanner.Scan() {
		words = append(words, caser.String(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var passPhrase []string
	for i := 0; i < nb; i++ {
		passPhrase = append(passPhrase, words[rand.Intn(len(words))])
	}

	// Add random number at the end (0-10)
	return strings.Join(passPhrase, separator) + strconv.Itoa(rand.Intn(10))
}
