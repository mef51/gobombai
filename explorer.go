package gobombai

import (
	"github.com/aybabtme/bomberman/player"
)

var _ player.Player = &ExplorerPlayer{}

type ExplorerPlayer struct {
	state   player.State
	update  chan player.State
	outMove chan player.Move
}

func NewExplorerPlayer(state player.State) *ExplorerPlayer {
	explorer := &ExplorerPlayer{
		state:   state,
		update:  make(chan player.State),
		outMove: make(chan player.Move, 1),
	}

	go func() {
		explorer.outMove <- player.Right
	}()

	return explorer
}

func (e *ExplorerPlayer) Name() string {
	return e.state.Name
}

func (e *ExplorerPlayer) Move() <-chan player.Move {
	return e.outMove
}

func (e *ExplorerPlayer) Update() chan<- player.State {
	return e.update
}
