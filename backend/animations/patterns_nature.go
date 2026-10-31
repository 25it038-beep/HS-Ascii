package animations

import (
	"fmt"
	"math"
)

func drawHeart(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h)/2.0 - 1.0

	// Heartbeat pulse animation
	pulse := 1.0 + 0.15*math.Sin(float64(frameNum)*0.4)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			nx := (float64(x) - cx) / (pulse * float64(w) * 0.18)
			ny := -(float64(y) - cy) / (pulse * float64(h) * 0.28)

			// Standard heart equation: (x^2 + y^2 - 1)^3 - x^2 * y^3 <= 0
			val := math.Pow(nx*nx+ny*ny-1.0, 3.0) - nx*nx*ny*ny*ny
			if val <= 0.0 {
				if math.Abs(val) < 0.2 {
					grid[y][x] = '♥'
				} else {
					grid[y][x] = '*'
				}
			}
		}
	}
}

func drawRain(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for x := 0; x < w; x++ {
		speed := mod(x*3+seed, 2) + 1
		y := mod(frameNum*speed+(x*7), h)
		if y >= 0 && y < h {
			grid[y][x] = '/'
		}
		// Splash on bottom
		if y == h-1 && x+1 < w {
			grid[h-1][x] = '.'
		}
	}
}

func drawFire(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	chars := []rune(" .:*#▲█")

	for y := 0; y < h; y++ {
		invY := h - 1 - y
		for x := 0; x < w; x++ {
			// Perlin-style flame flicker
			flame := math.Sin(float64(x)*0.3+float64(frameNum)*0.4) +
				math.Cos(float64(y)*0.4-float64(frameNum)*0.6) +
				math.Sin(float64(x+y+seed)*0.2)

			heat := (flame + 2.5) / 5.0 * float64(invY) / float64(h)
			idx := int(heat * float64(len(chars)))
			if idx >= len(chars) {
				idx = len(chars) - 1
			}
			if idx > 0 {
				grid[y][x] = chars[idx]
			}
		}
	}
}

func drawWaves(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	midY := float64(h) * 0.6

	chars := []rune("~-^≈~")

	for layer := 0; layer < 3; layer++ {
		phase := float64(frameNum)*0.15 + float64(layer)*1.2
		amp := 1.5 + float64(layer)*0.5

		for x := 0; x < w; x++ {
			y := int(midY + math.Sin(float64(x)*0.25+phase)*amp + float64(layer*2))
			if y >= 0 && y < h {
				grid[y][x] = chars[mod(layer+x, len(chars))]
			}
		}
	}
}

func drawParticles(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	count := (w * h) / 10
	if count < 10 {
		count = 10
	}

	chars := []rune(".+*°•")

	for i := 0; i < count; i++ {
		px := mod(frameNum*2+i*seed*3+i*7, w)
		py := mod(frameNum+i*seed*5+i*13, h)
		if py < h && px < w {
			grid[py][px] = chars[mod(i, len(chars))]
		}
	}
}

func drawRotate(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.4
	if float64(w)*0.2 < r {
		r = float64(w) * 0.2
	}

	angle := float64(frameNum) * 0.2
	for i := 0; i < 8; i++ {
		a := angle + float64(i)*math.Pi/4.0
		x := int(cx + r*math.Cos(a)/0.5)
		y := int(cy + r*math.Sin(a))
		if x >= 0 && x < w && y >= 0 && y < h {
			grid[y][x] = '#'
		}
	}
}

func drawStars(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	twinkle := []rune(".+*✦★")

	for i := 0; i < 40; i++ {
		sx := mod(i*23+seed*7, w)
		sy := mod(i*17+seed*13, h)
		t := mod(frameNum+i, len(twinkle))
		if sy < h && sx < w {
			grid[sy][sx] = twinkle[t]
		}
	}
}

func drawClock(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.45
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}

	// Clock circular rim
	for a := 0.0; a < 2*math.Pi; a += 0.2 {
		px := int(cx + r*math.Cos(a)/0.55)
		py := int(cy + r*math.Sin(a))
		if px >= 0 && px < w && py >= 0 && py < h {
			grid[py][px] = 'o'
		}
	}

	// Center pivot
	if int(cy) < h && int(cx) < w {
		grid[int(cy)][int(cx)] = '+'
	}

	// Seconds hand
	secAngle := float64(frameNum%60)/60.0*2*math.Pi - math.Pi/2
	for d := 1.0; d < r; d += 0.8 {
		sx := int(cx + d*math.Cos(secAngle)/0.55)
		sy := int(cy + d*math.Sin(secAngle))
		if sx >= 0 && sx < w && sy >= 0 && sy < h {
			grid[sy][sx] = '.'
		}
	}

	// Digital time banner at top
	timeStr := fmt.Sprintf("[%02d:%02d]", (frameNum/60)%24, frameNum%60)
	startX := int(cx) - len(timeStr)/2
	if startX >= 0 && int(cy)+int(r)+1 < h {
		for i, c := range timeStr {
			grid[int(cy)+int(r)+1][startX+i] = c
		}
	}
}

func drawBounce(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	if w <= 2 || h <= 2 {
		if h > 0 && w > 0 {
			grid[0][0] = 'O'
		}
		return
	}
	periodX := (w - 2) * 2
	periodY := (h - 2) * 2

	stepX := mod(frameNum+seed, periodX)
	x := 1 + stepX
	if stepX >= w-2 {
		x = w - 2 - (stepX - (w - 2))
	}

	stepY := mod(frameNum*2+seed, periodY)
	y := 1 + stepY
	if stepY >= h-2 {
		y = h - 2 - (stepY - (h - 2))
	}

	if y >= 0 && y < h && x >= 0 && x < w {
		grid[y][x] = 'O'
	}
}

func drawWings(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	wingPhase := mod(frameNum, 3)
	var wings []string
	switch wingPhase {
	case 0:
		wings = []string{
			" \\     ^     / ",
			"  \\   (o)   /  ",
			"   '--' '--'   ",
		}
	case 1:
		wings = []string{
			"               ",
			" ---- (o) ---- ",
			"   '--' '--'   ",
		}
	default:
		wings = []string{
			"               ",
			"   .--. .--.   ",
			"  /   (o)   \\  ",
			" /     v     \\ ",
		}
	}

	startY := cy - len(wings)/2
	for i, line := range wings {
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

func drawSymbol(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.4
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}

	rot := float64(frameNum) * 0.15

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				angle := math.Atan2(dy, dx) + rot
				if math.Sin(angle) > 0 {
					grid[y][x] = '☯'
				} else {
					grid[y][x] = '○'
				}
			}
		}
	}
}
