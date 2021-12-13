package main

type Ball struct {
	Entity

	frames map[string]FloatRect
}

func makeBall() (*Ball, error) {
	e := &Ball{
		Entity: NewEntity(),
	}

	e.Movement.Speed = Vec2f{90, 90}

	img, _, err := NewImageFromFile("data/images/ball.png")
	if err != nil {
		return nil, err
	}

	e.Graphics.Sprite = NewSprite(img)

	e.frames = map[string]FloatRect{
		"up.right":   {0, 0, 16, 16},
		"up.left":    {16, 0, 16, 16},
		"down.right": {0, 16, 16, 16},
		"down.left":  {16, 16, 16, 16},
		"idle":       {0, 32, 16, 16},
	}

	e.updateFrame()

	return e, nil
}

func (e *Ball) setFrame(fn string) {
	f, ok := e.frames[fn]
	if !ok {
		return
	}
	e.Graphics.Sprite.SetTextureRect(f.ToImageRect())
}

func (e *Ball) updateFrame() {
	m := e.Movement
	if m.Velocity == (Vec2f{}) {
		e.setFrame("idle")
		return
	}

	fn := ""

	if m.Velocity.Y < 0 {
		fn += "up."
	} else {
		fn += "down."
	}
	if m.Velocity.X < 0 {
		fn += "left"
	} else {
		fn += "right"
	}

	e.setFrame(fn)
}
