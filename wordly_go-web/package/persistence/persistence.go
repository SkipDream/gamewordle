package persistence

import (
	"bufio"
	"math/rand"
	"mygame/package/constants"
	"os"
	"time"
)

// IsWordValid проверяет, существует ли слово в словаре.
func IsWordValid(word string) bool {
	file, err := os.Open("bd.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true
		}
	}
	return false
}

// GetRandomWord возвращает случайное слово из словаря.
func GetRandomWord() string {
	file, err := os.Open("bd.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(constants.WordQuantity)

	scanner := bufio.NewScanner(file)
	for i := 0; i <= randomIndex; i++ {
		scanner.Scan()
		if i == randomIndex {
			return scanner.Text()
		}
	}
	return ""
}
