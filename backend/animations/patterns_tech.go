package animations

import (
	"fmt"
	"math"
)

func drawMatrix(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	chars := []rune("0123456789ABCDEF$#@%&*+-/=")

	for x := 0; x < w; x++ {
		colSpeed := mod(x*7+seed, 3) + 1
		colOffset := mod(x*13, h)
		headY := mod(frameNum*colSpeed+colOffset, h+8)

		for trail := 0; trail < 7; trail++ {
			y := headY - trail
			if y >= 0 && y < h {
				if trail == 0 {
					grid[y][x] = chars[mod(frameNum+x, len(chars))]
				} else if trail < 3 {
					grid[y][x] = chars[mod(x*3+trail+frameNum, len(chars))]
				} else if trail < 5 {
					grid[y][x] = '|'
				} else {
					grid[y][x] = ':'
				}
			}
		}
	}
}

func drawEqualizer(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	numBars := w / 3
	if numBars < 2 {
		numBars = 2
	}

	for b := 0; b < numBars; b++ {
		bx := b * 3
		freq := float64(b+1) * 0.4
		phase := float64(frameNum)*0.25 + float64(seed)*0.1
		val := math.Sin(phase*freq) + math.Cos(phase*0.7)
		barHeight := int((val + 2.0) / 4.0 * float64(h-2))
		if barHeight < 1 {
			barHeight = 1
		}
		if barHeight >= h {
			barHeight = h - 1
		}

		for y := h - 1; y >= h-barHeight; y-- {
			if y >= 0 && bx < w {
				grid[y][bx] = '█'
			}
			if y >= 0 && bx+1 < w {
				grid[y][bx+1] = '█'
			}
		}

		peakY := h - barHeight - 1
		if peakY >= 0 && bx < w {
			grid[peakY][bx] = '¯'
		}
		if peakY >= 0 && bx+1 < w {
			grid[peakY][bx+1] = '¯'
		}
	}
}

func drawRocket(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	rocket := []string{
		"   /\\   ",
		"  /  \\  ",
		"  |  |  ",
		"  |HS|  ",
		" /|  |\\ ",
		"/_|__|_\\",
	}

	flames := [][]string{
		{
			"  /||\\  ",
			"   \\/   ",
			"   **   ",
		},
		{
			"  ||||  ",
			"  \\  /  ",
			"   vv   ",
		},
		{
			"  \\||/  ",
			"   ||   ",
			"   ..   ",
		},
	}

	startY := h/2 - len(rocket)/2 - 1
	if startY < 0 {
		startY = 0
	}

	for i, line := range rocket {
		y := startY + i
		if y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	flame := flames[mod(frameNum, len(flames))]
	flameStartY := startY + len(rocket)
	for i, line := range flame {
		y := flameStartY + i
		if y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	for i := 0; i < 8; i++ {
		sx := mod(seed*11+i*17, w)
		sy := mod(frameNum*3+i*7, h)
		if sy >= 0 && sy < h && sx >= 0 && sx < w {
			grid[sy][sx] = '|'
		}
	}
}

func drawNeuralNetwork(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	layers := []int{2, 4, 4, 2}
	layerX := make([]int, len(layers))
	spacing := w / (len(layers) + 1)
	if spacing < 1 {
		spacing = 1
	}
	for i := range layers {
		layerX[i] = spacing * (i + 1)
	}

	type nodePos struct{ x, y int }
	allNodes := make([][]nodePos, len(layers))

	for l, count := range layers {
		allNodes[l] = make([]nodePos, count)
		nodeSpacing := h / (count + 1)
		if nodeSpacing < 1 {
			nodeSpacing = 1
		}
		for n := 0; n < count; n++ {
			ny := nodeSpacing * (n + 1)
			allNodes[l][n] = nodePos{layerX[l], ny}
			if ny < h && layerX[l] < w {
				pulse := mod(frameNum+l+n, 3)
				if pulse == 0 {
					grid[ny][layerX[l]] = '●'
				} else {
					grid[ny][layerX[l]] = '○'
				}
			}
		}
	}

	for l := 0; l < len(layers)-1; l++ {
		for _, n1 := range allNodes[l] {
			for _, n2 := range allNodes[l+1] {
				midX := (n1.x + n2.x) / 2
				midY := (n1.y + n2.y) / 2
				if midY >= 0 && midY < h && midX >= 0 && midX < w {
					if mod(frameNum+l, 2) == 0 {
						grid[midY][midX] = '·'
					} else {
						grid[midY][midX] = '-'
					}
				}
			}
		}
	}
}

func drawRadar(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	radius := float64(h) * 0.45
	if float64(w)*0.25 < radius {
		radius = float64(w) * 0.25
	}
	if radius < 3 {
		radius = 3
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if math.Abs(dist-radius) < 0.6 || math.Abs(dist-radius*0.5) < 0.6 {
				grid[y][x] = '.'
			}
		}
	}

	for x := int(cx - radius/0.55); x <= int(cx+radius/0.55); x++ {
		if x >= 0 && x < w && int(cy) >= 0 && int(cy) < h {
			grid[int(cy)][x] = '-'
		}
	}
	for y := int(cy - radius); y <= int(cy+radius); y++ {
		if y >= 0 && y < h && int(cx) >= 0 && int(cx) < w {
			grid[y][int(cx)] = '|'
		}
	}

	angle := float64(frameNum) * 0.2
	for r := 0.0; r < radius; r += 0.5 {
		sx := int(cx + (r*math.Cos(angle))/0.55)
		sy := int(cy + r*math.Sin(angle))
		if sy >= 0 && sy < h && sx >= 0 && sx < w {
			grid[sy][sx] = '#'
		}
	}

	blips := [][2]int{
		{int(cx + (radius*0.6*math.Cos(1.2))/0.55), int(cy + radius*0.6*math.Sin(1.2))},
		{int(cx + (radius*0.8*math.Cos(3.4))/0.55), int(cy + radius*0.8*math.Sin(3.4))},
	}
	for _, b := range blips {
		if b[1] >= 0 && b[1] < h && b[0] >= 0 && b[0] < w {
			grid[b[1]][b[0]] = 'X'
		}
	}
}

func drawBlackhole(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	rEH := 2.5

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.5
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < rEH {
				grid[y][x] = ' '
			} else if dist < rEH+0.8 {
				grid[y][x] = '#'
			} else if dist < rEH*3.5 {
				angle := math.Atan2(dy, dx) + float64(frameNum)*0.25 - dist*0.8
				spiral := math.Sin(angle * 3)
				if spiral > 0.4 {
					grid[y][x] = '='
				} else if spiral > -0.2 {
					grid[y][x] = '~'
				} else {
					grid[y][x] = '.'
				}
			}
		}
	}
}

func drawVehicle(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	car := []string{
		"   .--------.   ",
		"  /  _    _  \\  ",
		" -o-(_)--(_)-o- ",
	}

	startY := cy - len(car)/2
	for i, line := range car {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					if r == '_' && i == 2 {
						spokes := []rune{'|', '/', '-', '\\'}
						r = spokes[mod(frameNum, len(spokes))]
					}
					grid[y][x] = r
				}
			}
		}
	}

	roadY := startY + len(car)
	if roadY < h {
		for x := 0; x < w; x++ {
			grid[roadY][x] = '='
		}
	}
	subRoadY := roadY + 1
	if subRoadY < h {
		for x := 0; x < w; x++ {
			if mod(x+frameNum*3, 4) == 0 {
				grid[subRoadY][x] = '-'
			}
		}
	}

	smokeY := startY + 2
	for s := 1; s <= 4; s++ {
		sx := cx - len(car[0])/2 - s*2 - mod(frameNum, 2)
		if sx >= 0 && sx < w && smokeY < h {
			grid[smokeY][sx] = 'o'
		}
	}
}

func drawHsAiLogo(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	status := ">>> ONLINE <<<"
	if mod(frameNum, 4) < 2 {
		status = ">>> ACTIVE <<<"
	}

	logo := []string{
		"+---------------------+",
		"|   HS AI SOLUTIONS   |",
		"|    .----. .----.    |",
		"|   /  H   \\  S   \\   |",
		"|   \\      /      /   |",
		"|    '----' '----'    |",
		fmt.Sprintf("|   %14s    |", status),
		"+---------------------+",
	}

	startY := cy - len(logo)/2
	for i, line := range logo {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w {
					grid[y][x] = r
				}
			}
		}
	}

	for y := 0; y < h; y++ {
		if mod(y+frameNum, 3) == 0 {
			if w > 2 {
				grid[y][1] = '0' + rune(mod(y+frameNum, 2))
				grid[y][w-2] = '0' + rune(mod(y+frameNum+1, 2))
			}
		}
	}
}
