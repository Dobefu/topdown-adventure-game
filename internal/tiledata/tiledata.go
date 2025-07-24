// Package tiledata houses several constants that identify tiles by their ID.
package tiledata

const (
	// TileCollisionWall collision tile.
	TileCollisionWall = iota + 257
	// TileCollisionLedgeVertical collision tile.
	TileCollisionLedgeVertical
	// TileCollisionLedgeHorizontal collision tile.
	TileCollisionLedgeHorizontal
	// TileCollisionWater collision tile.
	TileCollisionWater
)
