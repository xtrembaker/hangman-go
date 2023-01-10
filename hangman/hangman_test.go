package hangman

import (
	"fmt"
	"hangman/src/github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	ass := assert.New(t)

	_, err := New(3, "")
	ass.NotNil(err)
	ass.Equal(err, fmt.Errorf("word '' must be at least 1 character"))
}

func TestMakeAGuessChangeStateToAlreadyGuessWhenUserTypeAlreadyTypedChar(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("a")
	g.MakeAGuess("a")

	assert.Equal(t, g.State, AlreadyGuess)
}

func TestMakeAGuessChangeStateToGoodGuessWhenUserTypeLetterThatBelongsToWord(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")

	assert.Equal(t, g.State, GoodGuess)
}
