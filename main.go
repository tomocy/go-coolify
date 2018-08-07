package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := coolifyTextIfPossible(scanner.Text())
		fmt.Println(text)
	}
}

func coolifyTextIfPossible(plainText string) string {
	if !canCoolify(plainText) {
		return plainText
	}

	coolifiedText, _ := coolifyText(plainText)
	return coolifiedText
}

func canCoolify(plainText string) bool {
	for _, r := range plainText {
		if isVowel(r) {
			return true
		}
	}

	return false
}

func coolifyText(plainText string) (string, error) {
	if !canCoolify(plainText) {
		return "", fmt.Errorf("cannot coolify %s because it has no vowels\n", plainText)
	}

	for {
		for i, r := range plainText {
			if !isVowel(r) {
				continue
			}
			if !wantToCoolify() {
				continue
			}

			if wantToDuplicateVowelThanRemove() {
				return duplicateVowel(plainText, i), nil
			}

			return removeVowel(plainText, i), nil
		}
	}
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aiueoAIUEO", r)
}

func wantToCoolify() bool {
	return rand.Intn(2)+1 <= 1
}

func wantToDuplicateVowelThanRemove() bool {
	return rand.Intn(2) <= 1
}

func duplicateVowel(text string, i int) string {
	b := []byte(text)
	b = append(b[:i+1], b[i:]...)
	return string(b)
}

func removeVowel(text string, i int) string {
	b := []byte(text)
	b = append(b[:i], b[i+1:]...)
	return string(b)
}
