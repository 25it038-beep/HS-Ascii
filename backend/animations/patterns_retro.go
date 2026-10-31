package animations

import (
	"fmt"
	"strings"
)

func drawPacman(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2

	period := w + 12
	if period < 1 {
		period = 20
	}
	px := mod(frameNum*2, period) - 6

	for x := 0; x < w; x += 3 {
		if x > px+2 && cy < h && x < w {
			grid[cy][x] = '·'
		}
	}

	mouthOpen := mod(frameNum, 4) < 2

	if px >= 0 && px < w && cy < h {
		if mouthOpen {
			grid[cy][px] = 'C'
		} else {
			grid[cy][px] = 'O'
		}
		if cy > 0 && px >= 0 && px < w {
			grid[cy-1][px] = '-'
		}
		if cy+1 < h && px >= 0 && px < w {
			grid[cy+1][px] = '-'
		}
	}

	gx := px - 6
	if gx >= 0 && gx < w && cy < h {
		grid[cy][gx] = 'M'
		if cy > 0 && gx < w {
			grid[cy-1][gx] = 'n'
		}
	}
}

func drawCat(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "pokeball") {
		ball := []string{
			"    .---.    ",
			"   / (o) \\   ",
			"  |===O===|  [GOTCHA!]",
			"   \\     /   ",
			"    '---'    ",
		}
		startY := cy - len(ball)/2
		for i, line := range ball {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	tailChar := '/'
	switch mod(frameNum, 4) {
	case 1:
		tailChar = '~'
	case 2:
		tailChar = '\\'
	case 3:
		tailChar = '|'
	}

	eyeChar := 'o'
	if mod(frameNum, 6) == 0 {
		eyeChar = '-'
	}

	cat := []string{
		"  /\\_/\\  ",
		fmt.Sprintf(" ( %c.%c ) ", eyeChar, eyeChar),
		"  > ^ <  ",
		fmt.Sprintf("  /| |\\%c ", tailChar),
	}

	startY := cy - len(cat)/2
	for i, line := range cat {
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

func drawDancer(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	var poses [][]string

	if strings.Contains(animName, "rick") {
		lyrics := []string{
			"♪ Never gonna give you up ♪",
			"♪ Never gonna let you down ♪",
			"♪ Never gonna run around ♪",
			"♪ and desert you! ♪",
		}
		lyric := lyrics[mod(frameNum/4, len(lyrics))]
		lx := cx - len(lyric)/2
		ly := cy - 4
		if ly >= 0 && ly < h {
			for j, r := range lyric {
				x := lx + j
				if x >= 0 && x < w {
					grid[ly][x] = r
				}
			}
		}
		poses = [][]string{
			{
				"  [o]   ",
				" ~(|)=o ",
				"  / \\   ",
			},
			{
				"  [o]   ",
				" o=(|)~ ",
				"  / \\   ",
			},
			{
				"  [o]   ",
				" \\(|)/  ",
				"  < >   ",
			},
		}
	} else if strings.Contains(animName, "naruto") {
		poses = [][]string{
			{
				"   _o__  ",
				" =/|     ",
				"  / \\    ",
			},
			{
				"   _o__  ",
				" ==/     ",
				"  < >    ",
			},
		}
	} else {
		poses = [][]string{
			{
				"   \\o/   ",
				"    |    ",
				"   / \\   ",
			},
			{
				"    o/   ",
				"   /|    ",
				"   / \\   ",
			},
			{
				"   \\o    ",
				"    |\\   ",
				"   / \\   ",
			},
			{
				"   _o_   ",
				"    |    ",
				"   < >   ",
			},
		}
	}

	pose := poses[mod(frameNum+seed, len(poses))]
	startY := cy - len(pose)/2

	if startY+len(pose)+1 < h {
		floorY := startY + len(pose) + 1
		for x := 0; x < w; x++ {
			if mod(x+frameNum+seed, 2) == 0 {
				grid[floorY][x] = '='
			} else {
				grid[floorY][x] = '-'
			}
		}
	}

	for i, line := range pose {
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

	if strings.Contains(animName, "disco") {
		sparkles := []rune{'*', '+', '.', 'o', 'x'}
		for s := 0; s < 6; s++ {
			sx := mod(s*17+frameNum*3+seed*7, w)
			sy := mod(s*7+frameNum+seed*3, h)
			if grid[sy][sx] == ' ' {
				grid[sy][sx] = sparkles[mod(frameNum+s, len(sparkles))]
			}
		}
	}
}

func drawGhost(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	floatY := cy + (mod(frameNum+seed, 3) - 1)
	floatX := cx + (mod(frameNum+seed*2, 5) - 2)

	var ghost []string
	if strings.Contains(animName, "alien") {
		ghost = []string{
			" .---. ",
			"( O.O )",
			" \\_!_/ ",
			"  / \\  ",
		}
	} else if strings.Contains(animName, "pacman") {
		eyes := "o o"
		if mod(frameNum, 4) < 2 {
			eyes = "O O"
		}
		ghost = []string{
			" .---. ",
			fmt.Sprintf("| %s |", eyes),
			"|     |",
			"~V~V~V~",
		}
	} else {
		ghost = []string{
			"  .-.  ",
			" (o o) ",
			" | O | ",
			" ~~~~~ ",
		}
	}

	for i, line := range ghost {
		y := floatY - len(ghost)/2 + i
		if y >= 0 && y < h {
			startX := floatX - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}
}

func drawSpaceInvaders(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	shiftX := 0
	if w > 20 {
		shiftX = mod(frameNum*2, w-18)
	}

	wing := "\\"
	if mod(frameNum, 2) == 1 {
		wing = "/"
	}

	invaders := []string{
		fmt.Sprintf(" [o_o]  [o_o]  [o_o] "),
		fmt.Sprintf("  %s %s    %s %s    %s %s  ", wing, wing, wing, wing, wing, wing),
	}

	for i, line := range invaders {
		y := 1 + i
		if y < h {
			for j, r := range line {
				x := shiftX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	if h-2 >= 0 {
		cannonX := (w / 2) + mod(frameNum*3, 7) - 3
		ship := "_/\\_"
		for j, r := range ship {
			x := cannonX + j
			if x >= 0 && x < w {
				grid[h-2][x] = r
			}
		}
		if h-4 > 0 {
			missileY := h - 3 - mod(frameNum, h-4)
			if missileY >= 0 && cannonX+1 >= 0 && cannonX+1 < w {
				grid[missileY][cannonX+1] = '|'
			}
		}
	}
}

func drawPong(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	if h <= 4 || w <= 6 {
		return
	}

	midX := w / 2
	for y := 0; y < h; y++ {
		if y%2 == 0 {
			grid[y][midX] = '|'
		}
	}

	p1Y := mod(frameNum+seed, h-3)
	if p1Y < 1 {
		p1Y = 1
	}
	for i := 0; i < 3; i++ {
		if p1Y+i < h && w > 2 {
			grid[p1Y+i][2] = '█'
		}
	}

	p2Y := mod(frameNum*2+seed, h-3)
	if p2Y < 1 {
		p2Y = 1
	}
	for i := 0; i < 3; i++ {
		if p2Y+i < h && w > 3 {
			grid[p2Y+i][w-3] = '█'
		}
	}

	ballCycle := (w - 6) * 2
	if ballCycle < 1 {
		ballCycle = 10
	}
	ballStep := mod(frameNum, ballCycle)
	ballX := 3 + ballStep
	if ballStep >= w-6 {
		ballX = w - 3 - (ballStep - (w - 6))
	}
	ballY := mod(frameNum*3, h)
	if ballY >= 0 && ballY < h && ballX >= 0 && ballX < w {
		grid[ballY][ballX] = 'O'
	}
}

func drawTetris(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	left := w/2 - 6
	right := w/2 + 6
	if left < 0 {
		left = 0
	}
	if right >= w {
		right = w - 1
	}

	for y := 0; y < h; y++ {
		grid[y][left] = '<'
		grid[y][right] = '>'
	}
	for x := left; x <= right; x++ {
		grid[h-1][x] = '='
	}

	for x := left + 1; x < right; x++ {
		if mod(x+seed, 3) != 0 && h-2 >= 0 {
			grid[h-2][x] = '#'
		}
		if mod(x+seed, 2) == 0 && h-3 >= 0 {
			grid[h-3][x] = '#'
		}
	}

	if h-4 > 0 {
		pieceY := mod(frameNum*2, h-4)
		pieceX := left + 4
		if pieceY >= 0 && pieceY < h-1 && pieceX-1 >= left && pieceX+1 <= right {
			grid[pieceY][pieceX-1] = '['
			grid[pieceY][pieceX] = ']'
			grid[pieceY][pieceX+1] = '['
			if pieceY+1 < h-1 {
				grid[pieceY+1][pieceX] = ']'
			}
		}
	}
}

func drawSnake(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if w <= 6 || h <= 6 {
		return
	}

	foodX := mod(seed*7+5, w-4) + 2
	foodY := mod(seed*3+2, h-4) + 2
	if foodY < h && foodX < w {
		grid[foodY][foodX] = '*'
	}

	length := 6
	for i := 0; i < length; i++ {
		seg := frameNum - i
		segX := mod(cx+seg*2, w-2)
		segY := mod(cy+(seg/3)%3, h-2)
		if segX <= 0 {
			segX += w - 2
		}
		if segY <= 0 {
			segY += h - 2
		}
		if segY < h && segX < w {
			if i == 0 {
				grid[segY][segX] = 'O'
			} else {
				grid[segY][segX] = 'o'
			}
		}
	}
}
