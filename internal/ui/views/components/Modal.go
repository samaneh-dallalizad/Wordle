package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

var BkgColorBtn = []string{"is-primary", "is-link", "is-info", "is-success", "is-warning", "is-danger"}

type ModalInfo struct {
	ElementId string
	Title     string
	Content   string
}

func (m *ModalInfo) Modal() g.Node {
	return Div(Class("modal"), ID(m.ElementId),

		Div(Class("modal-background")),
		Div(Class("modal-card"),
			ModalCardHead(m.Title),
			ModalCardBody(m.Content),
			ModalCardFoot("is-primary", "save"),
		),
	)
}

func ModalCardHead(title string) g.Node {
	return Header(Class("modal-card-head"),
		P(Class("modal-card-title"), g.Text(title)),
		Button(Class("delete")),
	)
}
func ModalCardFoot(bkgbtn string, titlebtn string) g.Node {
	return Footer(Class("modal-card-head"),
		Button(Class(fmt.Sprintf("button%s", bkgbtn)), g.Text(titlebtn)),
		Button(Class("button"), g.Text("cancel")),
	)
}

func ModalCardBody(content string) g.Node {
	return Section(Class("modal-card-body"), g.Text(content))
}
