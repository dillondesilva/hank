package render

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func RenderText(text string) {
	out, _ := glamour.Render(text, "dark")
	fmt.Print(out)
}
