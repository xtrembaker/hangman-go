package hangman

import (
	"fmt"
	"strings"
)

const (
	GoodGuess    = "GOOD_GUESS"
	AlreadyGuess = "ALREADY_GUESS"
	BadGuess     = "BAD_GUESS"
	Lost         = "Lost"
	Won          = "Won"
)

type Game struct {
	State        string   // Game state
	Letters      []string // Letters in the word to find
	FoundLetters []string // Good guesses letters
	UsedLetters  []string // Used letters
	TurnsLeft    int      // Remaining Attempts
}

func New(turns int, word string) (*Game, error) {
	if len(word) == 0 {
		return nil, fmt.Errorf("word '%s' must be at least 1 character", word)
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	return &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = AlreadyGuess
	} else if letterInWord(guess, g.Letters) {
		g.State = GoodGuess
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = Won
		}
	} else {
		g.State = BadGuess
		g.LoseTurn(guess)
		fmt.Println(g.TurnsLeft)
		if g.TurnsLeft <= 0 {
			g.State = Lost
		}
	}
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LoseTurn(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft--
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}
