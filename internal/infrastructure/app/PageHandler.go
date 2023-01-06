package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"
	"wordleGame/internal/infrastructure/domain/wordlesite"
	"wordleGame/internal/ui/views"

	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) HomePageHandler() func(*gin.Context) {

	return func(c *gin.Context) {
		gameCookie, err := c.Cookie("game")
		if err != nil {
			word, startErr := wordlesite.StartGame()
			if startErr != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			grid := &wordlesite.Grid{
				Word: *word,
			}
			game, jsonErr := json.Marshal(grid)
			if jsonErr != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			gameCookie = string(game)
			c.SetCookie("game", gameCookie, 3600, "/", "localhost", false, true)
		}
		grid := wordlesite.Grid{}
		err = json.Unmarshal([]byte(gameCookie), &grid)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
		_ = views.HomeView(&grid).Render(c.Writer)
	}
}

func (s *ApplicationServer) SubmitGuessHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		gameCookie, err := c.Cookie("game")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		grid := wordlesite.Grid{}
		err = json.Unmarshal([]byte(gameCookie), &grid)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		input, ok := c.GetPostFormArray("char")
		if !ok || len(input) != 5 {
			c.Status(http.StatusBadRequest)
			return
		}
		for i := 0; i < len(input); i++ {
			if len(input[i]) != 1 || !unicode.IsLetter(rune(input[i][0])) {
				c.Status(http.StatusBadRequest)
				return
			}
		}
		activeRow := 0
		for i := 0; i < len(grid.State); i++ {
			if grid.State[i][0].Letter != "" {
				activeRow += 1
			}
		}
		word := strings.ToLower(strings.Join(input, ""))
		result, err := grid.Word.Guess(word)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		grid.State[activeRow] = *result

		game, jsonErr := json.Marshal(grid)
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
		c.SetCookie("game", "", -1, "/", "localhost", false, true)

		c.Redirect(http.StatusSeeOther, "/")

	}
}
