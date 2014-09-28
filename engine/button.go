package engine

type button struct {
	image          *image
	onClickHandler *MouseHandler
}

// Button constructs a button from an image.
func Button(buttonImage *image) *button {
	b := button{image: buttonImage}
	return &b
}

func (btn *button) Position() Pos {
	return btn.image.Position()
}

func (btn *button) SetPosition(pos Pos) {
	btn.image.SetPosition(pos)
}

func (btn *button) Visible() bool {
	return btn.image.Visible()
}

func (btn *button) SetVisible(visible bool) {
	btn.image.SetVisible(visible)
}

// Implement paintSrc
func (btn *button) PaintTo(dest paintDest) {
	btn.image.PaintTo(dest)
}

func (btn *button) RequiresRedraw() bool {
	return btn.image.redraw
}

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
