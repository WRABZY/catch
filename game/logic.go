package catch

type Game struct {
	id             *string
	field          *[FIELD_SIDE][FIELD_SIDE]int8
	playerHP       int8
	playerSteps    int
	playerScore    int
	lastDirection  int8
	directionSteps int8
	x              int8
	y              int8
	enemiesXY      *map[int8]map[int8]struct{}
}

func New(id *string) *Game {
	var game = &Game{
		id:             id,
		field:          &[FIELD_SIDE][FIELD_SIDE]int8{},
		playerHP:       PLAYER_HP_MAX,
		playerSteps:    ZERO,
		playerScore:    ZERO,
		lastDirection:  NO_DIRECTION,
		directionSteps: ZERO,
		x:              NOWHERE,
		y:              NOWHERE,
	}

	for i := 0; i < FIELD_SIDE; i++ {
		game.field[i] = [FIELD_SIDE]int8{}
	}

	m := make(map[int8]map[int8]struct{})
	game.enemiesXY = &m

	return game
}

func (g *Game) insertPlayer(x, y int8) (ok bool) {
	if x > NOWHERE &&
		y > NOWHERE &&
		x < FIELD_SIDE &&
		y < FIELD_SIDE &&
		g.x == NOWHERE &&
		g.y == NOWHERE &&
		(g.field[y][x] == EMPTY || g.field[y][x] > KING) {

		g.field[y][x] = PLAYER
		g.x = x
		g.y = y
		return true
	}
	return false
}

func (g *Game) insertPawnNorth(x, y int8) (ok bool) {
	if x > NOWHERE &&
		x < FIELD_SIDE &&
		y == ZERO &&
		(g.field[y][x] == EMPTY || g.field[y][x] > KING) {

		g.field[y][x] = PAWN_NORTH
		if _, ok := (*g.enemiesXY)[x]; !ok {
			(*g.enemiesXY)[x] = make(map[int8]struct{})
		}
		(*g.enemiesXY)[x][y] = struct{}{}

		return true
	}
	return false
}

func (g *Game) insertPawnEast(x, y int8) (ok bool) {
	if x == LAST_INDEX &&
		y > NOWHERE &&
		y < FIELD_SIDE &&
		(g.field[y][x] == EMPTY || g.field[y][x] > KING) {

		g.field[y][x] = PAWN_EAST
		if _, ok := (*g.enemiesXY)[x]; !ok {
			(*g.enemiesXY)[x] = make(map[int8]struct{})
		}
		(*g.enemiesXY)[x][y] = struct{}{}

		return true
	}
	return false
}

func (g *Game) insertPawnSouth(x, y int8) (ok bool) {
	if x > NOWHERE &&
		x < FIELD_SIDE &&
		y == LAST_INDEX &&
		(g.field[y][x] == EMPTY || g.field[y][x] > KING) {

		g.field[y][x] = PAWN_SOUTH
		if _, ok := (*g.enemiesXY)[x]; !ok {
			(*g.enemiesXY)[x] = make(map[int8]struct{})
		}
		(*g.enemiesXY)[x][y] = struct{}{}

		return true
	}
	return false
}

func (g *Game) insertPawnWest(x, y int8) (ok bool) {
	if x == ZERO &&
		y > NOWHERE &&
		y < FIELD_SIDE &&
		(g.field[y][x] == EMPTY || g.field[y][x] > KING) {

		g.field[y][x] = PAWN_WEST
		if _, ok := (*g.enemiesXY)[x]; !ok {
			(*g.enemiesXY)[x] = make(map[int8]struct{})
		}
		(*g.enemiesXY)[x][y] = struct{}{}

		return true
	}
	return false
}
