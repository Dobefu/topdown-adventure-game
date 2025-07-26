package gameobject

import (
	"math"
	"slices"

	"github.com/Dobefu/topdown-adventure-game/internal/tiledata"
	"github.com/Dobefu/vectors"
)

// MoveWithCollisionRect moves a game object, with a given collision rectangle.
func (g *GameObject) MoveWithCollisionRect(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) (newVelocity vectors.Vector3, hasCollided bool, collidedTiles []int) {
	pos := g.GetPosition()

	pos.Add(vectors.Vector3{Z: velocity.Z})

	if pos.Z < 0 {
		pos.Z = 0
		velocity.Z = 0
	}

	g.SetPosition(*pos)

	var allCollidedTiles []int

	hasCollided, tilesCollidedWith := g.canMoveTo(vectors.Vector3{X: velocity.X, Y: 0, Z: 0}, x1, y1, x2, y2)
	if !hasCollided {
		pos.Add(vectors.Vector3{X: velocity.X, Y: 0, Z: 0})
		g.SetPosition(*pos)
		allCollidedTiles = append(allCollidedTiles, tilesCollidedWith...)
	} else if velocity.X != 0 {
		maxX := g.findMaxMovement(
			vectors.Vector3{X: velocity.X, Y: 0, Z: 0},
			x1,
			y1,
			x2,
			y2,
		)

		if maxX != 0 {
			pos.Add(vectors.Vector3{X: maxX, Y: 0, Z: 0})
			g.SetPosition(*pos)
		}

		allCollidedTiles = append(allCollidedTiles, tilesCollidedWith...)
	}

	pos = g.GetPosition()

	hasCollided, tilesCollidedWith = g.canMoveTo(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0}, x1, y1, x2, y2)
	if !hasCollided {
		pos.Add(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0})
		allCollidedTiles = append(allCollidedTiles, tilesCollidedWith...)
	} else if velocity.Y != 0 {
		maxY := g.findMaxMovement(
			vectors.Vector3{X: 0, Y: velocity.Y, Z: 0},
			x1,
			y1,
			x2,
			y2,
		)

		if maxY != 0 {
			pos.Add(vectors.Vector3{X: 0, Y: maxY, Z: 0})
		}

		allCollidedTiles = append(allCollidedTiles, tilesCollidedWith...)
	}

	g.SetPosition(*pos)

	return velocity, slices.Contains(allCollidedTiles, tiledata.TileCollisionWall), allCollidedTiles
}

func (g *GameObject) findMaxMovement(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) float64 {
	var target float64

	if velocity.X != 0 {
		target = velocity.X
	} else {
		target = velocity.Y
	}

	var minDistance float64
	maxDistance := math.Abs(target)
	threshold := 0.1

	for maxDistance-minDistance > threshold {
		center := (minDistance + maxDistance) / 2

		testVelocity := getTestVelocityFromVelocity(velocity, target, center)
		hasCollided, _ := g.canMoveTo(testVelocity, x1, y1, x2, y2)

		if hasCollided {
			maxDistance = center
		} else {
			minDistance = center
		}
	}

	if target > 0 {
		return minDistance
	}

	return -minDistance
}

func (g *GameObject) getCollidedTile(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) []int {
	pos := g.GetPosition()
	scene := *g.GetScene()

	target := vectors.Vector2{
		X: pos.X + velocity.X,
		Y: pos.Y + velocity.Y,
	}

	var corners []int

	topLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x1, Y: target.Y + y1})
	if topLeft != 0 {
		corners = append(corners, topLeft)
	}

	topRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x2, Y: target.Y + y1})
	if topRight != 0 {
		corners = append(corners, topRight)
	}

	bottomLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x1, Y: target.Y + y2})
	if bottomLeft != 0 {
		corners = append(corners, bottomLeft)
	}

	bottomRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x2, Y: target.Y + y2})
	if bottomRight != 0 {
		corners = append(corners, bottomRight)
	}

	uniqueTiles := make(map[int]bool)
	var result []int

	for _, corner := range corners {
		if uniqueTiles[corner] {
			continue
		}

		uniqueTiles[corner] = true
		result = append(result, corner)
	}

	return result
}

func (g *GameObject) canMoveTo(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) (hasCollided bool, collidedTiles []int) {
	collidedTiles = g.getCollidedTile(velocity, x1, y1, x2, y2)

	return slices.Contains(collidedTiles, tiledata.TileCollisionWall), collidedTiles
}

func getTestVelocityFromVelocity(
	velocity vectors.Vector3,
	target float64,
	center float64,
) (testVelocity vectors.Vector3) {
	if velocity.X != 0 {
		if target > 0 {
			testVelocity = vectors.Vector3{X: center, Y: 0, Z: 0}
		} else {
			testVelocity = vectors.Vector3{X: -center, Y: 0, Z: 0}
		}
	} else {
		if target > 0 {
			testVelocity = vectors.Vector3{X: 0, Y: center, Z: 0}
		} else {
			testVelocity = vectors.Vector3{X: 0, Y: -center, Z: 0}
		}
	}

	return testVelocity
}
