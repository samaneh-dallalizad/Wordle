package wordlesite

import (
	"errors"
)

type Word struct {
	Id     int
	Key    string
	WordID int
}

type Guess struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Guess string `json:"guess"`
}

type Block struct {
	Letter string
	State  int
}

type GuessResult [5]Block

type Game struct {
	Word      Word
	ActiveRow int
	Grid      [6]GuessResult
}

var NotAWord = errors.New("Not a word")
