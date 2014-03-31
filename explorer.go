package gobombai

import (
	"github.com/aybabtme/bomberman/cell"
	"github.com/aybabtme/bomberman/logger"
	"github.com/aybabtme/bomberman/player"
)

var (
	_   player.Player = &ExplorerPlayer{} // ExplorerPlayer implements play.Player
	log               = logger.New("", "/home/mohammedc/dev/gocode/src/github.com/mef51/gobombai/player.log", logger.Debug)
)

type ExplorerPlayer struct {
	state   player.State
	update  chan player.State
	outMove chan player.Move
}

type point struct {
	x int
	y int
}

func (p *point) toString() string {
	return string(p.x) + ", " + string(p.y)
}

func NewExplorerPlayer(state player.State) *ExplorerPlayer {
	log.Infof("Starting Explorer Player")
	log.Debugf("What's up!")
	explorer := &ExplorerPlayer{
		state:   state,
		update:  make(chan player.State, 1),
		outMove: make(chan player.Move, 1),
	}

	go func() {
		for {
			newState := <-explorer.update
			board := newState.Board
			x, y := newState.X, newState.Y

			rocks := findRocks(x, y, board, 0)
			for _, rock := range rocks {
				log.Debugf("here")
				log.Debugf(rock.toString())
			}

			// explorer.outMove <- player.Up
		}
	}()

	return explorer
}

///////////////
// Find where the reachable rocks are and return a list of them
func findRocks(x, y int, board [][]*cell.Exported, steps int) (rocks []*point) {
	if steps > 10 {
		return rocks
	}

	// find all the breakable blocks that are reachable, then pick which one to break.
	directions := []player.Move{player.Up, player.Down, player.Left, player.Right}

	for _, dir := range directions {
		log.Debugf("%d", steps)
		log.Debugf("Direction:" + string(dir))
		nextx := x
		nexty := y
		switch dir {
		case player.Up:
			log.Debugf("Case Up")
			nexty -= 1
		case player.Down:
			log.Debugf("Case Down")
			nexty += 1
		case player.Left:
			log.Debugf("Case Left")
			nextx -= 1
		case player.Right:
			log.Debugf("Case Right")
			nextx += 1
		}
		// log.Debugf("Pos: %d, %d", x, y)
		// log.Debugf("NextPos: %d, %d", nextx, nexty)
		// log.Debugf(board[nextx][nexty].Name)
		if board[nextx][nexty].Name == "Ground" {
			log.Debugf("Found Ground")
			findRocks(nextx, nexty, board, steps+1)
		} else if board[nextx][nexty].Name == "Rock" {
			log.Debugf("Found Rock")
			return append(rocks, &point{nextx, nexty})
		} else {
			log.Debugf("Found Nothing")
			return rocks
		}
	}
	return rocks
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
