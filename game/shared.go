package catch

// Entity codes
const (
	Nobody = iota
	Player
	PawnNorth
	PawnEast
	PawnSouth
	PawnWest
	DamagedRook
	Rook
	DamagedBishop2
	DamagedBishop1
	Bishop
	DamagedKnight3
	DamagedKnight2
	DamagedKnight1
	Knight
	DamagedQueen4
	DamagedQueen3
	DamagedQueen2
	DamagedQueen1
	Queen
	King
	ShadowPlayer
	ShadowPawn
	ShadowRook
	ShadowBishop
	ShadowKnight
	ShadowQueen
	ShadowKing
)

const Zero = Nobody

// Placing
const (
	FieldSide      = DamagedBishop1
	FieldSize      = FieldSide * FieldSide
	FieldPerimeter = 4 * (FieldSide - 1)
	Nowhere        = -1
	Center         = FieldSide / 2
	LastIndex      = FieldSide - 1

	DirectionsNumber = PawnSouth
	North            = PawnNorth
	East             = PawnEast
	South            = PawnSouth
	West             = PawnWest
	NoDirection      = Nobody
)

const HPMax = DamagedBishop1
