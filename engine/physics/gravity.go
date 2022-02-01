package physics

type Gravity struct {
	Acceleration float64
	Velocity     float64
}

func NewGravity(acc float64) Gravity {
	return Gravity{
		Acceleration: acc,
	}
}

func (g *Gravity) CalculateVelocity() float64 {
	g.Velocity += g.Acceleration
	return g.Velocity
}

func (g *Gravity) ResetVelocity() {
	g.Velocity = 0
}
