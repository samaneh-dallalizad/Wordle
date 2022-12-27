package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func GameArea() g.Node {
	var count [6]int
	return Div(
		Ul(
			g.Group(g.Map(count[:], func(i int) g.Node {
				return Li(Class("colums  is-multiline"),
					g.Group(g.Map(count[0:5], func(i int) g.Node {
						return Div(Class("column is-5"),
							Input(Type("text"), MaxLength("10")),
						)
					})),
				)
			})),
		),
		Button(g.Text("save")),
	)
}
