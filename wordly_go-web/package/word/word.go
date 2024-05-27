package word

import (
	"mygame/package/constants"
	"mygame/package/letter"
)

// Word структура для слова
type Word struct {
	Letters []*letter.Letter
}

// Equals проверяет, совпадает ли слово с другим словом
func (w *Word) Equals(other string) bool {
	for i, letter := range w.Letters {
		if letter.Char != []rune(other)[i] {
			return false
		}
	}
	return true
}

// ChangeColor изменяет цвет буквы в слове по индексу
func (w *Word) ChangeColor(index int, color string) {
	w.Letters[index].Color = color
}

// NewWord создает новое слово
func NewWord(word string) *Word {
	w := &Word{}
	for _, char := range word {
		w.Letters = append(w.Letters, letter.NewLetter(char, constants.Gray))
	}
	return w
}
