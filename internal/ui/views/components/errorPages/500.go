package errorPages

import (
	g "github.com/maragudk/gomponents"
	// . "github.com/maragudk/gomponents/html"
)

func Error_500() g.Node {
	return g.Text("Internal Server error")
}
