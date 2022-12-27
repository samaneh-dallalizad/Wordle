package views

import (
	g "github.com/maragudk/gomponents"
	//nolint:stylecheck,golint
	"wordleGame/internal/ui/views/components"

	. "github.com/maragudk/gomponents/html"
)

func HomeView(n string) g.Node {
	componentsView := Doctype(
		HTML(
			Head(
				Meta(Charset("utf-8")),
				Link(
					Href("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"),
					Rel("stylesheet"),
				),
				Script(
					Type("text/javascript"),
					Src("/scripts/delete-notification.js"),
				),
			),
			Lang("en"),

			Body(

				Div(Class("container"),
					H1(g.Text(n)),
					components.HowPlayGame(),
					components.GameArea(),
				),
			),
		),
	)

	return componentsView
}
