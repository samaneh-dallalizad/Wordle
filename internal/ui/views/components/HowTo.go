package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HowPlayGame() g.Node {
	return Div(Class("notification is-link"),

		Button(Class("delete")),
		P(g.Text("Guess the WORDLE in "), Strong(g.Text("6")), g.Text(" tries.")),
		P(g.Text("Each guess must be a valid 5 letter word. Hit the enter button to submit.")),
		P(g.Text("After each guess, the color of the tiles will change to show how close your guess was to the word.")),
	)
}
