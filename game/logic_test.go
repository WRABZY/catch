package catch

import "testing"

func TestNew(t *testing.T) {
	id := "abc"
	game := New(&id)

	if *game.id != id {
		t.Errorf(`game.id = %q, but must be %q`, *game.id, id)
	}

	for i := 0; i < FIELD_SIZE; i++ {
		if game.field[i] != EMPTY {
			t.Errorf(`game.field[%d] = %d, but must be %d (EMPTY)`, i, game.field[i], EMPTY)
		}
		if game.oldField[i] != EMPTY {
			t.Errorf(`game.fioldFieldeld[%d] = %d, but must be %d (EMPTY)`, i, game.oldField[i], EMPTY)
		}
	}

	if game.playerHP != PLAYER_HP_MAX {
		t.Errorf(`game.playerHP = %d, but must be %d (PLAYER_HP_MAX)`, game.playerHP, PLAYER_HP_MAX)
	}

	if game.oldPlayerHP != PLAYER_HP_MAX {
		t.Errorf(`game.oldPlayerHP = %d, but must be %d (PLAYER_HP_MAX)`, game.oldPlayerHP, PLAYER_HP_MAX)
	}

	if game.playerSteps != ZERO {
		t.Errorf(`game.playerSteps = %d, but must be 0`, game.playerSteps)
	}

	if game.playerScore != ZERO {
		t.Errorf(`game.playerScore = %d, but must be 0`, game.playerScore)
	}
}
