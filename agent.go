package main

import (
	"math/rand"
	"time"
)

var StickRelease = Vec2f{0, -1}

var StickChoices = []Vec2f{
	{1, 0},
	{0, 0},
	{-1, 0},
}

type AgentAlgo int

const (
	AgentAlgoFollow AgentAlgo = iota
	AgentAlgoCenter
	AgentAlgoRandom
)

const AgentAlgoCount = 3

type Agent struct {
	ticks  int
	active bool
	runs   int
	wins   int
	game   *Game
	algo   AgentAlgo
}

func makeAgent(g *Game) (*Agent, error) {
	a := &Agent{
		ticks:  0,
		active: false,
		game:   g,
		runs:   0,
		wins:   0,
		algo:   AgentAlgoFollow,
	}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	return a, nil
}

func (a *Agent) Win() error {
	a.ticks = 0
	a.runs += 1
	a.wins += 1

	return nil
}

func (a *Agent) Toggle() error {
	a.active = !a.active

	return nil
}

func (a *Agent) NextAlgo() error {
	a.algo = (a.algo + 1) % AgentAlgoCount

	return nil
}

func (a *Agent) Lose() error {
	a.ticks = 0
	a.runs += 1

	return nil
}

func (a *Agent) Tick() error {
	a.ticks += 1

	return nil
}

func (a *Agent) getStickStateRandom() Vec2f {
	if a.game.attachBallToPaddle {
		return StickRelease
	}

	return StickChoices[rand.Intn(len(StickChoices))]
}

func (a *Agent) getStickStateCenter() Vec2f {
	if a.game.attachBallToPaddle {
		return StickRelease
	}

	paddleCenter := a.game.paddle.Transform.Position.X + a.game.paddle.Graphics.Sprite.GetSize().X/2
	ballCenter := a.game.ball.Transform.Position.X + a.game.ball.Graphics.Sprite.GetSize().X/2

	stick := Vec2f{}

	if paddleCenter > ballCenter {
		stick.X -= 1
	} else if paddleCenter < ballCenter {
		stick.X += 1
	} else {
		if a.game.ball.Movement.Velocity.X > 0 {
			stick.X += 1
		} else {
			stick.X -= 1
		}
	}

	return stick
}

func (a *Agent) getStickStateFollow() Vec2f {
	if a.game.attachBallToPaddle {
		return StickRelease
	}

	stick := Vec2f{}

	if a.game.paddle.Transform.Position.X > a.game.ball.Transform.Position.X {
		stick.X -= 1
	} else if a.game.paddle.Transform.Position.X < a.game.ball.Transform.Position.X {
		stick.X += 1
	} else {
		if a.game.ball.Movement.Velocity.X > 0 {
			stick.X += 1
		} else {
			stick.X -= 1
		}
	}

	return stick
}

func (a *Agent) GetStickState() Vec2f {
	switch a.algo {
	case AgentAlgoRandom:
		return a.getStickStateRandom()
	case AgentAlgoFollow:
		return a.getStickStateFollow()
	case AgentAlgoCenter:
		return a.getStickStateCenter()
	default:
		panic("unexpected algo")
	}
}
