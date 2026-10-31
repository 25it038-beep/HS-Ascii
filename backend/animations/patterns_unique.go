package animations

import (
	"fmt"
	"math"
	"strings"
)

func drawMario(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	// Ground
	groundY := h - 3
	for x := 0; x < w; x++ {
		setCell(grid, x, groundY, '=')
		setCell(grid, x, groundY+1, '#')
	}

	if strings.Contains(animName, "bowser") {
		// Bowser breathing fire
		fireLen := mod(frameNum*3, 10) + 3
		for f := 0; f < fireLen; f++ {
			fx := cx + 8 + f
			fy := groundY - 4 + mod(f+frameNum, 3) - 1
			flames := []rune{'*', '~', '>', '^', '#'}
			setCell(grid, fx, fy, flames[mod(f+frameNum, len(flames))])
		}
		bowser := []string{
			"    /\\____/\\     ",
			"   ( > 皿 < )===~",
			"  /|  [##]  |\\   ",
			" (_|________|_)  ",
		}
		startY := groundY - len(bowser)
		for i, line := range bowser {
			y := startY + i
			startX := cx - 4
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// Question mark coin block
	blockY := h / 3
	blockX := cx - 3
	hit := mod(frameNum, 8) >= 4
	blockOffset := 0
	if hit {
		blockOffset = -1
	}

	box := []string{
		"[???]",
	}
	if blockY+blockOffset >= 0 && blockY+blockOffset < h {
		for j, r := range box[0] {
			if blockX+j >= 0 && blockX+j < w {
				grid[blockY+blockOffset][blockX+j] = r
			}
		}
	}

	// Coin pops out when hit
	if hit {
		coinY := blockY - 2 - mod(frameNum, 3)
		if coinY >= 0 && coinY < h && cx-1 >= 0 && cx-1 < w {
			grid[coinY][cx-1] = '$'
		}
	}

	// Mario jumping trajectory
	jumpCycle := mod(frameNum, 8)
	marioY := groundY - 4
	if jumpCycle >= 2 && jumpCycle <= 5 {
		marioY = blockY + 1 // at the block
	}

	mario := []string{
		"  .-'\"'-.  ",
		" / _   _ \\ ",
		"| (o) (o) |",
		" \\   ^   / ",
		"  /|===|\\  ",
	}

	for i, line := range mario {
		y := marioY + i
		if y >= 0 && y < groundY {
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

func drawSonic(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2
	cx := w / 2

	if strings.Contains(animName, "knuckles") {
		// Knuckles punching
		punch := []string{
			"   /\\__/\\   ",
			"  ( `皿´ )  ",
			" =([##]  )====> [POW!]",
			"   /    \\   ",
		}
		startY := cy - len(punch)/2
		for i, line := range punch {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "donkey") || strings.Contains(animName, "barrel") {
		dk := []string{
			"    .---.    ",
			"   ( o o )   ",
			"  ==(  v  )== ",
			"   / [DK]\\   ",
		}
		startY := cy - len(dk)/2
		for i, line := range dk {
			y := startY + i
			startX := cx - 12
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		bx := cx + mod(frameNum*3, 14)
		setCell(grid, bx, cy, '(')
		setCell(grid, bx+1, cy, '#')
		setCell(grid, bx+2, cy, ')')
		return
	}

	if strings.Contains(animName, "saxophone") || strings.Contains(animName, "jazz") {
		sax := []string{
			"     __       ♪",
			"    /  \\     ♫ ",
			"   | () |====~ ",
			"    \\__/       ",
		}
		startY := cy - len(sax)/2
		for i, line := range sax {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "samus") || strings.Contains(animName, "metroid") {
		beamLen := mod(frameNum*4, 16) + 4
		for b := 0; b < beamLen; b++ {
			setCell(grid, cx+8+b, cy, '=')
		}
		setCell(grid, cx+8+beamLen, cy, '>')
		samus := []string{
			"    .---.    ",
			"   / [o] \\   ",
			"  ==( [O] )=====> [CHARGE-BEAM]",
			"   /  |  \\   ",
		}
		startY := cy - len(samus)/2
		for i, line := range samus {
			y := startY + i
			startX := cx - 8
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// Spin dash cycle
	isSpinning := mod(frameNum, 6) < 4
	sonicX := mod(frameNum*4, w+20) - 10

	// Road
	roadY := cy + 3
	for x := 0; x < w; x++ {
		setCell(grid, x, roadY, '=')
	}

	if isSpinning {
		// Spinning blue ball
		spinChars := []rune{'@', '%', '&', '#'}
		ch := spinChars[mod(frameNum, len(spinChars))]
		setCell(grid, sonicX, cy, ch)
		setCell(grid, sonicX, cy+1, ch)
		setCell(grid, sonicX+1, cy, ch)
		setCell(grid, sonicX+1, cy+1, ch)
		// Dust trails behind
		for d := 1; d <= 4; d++ {
			setCell(grid, sonicX-d*2, cy+1, '.')
		}
	} else {
		// Sonic running
		running := []string{
			" /\\_/\\ ",
			"( o.o )",
			" > ^ < ",
		}
		for i, line := range running {
			y := cy - 1 + i
			if y >= 0 && y < roadY {
				for j, r := range line {
					x := sonicX + j
					if x >= 0 && x < w && r != ' ' {
						grid[y][x] = r
					}
				}
			}
		}
	}
}

func drawPikachu(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "charizard") {
		flame := mod(frameNum*2, 8) + 2
		for f := 0; f < flame; f++ {
			flameChars := []rune{'*', '~', '>', '^'}
			setCell(grid, cx+10+f, cy+mod(f, 2), flameChars[mod(f+frameNum, len(flameChars))])
		}
		charizard := []string{
			"   /\\        /\\   ",
			"  /  \\______/  \\  ",
			" (  ( o )  ( o ) ) ",
			"  \\      ^      /  ===> [FLAME]",
			"   )           (   ",
			"  /  /\\_____/\\  \\  ",
		}
		startY := cy - len(charizard)/2
		for i, line := range charizard {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "gameboy") {
		gb := []string{
			" .------------. ",
			" | .--------. | ",
			" | | >PLAY  | | ",
			" | '--------' | ",
			" |     _      | ",
			" |   _| |_ (B)| ",
			" |  |_   _|(A)| ",
			" '------------' ",
		}
		startY := cy - len(gb)/2
		for i, line := range gb {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	pika := []string{
		"  |\\      /|  ",
		"  | \\    / |  ",
		"   ) \\  / (   ",
		"  /  o  o  \\  ",
		" (    ^     ) ",
		"  \\  '--'  /  ",
		"   |      |  ⚡",
		"   /\\----/\\   ",
	}

	sparkle := mod(frameNum, 3) == 0

	startY := cy - len(pika)/2
	for i, line := range pika {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					if r == '⚡' && sparkle {
						grid[y][x] = '*'
					} else {
						grid[y][x] = r
					}
				}
			}
		}
	}

	// Electric thunder bolts flashing around
	if sparkle {
		bolts := [][2]int{
			{cx - 10, cy - 2},
			{cx + 10, cy + 1},
			{cx - 8, cy + 3},
			{cx + 9, cy - 3},
		}
		for _, b := range bolts {
			if b[1] >= 0 && b[1] < h && b[0] >= 0 && b[0] < w {
				grid[b[1]][b[0]] = '⚡'
			}
		}
	}
}

func drawDog(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	tail := '/'
	if mod(frameNum, 2) == 1 {
		tail = '\\'
	}

	tongue := ' '
	if mod(frameNum, 4) >= 2 {
		tongue = 'v'
	}

	dog := []string{
		"   /^ ^\\   ",
		"  / 0 0 \\  ",
		fmt.Sprintf("  V\\ Y /V %c", tail),
		fmt.Sprintf("   / - \\%c  ", tongue),
		"  /|   |\\  ",
		" (_|---|_) ",
	}

	startY := cy - len(dog)/2
	for i, line := range dog {
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

	// Bark text popping up
	if mod(frameNum, 6) < 3 {
		bark := "WOOF!"
		barkY := startY - 1
		barkX := cx + 6
		if barkY >= 0 && barkY < h {
			for j, r := range bark {
				if barkX+j < w {
					grid[barkY][barkX+j] = r
				}
			}
		}
	}
}

func drawAmongUs(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2

	walkCycle := mod(frameNum*2, w+16) - 8

	leg1 := '|'
	leg2 := '|'
	if mod(frameNum, 2) == 1 {
		leg1 = '/'
		leg2 = '\\'
	}

	crewmate := []string{
		"   .-----.   ",
		"  /  [===]\\  ",
		" |   [===] | ",
		" |         | ",
		fmt.Sprintf("  \\_ %c  %c _/ ", leg1, leg2),
	}

	startY := cy - len(crewmate)/2
	for i, line := range crewmate {
		y := startY + i
		if y >= 0 && y < h {
			for j, r := range line {
				x := walkCycle + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}
}

func drawSkull(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "glitch") {
		// Glitched skull with jittering scanline displacement
		glitchChars := []rune{'#', '%', '&', '@', '$', '!', '?'}
		glitchLine := mod(frameNum, 5)
		skull := []string{
			"     .-------.     ",
			"   /   (X) (X) \\   ",
			"  |       ^       |  [GLITCH]",
			"   \\    '---'    /   ",
			"    '--'-'-'-'--'    ",
		}
		startY := cy - len(skull)/2
		for i, line := range skull {
			y := startY + i
			offset := 0
			if i == glitchLine {
				offset = (mod(frameNum, 3) - 1) * 2
			}
			startX := cx - len(line)/2 + offset
			for j, r := range line {
				if i == glitchLine && mod(j+frameNum, 3) == 0 {
					r = glitchChars[mod(j+frameNum, len(glitchChars))]
				}
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "cyber") {
		// Cyber skull with cybernetic monocle & circuits
		skull := []string{
			"     .-------.     ",
			"   /   [●] (o) \\   ", // cyber eye
			"  |--[CYBER-OS]---|  ",
			"   \\    '---'    /   ",
			"    '--'-'-'-'--'    ",
		}
		startY := cy - len(skull)/2
		for i, line := range skull {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	eye := 'O'
	if mod(frameNum, 4) == 0 {
		eye = 'X' // glowing sinister eye
	}

	skull := []string{
		"     .-------.     ",
		"   /           \\   ",
		fmt.Sprintf("  |   (%c)   (%c)   |  ", eye, eye),
		"  |       ^       |  ",
		"   \\    '---'    /   ",
		"    |  | | | |  |    ",
		"    '--'-'-'-'--'    ",
	}

	startY := cy - len(skull)/2
	for i, line := range skull {
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

func drawBatman(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "joker") {
		ha := []string{"HA!", "HAHA!", "HA HA HA!", "WHY SO SERIOUS?"}
		say := ha[mod(frameNum/3, len(ha))]
		for j, r := range say {
			setCell(grid, cx-len(say)/2+j, cy-5, r)
		}
		joker := []string{
			"   .-\"\"\"\"-.   ",
			"  /        \\  ",
			" |  ^    ^  | ",
			" (  o    o  ) ",
			"  \\  \\__/  /  ",
			"   '------'   ",
		}
		startY := cy - len(joker)/2
		for i, line := range joker {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "gandalf") {
		quote := "YOU SHALL NOT PASS!"
		if mod(frameNum/4, 2) == 1 {
			quote = "RUN, YOU FOOLS!"
		}
		for j, r := range quote {
			setCell(grid, cx-len(quote)/2+j, cy-5, r)
		}
		gandalf := []string{
			"     /\\     ",
			"    /  \\    ",
			"   /____\\   ",
			"  (  o.o )  ",
			"  /|  |  |==[✦]",
			"   |  |  |  ",
			"  /       \\ ",
		}
		startY := cy - len(gandalf)/2
		for i, line := range gandalf {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// Gotham City Skyline at bottom
	cityY := h - 4
	for x := 0; x < w; x++ {
		if x%7 == 0 || x%7 == 5 {
			for y := cityY; y < h; y++ {
				setCell(grid, x, y, '|')
			}
		}
		setCell(grid, x, cityY, '_')
	}

	// Bat-Signal spotlight beam from ground to sky
	beamAngle := math.Sin(float64(frameNum)*0.1) * 6.0
	spotCX := cx + int(beamAngle)
	spotCY := h / 3

	// Bat symbol inside circle
	bat := []string{
		"  .---.     .---.  ",
		" /     \\___/     \\ ",
		"|   /\\       /\\   |",
		" \\_/  \\__ __/  \\_/ ",
		"         V         ",
	}

	startY := spotCY - len(bat)/2
	for i, line := range bat {
		y := startY + i
		if y >= 0 && y < cityY {
			startX := spotCX - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	// Spotlight oval circle
	for a := 0.0; a < 2*math.Pi; a += 0.3 {
		sx := int(float64(spotCX) + 12.0*math.Cos(a))
		sy := int(float64(spotCY) + 4.0*math.Sin(a))
		if sy >= 0 && sy < cityY && sx >= 0 && sx < w {
			grid[sy][sx] = '.'
		}
	}
}

func drawHeroFlight(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2
	cx := w / 2

	if strings.Contains(animName, "hadouken") {
		// Ryu shooting Hadouken energy wave
		energy := []string{
			"    .----.    ",
			"   ( *o*  )   ======>>>>> [HADOUKEN!]",
			"    '----'    ",
		}
		startY := cy - len(energy)/2
		for i, line := range energy {
			y := startY + i
			startX := cx - 10
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "subzero") {
		// Sub-Zero ice freeze blast
		frost := []string{
			"   /\\   /\\   ",
			"  <  ICE-BLAST > ===❄===❄===> [FREEZE]",
			"   \\/   \\/   ",
		}
		startY := cy - len(frost)/2
		for i, line := range frost {
			y := startY + i
			startX := cx - 10
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "dragon") {
		dragon := []string{
			"   /\\        /\\   ",
			"  /  \\______/  \\  ",
			" (   (>o<)      ) ===~ (DRAGON-FLAME)",
			"  \\____________/  ",
			"    /\\      /\\    ",
		}
		startY := cy - len(dragon)/2
		for i, line := range dragon {
			y := startY + i
			startX := cx - len(dragon)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "iss") || strings.Contains(animName, "spacestation") {
		iss := []string{
			"  [#][#][#]====[ O ]====[#][#][#]  ", // solar wings
			"               | |                 ",
			"  [#][#][#]====[___]====[#][#][#]  ",
		}
		startY := cy - len(iss)/2
		for i, line := range iss {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "stealth") || strings.Contains(animName, "bomber") {
		b2 := []string{
			"           /\\           ",
			"         /    \\         ",
			"      /          \\      ",
			"   /       B-2      \\   ",
			" <____________________> ",
		}
		startY := cy - len(b2)/2
		for i, line := range b2 {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "thor") {
		// Mjolnir hammer spinning with lightning
		mjolnir := []string{
			"  [===THOR===]  ",
			"       ||   /\\/\\",
			"       ||  // //",
			"       ||       ",
		}
		startY := cy - len(mjolnir)/2
		for i, line := range mjolnir {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "ironman") {
		// Iron Man repulsor beam blast
		blastLen := mod(frameNum*4, 18) + 4
		for b := 0; b < blastLen; b++ {
			setCell(grid, cx+10+b, cy, '=')
			setCell(grid, cx+10+b, cy-1, '-')
			setCell(grid, cx+10+b, cy+1, '-')
		}
		setCell(grid, cx+10+blastLen, cy, '>')
		ironman := []string{
			"    .---.    ",
			"   / (o) \\   ",
			"  ==( [O] )=====> [REPULSOR]",
			"   /  |  \\   ",
		}
		startY := cy - len(ironman)/2
		for i, line := range ironman {
			y := startY + i
			startX := cx - 6
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	heroX := mod(frameNum*3, w+20) - 10

	cape := '~'
	if mod(frameNum, 2) == 1 {
		cape = '='
	}

	hero := []string{
		fmt.Sprintf("   %c%c%c.==.=======>", cape, cape, cape),
		"     /  o  \\       ",
		fmt.Sprintf("   %c%c%c'=='=======>", cape, cape, cape),
	}

	startY := cy - len(hero)/2
	for i, line := range hero {
		y := startY + i
		if y >= 0 && y < h {
			for j, r := range line {
				x := heroX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	// Clouds passing by
	for i := 0; i < 5; i++ {
		cloudX := mod(i*23-frameNum*2, w)
		cloudY := mod(i*7+2, h)
		if cloudY < h && cloudX < w {
			grid[cloudY][cloudX] = '('
			if cloudX+1 < w {
				grid[cloudY][cloudX+1] = ')'
			}
		}
	}
}

func drawLightsaber(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "samurai") || strings.Contains(animName, "ninja") {
		katana := []string{
			"    \\      /    ",
			"     \\    /     ",
			"      \\  /      ",
			"       \\/       ",
			"       /\\       ",
			"      /  \\      ",
			"   ==[KATANA]== ",
		}
		startY := cy - len(katana)/2
		for i, line := range katana {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "yoda") {
		yoda := []string{
			"  \\_    _/  ",
			"   (o..o)   ",
			"   / || \\===[*] (GREEN SABER)",
			"   \\ || /   ",
		}
		startY := cy - len(yoda)/2
		for i, line := range yoda {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "vader") {
		vader := []string{
			"     .---.     ",
			"    / |^| \\    ",
			"   | (O|O) |   ",
			"  /|  ===  |\\===[/] (RED SABER)",
			"   \\_______/   ",
		}
		startY := cy - len(vader)/2
		for i, line := range vader {
			y := startY + i
			startX := cx - 10
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "death-star") {
		// Death Star superlaser
		ds := []string{
			"     .--------.     ",
			"   /   ( O )    \\   =====",
			"  |              |       ======",
			"   \\            /              ======>>>>> [LASER]",
			"     '--------'     ",
		}
		startY := cy - len(ds)/2
		for i, line := range ds {
			y := startY + i
			startX := cx - 12
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "falcon") {
		// Millennium Falcon hyperdrive
		falcon := []string{
			"      /\\====/\\      ",
			"    /            \\   ",
			"===|   [O]  [O]   |===>>> [HYPERSPACE]",
			"    \\____________/   ",
		}
		startY := cy - len(falcon)/2
		for i, line := range falcon {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// Two sabers clashing in center
	saberLen := 12
	clash := mod(frameNum, 4) == 0

	// Red saber (Vader) from left
	for i := 0; i < saberLen; i++ {
		sx := cx - saberLen + i
		sy := cy - saberLen/2 + i
		if sy >= 0 && sy < h && sx >= 0 && sx < w {
			if i < 3 {
				grid[sy][sx] = '=' // hilt
			} else {
				grid[sy][sx] = '/' // blade
			}
		}
	}

	// Blue/Green saber (Yoda/Luke) from right
	for i := 0; i < saberLen; i++ {
		sx := cx + saberLen - i
		sy := cy - saberLen/2 + i
		if sy >= 0 && sy < h && sx >= 0 && sx < w {
			if i < 3 {
				grid[sy][sx] = '=' // hilt
			} else {
				grid[sy][sx] = '\\' // blade
			}
		}
	}

	// Clash sparks in the middle
	if clash {
		setCell(grid, cx, cy, '*')
		setCell(grid, cx, cy-1, '+')
		setCell(grid, cx, cy+1, '+')
		setCell(grid, cx-1, cy, '+')
		setCell(grid, cx+1, cy, '+')
	}
}

func drawSpider(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "venom") {
		// Venom symbiote with white spider eyes and teeth
		venom := []string{
			"    /\\      /\\    ",
			"   /  \\____/  \\   ",
			"  | (O)    (O) |  ",
			"  |  \\ vvvv /  |==~ (SYMBIO-TONGUE)",
			"   \\  '----'  /   ",
			"    '--------'    ",
		}
		startY := cy - len(venom)/2
		for i, line := range venom {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// Thread hanging from ceiling
	spiderRange := h - 6
	if spiderRange < 1 {
		spiderRange = 1
	}
	spiderY := mod(frameNum, spiderRange) + 2
	for y := 0; y < spiderY; y++ {
		setCell(grid, cx, y, '|')
	}

	leg := '/'
	if mod(frameNum, 2) == 1 {
		leg = '\\'
	}

	spider := []string{
		fmt.Sprintf("  %c\\ (..) /%c  ", leg, leg),
		"  =-( oo )-=  ",
		fmt.Sprintf("  %c/ (..) \\%c  ", leg, leg),
	}

	for i, line := range spider {
		y := spiderY + i
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
}

func drawBasketball(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	// Basketball hoop & backboard on right
	hoopX := w - 10
	hoopY := h / 3
	for y := hoopY - 3; y <= hoopY+4; y++ {
		setCell(grid, hoopX+4, y, '|')
	}
	// Rim & Net
	for x := hoopX; x <= hoopX+3; x++ {
		setCell(grid, x, hoopY, '=')
	}
	setCell(grid, hoopX+1, hoopY+1, '\\')
	setCell(grid, hoopX+2, hoopY+1, '/')

	// Ball arcing through the air towards hoop
	ballCycle := w - 14
	if ballCycle < 1 {
		ballCycle = 10
	}
	step := mod(frameNum*3, ballCycle)
	ballX := 4 + step
	t := float64(step) / float64(ballCycle)
	ballY := int(float64(h)*0.7 - math.Sin(t*math.Pi)*float64(h)*0.55)

	setCell(grid, ballX, ballY, '🏀')

	// Floor
	floorY := h - 2
	for x := 0; x < w; x++ {
		setCell(grid, x, floorY, '=')
	}
}

func drawTennis(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2

	// Net across the middle
	for x := 0; x < w; x++ {
		setCell(grid, x, cy, '-')
	}

	// Ball traveling back and forth across net
	ballRange := h - 4
	if ballRange < 1 {
		ballRange = 1
	}
	ballY := mod(frameNum*2, ballRange) + 2
	ballX := w/2 + int(math.Sin(float64(frameNum)*0.3)*float64(w)*0.35)

	setCell(grid, ballX, ballY, '🎾')

	// Player racket 1 (top)
	p1X := w/2 + int(math.Sin(float64(frameNum)*0.3)*float64(w)*0.3)
	setCell(grid, p1X, 1, 'D')

	// Player racket 2 (bottom)
	p2X := w/2 - int(math.Sin(float64(frameNum)*0.3)*float64(w)*0.3)
	setCell(grid, p2X, h-2, 'D')
}

func drawBowling(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2

	// Bowling lane gutters
	for x := 0; x < w; x++ {
		setCell(grid, x, cy-4, '=')
		setCell(grid, x, cy+4, '=')
	}

	// 10 Bowling pins on right
	pinsX := w - 8
	hit := mod(frameNum, 12) >= 8
	pinChar := '!'
	if hit {
		pinChar = '/' // flying pins
	}

	pins := [][]int{
		{pinsX, cy},
		{pinsX + 2, cy - 1}, {pinsX + 2, cy + 1},
		{pinsX + 4, cy - 2}, {pinsX + 4, cy}, {pinsX + 4, cy + 2},
	}
	for _, p := range pins {
		setCell(grid, p[0], p[1], pinChar)
	}

	// Bowling ball rolling down the lane
	laneLen := w - 10
	if laneLen < 1 {
		laneLen = 1
	}
	ballStep := mod(frameNum*3, laneLen)
	ballX := 2 + ballStep
	if hit {
		setCell(grid, pinsX, cy, '*')
	} else {
		setCell(grid, ballX, cy, '●')
	}
}

func drawBurger(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	steam := '~'
	if mod(frameNum, 2) == 1 {
		steam = ' '
	}

	burger := []string{
		fmt.Sprintf("    %c   %c   %c    ", steam, steam, steam),
		"  .------------.  ",
		" (  *  *  *  *  ) ", // sesame top bun
		"  \\~~~~~~~~~~~~/  ", // lettuce
		"   [==========]   ", // tomato & cheese
		"  (############)  ", // beef patty
		"   '----------'   ", // bottom bun
		"==================", // plate
	}

	startY := cy - len(burger)/2
	for i, line := range burger {
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

func drawPizza(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	pizza := []string{
		"  /================\\  ", // crust
		" /   (O)    (O)     \\ ", // pepperonis
		"/      (O)      (O)  \\",
		"\\   (O)     (O)      /",
		" \\      (O)         / ",
		"  \\================/  ",
	}

	startY := cy - len(pizza)/2
	for i, line := range pizza {
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

func drawHelicopter(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	// Spinning main rotor
	rotor := "-------[+]-------"
	if mod(frameNum, 2) == 1 {
		rotor = "   ====[+]====   "
	}

	heli := []string{
		rotor,
		"       ---       ",
		"  .---/   \\----. ",
		" /    (o)       \\",
		"|       [HS]    /=",
		" \\-------------/ ",
		"     === ===     ", // skids
	}

	startY := cy - len(heli)/2
	for i, line := range heli {
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

func drawSubmarine(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	sub := []string{
		"        |__       ", // periscope
		"        |         ",
		"    .---|---.     ",
		"---/ (o) (o) \\---.",
		"==|   HS-AI       >", // propeller
		"---\\             /---'",
		"    '-----------' ",
	}

	startY := cy - len(sub)/2
	for i, line := range sub {
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

	// Propeller bubbles trailing to the right
	bubbleChars := []rune{'o', 'O', '°', '.'}
	for b := 1; b <= 6; b++ {
		bx := cx + len(sub[0])/2 + b*2 + mod(frameNum, 2)
		by := cy + 1 - mod(b+frameNum, 3)
		if by >= 0 && by < h && bx < w {
			grid[by][bx] = bubbleChars[mod(b+frameNum, len(bubbleChars))]
		}
	}
}

func drawTrain(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	train := []string{
		"       (@@@)  (@@)       ",
		"        |~|              ",
		"    .---| |------------. ",
		"   / |  | |            | ",
		"==|  |  |_|  [HS-EXP]  | ",
		"  \\====================/ ",
		"   (O)(O)  (O)(O) (O)(O) ", // drive wheels
	}

	startY := cy - len(train)/2
	for i, line := range train {
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

	// Railroad tracks beneath
	trackY := startY + len(train)
	for x := 0; x < w; x++ {
		setCell(grid, x, trackY, '=')
		if mod(x+frameNum*2, 3) == 0 {
			setCell(grid, x, trackY+1, '|')
		}
	}
}

func drawSailingShip(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "zeppelin") || strings.Contains(animName, "airship") {
		zep := []string{
			"     .-----------------------.     ",
			"   (      ZEPPELIN-AIRSHIP     )   ",
			"     '-----------------------'     ",
			"              [=====]              ", // gondola
		}
		startY := cy - len(zep)/2
		for i, line := range zep {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		// Passing clouds
		for i := 0; i < 4; i++ {
			cRow := cy + 3 + i
			setCell(grid, mod(i*23-frameNum*2, w), cRow, '(')
			setCell(grid, mod(i*23-frameNum*2+1, w), cRow, ')')
		}
		return
	}

	if strings.Contains(animName, "viking") {
		viking := []string{
			"           /\\            ",
			"          /  \\           ",
			"         /____\\          ", // sail
			"          |  |           ",
			"   <====| (O)(O)(O) |===>", // dragon prow & shields
			"  ~~~~~~~~~~~~~~~~~~~~~~~~",
		}
		startY := cy - len(viking)/2
		for i, line := range viking {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "pirate") || strings.Contains(animName, "galleon") {
		flag := "[X]"
		if mod(frameNum, 2) == 1 {
			flag = "[*]"
		}
		pirate := []string{
			fmt.Sprintf("       %s |\\        ", flag), // jolly roger
			"          | \\       ",
			"          |__\\      ",
			"     \\             /",
			"   ~~~'-----------'~~~",
		}
		startY := cy - len(pirate)/2
		for i, line := range pirate {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	ship := []string{
		"      |\\       |\\      ",
		"      | \\      | \\     ",
		"      |__\\     |__\\    ",
		"     /|       /|       ",
		"    /_|______/_|______ ",
		"  \\                 /  ",
		"~~~'---------------'~~~",
	}

	startY := cy - len(ship)/2
	for i, line := range ship {
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

	// Rolling ocean waves below
	waterY := startY + len(ship)
	for x := 0; x < w; x++ {
		wy := waterY + int(math.Sin(float64(x+frameNum*2)*0.2)*1.5)
		setCell(grid, x, wy, '~')
	}
}

func drawMotorcycle(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "cycling") {
		// Bicycle rider
		bike := []string{
			"     __o    ",
			"   _`\\<,_   ",
			"  (O)/ (O)  ",
		}
		startY := cy - len(bike)/2
		for i, line := range bike {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		for x := 0; x < w; x++ {
			setCell(grid, x, startY+len(bike), '=')
		}
		return
	}

	if strings.Contains(animName, "hover") {
		// Futuristic hover bike
		hb := []string{
			"     __o            ",
			"   _`\\<,_   ~======>",
			"  [==HOVER-BIKE==]~ ",
		}
		startY := cy - len(hb)/2
		for i, line := range hb {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	wheel := 'O'
	if mod(frameNum, 2) == 1 {
		wheel = '0'
	}

	bike := []string{
		"     __o     ",
		"   _`\\<,_    ", // rider
		fmt.Sprintf("  (%c)/ (%c)  ", wheel, wheel),
	}

	startY := cy - len(bike)/2
	for i, line := range bike {
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

	// Speed road
	roadY := startY + len(bike)
	for x := 0; x < w; x++ {
		setCell(grid, x, roadY, '=')
		if mod(x+frameNum*3, 4) == 0 {
			setCell(grid, x, roadY+1, '-')
		}
	}
}

func drawEclipse(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.42

	// Moon slides across sun
	moonOffset := math.Sin(float64(frameNum)*0.1) * r * 1.5

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dxSun := (float64(x) - cx) * 0.55
			dySun := float64(y) - cy
			distSun := math.Sqrt(dxSun*dxSun + dySun*dySun)

			dxMoon := (float64(x) - (cx + moonOffset/0.55)) * 0.55
			dyMoon := float64(y) - cy
			distMoon := math.Sqrt(dxMoon*dxMoon + dyMoon*dyMoon)

			if distSun <= r {
				if distMoon < r*0.95 {
					grid[y][x] = ' ' // eclipsed by dark moon
				} else {
					// Glowing solar corona ring
					grid[y][x] = '☼'
				}
			} else if distSun < r*1.3 {
				// Coronal rays
				if mod(x+y+frameNum, 3) == 0 {
					grid[y][x] = '*'
				}
			}
		}
	}
}

func drawHourglass(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	sand := '.'
	if mod(frameNum, 2) == 1 {
		sand = ':'
	}

	glass := []string{
		"  +=======+  ",
		fmt.Sprintf("  | %c%c%c%c%c |  ", sand, sand, sand, sand, sand), // sand top
		"   \\     /   ",
		"    \\   /    ",
		"     ) (     ",
		"    /   \\    ",
		"   /     \\   ",
		fmt.Sprintf("  | ..%c.. |  ", sand), // sand piling up bottom
		"  +=======+  ",
	}

	startY := cy - len(glass)/2
	for i, line := range glass {
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

func drawCountdown(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	countStep := mod(frameNum/5, 4)
	digits := [][]string{
		{
			"  ====  ",
			"      = ",
			"   ===  ",
			"      = ",
			"  ====  ",
		}, // 3
		{
			"  ====  ",
			"      = ",
			"  ====  ",
			" =      ",
			"  ====  ",
		}, // 2
		{
			"    =   ",
			"   ==   ",
			"    =   ",
			"    =   ",
			"  ===== ",
		}, // 1
		{
			"  ==  ==   ",
			" =  = =  = ",
			" = == =  = ",
			" =  = =  = ",
			"  ===  ==  ",
		}, // GO
	}

	d := digits[countStep]
	startY := cy - len(d)/2
	for i, line := range d {
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

func drawVolcano(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	// Volcano cone
	coneH := h / 2
	craterY := h - coneH - 2

	for y := craterY; y < h; y++ {
		t := float64(y-craterY) / float64(coneH+2)
		left := int(float64(cx) - t*float64(w)*0.35)
		right := int(float64(cx) + t*float64(w)*0.35)
		for x := left; x <= right; x++ {
			if x >= 0 && x < w {
				if x == left {
					grid[y][x] = '/'
				} else if x == right {
					grid[y][x] = '\\'
				} else if mod(x+y, 4) == 0 {
					grid[y][x] = '^' // lava flow streak
				} else {
					grid[y][x] = '#'
				}
			}
		}
	}

	// Flying molten lava bombs & sparks from crater
	for b := 0; b < 10; b++ {
		bx := cx + int(math.Sin(float64(b+frameNum)*0.5)*float64(b*2))
		by := craterY - 1 - mod(frameNum*2+b*3, craterY)
		if by >= 0 && by < craterY && bx >= 0 && bx < w {
			grid[by][bx] = '▲'
		}
	}
}

func drawWaterfall(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	// Cliff ledge
	cliffY := h / 4
	for x := 0; x < w; x++ {
		if x < cx-8 || x > cx+8 {
			setCell(grid, x, cliffY, '=')
		}
	}

	// Cascading water stream
	waterChars := []rune{'|', '!', '¦', ':'}
	for y := cliffY; y < h-2; y++ {
		for x := cx - 7; x <= cx+7; x++ {
			idx := mod(x+y+frameNum*2, len(waterChars))
			setCell(grid, x, y, waterChars[idx])
		}
	}

	// Splashing foam & mist pool at bottom
	mistY := h - 2
	for x := cx - 12; x <= cx+12; x++ {
		ch := 'o'
		if mod(x+frameNum, 2) != 0 {
			ch = '~'
		}
		setCell(grid, x, mistY, ch)
	}
}

func drawRainbow(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	baseY := float64(h) - 2.0

	// 5 Arching bands of the rainbow
	bands := []rune{'=', '-', '~', '.', '°'}
	for b, ch := range bands {
		r := float64(h)*0.85 - float64(b)*1.5
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			radSq := r*r - dx*dx
			if radSq > 0 {
				dy := math.Sqrt(radSq)
				y := int(baseY - dy)
				if y >= 0 && y < h {
					grid[y][x] = ch
				}
			}
		}
	}

	// Cloud puffs on ground at rainbow ends
	clouds := []string{"(@@@@)", " (@@) "}
	for _, c := range clouds {
		cy := int(baseY) - 1
		for j, r := range c {
			setCell(grid, 4+j, cy, r)
			setCell(grid, w-12+j, cy, r)
		}
	}
}

func drawTerminalHacker(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	lines := []string{
		"[HS-KERNEL v4.19.0] INITIALIZING SECURE SHELL...",
		"root@hs-ai:~# access --override --bypass-auth",
		"[+] DECRYPTING QUANTUM KEY: 0x9AF83E1B4C...",
		"[+] CONNECTION ESTABLISHED: 192.168.1.1:443",
		"[+] NEURAL WEIGHTS SYNCHRONIZED: 100% OK",
		"00000000  7f 45 4c 46 02 01 01 00  00 00 00 00 00 00 00 00  |.ELF............|",
		"00000010  03 00 3e 00 01 00 00 00  30 1e 00 00 00 00 00 00  |..>.....0.......|",
		"root@hs-ai:~# execute --mode=LIVE_STREAM",
		"[>>> STREAM ACTIVE: 349 ANIMATIONS RUNNING <<<]",
	}

	for i, l := range lines {
		y := 1 + i
		if y < h {
			offset := mod(frameNum, 2)
			for j, r := range l {
				if j < w-2 {
					if y == len(lines) && offset == 0 && j == len(l)-1 {
						grid[y][j] = '█' // blinking prompt cursor
					} else {
						grid[y][j] = r
					}
				}
			}
		}
	}
}

func drawGuitar(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cy := h / 2

	vibe := '-'
	if mod(frameNum, 2) == 1 {
		vibe = '~'
	}

	guitar := []string{
		"  .---.                                               ",
		" ( O O )======================================[HS-AMP]",
		fmt.Sprintf("  \\===/%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c%c/       ", vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe, vibe),
		"  (   )                                               ",
		"   '='                                                ",
	}

	startY := cy - len(guitar)/2
	for i, line := range guitar {
		y := startY + i
		if y >= 0 && y < h {
			startX := 4
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}
}

func drawFirework(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) * 0.45

	cycle := mod(frameNum, 10)
	if cycle < 3 {
		// Rocket flying up
		rocketY := h - 2 - cycle*3
		if rocketY >= 0 && rocketY < h && int(cx) < w {
			grid[rocketY][int(cx)] = '^'
			if rocketY+1 < h {
				grid[rocketY+1][int(cx)] = '|'
			}
		}
	} else {
		// Burst explosion expanding outward
		radius := float64(cycle-3) * 2.2
		particles := []rune{'*', '+', '✦', '.', '°'}
		for a := 0.0; a < 2*math.Pi; a += 0.4 {
			px := int(cx + radius*math.Cos(a)/0.55)
			py := int(cy + radius*math.Sin(a))
			if py >= 0 && py < h && px >= 0 && px < w {
				grid[py][px] = particles[mod(int(a*5)+cycle, len(particles))]
			}
		}
	}
}

func drawAurora(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])

	chars := []rune("~≈~-^")

	for layer := 0; layer < 4; layer++ {
		phase := float64(frameNum)*0.1 + float64(layer)*1.5
		amp := 3.0 + float64(layer)*1.2
		midY := float64(h)*0.35 + float64(layer*2)

		for x := 0; x < w; x++ {
			y := int(midY + math.Sin(float64(x)*0.15+phase)*amp)
			if y >= 0 && y < h {
				grid[y][x] = chars[mod(layer+x, len(chars))]
			}
		}
	}

	// Stars shining through the aurora curtains
	for i := 0; i < 20; i++ {
		sx := mod(i*29+seed*13, w)
		sy := mod(i*11+seed*7, h/2)
		if sy < h && sx < w {
			grid[sy][sx] = '✦'
		}
	}
}


func drawWolf(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	if strings.Contains(animName, "bear") {
		// Polar bear on arctic ice
		bear := []string{
			"     __  __      ",
			"    ( o)(o )     ",
			"   /        \\    ",
			"  |  POLAR   |   ",
			"  |   BEAR   |   ",
			"   \\________/    ",
			"    ||    ||     ",
		}
		startY := h - 4 - len(bear)
		for i, line := range bear {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		iceY := h - 3
		for x := 0; x < w; x++ {
			setCell(grid, x, iceY, '=')
		}
		return
	}

	// Full moon in top right
	moonX := cx + 14
	moonY := h / 4
	for a := 0.0; a < 2*math.Pi; a += 0.4 {
		mx := int(float64(moonX) + 5.0*math.Cos(a)/0.55)
		my := int(float64(moonY) + 3.0*math.Sin(a))
		if my >= 0 && my < h && mx >= 0 && mx < w {
			grid[my][mx] = 'O'
		}
	}

	// Howling wolf on rocky cliff ledge
	cliffY := h - 4
	if cliffY >= 0 && cliffY < h {
		for x := 0; x < cx+6 && x < w; x++ {
			grid[cliffY][x] = '#'
		}
	}

	wolf := []string{
		"       /\\       ",
		"      /  \\___   ",
		"     /  /\\__ `> ", // howling snout pointed up
		"    /  /    `   ",
		"   /  /         ",
		" _/  /          ",
		"(____)          ",
	}

	startY := cliffY - len(wolf)
	for i, line := range wolf {
		y := startY + i
		if y >= 0 && y < h {
			startX := cx - 8
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}

	// Howl soundwaves radiating towards moon
	if mod(frameNum, 4) < 2 {
		waves := []string{")", "))", ")))"}
		wStr := waves[mod(frameNum, len(waves))]
		hy := startY + 2
		hx := cx + 3
		for j, r := range wStr {
			if hx+j >= 0 && hx+j < w && hy >= 0 && hy < h {
				grid[hy][hx+j] = r
			}
		}
	}
}

func drawFrog(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	// Water lily pads on water
	waterY := h - 3
	for x := 0; x < w; x++ {
		if waterY < h {
			grid[waterY][x] = '~'
		}
	}

	// Frog leaping in an arc
	cycle := mod(frameNum, 8)
	frogX := cx - 12 + cycle*3
	frogY := waterY - 3
	if cycle >= 2 && cycle <= 5 {
		frogY = waterY - 7 // leaping up
	}

	frog := []string{
		"  (o)_(o)  ",
		" (___ ___ )",
		" /       \\ ",
		"(_/|   |\\_)",
	}

	for i, line := range frog {
		y := frogY + i
		if y >= 0 && y < waterY {
			startX := frogX - len(line)/2
			for j, r := range line {
				x := startX + j
				if x >= 0 && x < w && r != ' ' {
					grid[y][x] = r
				}
			}
		}
	}
}

func drawWhale(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	waterY := h / 2

	// Water line across middle
	for x := 0; x < w; x++ {
		wy := waterY + int(math.Sin(float64(x+frameNum*2)*0.2)*1.2)
		setCell(grid, x, wy, '~')
	}

	// Water spout shooting up into the air
	spoutH := mod(frameNum, 5) + 2
	for s := 0; s < spoutH; s++ {
		sy := waterY - 2 - s
		setCell(grid, cx, sy, '|')
		if s == spoutH-1 {
			setCell(grid, cx-1, sy, '(')
			setCell(grid, cx+1, sy, ')')
		}
	}

	// Giant whale body
	whale := []string{
		"       .------------.         ",
		"     /                \\____   ",
		"====|  (o)                 \\  ",
		"     \\                     /  ",
		"       '------------------'   ",
		"              /\\              ",
		"             /__\\             ", // fluke tail
	}

	if strings.Contains(animName, "docker") {
		// Stacked shipping containers on whale back
		containers := []string{
			"        [#][#][#]        ",
			"       [#][#][#][#]      ",
			"     .------------.      ",
			"   /   (o)        \\____  ",
			"===|                   \\ ",
			"    \\                 /  ",
			"     '---------------'   ",
		}
		whale = containers
	}

	startY := waterY - 1
	for i, line := range whale {
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

func drawHorse(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	groundY := h - 3

	for x := 0; x < w; x++ {
		if groundY < h {
			grid[groundY][x] = '='
		}
	}

	// Galloping legs animation
	leg1 := '/'
	leg2 := '\\'
	if mod(frameNum, 2) == 1 {
		leg1 = '\\'
		leg2 = '/'
	}

	horse := []string{
		"     (\\(\\         ",
		"    ( -.-)        ",
		"   o_(\")(\")       ",
		"     /    \\       ",
		fmt.Sprintf("    %c      %c      ", leg1, leg2),
		fmt.Sprintf("   %c        %c     ", leg2, leg1),
	}

	startY := groundY - len(horse)
	for i, line := range horse {
		y := startY + i
		if y >= 0 && y < groundY {
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

func drawSaturn(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.35

	// Central planet sphere
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				// Planet atmospheric bands
				if mod(y, 2) == 0 {
					grid[y][x] = '='
				} else {
					grid[y][x] = '~'
				}
			}
		}
	}

	// Huge tilted planetary rings
	tilt := 0.25
	ringOuter := r * 2.2
	ringInner := r * 1.3

	for rx := -ringOuter / 0.55; rx <= ringOuter/0.55; rx += 0.5 {
		x := int(cx + rx)
		y := int(cy + rx*tilt*0.55)
		distX := rx * 0.55
		dist := math.Abs(distX)

		if dist >= ringInner && dist <= ringOuter {
			if y >= 0 && y < h && x >= 0 && x < w {
				if mod(frameNum+int(dist), 2) == 0 {
					grid[y][x] = '='
				} else {
					grid[y][x] = '-'
				}
			}
		}
	}
}

func drawBoba(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	if strings.Contains(animName, "ice-cream") {
		iceCream := []string{
			"        (●)        ", // cherry
			"     .-------.     ",
			"   /   ( o )   \\   ", // scoops
			"  (    ( o )    )  ",
			"   \\___________/   ",
			"    \\         /    ", // waffle cone
			"     \\       /     ",
			"      \\     /      ",
			"       \\___/       ",
		}
		startY := cy - len(iceCream)/2
		for i, line := range iceCream {
			y := startY + i
			startX := cx - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	cup := []string{
		"        ||        ", // straw
		"     .__||__.     ", // lid
		"    /   ||   \\    ",
		"   |    ||    |   ",
		"   |  BOBA    |   ", // tea
		"   | (o)(o)(o)|   ", // tapioca pearls
		"   | (o)(o)(o)|   ",
		"    \\________/    ",
	}

	startY := cy - len(cup)/2
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

func drawSushi(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	chopsticks := []string{
		"   \\    /   ",
		"    \\  /    ",
		"     \\/     ",
	}
	for i, line := range chopsticks {
		y := cy - 5 + i
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

	sushiRoll := []string{
		"  .------------.  ",
		" /   [SUSHI]    \\ ",
		"|  (O)  (O)  (O) |", // fish roe / filling
		" \\~~~~~~~~~~~~~~/ ", // seaweed nori wrap
		"  '------------'  ",
		"==================", // wooden geta tray
	}

	startY := cy - 1
	for i, line := range sushiRoll {
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

func drawBoombox(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	speaker := 'O'
	if mod(frameNum, 2) == 1 {
		speaker = 'o'
	}

	boom := []string{
		"      .-----[HS-AUDIO]-----.      ",
		"     /                      \\     ", // handle
		"  .----------------------------.  ",
		fmt.Sprintf("  | ( %c )   [====]   ( %c ) |  ", speaker, speaker),
		fmt.Sprintf("  | ((%c))   |oo::|   ((%c)) |  ", speaker, speaker), // cassette deck
		"  '----------------------------'  ",
	}

	startY := cy - len(boom)/2
	for i, line := range boom {
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

	// Soundwaves blasting out from speakers
	if mod(frameNum, 2) == 0 {
		for y := startY + 3; y <= startY+4; y++ {
			setCell(grid, cx-18, y, ')')
			setCell(grid, cx+18, y, '(')
		}
	}
}

func drawCassette(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	spool := '*'
	if mod(frameNum, 2) == 1 {
		spool = '+'
	}

	tape := []string{
		"  .------------------------.  ",
		"  |  [A]  MIXTAPE VOL. 1   |  ",
		fmt.Sprintf("  |  ( %c )======( %c )  |  ", spool, spool),
		"  |  '--' /====\\ '--'   |  ",
		"  '------/======\\------'  ",
	}

	startY := cy - len(tape)/2
	for i, line := range tape {
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

func drawF1Racer(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2
	cy := h / 2

	f1 := []string{
		"         ___           ",
		"   __   /   \\   __    ",
		"==|  |=[HS-F1]=|  |===>",
		"  (OO)         (OO)    ", // slicks
	}

	startY := cy - len(f1)/2
	for i, line := range f1 {
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

	// Curbs / track below
	roadY := startY + len(f1)
	for x := 0; x < w; x++ {
		ch := '='
		if mod(x+frameNum*4, 6) < 3 {
			ch = '#'
		}
		setCell(grid, x, roadY, ch)
	}
}
