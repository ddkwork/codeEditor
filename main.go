package main

import (
	"codeEditor"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("codeEditor", func(w *unison.Window) {
		w.Content().AddChild(codeEditor.New().Layout())
	})
}
