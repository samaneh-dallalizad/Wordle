package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type ModalInfo struct {
	ElementId string
	Title     string
	Content   string
	IsActive  string
}

func (m *ModalInfo) Modal() g.Node {
	return Div(Class(fmt.Sprintf("modal %s", m.IsActive)), ID(m.ElementId),

		Div(Class("modal-background")),
		Div(Class("modal-card"),
			ModalCardHead(m.Title),
			ModalCardBody(m.Content),
		),
	)
}

func ModalCardHead(title string) g.Node {
	return Header(Class("modal-card-head"),
		P(Class("modal-card-title"), g.Text(title)),
		Button(Class("delete")),
	)
}

func ModalCardBody(content string) g.Node {
	return Section(Class("modal-card-body"), g.Text(content))
}
