package animations

import (
	"hs-ascii-live/engine"
	"strings"
)

func RegisterAll(r *engine.Registry) {
	// Categories and names
	categories := map[string][]string{
		"characters": {
			"rickroll", "dancing-rick", "pikachu", "cat", "dog", "among-us", "pacman", "mario", "sonic", "skull", "ghost", "dancing-stickman", "robot", "alien", "spider", "ninja", "samurai", "viking", "pirate", "wizard",
		},
		"space": {
			"galaxy", "milky-way", "blackhole", "solar-system", "rotating-earth", "moon-phases", "saturn", "shooting-stars", "meteor-shower", "supernova", "wormhole", "spaceship", "rocket-launch", "astronaut", "ufo", "comet", "asteroid-belt", "nebula", "constellation", "milky-way2",
		},
		"nature": {
			"fire", "waterfall", "ocean-waves", "rain", "thunderstorm", "tornado", "snowfall", "forest", "blooming-flowers", "sunrise", "sunset", "volcano", "lightning", "wind", "geyser", "aurora", "cave", "river",
		},
		"technology": {
			"matrix-rain", "hacking-terminal", "computer-boot", "digital-glitch", "network-connections", "binary-rain", "ai-neural-network", "circuit-board", "cpu-processing", "gpu-rendering", "data-streams", "cyber-grid", "digital-face", "quantum-computing", "ai-brain", "encryption", "quantum-teleport", "hologram", "server-farm", "drone-swarm",
		},
		"abstract": {
			"rotating-cube", "3d-sphere", "spinning-donut", "dna-helix", "fractal", "tunnel", "particle-explosion", "kaleidoscope", "infinity-loop", "waves", "spiral", "morphing-shapes", "vortex", "optical-illusion", "meta-cube", "fractal-2", "spiral-2",
		},
		"symbols": {
			"animated-heart", "broken-heart", "smiley", "crying-face", "fire-heart", "music-notes", "equalizer", "hourglass", "loading-animation", "countdown", "digital-clock", "pulse", "blink",
		},
		"hsai": {
			"hs-ai-solutions-logo", "hs-ai-studio", "gravity-ai", "galaxy-vision", "hs-ai-neural-network", "ai-model-training", "ai-agent", "hs-ai-core", "hs-digital-universe", "hs-ascii-live-intro", "hs-core-v2", "vision-alpha",
		},
	}

	for cat, names := range categories {
		for _, n := range names {
			anim := NewProceduralAnimation(n, cat)
			r.Register(anim)
		}
	}
}

type ProceduralAnimation struct {
	name        string
	category    string
	description string
}

func NewProceduralAnimation(name, category string) *ProceduralAnimation {
	return &ProceduralAnimation{
		name:        name,
		category:    category,
		description: "Animated ASCII art - " + strings.ReplaceAll(name, "-", " "),
	}
}

func (a *ProceduralAnimation) Name() string { return a.name }
func (a *ProceduralAnimation) Category() string { return a.category }
func (a *ProceduralAnimation) Description() string { return a.description }

func (a *ProceduralAnimation) NextFrame(opts engine.AnimationOptions, frameNum int) engine.Frame {
	seed := hashString(a.name)
	w := 40
	h := 15
	if opts.Width > 0 {
		w = opts.Width
	}
	if opts.Height > 0 {
		h = opts.Height
	}
	pattern := patternForName(a.name)
	content := generatePattern(pattern, w, h, frameNum, seed)
	return engine.Frame{Content: content, Delay: 0}
}

func hashString(s string) int {
	h := 0
	for _, c := range s {
		h = h*31 + int(c)
	}
	if h < 0 {
		h = -h
	}
	return h
}

func patternForName(name string) string {
	// priority mappings
	switch {
	case strings.Contains(name, "heart"):
		return "heart"
	case strings.Contains(name, "matrix") || strings.Contains(name, "binary") || strings.Contains(name, "rain"):
		return "rain"
	case strings.Contains(name, "fire") || strings.Contains(name, "volcano"):
		return "fire"
	case strings.Contains(name, "water") || strings.Contains(name, "ocean") || strings.Contains(name, "wave"):
		return "waves"
	case strings.Contains(name, "snow") || strings.Contains(name, "rain"):
		return "particles"
	case strings.Contains(name, "spiral") || strings.Contains(name, "vortex") || strings.Contains(name, "tunnel"):
		return "spiral"
	case strings.Contains(name, "cube") || strings.Contains(name, "sphere") || strings.Contains(name, "donut"):
		return "rotate"
	case strings.Contains(name, "galaxy") || strings.Contains(name, "milky") || strings.Contains(name, "star"):
		return "stars"
	case strings.Contains(name, "clock") || strings.Contains(name, "countdown"):
		return "clock"
	}
	// fallback: assign a pattern based on the name hash so every animation gets a distinct look
	patterns := []string{"heart", "rain", "fire", "waves", "particles", "spiral", "rotate", "stars", "clock", "bounce"}
	return patterns[hashString(name)%len(patterns)]
}

func generatePattern(pattern string, w, h, frameNum, seed int) string {
	grid := make([][]rune, h)
	for i := range grid {
		grid[i] = make([]rune, w)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	switch pattern {
	case "heart":
		drawHeart(grid, frameNum, seed)
	case "rain":
		drawRain(grid, frameNum, seed)
	case "fire":
		drawFire(grid, frameNum, seed)
	case "waves":
		drawWaves(grid, frameNum, seed)
	case "particles":
		drawParticles(grid, frameNum, seed)
	case "spiral":
		drawSpiral(grid, frameNum, seed)
	case "rotate":
		drawRotate(grid, frameNum, seed)
	case "stars":
		drawStars(grid, frameNum, seed)
	case "clock":
		drawClock(grid, frameNum, seed)
	default:
		drawBounce(grid, frameNum, seed)
	}
	// Add title
	title := "HS ASCII LIVE - " + pattern
	var sb strings.Builder
	for i := 0; i < h; i++ {
		sb.WriteString(string(grid[i]))
		if i < h-1 {
			sb.WriteString("\n")
		}
	}
	// prepend title line
	return title + "\n" + sb.String()
}

func drawHeart(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w/2
	cy := h/2
	offset := frameNum % 4
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := float64(x-cx+offset)
			dy := float64(y-cy)
			if (dx*dx+dy*dy-1)*(dx*dx+dy*dy-1) < 0.3 {
				grid[y][x] = '♥'
			}
		}
	}
}

func drawRain(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for x := 0; x < w; x++ {
		y := (frameNum + (x+seed)%h) % h
		if y >=0 && y < h {
			grid[y][x] = '|'
		}
	}
}

func drawFire(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for y := h-1; y >=0; y-- {
		for x := 0; x < w; x++ {
			if (x+y+frameNum+seed)%7 == 0 {
				grid[y][x] = '▲'
			} else if (x+y+frameNum+seed)%5 == 0 {
				grid[y][x] = '*'
			}
		}
	}
}

func drawWaves(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for x := 0; x < w; x++ {
		offset := int(float64(frameNum+seed+x)*0.2)
		if h/2 > 0 {
			offset = offset % (h/2)
		}
		y := h/2 + offset
		if y >=0 && y < h {
			grid[y][x] = '~'
		}
	}
}

func drawParticles(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for i:=0;i<w*h/10;i++ {
		x := (frameNum*3 + i*seed) % w
		y := (frameNum*2 + i*seed*7) % h
		grid[y][x] = '.'
	}
}

func drawSpiral(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w/2
	cy := h/2
	for i:=0;i<50;i++ {
		angle := float64(i)*0.2 + float64(frameNum)*0.1
		r := float64(i)*0.5
		x := int(float64(cx) + r*float64(i%2))
		y := int(float64(cy) + r*float64((i+seed)%2))
		// simple approximation
		_ = angle
		if y>=0 && y<h && x>=0 && x<w {
			grid[y][x] = '*'
		}
	}
}

func drawRotate(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	offset := frameNum % w
	for y:=0;y<h;y++ {
		for x:=0;x<w;x++ {
			if (x+offset+seed)%(w/2)==0 || (y+offset+seed)%(h/2)==0 {
				grid[y][x] = '#'
			}
		}
	}
}

func drawStars(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	for i:=0;i<30;i++ {
		x := (frameNum*5 + i*seed*3) % w
		y := (frameNum*3 + i*seed*7) % h
		grid[y][x] = '*'
	}
}

func drawClock(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	sec := frameNum % 60
	for i:=0;i<w;i++ {
		grid[0][i] = '-'
	}
	for i:=0;i<h;i++ {
		grid[i][0] = '|'
		grid[i][w-1] = '|'
	}
	grid[h-1][w/2] = ' '
	// seconds hand
	angle := int(float64(sec)/60*360) % 360
	grid[h/2+angle%h][w/2+angle%w] = '>'
}

func drawBounce(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	x := (frameNum*2 + seed) % w
	y := (frameNum*3) % h
	if y>=0 && y<h && x>=0 && x<w {
		grid[y][x] = 'O'
	}
}
