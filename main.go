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
		text := leaveTextOrCoolifyIfPossible(scanner.Text())
		fmt.Println(text)
	}
}

func leaveTextOrCoolifyIfPossible(plainText string) string {
	if !wantToCoolify() {
		return plainText
	}

	return coolifyTextIfPossible(plainText)
}

func coolifyTextIfPossible(plainText string) string {
	if !canCoolify(plainText) {
		return plainText
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
				return duplicateVowel(plainText, i)
			}

			return removeVowel(plainText, i)
		}
	}
}

func canCoolify(plainText string) bool {
	for _, r := range plainText {
		if isVowel(r) {
			return true
		}
	}

	return false
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
