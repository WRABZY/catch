package catch

import (
	"math/rand/v2"
	"time"
)

type Game struct {
	id          *string
	field       *[FieldSide][FieldSide]int8
	enemiesXY   *map[int8]map[int8]struct{}
	hp          int8
	steps       uint64
	score       uint64
	direction   int8
	dirSteps    int8
	x           int8
	y           int8
	yNorth      int8
	xEast       int8
	ySouth      int8
	xWest       int8
	random      *rand.Rand
	nextPawnDir int8
	nextPawnPos int8
}

type Randomizer struct{}

func (r *Randomizer) Uint64() uint64 {
	return uint64(time.Now().UnixMilli())
}

func newGame(id *string) *Game {
	var game = &Game{
		id:        id,
		field:     &[FieldSide][FieldSide]int8{},
		hp:        HPMax,
		steps:     Zero,
		score:     Zero,
		direction: NoDirection,
		dirSteps:  Zero,
		x:         Center,
		y:         Center,
		yNorth:    Center - 1,
		xEast:     Center + 1,
		ySouth:    Center + 1,
		xWest:     Center - 1,
	}

	m := make(map[int8]map[int8]struct{})
	game.enemiesXY = &m

	var i int8
	for i = Zero; i < FieldSide; i++ {
		game.field[i] = [FieldSide]int8{}
		m[i] = make(map[int8]struct{})
	}
	game.field[game.y][game.x] = Player

	game.random = rand.New(new(Randomizer))
	game.nextPawnDir = int8(game.random.IntN(DirectionsNumber) + North)
	game.nextPawnPos = int8(game.random.IntN(FieldSide))

	game.insertPawn()

	return game
}

func (g *Game) insertPawn() { // pawns - clockwise
	var x *int8
	var y *int8
	var zl int8
	var pawnToInsert int8
	var inserted = false

	for i := Zero; i < FieldPerimeter && !inserted; i++ {
		switch g.nextPawnDir {
		case North:
			x = &g.nextPawnPos
			zl = Zero
			y = &zl
			pawnToInsert = PawnNorth
			g.nextPawnDir = East
		case East:
			y = &g.nextPawnPos
			zl = LastIndex
			x = &zl
			pawnToInsert = PawnEast
			g.nextPawnDir = South
		case South:
			x = &g.nextPawnPos
			zl = LastIndex
			y = &zl
			pawnToInsert = PawnSouth
			g.nextPawnDir = West
		case West:
			y = &g.nextPawnPos
			zl = Zero
			x = &zl
			pawnToInsert = PawnWest
			g.nextPawnDir = North
		}

		if g.field[*y][*x] < Player || g.field[*y][*x] > King {
			g.field[*y][*x] = pawnToInsert
			(*g.enemiesXY)[*x][*y] = struct{}{}
			inserted = true
		}

		if g.nextPawnPos == LastIndex {
			g.nextPawnPos = Zero
		} else {
			g.nextPawnPos++
		}
	}

}
