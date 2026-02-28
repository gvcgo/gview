package main

import (
	"fmt"

	"github.com/gvcgo/gview"
	"github.com/gdamore/tcell/v3"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content cview.Primitive) {
	textView := cview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/gvcgo/gview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
