package main

import (
	"codeEditor"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("codeEditor", func(w *unison.Window) {
		w.Content().AddChild(codeEditor.New().Layout())
	})
}
