package catch

type Game struct {
	id          *string
	field       *[FIELD_SIZE]int8
	oldField    *[FIELD_SIZE]int8
	playerHP    int8
	oldPlayerHP int8
	playerSteps int
	playerScore int
}

func New(id *string) *Game {
	var game = &Game{
		id:          id,
		field:       &[FIELD_SIZE]int8{},
		oldField:    &[FIELD_SIZE]int8{},
		playerHP:    PLAYER_HP_MAX,
		oldPlayerHP: PLAYER_HP_MAX,
		playerSteps: ZERO,
		playerScore: ZERO,
	}

	return game
}
