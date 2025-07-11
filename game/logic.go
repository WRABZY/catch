package catch

type Game struct {
	id        *string
	field     *[FIELD_SIDE][FIELD_SIDE]int8
	enemiesXY *map[int8]map[int8]struct{}
	hp        int8
	steps     uint64
	score     uint64
	direction int8
	dirSteps  int8
	x         int8
	y         int8
	yNorth    int8
	xEast     int8
	ySouth    int8
	xWest     int8
}

func newGame(id *string) *Game {
	var game = &Game{
		id:        id,
		field:     &[FIELD_SIDE][FIELD_SIDE]int8{},
		hp:        START_HP,
		steps:     ZERO,
		score:     ZERO,
		direction: NO_DIRECTION,
		dirSteps:  ZERO,
		x:         START_X_POS,
		y:         START_Y_POS,
		yNorth:    START_Y_POS - 1,
		xEast:     START_X_POS + 1,
		ySouth:    START_Y_POS + 1,
		xWest:     START_X_POS - 1,
	}

	m := make(map[int8]map[int8]struct{})
	game.enemiesXY = &m

	var i int8
	for i = ZERO; i < FIELD_SIDE; i++ {
		game.field[i] = [FIELD_SIDE]int8{}
		m[i] = make(map[int8]struct{})
	}

	return game
}
