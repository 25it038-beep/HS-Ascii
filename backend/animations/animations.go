package animations

import (
	"fmt"
	"hs-ascii-live/engine"
	"strings"
)

func RegisterAll(r *engine.Registry) {
	// Categories and names
	categories := map[string][]string{
		"characters": {
			"rickroll", "dancing-rick", "pikachu", "cat", "dog", "among-us", "pacman", "mario", "sonic", "skull", "ghost", "dancing-stickman", "robot", "alien", "spider", "ninja", "samurai", "viking", "pirate", "wizard",
			"batman", "superman", "ironman", "spiderman", "yoda", "darth-vader", "gandalf", "cyber-ninja",
		},
		"space": {
			"galaxy", "milky-way", "blackhole", "solar-system", "rotating-earth", "moon-phases", "saturn", "shooting-stars", "meteor-shower", "supernova", "wormhole", "spaceship", "rocket-launch", "astronaut", "ufo", "comet", "asteroid-belt", "nebula", "constellation", "milky-way2",
			"jupiter-storm", "pulsar", "quasars", "rings-of-uranus", "solar-eclipse", "deep-space-probe",
		},
		"nature": {
			"fire", "waterfall", "ocean-waves", "rain", "thunderstorm", "tornado", "snowfall", "forest", "blooming-flowers", "sunrise", "sunset", "volcano", "lightning", "wind", "geyser", "aurora", "cave", "river",
			"blizzard", "autumn-leaves", "desert-dunes", "coral-reef", "bamboo-forest", "meteor-impact",
		},
		"technology": {
			"matrix-rain", "hacking-terminal", "computer-boot", "digital-glitch", "network-connections", "binary-rain", "ai-neural-network", "circuit-board", "cpu-processing", "gpu-rendering", "data-streams", "cyber-grid", "digital-face", "quantum-computing", "ai-brain", "encryption", "quantum-teleport", "hologram", "server-farm", "drone-swarm",
			"blockchain-mining", "cyber-skull", "radar-scanner", "synthwave-sun", "laser-grid", "terminal-matrix",
		},
		"abstract": {
			"rotating-cube", "3d-sphere", "spinning-donut", "dna-helix", "fractal", "tunnel", "particle-explosion", "kaleidoscope", "infinity-loop", "waves", "spiral", "morphing-shapes", "vortex", "optical-illusion", "meta-cube", "fractal-2", "spiral-2",
			"hypercube-tesseract", "lissajous-curve", "torus-knot", "sine-wave-field", "sacred-geometry", "quantum-foam",
		},
		"symbols": {
			"animated-heart", "broken-heart", "smiley", "crying-face", "fire-heart", "music-notes", "equalizer", "hourglass", "loading-animation", "countdown", "digital-clock", "pulse", "blink",
			"yin-yang", "peace-sign", "radioactive-hazard", "biohazard", "infinity-symbol", "audio-bars",
		},
		"hsai": {
			"hs-ai-solutions-logo", "hs-ai-studio", "gravity-ai", "galaxy-vision", "hs-ai-neural-network", "ai-model-training", "ai-agent", "hs-ai-core", "hs-digital-universe", "hs-ascii-live-intro", "hs-core-v2", "vision-alpha",
			"hs-superintelligence", "hs-agentic-flow", "hs-neural-matrix", "hs-deep-thought", "hs-galaxy-core",
		},
		"gaming": {
			"space-invaders", "pong", "tetris", "flappy-bird", "dino-runner", "snake-game", "doom-fire", "arcade-joystick", "gameboy", "zelda-triforce",
		},
		"vehicles": {
			"cybertruck", "speeding-car", "steam-locomotive", "airplane-flight", "submarine", "motorcycle", "helicopter", "hovercraft", "sailing-ship", "space-shuttle",
		},
		"creatures": {
			"roaring-lion", "dragon-flight", "phoenix-rising", "flying-bat", "swimming-whale", "hopping-frog", "galloping-horse", "slithering-snake", "flapping-butterfly", "howling-wolf",
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

func (a *ProceduralAnimation) Name() string        { return a.name }
func (a *ProceduralAnimation) Category() string    { return a.category }
func (a *ProceduralAnimation) Description() string { return a.description }

func (a *ProceduralAnimation) NextFrame(opts engine.AnimationOptions, frameNum int) engine.Frame {
	seed := hashString(a.name)
	w := 80
	h := 24
	if opts.Width > 0 {
		w = opts.Width
	}
	if opts.Height > 0 {
		h = opts.Height
	}
	if w < 10 {
		w = 10
	}
	if h < 5 {
		h = 5
	}
	pattern := patternForName(a.name)
	content := generatePattern(pattern, w, h, frameNum, seed, a.name)
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
	if h < 0 { // math.MinInt
		h = 0
	}
	return h
}

func mod(n, m int) int {
	if m <= 0 {
		return 0
	}
	r := n % m
	if r < 0 {
		r += m
	}
	return r
}

func patternForName(name string) string {
	n := strings.ToLower(name)
	switch {
	case strings.Contains(n, "pacman"):
		return "pacman"
	case strings.Contains(n, "cat"):
		return "cat"
	case strings.Contains(n, "stickman") || strings.Contains(n, "rick") || strings.Contains(n, "dance"):
		return "dancer"
	case strings.Contains(n, "skull") || strings.Contains(n, "ghost"):
		return "ghost"
	case strings.Contains(n, "invader"):
		return "space-invaders"
	case strings.Contains(n, "pong"):
		return "pong"
	case strings.Contains(n, "tetris"):
		return "tetris"
	case strings.Contains(n, "snake"):
		return "snake"
	case strings.Contains(n, "donut") || strings.Contains(n, "torus"):
		return "donut"
	case strings.Contains(n, "cube") || strings.Contains(n, "tesseract"):
		return "cube"
	case strings.Contains(n, "dna") || strings.Contains(n, "helix"):
		return "dna"
	case strings.Contains(n, "matrix") || strings.Contains(n, "terminal") || strings.Contains(n, "binary") || strings.Contains(n, "code"):
		return "matrix"
	case strings.Contains(n, "equalizer") || strings.Contains(n, "music") || strings.Contains(n, "audio"):
		return "equalizer"
	case strings.Contains(n, "rocket") || strings.Contains(n, "spaceship") || strings.Contains(n, "shuttle") || strings.Contains(n, "probe"):
		return "rocket"
	case strings.Contains(n, "earth") || strings.Contains(n, "globe") || strings.Contains(n, "sphere") || strings.Contains(n, "moon") || strings.Contains(n, "saturn") || strings.Contains(n, "planet"):
		return "earth"
	case strings.Contains(n, "car") || strings.Contains(n, "truck") || strings.Contains(n, "vehicle") || strings.Contains(n, "locomotive") || strings.Contains(n, "motorcycle") || strings.Contains(n, "submarine") || strings.Contains(n, "flight") || strings.Contains(n, "helicopter") || strings.Contains(n, "ship"):
		return "vehicle"
	case strings.Contains(n, "neural") || strings.Contains(n, "brain") || strings.Contains(n, "network") || strings.Contains(n, "agent") || strings.Contains(n, "model") || strings.Contains(n, "connect"):
		return "neural"
	case strings.Contains(n, "radar") || strings.Contains(n, "scan"):
		return "radar"
	case strings.Contains(n, "blackhole") || strings.Contains(n, "wormhole") || strings.Contains(n, "vortex") || strings.Contains(n, "singularity"):
		return "blackhole"
	case strings.Contains(n, "hs-") || strings.Contains(n, "hsai") || strings.Contains(n, "gravity") || strings.Contains(n, "vision"):
		return "hsai"
	case strings.Contains(n, "dragon") || strings.Contains(n, "phoenix") || strings.Contains(n, "bird") || strings.Contains(n, "bat") || strings.Contains(n, "butterfly") || strings.Contains(n, "wing"):
		return "wings"
	case strings.Contains(n, "yin-yang") || strings.Contains(n, "peace") || strings.Contains(n, "hazard") || strings.Contains(n, "biohazard") || strings.Contains(n, "infinity"):
		return "symbol"
	case strings.Contains(n, "heart"):
		return "heart"
	case strings.Contains(n, "fire") || strings.Contains(n, "volcano") || strings.Contains(n, "doom") || strings.Contains(n, "flame"):
		return "fire"
	case strings.Contains(n, "water") || strings.Contains(n, "ocean") || strings.Contains(n, "wave") || strings.Contains(n, "river") || strings.Contains(n, "waterfall"):
		return "waves"
	case strings.Contains(n, "rain") || strings.Contains(n, "storm"):
		return "rain"
	case strings.Contains(n, "snow") || strings.Contains(n, "blizzard"):
		return "particles"
	case strings.Contains(n, "spiral") || strings.Contains(n, "tunnel"):
		return "spiral"
	case strings.Contains(n, "galaxy") || strings.Contains(n, "milky") || strings.Contains(n, "star") || strings.Contains(n, "supernova") || strings.Contains(n, "nebula") || strings.Contains(n, "constellation") || strings.Contains(n, "meteor") || strings.Contains(n, "pulsar") || strings.Contains(n, "quasar") || strings.Contains(n, "eclipse"):
		return "stars"
	case strings.Contains(n, "clock") || strings.Contains(n, "countdown") || strings.Contains(n, "time") || strings.Contains(n, "hourglass"):
		return "clock"
	}

	patterns := []string{
		"pacman", "cat", "dancer", "ghost", "space-invaders", "pong", "tetris", "snake",
		"cube", "donut", "dna", "earth", "spiral", "matrix", "equalizer", "rocket",
		"neural", "radar", "blackhole", "vehicle", "hsai", "wings", "symbol",
		"heart", "fire", "waves", "rain", "particles", "stars", "clock", "rotate", "bounce",
	}
	return patterns[hashString(name)%len(patterns)]
}

func generatePattern(pattern string, w, h, frameNum, seed int, animName string) string {
	gridH := h - 1
	if gridH < 1 {
		gridH = 1
	}
	grid := make([][]rune, gridH)
	for i := range grid {
		grid[i] = make([]rune, w)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	switch pattern {
	case "pacman":
		drawPacman(grid, frameNum, seed)
	case "cat":
		drawCat(grid, frameNum, seed)
	case "dancer":
		drawDancer(grid, frameNum, seed)
	case "ghost":
		drawGhost(grid, frameNum, seed)
	case "space-invaders":
		drawSpaceInvaders(grid, frameNum, seed)
	case "pong":
		drawPong(grid, frameNum, seed)
	case "tetris":
		drawTetris(grid, frameNum, seed)
	case "snake":
		drawSnake(grid, frameNum, seed)
	case "cube":
		drawCube(grid, frameNum, seed)
	case "donut":
		drawDonut(grid, frameNum, seed)
	case "dna":
		drawDNA(grid, frameNum, seed)
	case "earth":
		drawEarth(grid, frameNum, seed)
	case "spiral":
		drawSpiral(grid, frameNum, seed)
	case "matrix":
		drawMatrix(grid, frameNum, seed)
	case "equalizer":
		drawEqualizer(grid, frameNum, seed)
	case "rocket":
		drawRocket(grid, frameNum, seed)
	case "neural":
		drawNeuralNetwork(grid, frameNum, seed)
	case "radar":
		drawRadar(grid, frameNum, seed)
	case "blackhole":
		drawBlackhole(grid, frameNum, seed)
	case "vehicle":
		drawVehicle(grid, frameNum, seed)
	case "hsai":
		drawHsAiLogo(grid, frameNum, seed)
	case "wings":
		drawWings(grid, frameNum, seed)
	case "symbol":
		drawSymbol(grid, frameNum, seed)
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
	case "rotate":
		drawRotate(grid, frameNum, seed)
	case "stars":
		drawStars(grid, frameNum, seed)
	case "clock":
		drawClock(grid, frameNum, seed)
	default:
		drawBounce(grid, frameNum, seed)
	}

	nameDisplay := animName
	if nameDisplay == "" {
		nameDisplay = pattern
	}
	title := fmt.Sprintf("── HS ASCII LIVE: %s [%s] ──", strings.ToUpper(strings.ReplaceAll(nameDisplay, "-", " ")), strings.ToUpper(pattern))
	if len(title) > w {
		title = fmt.Sprintf("── HS ASCII LIVE: %s ──", strings.ToUpper(nameDisplay))
	}
	if len(title) > w {
		title = "HS ASCII LIVE - " + strings.ToUpper(nameDisplay)
	}
	if len(title) > w && w >= 13 {
		title = "HS ASCII LIVE"
	} else if len(title) > w {
		title = title[:w]
	}

	headerRow := make([]rune, w)
	for j := range headerRow {
		headerRow[j] = ' '
	}
	startCol := (w - len(title)) / 2
	if startCol < 0 {
		startCol = 0
	}
	for j, r := range title {
		if startCol+j < w {
			headerRow[startCol+j] = r
		}
	}

	var sb strings.Builder
	sb.WriteString(string(headerRow))
	for i := 0; i < gridH; i++ {
		sb.WriteString("\r\n")
		sb.WriteString(string(grid[i]))
	}
	return sb.String()
}
