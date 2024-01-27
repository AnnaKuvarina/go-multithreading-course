package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rWLock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	updateBoidMapItem(b, -1)
	b.position = b.position.Add(b.velocity)
	updateBoidMapItem(b, b.id)
	/*	next := b.position.Add(b.velocity)
		if next.x >= screenWidth || next.x < 0 {
			b.velocity = Vector2d{-b.velocity.x, b.velocity.y}
		}
		if next.y >= screenHeight || next.y < 0 {
			b.velocity = Vector2d{b.velocity.x, -b.velocity.y}
		}*/

	rWLock.Unlock()
}

func (b *Boid) calcAcceleration() Vector2d {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition := Vector2d{0, 0}
	avgVelocity := Vector2d{0, 0}
	separation := Vector2d{0, 0}

	count := 0.0

	rWLock.RLock()

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				anotherBoid := boids[otherBoidId]
				if dist := anotherBoid.position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(anotherBoid.velocity)
					avgPosition = avgPosition.Add(anotherBoid.position)
					separation = separation.Add(b.position.Subtract(anotherBoid.position).DivisionV(dist))
				}
			}
		}
	}

	rWLock.RUnlock()

	accel := Vector2d{b.borderBounce(b.position.x, screenWidth),
		b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		avgPosition = avgPosition.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	}
	if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}

	return 0
}

func createBoid(boidID int) *Boid {
	return &Boid{
		position: Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2d{(rand.Float64() * 2) - 1, (rand.Float64() * 2) - 1},
		id:       boidID,
	}
}

func updateBoidMapItem(boid *Boid, value int) {
	boidMap[int(boid.position.x)][int(boid.position.y)] = value
}
