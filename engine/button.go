package engine

type button struct {
	*image
	onClickHandler *MouseHandler
}

// Button constructs a button from an image.
func Button(buttonImage *image) *button {
	b := button{image: buttonImage}
	return &b
}

// OnClick sets a callback activated when the button
// is clicked
func (btn *button) OnClick(callback func()) {
	handler := MouseHandler{
		RespondsTo: func(pos Pos, buttons MouseButtons) bool {
			return pos.In(btn.image.Rect())
		},
		OnClick: callback,
	}

	btn.onClickHandler = &handler

	MouseHandlers = append(MouseHandlers, btn.onClickHandler)
}
