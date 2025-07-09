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
		testCase{
			game:        New(&id),
			x:           0,
			y:           0,
			expectation: true,
		},
		testCase{
			game:        New(&id),
			x:           FIELD_SIDE - 1,
			y:           FIELD_SIDE - 1,
			expectation: true,
		},
		testCase{
			game:        New(&id),
			x:           FIELD_SIDE,
			y:           FIELD_SIDE - 1,
			expectation: false,
		},
		testCase{
			game:        New(&id),
			x:           0,
			y:           -5,
			expectation: false,
		},
		testCase{
			game:        New(&id),
			x:           FIELD_SIDE + 1,
			y:           1,
			expectation: false,
		},
		testCase{
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
