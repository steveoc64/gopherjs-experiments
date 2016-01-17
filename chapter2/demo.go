package main

import (
	"honnef.co/go/js/dom"
	"strconv"
)

func main() {
	w := dom.GetWindow()
	d := w.Document()
	h := d.GetElementByID("foo")
	k := d.GetElementByID("keycode")

	h.AddEventListener("click", false, func(event dom.Event) {
		event.PreventDefault()
		h.SetInnerHTML("I am Clicked")
		println("This message is printed in browser console")
	})

	w.AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		k.SetInnerHTML(strconv.Itoa(ke.KeyCode))
	})
}
