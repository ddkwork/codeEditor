package main

import (
	"codeEditor"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("codeEditor", func(w *unison.Window) {
		codeEditor.New().Layout(w.Content())
	})
}
