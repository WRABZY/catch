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
