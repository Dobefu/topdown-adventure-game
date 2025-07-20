package fastrand

import "math/rand"

/**
 * This code was taken directly from the Ebitengine examples:
 * https://github.com/hajimehoshi/ebiten/blob/2.8/examples/noise/main.go#L35
 */

var (
	Rand fastrand
)

func init() {
	Rand = fastrand{
		x: rand.Uint32(), // #nosec G404
		y: rand.Uint32(), // #nosec G404
		z: rand.Uint32(), // #nosec G404
		w: rand.Uint32(), // #nosec G404
	}
}

type fastrand struct {
	x uint32
	y uint32
	z uint32
	w uint32
}

func (r *fastrand) Next() uint32 {
	t := r.x ^ (r.x << 11)
	r.x, r.y, r.z = r.y, r.z, r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))

	return r.w
}
