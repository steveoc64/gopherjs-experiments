package main

import (
	"honnef.co/go/js/dom"
	"strconv"
)

func inputKeyUp(event dom.Event) {
	input := event.Target().(*dom.HTMLInputElement)

	span := dom.GetWindow().Document().GetElementByID("inputvalue")
	span.SetInnerHTML(input.Value)
}

func main() {
	w := dom.GetWindow()
	d := w.Document()
	h := d.GetElementByID("foo")
	k := d.GetElementByID("keycode")
	i := d.GetElementByID("input").(*dom.HTMLInputElement)
	vv := d.GetElementByID("vv").(*dom.HTMLDivElement)

	print(vv.Dataset()["what"])
	h.AddEventListener("click", false, func(event dom.Event) {
		event.PreventDefault()
		h.SetInnerHTML("I am Clicked")
		println("This message is printed in browser console")
	})

	w.AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		k.SetInnerHTML(strconv.Itoa(ke.KeyCode))
	})

	i.Focus()
	i.AddEventListener("keyup", false, inputKeyUp)

}
