package catch

import (
	"testing"
)

func TestNew(t *testing.T) {
	id := "TestNew"
	game := New(&id)

	if *game.id != id {
		t.Errorf(`game.id = %q, but must be %q`, *game.id, id)
	}

	for i := 0; i < FIELD_SIDE; i++ {
		for j := 0; j < FIELD_SIDE; j++ {
			if game.field[i][j] != EMPTY {
				t.Errorf(`game.field[%d][%d] = %d, but must be %d (EMPTY)`, i, j, game.field[i][j], EMPTY)
			}
		}

	}

	if game.playerHP != PLAYER_HP_MAX {
		t.Errorf(`game.playerHP = %d, but must be %d (PLAYER_HP_MAX)`, game.playerHP, PLAYER_HP_MAX)
	}

	if game.playerSteps != ZERO {
		t.Errorf(`game.playerSteps = %d, but must be 0`, game.playerSteps)
	}

	if game.playerScore != ZERO {
		t.Errorf(`game.playerScore = %d, but must be 0`, game.playerScore)
	}

	if game.lastDirection != NO_DIRECTION {
		t.Errorf(`game.lastDirection = %d, but must be 0 (NO_DIRECTION)`, game.lastDirection)
	}

	if game.directionSteps != ZERO {
		t.Errorf(`game.directionSteps = %d, but must be 0`, game.directionSteps)
	}

	if game.x != NOWHERE {
		t.Errorf(`game.x = %d, but must be %d`, game.x, NOWHERE)
	}

	if game.y != NOWHERE {
		t.Errorf(`game.y = %d, but must be %d`, game.y, NOWHERE)
	}

	if len(*game.enemiesXY) != ZERO {
		t.Errorf(`len(*game.enemiesXY) = %d, but must be 0`, len(*game.enemiesXY))
	}
}

func TestInsertPlayer(t *testing.T) {
	type testCase struct {
		game        *Game
		x           int8
		y           int8
		expectation bool
	}

	id := "TestInsertPlayer"
	cases := []testCase{
		{
			game:        New(&id),
			x:           0,
			y:           0,
			expectation: true,
		},
		{
			game:        New(&id),
			x:           FIELD_SIDE - 1,
			y:           FIELD_SIDE - 1,
			expectation: true,
		},
		{
			game:        New(&id),
			x:           FIELD_SIDE,
			y:           FIELD_SIDE - 1,
			expectation: false,
		},
		{
			game:        New(&id),
			x:           0,
			y:           -5,
			expectation: false,
		},
		{
			game:        New(&id),
			x:           FIELD_SIDE + 1,
			y:           1,
			expectation: false,
		},
		{
			game:        New(&id),
			x:           3,
			y:           7,
			expectation: true,
		},
	}

	for i, c := range cases {
		if c.game.insertPlayer(c.x, c.y) != c.expectation {
			t.Errorf(`Case %d, insertion player to x = %d, y = %d must be %t, but was %t`, i, c.x, c.y, c.expectation, !c.expectation)
		}
		if c.expectation {
			if c.game.x != c.x {
				t.Errorf(`Case %d, game.x != case.x (%d != %d)`, i, c.game.x, c.x)
			}
			if c.game.y != c.y {
				t.Errorf(`Case %d, game.y != case.y (%d != %d)`, i, c.game.y, c.y)
			}
		}
	}

	// Extra
	if cases[0].game.insertPlayer(cases[0].x, cases[0].y) {
		t.Errorf(`Extra, re-insertion player to x = %d, y = %d must be false, but was true`, cases[0].game.x, cases[0].game.y)
	}

	extraCaseBaseIndex := 5
	for i, arr := range cases[extraCaseBaseIndex].game.field {
		for j, entity := range arr {
			if entity == PLAYER {
				if int8(i) != cases[extraCaseBaseIndex].y {
					t.Errorf(`Extra, the player had to be inserted into y = %d, but was inserted into y = %d`, cases[extraCaseBaseIndex].y, i)
				}
				if int8(j) != cases[extraCaseBaseIndex].x {
					t.Errorf(`Extra, the player had to be inserted into x = %d, but was inserted into x = %d`, cases[extraCaseBaseIndex].x, j)
				}
			}
		}
	}
}

func TestInsertPawns(t *testing.T) {
	id := "TestInsertPawns"
	game := New(&id)

	var x int8 = -128
	var y int8 = 0

	if game.insertPawnNorth(x, y) {
		t.Errorf(`North pawn was inserted into x = %d y = %d`, x, y)
	}

	x = 0
	y = 127

	if game.insertPawnNorth(x, y) {
		t.Errorf(`North pawn was inserted into x = %d y = %d`, x, y)
	}

	for x, y = 0, 0; x < FIELD_SIDE; x++ {
		if !game.insertPawnNorth(x, y) {
			t.Errorf(`North pawn wasn't inserted into x = %d y = %d`, x, y)
		}
	}

	if game.insertPawnNorth(x, y) {
		t.Errorf(`North pawn was inserted into x = %d y = %d`, x, y)
	}

	x--
	if game.insertPawnNorth(x, y) {
		t.Errorf(`North pawn was inserted into x = %d y = %d, but there is already a north pawn`, x, y)
	}

	if game.insertPawnEast(x, y) {
		t.Errorf(`East pawn was inserted into x = %d y = %d, but there is already a north pawn`, x, y)
	}

	for y++; y < FIELD_SIDE; y++ {
		if !game.insertPawnEast(x, y) {
			t.Errorf(`East pawn wasn't inserted into x = %d y = %d`, x, y)
		}
	}

	if game.insertPawnEast(x, y) {
		t.Errorf(`East pawn was inserted into x = %d y = %d`, x, y)
	}

	y--
	if game.insertPawnEast(x, y) {
		t.Errorf(`East pawn was inserted into x = %d y = %d, but there is already a east pawn`, x, y)
	}

	if game.insertPawnSouth(x, y) {
		t.Errorf(`South pawn was inserted into x = %d y = %d, but there is already a east pawn`, x, y)
	}

	for x--; x > NOWHERE; x-- {
		if !game.insertPawnSouth(x, y) {
			t.Errorf(`South pawn wasn't inserted into x = %d y = %d`, x, y)
		}
	}

	if game.insertPawnSouth(x, y) {
		t.Errorf(`South pawn was inserted into x = %d y = %d`, x, y)
	}

	x++
	if game.insertPawnSouth(x, y) {
		t.Errorf(`South pawn was inserted into x = %d y = %d, but there is already a south pawn`, x, y)
	}

	if game.insertPawnWest(x, y) {
		t.Errorf(`West pawn was inserted into x = %d y = %d, but there is already a south pawn`, x, y)
	}

	for y--; y > ZERO; y-- {
		if !game.insertPawnWest(x, y) {
			t.Errorf(`West pawn wasn't inserted into x = %d y = %d`, x, y)
		}
	}

	if game.insertPawnWest(x, y) {
		t.Errorf(`West pawn was inserted into x = %d y = %d, but there is already a west pawn`, x, y)
	}

	for x = ZERO; x < FIELD_SIDE; x++ {
		for y = ZERO; y < FIELD_SIDE; y++ {
			if game.insertPawnWest(x, y) {
				if x == ZERO || y == ZERO || x == LAST_INDEX || y == LAST_INDEX {
					t.Errorf(`Pawn was repeatedly inserted into x = %d y = %d`, x, y)
				}
				t.Errorf(`Pawn was inserted into x = %d y = %d`, x, y)
			}
		}
	}

	enemies := 0
	for x, row := range *game.enemiesXY {
		for y = range row {
			enemies++
			if x == ZERO || x == LAST_INDEX || y == ZERO || y == LAST_INDEX {
				continue
			}
			t.Errorf(`In game.enemiesXY found pawn with x = %d, y = %d`, x, y)
		}
	}

	mustBeEnemies := FIELD_SIDE*4 - 4
	if enemies != mustBeEnemies {
		t.Errorf(`Map game.enemiesXY contains %d items, but must be %d`, enemies, mustBeEnemies)
	}

	northPawns := 0
	eastPawns := 0
	southPawns := 0
	westPawns := 0
	for x = 0; x < FIELD_SIDE; x++ {
		for y = 0; y < FIELD_SIDE; y++ {
			switch game.field[y][x] {
			case ZERO:
				continue
			case PAWN_NORTH:
				northPawns++
				if y != ZERO {
					t.Errorf(`In game.field found north pawn with y = %d`, y)
				}
			case PAWN_EAST:
				eastPawns++
				if x != LAST_INDEX {
					t.Errorf(`In game.field found east pawn with x = %d`, x)
				}
			case PAWN_SOUTH:
				southPawns++
				if y != LAST_INDEX {
					t.Errorf(`In game.field found south pawn with y = %d`, y)
				}
			case PAWN_WEST:
				westPawns++
				if x != ZERO {
					t.Errorf(`In game.field found west pawn with x = %d`, x)
				}
			}
		}
	}

	pawns := northPawns + eastPawns + southPawns + westPawns
	if pawns != enemies {
		t.Errorf(`In game.field found %d pawns, in game.enemiesXY %d enemies`, pawns, enemies)
	}

	game = New(&id)
	if !game.insertPawnEast(LAST_INDEX, CENTER) {
		t.Errorf(`Can't insert east pawn into x = %d y = %d`, LAST_INDEX, CENTER)
	}

	if len(*game.enemiesXY) != 1 {
		t.Errorf(`len(*game.enemiesXY) = %d after inserting 1 east pawn`, len(*game.enemiesXY))
	}

	game = New(&id)
	if !game.insertPawnSouth(CENTER, LAST_INDEX) {
		t.Errorf(`Can't insert south pawn into x = %d y = %d`, CENTER, LAST_INDEX)
	}

	if len(*game.enemiesXY) != 1 {
		t.Errorf(`len(*game.enemiesXY) = %d after inserting 1 south pawn`, len(*game.enemiesXY))
	}

	game = New(&id)
	if !game.insertPawnWest(ZERO, CENTER) {
		t.Errorf(`Can't insert west pawn into x = %d y = %d`, ZERO, CENTER)
	}

	if len(*game.enemiesXY) != 1 {
		t.Errorf(`len(*game.enemiesXY) = %d after inserting 1 west pawn`, len(*game.enemiesXY))
	}
}
