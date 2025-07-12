package catch

import (
	"testing"
)

func TestInsertPawn(t *testing.T) {
	id := "TestInsertPawn"
	game := newGame(&id)

	/*
		fmt.Println("game.nextPawnDir", game.nextPawnDir)
		fmt.Println("game.nextPawnPos", game.nextPawnPos)
	*/

	for i := Zero; i < FieldPerimeter+1; i++ {
		game.insertPawn()
	}

	/*
		fmt.Println("game.field")
		for _, arr := range game.field {
			fmt.Printf("%v\n", arr)
		}
	*/

	/*
		fmt.Println("game.enemiesXY")
		for _, m := range *game.enemiesXY {
			fmt.Printf("%v\n", m)
		}
	*/

	enemies := Zero
	for y, arr := range game.field {
		for x, v := range arr {
			switch v {
			case Nobody:
				if x == Zero || x == LastIndex || y == Zero || y == LastIndex {
					t.Errorf("Empty cell x = %d, y = %d. Must be pawn", x, y)
				}
			case Player:
				if int8(x) != Center || int8(y) != Center {
					t.Errorf("The Player is shifted along the x-axis x = %d, y = %d. Must be in center: x = %d, y = %d", x, y, Center, Center)
				}
				if int8(x) != game.x || int8(y) != game.y {
					t.Errorf("The Player found on x = %d, y = %d. But in game his coodinates x = %d, y = %d", x, y, game.x, game.y)
				}
			case PawnNorth:
				if y > Zero {
					t.Errorf("North pawn found on y = %d (x = %d). But north pawn's y-coordinate can't be other than %d", y, x, Zero)
				}
				enemies++
			case PawnEast:
				if x < LastIndex {
					t.Errorf("East pawn found on x = %d (y = %d). But east pawn's x-coordinate can't be other than %d", x, y, LastIndex)
				}
				enemies++
			case PawnSouth:
				if y < LastIndex {
					t.Errorf("South pawn found on y = %d (x = %d). But south pawn's y-coordinate can't be other than %d", y, x, LastIndex)
				}
				enemies++
			case PawnWest:
				if x > Zero {
					t.Errorf("West pawn found on x = %d (y = %d). But west pawn's x-coordinate can't be other than %d", x, y, Zero)
				}
				enemies++
			}
		}
	}

	for x, m := range *game.enemiesXY {
		for y := range m {
			if x != Zero && x != LastIndex && y != Zero && y != LastIndex {
				t.Errorf("Coordinates x = %d, y = %d were found in game.enemiesXY", x, y)
			}
			enemies--
		}
	}

	if enemies > Zero {
		t.Errorf("More enemies found on game.field than in game.enemiesXY, dif: %d", enemies)
	} else if enemies < Zero {
		t.Errorf("More enemies found in game.enemiesXY than on game.field, dif: %d", enemies)
	}

}
