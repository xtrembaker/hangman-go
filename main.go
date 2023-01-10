package main

import (
	"fmt"
	"hangman/dictionnary"
	"hangman/hangman"
	"os"
)

func main() {
	err := dictionnary.LoadFromList("a-list.csv")
	if err != nil {
		fmt.Printf("Could not load dicitonnary: %v\n", err)
		os.Exit(1)
	}
	g := hangman.New(8, dictionnary.PickWord())

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case hangman.Won, hangman.Lost:
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		fmt.Println(l)
		g.MakeAGuess(guess)
	}

}
