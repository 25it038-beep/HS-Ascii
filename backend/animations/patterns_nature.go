package animations

import (
	"fmt"
	"math"
	"strings"
)

func drawHeart(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h)/2.0 - 1.0

	pulse := 1.0 + 0.15*math.Sin(float64(frameNum)*0.4)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			nx := (float64(x) - cx) / (pulse * float64(w) * 0.18)
			ny := -(float64(y) - cy) / (pulse * float64(h) * 0.28)

			val := math.Pow(nx*nx+ny*ny-1.0, 3.0) - nx*nx*ny*ny*ny
			if val <= 0.0 {
				if strings.Contains(animName, "broken") {
					// Jagged lightning fracture down the center
					crackOffset := int(math.Sin(float64(y)*1.5) * 2.0)
					if x == int(cx)+crackOffset || x == int(cx)+crackOffset+1 {
						continue // gap in the heart
					}
				}

				if strings.Contains(animName, "fire") && y < int(cy) {
					if mod(x+frameNum, 2) == 0 {
						grid[y][x] = '▲' // fire heart flames
					} else {
						grid[y][x] = '*'
					}
				} else if math.Abs(val) < 0.2 {
					grid[y][x] = '♥'
				} else {
					grid[y][x] = '*'
				}
			}
		}
	}
}

func drawRain(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	dropChar := '/'
	if strings.Contains(animName, "deluge") || strings.Contains(animName, "monsoon") {
		dropChar = '|'
	} else if strings.Contains(animName, "hail") {
		dropChar = 'o'
	}

	for x := 0; x < w; x++ {
		speed := mod(x*3+seed, 3) + 1
		y := mod(frameNum*speed+(x*7+seed*13), h)
		if y >= 0 && y < h {
			grid[y][x] = dropChar
		}
		if y == h-1 && x+1 < w {
			grid[h-1][x] = '.'
		}
	}
}

func drawFire(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	chars := []rune(" .:*#▲█")

	for y := 0; y < h; y++ {
		invY := h - 1 - y
		for x := 0; x < w; x++ {
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

func drawWaves(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	midY := float64(h) * 0.6

	chars := []rune("~-^≈~")
	if strings.Contains(animName, "bio") || strings.Contains(animName, "sea") {
		chars = []rune("~*°≈~")
	}

	waveSpeed := 0.15 + float64(seed%3)*0.05
	waveAmp := 1.5 + float64(seed%3)*0.5

	for layer := 0; layer < 3; layer++ {
		phase := float64(frameNum)*waveSpeed + float64(layer)*1.2 + float64(seed)*0.3
		amp := waveAmp + float64(layer)*0.5

		for x := 0; x < w; x++ {
			y := int(midY + math.Sin(float64(x)*0.25+phase)*amp + float64(layer*2))
			if y >= 0 && y < h {
				grid[y][x] = chars[mod(layer+x+seed, len(chars))]
			}
		}
	}
}

func drawParticles(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	count := (w * h) / 10
	if count < 10 {
		count = 10
	}

	chars := []rune(".+*°•")
	if strings.Contains(animName, "snow") || strings.Contains(animName, "frost") || strings.Contains(animName, "blizzard") {
		chars = []rune("*+❄°.")
	} else if strings.Contains(animName, "desert") || strings.Contains(animName, "sand") {
		chars = []rune(".,·:~")
	}

	for i := 0; i < count; i++ {
		px := mod(frameNum*2+i*seed*3+i*7, w)
		py := mod(frameNum+i*seed*5+i*13, h)
		if py < h && px < w {
			grid[py][px] = chars[mod(i+seed, len(chars))]
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

	angle := float64(frameNum)*0.2 + float64(seed)*0.1
	for i := 0; i < 8; i++ {
		a := angle + float64(i)*math.Pi/4.0
		x := int(cx + r*math.Cos(a)/0.5)
		y := int(cy + r*math.Sin(a))
		if x >= 0 && x < w && y >= 0 && y < h {
			grid[y][x] = '#'
		}
	}
}

func drawStars(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	twinkle := []rune(".+*✦★")

	// Starfield with unique layout per seed
	for i := 0; i < 45; i++ {
		sx := mod(i*29+seed*17, w)
		sy := mod(i*19+seed*11, h)
		t := mod(frameNum+i+seed, len(twinkle))
		setCell(grid, sx, sy, twinkle[t])
	}

	// For constellation animations: connect a constellation
	if strings.Contains(animName, "constellation") {
		constellation := [][2]int{
			{w/2 - 12, h/2 - 3},
			{w/2 - 4, h/2 - 1},
			{w / 2, h / 2},
			{w/2 + 8, h/2 - 2},
			{w/2 + 14, h/2 + 3},
		}
		for i, pt := range constellation {
			setCell(grid, pt[0], pt[1], '★')
			if i > 0 {
				prev := constellation[i-1]
				drawLine(grid, prev[0], prev[1], pt[0], pt[1], '-')
			}
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

	for a := 0.0; a < 2*math.Pi; a += 0.2 {
		px := int(cx + r*math.Cos(a)/0.55)
		py := int(cy + r*math.Sin(a))
		if px >= 0 && px < w && py >= 0 && py < h {
			grid[py][px] = 'o'
		}
	}

	setCell(grid, int(cx), int(cy), '+')

	secAngle := float64(frameNum%60)/60.0*2*math.Pi - math.Pi/2
	for d := 1.0; d < r; d += 0.8 {
		sx := int(cx + d*math.Cos(secAngle)/0.55)
		sy := int(cy + d*math.Sin(secAngle))
		setCell(grid, sx, sy, '.')
	}

	timeStr := fmt.Sprintf("[%02d:%02d]", (frameNum/60)%24, frameNum%60)
	startX := int(cx) - len(timeStr)/2
	for i, c := range timeStr {
		setCell(grid, startX+i, int(cy)+int(r)+1, c)
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

func drawWings(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	wingPhase := mod(frameNum, 3)

	if strings.Contains(animName, "dragon") {
		// Dragon with horns and breathing fire
		fire := "   "
		if mod(frameNum, 2) == 1 {
			fire = "▲*▲"
		}
		dragon := []string{
			"    /\\____/\\     ",
			"   /  o  o  \\    ",
			fmt.Sprintf("  (    ==    ) %s", fire), // breathing fire
			"   \\        /    ",
			"   / /----\\ \\    ",
		}
		startY := cy - len(dragon)/2
		for i, line := range dragon {
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
		return
	}

	if strings.Contains(animName, "bat") {
		// Bat with pointy ears and fangs
		bat := []string{
			"   /\\       /\\   ",
			"  /  \\ (..) /  \\ ",
			" / /\\ \\/vv\\/ /\\ \\",
		}
		startY := cy - len(bat)/2
		for i, line := range bat {
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
		return
	}

	if strings.Contains(animName, "butterfly") {
		// Fluttering monarch butterfly
		flut := mod(frameNum, 2)
		wing1 := " /|\\   /|\\ "
		wing2 := " \\_/   \\_/ "
		if flut == 1 {
			wing1 = "  |     |  "
			wing2 = "  \\     /  "
		}
		bfly := []string{
			"    \\   /    ",
			wing1,
			"     (o)     ",
			wing2,
		}
		startY := cy - len(bfly)/2
		for i, line := range bfly {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "phoenix") {
		// Phoenix rising from ashes
		sparks := []rune{'*', '^', '✦', '+', '~'}
		for s := 0; s < 8; s++ {
			sx := cx - 10 + mod(s*7+frameNum*3, 20)
			sy := cy + 3 - mod(s*3+frameNum, 8)
			setCell(grid, sx, sy, sparks[mod(s+frameNum, len(sparks))])
		}
		phoenix := []string{
			"  \\/\\    /\\/  ",
			"  /  \\  /  \\  ",
			" (    \\/    ) ",
			"  \\   (o)  /  ",
			"   \\/====\\/   ",
		}
		startY := cy - len(phoenix)/2
		for i, line := range phoenix {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "flappy") {
		// Flappy Bird and obstacle pipes
		birdY := cy + (mod(frameNum, 4) - 2)
		bird := []string{
			" .--. ",
			"( >o )===",
			" '--' ",
		}
		for i, line := range bird {
			y := birdY + i
			startX := cx - 8
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		// Green pipes scrolling left
		pipeX := mod(w-frameNum*2, w)
		for y := 0; y < h; y++ {
			if y < cy-3 || y > cy+3 {
				setCell(grid, pipeX, y, '[')
				setCell(grid, pipeX+1, y, '#')
				setCell(grid, pipeX+2, y, ']')
			}
		}
		return
	}

	if strings.Contains(animName, "eagle") {
		eagle := []string{
			"  \\______/\\______/  ",
			"    \\   ( >o)   /   ", // hooked beak
			"     \\___  ___/     ",
			"         \\/         ",
		}
		startY := cy - len(eagle)/2
		for i, line := range eagle {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

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

func drawSymbol(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	icx := w / 2
	icy := h / 2

	if strings.Contains(animName, "triforce") {
		// Zelda Golden Triforce
		triforce := []string{
			"      /\\      ",
			"     /__\\     ",
			"    /\\  /\\    ",
			"   /__\\/__\\   ",
		}
		startY := icy - len(triforce)/2
		for i, line := range triforce {
			y := startY + i
			startX := icx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "lambda") {
		// Half-Life Gordon Freeman Lambda
		lambda := []string{
			"     /\\     ",
			"    /  \\    ",
			"   / /\\ \\   ",
			"  / /  \\ \\  ",
			" / /    \\ \\ ",
			"/_/      \\_\\",
		}
		startY := icy - len(lambda)/2
		for i, line := range lambda {
			y := startY + i
			startX := icx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "hazard") || strings.Contains(animName, "radioactive") {
		hazard := []string{
			"     /\\     ",
			"    /  \\    ",
			"   / /\\ \\   ",
			"  / /! \\ \\  ",
			" / /____\\ \\ ",
			"/__________\\",
		}
		startY := icy - len(hazard)/2
		for i, line := range hazard {
			y := startY + i
			startX := icx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "infinity") {
		// Infinity Loop Lemniscate
		for t := 0.0; t < 2*math.Pi; t += 0.1 {
			scale := float64(w) * 0.35
			denom := 1.0 + math.Sin(t)*math.Sin(t)
			x := int(cx + (scale*math.Cos(t)/denom)/0.55)
			y := int(cy + (scale*0.45*math.Sin(t)*math.Cos(t))/denom)
			setCell(grid, x, y, '∞')
		}
		return
	}

	r := float64(h) * 0.4
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}

	rot := float64(frameNum)*0.15 + float64(seed)*0.2

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				angle := math.Atan2(dy, dx) + rot
				if math.Sin(angle) > 0 {
					setCell(grid, x, y, '☯')
				} else {
					setCell(grid, x, y, '○')
				}
			}
		}
	}
}
