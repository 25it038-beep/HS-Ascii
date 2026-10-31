package animations

import (
	"fmt"
	"math"
)

func drawPiano(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	// Piano keyboard on the bottom 6 rows
	keyH := 6
	if keyH > h/2 {
		keyH = h / 2
	}
	startY := h - keyH

	// Floating music notes in the air
	notes := []rune{'♪', '♫', '♩', '♬', '♭', '♮'}
	numNotes := w / 8
	if numNotes < 3 {
		numNotes = 3
	}
	for i := 0; i < numNotes; i++ {
		nx := mod(i*19+seed*7+frameNum*2, w)
		ny := mod(startY-1-(frameNum+i*3), startY)
		if ny >= 0 && ny < startY && nx < w {
			grid[ny][nx] = notes[mod(i+frameNum/4, len(notes))]
		}
	}

	// Draw keyboard outline
	if startY >= 0 && startY < h {
		for x := 0; x < w; x++ {
			grid[startY][x] = '='
		}
	}

	// White keys with black keys
	for y := startY + 1; y < h; y++ {
		for x := 0; x < w; x++ {
			isBlackCol := mod(x, 6) == 2 || mod(x, 6) == 4
			isBlackRow := y < startY+(keyH*2)/3

			if isBlackCol && isBlackRow {
				// Depressed/active key
				if mod(x+frameNum, 8) == 0 {
					grid[y][x] = '#'
				} else {
					grid[y][x] = '█'
				}
			} else {
				if mod(x, 6) == 0 {
					grid[y][x] = '|' // key separator
				} else {
					// Active white key ripple
					if mod(x-frameNum*2, 12) == 0 && y == h-1 {
						grid[y][x] = '-'
					} else {
						grid[y][x] = ' '
					}
				}
			}
		}
	}
}

func drawVinyl(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.45
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}
	if r < 3 {
		r = 3
	}

	rot := float64(frameNum) * 0.2

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				if dist < 1.0 {
					grid[y][x] = '+' // spindle center hole
				} else if dist < r*0.35 {
					// Center record label
					labelChars := []rune("HS AI LIVE")
					idx := mod(int(dist*4)+frameNum, len(labelChars))
					grid[y][x] = labelChars[idx]
				} else {
					// Concentric vinyl grooves with rotating light sheen
					angle := math.Atan2(dy, dx) + rot
					sheen := math.Sin(angle * 2)
					if math.Abs(sheen) > 0.8 {
						grid[y][x] = '*' // light reflection glint
					} else if mod(int(dist*3), 2) == 0 {
						grid[y][x] = '~'
					} else {
						grid[y][x] = '-'
					}
				}
			}
		}
	}

	// Turntable tonearm from top-right to record
	armEndX := int(cx + r*0.6/0.55)
	armEndY := int(cy - r*0.4)
	if armEndY >= 0 && armEndY < h && armEndX < w {
		grid[armEndY][armEndX] = '▼' // needle cartridge
	}
}

func drawCitySkyline(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	horizonY := int(float64(h) * 0.55)

	// Retro Synthwave Sun sinking on horizon
	sunR := float64(h) * 0.25
	sunCX := float64(w) / 2.0
	for y := 0; y < horizonY; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - sunCX) * 0.55
			dy := float64(y) - float64(horizonY)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist <= sunR {
				// Horizontal blind stripes across the sun
				if mod(y, 2) == 0 {
					grid[y][x] = '='
				}
			}
		}
	}

	// City skyscrapers skyline
	buildingWidth := 6
	numBuildings := w / buildingWidth
	for b := 0; b < numBuildings+1; b++ {
		bx := b * buildingWidth
		bSeed := (b*13 + seed) % 7
		bHeight := int(float64(horizonY) * (0.35 + float64(bSeed)*0.09))
		topY := horizonY - bHeight

		for y := topY; y < horizonY; y++ {
			if y >= 0 && y < h {
				for x := bx; x < bx+buildingWidth-1; x++ {
					if x < w {
						if x == bx || x == bx+buildingWidth-2 {
							grid[y][x] = '|'
						} else if y == topY {
							grid[y][x] = '_'
						} else {
							// Glowing windows in buildings
							if mod(x+y+frameNum/3+b, 3) == 0 {
								grid[y][x] = '□'
							} else {
								grid[y][x] = ' '
							}
						}
					}
				}
			}
		}
	}

	// 3D Perspective Grid ground moving towards camera
	for y := horizonY; y < h; y++ {
		// Horizontal perspective lines
		gridStep := (y - horizonY) * (y - horizonY)
		if mod(gridStep-frameNum*2, 14) < 2 {
			for x := 0; x < w; x++ {
				grid[y][x] = '='
			}
		}

		// Vanishing perspective vertical lines
		t := float64(y-horizonY) / float64(h-horizonY+1)
		for v := -5; v <= 5; v++ {
			vx := int(sunCX + float64(v)*(float64(w)*0.1)*t*2.0)
			if vx >= 0 && vx < w {
				if grid[y][vx] == '=' {
					grid[y][vx] = '+'
				} else {
					grid[y][vx] = '|'
				}
			}
		}
	}
}

func drawCoffee(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	// Rising steam wisps above cup
	steamH := 5
	steamY := cy - 3
	for s := 0; s < steamH; s++ {
		y := steamY - s
		if y >= 0 && y < h {
			offset1 := int(math.Sin(float64(s+frameNum)*0.4) * 2.0)
			offset2 := int(math.Cos(float64(s+frameNum)*0.4) * 2.0)
			if cx-3+offset1 >= 0 && cx-3+offset1 < w {
				grid[y][cx-3+offset1] = '('
			}
			if cx+3+offset2 >= 0 && cx+3+offset2 < w {
				grid[y][cx+3+offset2] = ')'
			}
			if cx+offset1 >= 0 && cx+offset1 < w {
				grid[y][cx+offset1] = '~'
			}
		}
	}

	// Coffee cup body
	cup := []string{
		"  .----------------.  ",
		" /      COFFEE      \\_ ",
		"|    HS AI ROAST    | )",
		"|                   | )",
		" \\                 /¯  ",
		"  '---------------'    ",
		"=======================",
	}

	startY := cy - 2
	for i, line := range cup {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}
}

func drawSports(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	// Parabolic ball trajectory
	ballCycle := w - 8
	if ballCycle < 1 {
		ballCycle = 10
	}
	ballX := mod(frameNum*3, ballCycle) + 4
	normX := (float64(ballX) - float64(w)/2.0) / (float64(w) / 2.0)
	ballY := int(float64(h)*0.25 + normX*normX*float64(h)*0.45)

	if ballY >= 0 && ballY < h && ballX >= 0 && ballX < w {
		grid[ballY][ballX] = '⚽'
	}

	// Ball motion trail
	for t := 1; t <= 3; t++ {
		tx := ballX - t*2
		tn := (float64(tx) - float64(w)/2.0) / (float64(w) / 2.0)
		ty := int(float64(h)*0.25 + tn*tn*float64(h)*0.45)
		if ty >= 0 && ty < h && tx >= 0 && tx < w {
			grid[ty][tx] = '.'
		}
	}

	// Goal post on the right
	goalX := w - 6
	goalTop := h / 3
	for y := goalTop; y < h-1; y++ {
		if goalX < w && y < h {
			grid[y][goalX] = '|'
			if goalX+3 < w {
				grid[y][goalX+3] = '|'
			}
			if mod(y, 2) == 0 && goalX+1 < w && goalX+2 < w {
				grid[y][goalX+1] = '#'
				grid[y][goalX+2] = '#'
			}
		}
	}
	for x := goalX; x <= goalX+3; x++ {
		if x < w && goalTop < h {
			grid[goalTop][x] = '='
		}
	}

	// Ground turf line
	if h-1 >= 0 {
		for x := 0; x < w; x++ {
			if mod(x+frameNum, 2) == 0 {
				grid[h-1][x] = '^'
			} else {
				grid[h-1][x] = '_'
			}
		}
	}
}

func drawTux(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	eye := 'o'
	if mod(frameNum, 8) == 0 {
		eye = '-' // blink
	}

	tux := []string{
		"     .--.     ",
		fmt.Sprintf("    |%c_ %c|    ", eye, eye),
		"    |:_/ |    ",
		"   //   \\ \\   ",
		"  (|     | )  ",
		" /'\\_   _/`\\  ",
		" \\___)=(___/  ",
	}

	startY := cy - len(tux)/2
	for i, line := range tux {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	// Iceberg platform beneath Tux
	iceY := startY + len(tux)
	if iceY < h {
		for x := cx - 12; x <= cx+12; x++ {
			if x >= 0 && x < w {
				grid[iceY][x] = '='
			}
		}
	}
}

func drawPokeball(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.42
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				if math.Abs(dist-r) < 0.5 {
					grid[y][x] = '#' // outer border
				} else if dist < r*0.25 {
					// Center button
					if dist < r*0.12 {
						if mod(frameNum, 4) < 2 {
							grid[y][x] = '★' // glowing button
						} else {
							grid[y][x] = 'o'
						}
					} else {
						grid[y][x] = 'O'
					}
				} else if math.Abs(dy) < 0.8 {
					grid[y][x] = '=' // dividing center belt
				} else if dy < 0 {
					grid[y][x] = '█' // red top half
				} else {
					grid[y][x] = ' ' // white bottom half
				}
			}
		}
	}

	// Energy capture spark particles when button flashes
	if mod(frameNum, 4) < 2 {
		for i := 0; i < 8; i++ {
			angle := float64(i) * math.Pi / 4.0
			sx := int(cx + (r+2.0)*math.Cos(angle)/0.55)
			sy := int(cy + (r+2.0)*math.Sin(angle))
			if sy >= 0 && sy < h && sx >= 0 && sx < w {
				grid[sy][sx] = '✦'
			}
		}
	}
}

func drawTornado(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0

	// Swaying funnel center
	sway := math.Sin(float64(frameNum)*0.2) * 5.0

	for y := 0; y < h; y++ {
		// Funnel widens from ground to sky
		t := float64(h-1-y) / float64(h)
		radius := 2.0 + t*float64(w)*0.22
		fcx := cx + sway*t

		left := int(fcx - radius)
		right := int(fcx + radius)

		for x := left; x <= right; x++ {
			if x >= 0 && x < w {
				swirl := mod(x+y+frameNum*3, 4)
				if x == left {
					grid[y][x] = '/'
				} else if x == right {
					grid[y][x] = '\\'
				} else if swirl == 0 {
					grid[y][x] = '@'
				} else if swirl == 1 {
					grid[y][x] = '%'
				} else {
					grid[y][x] = '~'
				}
			}
		}
	}

	// Flying debris in the air
	for i := 0; i < 15; i++ {
		dx := mod(int(cx)+i*seed*3+frameNum*5, w)
		dy := mod(i*seed*7+frameNum*2, h)
		if dy < h && dx < w {
			grid[dy][dx] = '*'
		}
	}
}

func drawDiscoBall(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) * 0.38
	r := float64(h) * 0.28
	if float64(w)*0.18 < r {
		r = float64(w) * 0.18
	}

	// Hanging chain from ceiling
	for y := 0; y < int(cy-r); y++ {
		if y < h && int(cx) < w {
			grid[y][int(cx)] = '|'
		}
	}

	facets := []rune("░▒▓█#*+")

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				angle := math.Atan2(dy, dx) + float64(frameNum)*0.2
				facetIdx := mod(int(angle*4)+int(dist*3), len(facets))
				grid[y][x] = facets[facetIdx]
			}
		}
	}

	// Radiating dance party light beams
	numBeams := 8
	for b := 0; b < numBeams; b++ {
		bAngle := float64(b)*(2*math.Pi/float64(numBeams)) + float64(frameNum)*0.12
		for dist := r + 1.0; dist < float64(h)*0.8; dist += 1.0 {
			bx := int(cx + dist*math.Cos(bAngle)/0.55)
			by := int(cy + dist*math.Sin(bAngle))
			if by >= 0 && by < h && bx >= 0 && bx < w {
				grid[by][bx] = '·'
			}
		}
	}
}
