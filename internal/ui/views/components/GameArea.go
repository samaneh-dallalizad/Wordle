package components

import (
	"fmt"
	"wordleGame/internal/infrastructure/domain/wordlesite"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

var classForResponse = []string{"has-background-dark", "has-background-warning", "has-background-success"}

func GameArea(d *wordlesite.Game) g.Node {
	var row = -1
	var activeRow = d.ActiveRow
	return FormEl(Method("post"), Action("/guess_result"), Div(
		Ul(Class("my-5"),
			g.Group(g.Map(d.Grid[:], func(r wordlesite.GuessResult) g.Node {
				row += 1
				return Li(Class(fmt.Sprintf("columns is-multiline row row-%d", row)),
					g.Group(g.Map(r[:], func(b wordlesite.Block) g.Node {
						return Div(Class("column is-1 row-block"),

							g.If(activeRow == row, Input(Class("input is-large has-text-centered is-uppercase"), Type("text"), MaxLength("1"), Name("char"))),
							g.If(activeRow > row, Input(Class(fmt.Sprintf("input is-large has-text-centered is-uppercase %s", classForResponse[b.State])), Type("text"), MaxLength("1"), Disabled(), Value(b.Letter))),
							g.If(activeRow < row, Input(Class("input is-large has-text-centered is-uppercase"), Type("text"), MaxLength("1"), Disabled())),
						)
					})),
				)
			})),
		),

		Button(Class("button is-primary"), g.Text("submit"), Type("submit")),
		A(Class("button is-danger ml-2"), g.Text("start new game"), Href("/new_game")),
	),
	)
}
