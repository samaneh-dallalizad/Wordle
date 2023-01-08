package errorPages

import (
	g "github.com/maragudk/gomponents"
	// . "github.com/maragudk/gomponents/html"
)

func Error_404() g.Node {
	return g.Text("Not Found")
}
