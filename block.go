// Copyright (c) 2022 Elias Daler
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Block struct {
	Entity

	IsBroken bool
}

func makeBlock(pos Vec2f, img *ebiten.Image) (*Block, error) {
	e := &Block{
		Entity: NewEntity(),
	}

	e.Transform.Position = pos

	e.Graphics.Sprite = NewSprite(img)
	return e, nil
}
