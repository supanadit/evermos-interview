package main

import "fmt"

type sprite struct {
	Ascii     byte
	Forbidden bool
}

func wall() sprite {
	return sprite{
		Ascii:     '#',
		Forbidden: true,
	}
}

func floor() sprite {
	return sprite{
		Ascii:     '.',
		Forbidden: false,
	}
}

func player() sprite {
	return sprite{
		Ascii:     'X',
		Forbidden: false,
	}
}

type area struct {
	Sprite          [][]sprite
	Player          sprite
	LocationXPlayer int
	LocationYPlayer int
	Width           int
	Height          int
}

func (r *area) createArea(height int, width int) {
	a := make([][]sprite, height)
	for i := 0; i < height; i++ {
		a[i] = make([]sprite, width)
		for j := 0; j < width; j++ {
			if i != 0 && i != (height-1) && j != (width-1) && j != 0 {
				a[i][j] = floor()
			} else {
				a[i][j] = wall()
			}
		}
	}
	r.Width = width
	r.Height = height
	r.Sprite = a
}

func (r *area) setPlayer(s sprite, locationX int, locationY int) {
	r.Sprite[locationY][locationX] = s
	r.LocationXPlayer = locationX
	r.LocationYPlayer = locationY
	r.Player = s
}

func (r *area) setObject(s sprite, locationX int, locationY int) {
	r.Sprite[locationY][locationX] = s
}

func (r *area) draw() {
	for i := 0; i < len(r.Sprite); i++ {
		for j := 0; j < len(r.Sprite[i]); j++ {
			fmt.Print(string(r.Sprite[i][j].Ascii))
		}
		fmt.Printf("\n")
	}
}

func (r *area) playerMoveToNorth() (s bool) {
	if r.isPlayerCanMoveToLocation(r.LocationXPlayer, r.LocationYPlayer-1) {
		r.replaceWithFloor(r.LocationXPlayer, r.LocationYPlayer)
		r.LocationYPlayer -= 1
		r.Sprite[r.LocationYPlayer][r.LocationXPlayer] = r.Player
		s = true
	}
	return s
}

func (r *area) playerMoveToEast() (s bool) {
	if r.isPlayerCanMoveToLocation(r.LocationXPlayer+1, r.LocationYPlayer) {
		r.replaceWithFloor(r.LocationXPlayer, r.LocationYPlayer)
		r.LocationXPlayer += 1
		r.Sprite[r.LocationYPlayer][r.LocationXPlayer] = r.Player
		s = true
	}
	return s
}

func (r *area) playerMoveToWest() (s bool) {
	if r.isPlayerCanMoveToLocation(r.LocationXPlayer-1, r.LocationYPlayer) {
		r.replaceWithFloor(r.LocationXPlayer, r.LocationYPlayer)
		r.LocationXPlayer -= 1
		r.Sprite[r.LocationYPlayer][r.LocationXPlayer] = r.Player
		s = true
	}
	return s
}

func (r *area) playerMoveToSouth() (s bool) {
	if r.isPlayerCanMoveToLocation(r.LocationXPlayer, r.LocationYPlayer+1) {
		r.replaceWithFloor(r.LocationXPlayer, r.LocationYPlayer)
		r.LocationYPlayer += 1
		r.Sprite[r.LocationYPlayer][r.LocationXPlayer] = r.Player
		s = true
	}
	return s
}

func (r area) isPlayerCanMoveToLocation(locationX int, locationY int) bool {
	return !r.Sprite[locationY][locationX].Forbidden
}

func (r *area) replaceWithFloor(locationX int, locationY int) {
	r.Sprite[locationY][locationX] = floor()
}

func (r *area) replaceWithWall(locationX int, locationY int) {
	r.Sprite[locationY][locationX] = wall()
}

func (r *area) countTotalFloor() int {
	totalFloor := 0
	for i := 0; i < len(r.Sprite); i++ {
		for j := 0; j < len(r.Sprite[i]); j++ {
			if !r.Sprite[i][j].Forbidden {
				totalFloor += 1
			}
		}
	}
	return totalFloor
}

func main() {
	r := area{}
	r.createArea(6, 8)
	r.setPlayer(player(), 1, 4)
	// Dinding di samping player
	r.setObject(wall(), 2, 4)

	// Dinding di baris ke 3
	r.setObject(wall(), 4, 3)
	r.setObject(wall(), 6, 3)

	// Dinding di baris ke 2
	r.setObject(wall(), 2, 2)
	r.setObject(wall(), 3, 2)
	r.setObject(wall(), 4, 2)

	northStep := 0
	// Move to north till finish
	for r.playerMoveToNorth() {
		northStep += 1
	}
	eastStep := 0
	// Move to east till finish
	for r.playerMoveToEast() {
		eastStep += 1
	}
	southStep := 0
	// Move to south till finish
	for r.playerMoveToSouth() {
		southStep += 1
	}

	r.draw()
	fmt.Printf("Walk to the north `A` step(s) = %d\n", northStep)
	fmt.Printf("Then walk to the east `B` step(s) = %d\n", eastStep)
	fmt.Printf("Last, walk to the south `C` step(s) = %d\n", southStep)
	fmt.Printf("Total step = %d\n", northStep+eastStep+southStep)
	fmt.Printf("Total floor available to walk = %d\n", r.countTotalFloor())
}
