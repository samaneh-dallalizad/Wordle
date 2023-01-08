package views

import (
	g "github.com/maragudk/gomponents"

	"wordleGame/internal/infrastructure/domain/wordlesite"
	"wordleGame/internal/ui/views/components"

	. "github.com/maragudk/gomponents/html"
)

func HomeView(w *wordlesite.Game, e string) g.Node {
	modalInfo := components.ModalInfo{
		IsActive:  "",
		ElementId: "mod",
		Content:   e,
		Title:     "Error",
	}
	if e != "" {
		modalInfo.IsActive = "is-active"
	}

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

				Script(
					Type("text/javascript"),
					Src("/scripts/modal.js"),
				),
			),
			Lang("en"),

			Body(

				Div(Class("container"),
					components.HowPlayGame(),
					components.GameArea(w),
					modalInfo.Modal(),
				),
			),
		),
	)

	return componentsView
}
