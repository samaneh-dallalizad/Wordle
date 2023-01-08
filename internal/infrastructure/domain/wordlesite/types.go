package wordlesite

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

type Grid struct {
	Word  Word
	State [6]GuessResult
}

var NotAWord = errors.New("Not a word")

func (w *Word) Guess(guess string) (*GuessResult, error) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}

	input := Guess{
		Id:    w.Id,
		Key:   w.Key,
		Guess: guess,
	}
	inputJson, err := json.Marshal(input)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	fmt.Printf("Input %s", inputJson)
	req, err := http.NewRequest("POST", "https://word.digitalnook.net/api/v1/guess/", bytes.NewReader(inputJson))

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	req.Header.Add("Content-Type", `application/json`)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	fmt.Printf("Body : %s", body)

	if resp.StatusCode == 400 {
		fmt.Printf("Input is not a word: %s", guess)
		return nil, NotAWord
	}
	ret := GuessResult{}
	jsonErr := json.Unmarshal(body, &ret)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil, jsonErr
	}

	return &ret, nil

}

func StartGame() (*Word, error) {

	c := http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := http.NewRequest("POST", "https://word.digitalnook.net/api/v1/start_game/", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	fmt.Printf("Body : %s", body)

	ret := Word{}
	jsonErr := json.Unmarshal(body, &ret)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil, jsonErr
	}

	return &ret, nil
}
