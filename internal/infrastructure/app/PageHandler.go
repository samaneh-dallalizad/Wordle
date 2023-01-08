package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"unicode"
	"wordleGame/internal/infrastructure/domain/wordlesite"
	"wordleGame/internal/ui/views"

	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) HomePageHandler() func(*gin.Context) {

	return func(c *gin.Context) {
		//get any error from query string
		e := c.DefaultQuery("error", "")
		//get current state from cookie
		gameCookie, err := c.Cookie("game")
		//if we do not have game cookie(we need to start a new game)/only happen in the first time or user click on start new game
		if err != nil {
			startGameWord, startErr := wordlesite.StartGame()
			if startErr != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			//set word in game table
			gameTable := wordlesite.Game{
				Word:      *startGameWord,
				ActiveRow: 0,
			}
			//json encode
			game, jsonErr := json.Marshal(&gameTable)
			if jsonErr != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			//store game table in cookie
			gameCookie = string(game)
			c.SetCookie("game", gameCookie, 3600, "/", "localhost", false, true)
		}
		//if gamecookie is available we just show it
		gameTable := wordlesite.Game{}
		//json decode
		err = json.Unmarshal([]byte(gameCookie), &gameTable)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
		_ = views.HomeView(&gameTable, e).Render(c.Writer)
	}
}

func (s *ApplicationServer) SubmitGuessHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		// get cookie
		gameCookie, err := c.Cookie("game")
		//if cookie is not present
		if err != nil {
			//create error query string
			v := url.Values{
				"error": {"Game cookie is not present"},
			}
			c.Redirect(http.StatusSeeOther, "/?"+v.Encode())
			return
		}
		//cookie is present
		gameTable := wordlesite.Game{}
		err = json.Unmarshal([]byte(gameCookie), &gameTable)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		//get inputs
		input, ok := c.GetPostFormArray("char")
		if !ok || len(input) != 5 {
			v := url.Values{
				"error": {"Invalid Input"},
			}
			c.Redirect(http.StatusSeeOther, "/?"+v.Encode())
			return
		}

		//validate inputs just be letter
		for i := 0; i < len(input); i++ {
			if len(input[i]) != 1 || !unicode.IsLetter(rune(input[i][0])) {
				v := url.Values{
					"error": {"Only Letters are accepted"},
				}
				c.Redirect(http.StatusSeeOther, "/?"+v.Encode())
			}
		}

		word := strings.ToLower(strings.Join(input, ""))
		result, err := gameTable.Word.Guess(word)
		if err != nil {
			if errors.Is(err, wordlesite.NotAWord) {
				v := url.Values{
					"error": {"Input it not a valid word"},
				}
				c.Redirect(http.StatusSeeOther, "/?"+v.Encode())
				return
			} else {
				c.Status(http.StatusInternalServerError)
				return
			}
		}

		gameTable.Grid[gameTable.ActiveRow] = *result
		gameTable.ActiveRow = gameTable.ActiveRow + 1

		game, jsonErr := json.Marshal(gameTable)
		if jsonErr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.SetCookie("game", string(game), 3600, "/", "localhost", false, true)

		c.Redirect(http.StatusSeeOther, "/")

	}
}

func (s *ApplicationServer) StartNewGame() func(*gin.Context) {
	return func(c *gin.Context) {
		//remove cookie
		c.SetCookie("game", "", -1, "/", "localhost", false, true)
		//redirect to homepage
		c.Redirect(http.StatusSeeOther, "/")

	}
}
