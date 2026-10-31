package animations

import (
	"fmt"
	"hs-ascii-live/engine"
	"strings"
)

func RegisterAll(r *engine.Registry) {
	categories := map[string][]string{
		"characters": {
			"rickroll", "dancing-rick", "pikachu", "cat", "dog", "among-us", "pacman", "mario", "sonic", "skull", "ghost", "dancing-stickman", "robot", "alien", "spider", "ninja", "samurai", "viking", "pirate", "wizard",
			"batman", "superman", "ironman", "spiderman", "yoda", "darth-vader", "gandalf", "cyber-ninja",
			"naruto-run", "goku-kamehameha", "luffy-gear", "pikachu-thunder", "super-mario-jump", "sonic-spin", "megaman-buster", "ironman-flight", "batman-signal", "spider-web-swing", "darth-vader-choke", "yoda-force",
		},
		"space": {
			"galaxy", "milky-way", "blackhole", "solar-system", "rotating-earth", "moon-phases", "saturn", "shooting-stars", "meteor-shower", "supernova", "wormhole", "spaceship", "rocket-launch", "astronaut", "ufo", "comet", "asteroid-belt", "nebula", "constellation", "milky-way2",
			"jupiter-storm", "pulsar", "quasars", "rings-of-uranus", "solar-eclipse", "deep-space-probe",
			"james-webb-telescope", "mars-rover-curiosity", "iss-spacestation", "hubble-deepfield", "trappist-system", "andromeda-collision", "europa-geysers", "titan-methane-lakes", "oort-cloud", "pulsar-beacon",
		},
		"nature": {
			"fire", "waterfall", "ocean-waves", "rain", "thunderstorm", "tornado", "snowfall", "forest", "blooming-flowers", "sunrise", "sunset", "volcano", "lightning", "wind", "geyser", "aurora", "cave", "river",
			"blizzard", "autumn-leaves", "desert-dunes", "coral-reef", "bamboo-forest", "meteor-impact",
			"coral-spawning", "whale-breech", "geyser-eruption", "sequoia-tree", "bioluminescence-sea", "crystal-cavern", "iceberg-calving", "desert-oasis", "wildfire-sweep", "canyon-river",
		},
		"technology": {
			"matrix-rain", "hacking-terminal", "computer-boot", "digital-glitch", "network-connections", "binary-rain", "ai-neural-network", "circuit-board", "cpu-processing", "gpu-rendering", "data-streams", "cyber-grid", "digital-face", "quantum-computing", "ai-brain", "encryption", "quantum-teleport", "hologram", "server-farm", "drone-swarm",
			"blockchain-mining", "cyber-skull", "radar-scanner", "synthwave-sun", "laser-grid", "terminal-matrix",
			"docker-whale", "linux-tux-penguin", "kubernetes-wheel", "git-branch-merge", "vim-editor", "crypto-wallet", "microchip-die", "fiber-optic-pulse", "mainframe-tape", "satellite-uplink",
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
			"hs-quantum-leap", "hs-autonomous-agent", "hs-omni-intelligence", "hs-vision-core-v3", "hs-neural-synapse", "hs-sentient-stream", "hs-apex-model", "hs-matrix-titan", "hs-infinite-scale", "hs-hyper-drive",
		},
		"gaming": {
			"space-invaders", "pong", "tetris", "flappy-bird", "dino-runner", "snake-game", "doom-fire", "arcade-joystick", "gameboy", "zelda-triforce",
			"minecraft-steve", "portal-aperture", "half-life-lambda", "doom-slayer", "street-fighter-hadouken", "mortal-kombat-subzero", "zelda-sword-master", "pokeball-catch", "donkey-kong-barrels", "contra-spreadgun", "galaga-fighter", "asteroids-ship", "dig-dug", "frogger-crossing", "mega-man",
		},
		"vehicles": {
			"cybertruck", "speeding-car", "steam-locomotive", "airplane-flight", "submarine", "motorcycle", "helicopter", "hovercraft", "sailing-ship", "space-shuttle",
			"formula1-racer", "bullet-train-shinkansen", "stealth-bomber", "apollo-saturn-v", "harley-chopper", "monster-truck", "zeppelin-airship", "pirate-galleon", "u-boat-periscope", "hover-bike",
		},
		"creatures": {
			"roaring-lion", "dragon-flight", "phoenix-rising", "flying-bat", "swimming-whale", "hopping-frog", "galloping-horse", "slithering-snake", "flapping-butterfly", "howling-wolf",
			"majestic-eagle", "stealth-panther", "charging-rhino", "leaping-dolphin", "chameleon-blend", "octopus-ink", "owl-stare", "peacock-feather", "mantis-strike", "polar-bear",
		},
		"music": {
			"piano-keys", "vinyl-record", "drum-kit", "electric-guitar", "saxophone-jazz", "synthesizer-wave", "cassette-tape", "boombox", "audio-spectrum", "turntable-scratch", "soundwave-pulse", "disco-ball", "metronome", "microphone-glow", "headphones-beat",
		},
		"sports": {
			"soccer-kick", "basketball-dunk", "tennis-match", "baseball-swing", "skater-kickflip", "surfer-barrel", "skiing-downhill", "boxing-punch", "swimming-freestyle", "bowling-strike", "archery-bullseye", "cycling-sprint", "golf-swing", "gymnast-flip", "weightlifter-snatch",
		},
		"cyberpunk": {
			"neon-city-skyline", "cyber-samurai", "synthwave-highway", "hologram-girl", "glitch-skull", "neon-sunset", "cyber-monolith", "wireframe-landscape", "cyberspace-portal", "megacity-spire", "data-core-breach", "quantum-flux", "cyber-visor", "neural-sync", "orbital-elevator",
		},
		"food": {
			"steaming-coffee", "sizzling-pizza", "burger-stack", "ramen-noodles", "popping-corn", "ice-cream-cone", "sushi-roll", "boba-tea", "taco-tuesday", "donut-glaze", "champagne-toast", "birthday-cake",
		},
		"weather": {
			"typhoon-cyclone", "sandstorm-desert", "monsoon-deluge", "rainbow-arc", "hailstorm", "solar-flare", "foggy-morning", "lightning-strike", "autumn-gale", "frost-crystals", "dust-devil", "spring-breeze",
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
	case strings.Contains(n, "piano") || strings.Contains(n, "keyboard"):
		return "piano"
	case strings.Contains(n, "vinyl") || strings.Contains(n, "record") || strings.Contains(n, "turntable"):
		return "vinyl"
	case strings.Contains(n, "city") || strings.Contains(n, "skyline") || strings.Contains(n, "synthwave") || strings.Contains(n, "megacity") || strings.Contains(n, "landscape"):
		return "city"
	case strings.Contains(n, "coffee") || strings.Contains(n, "tea") || strings.Contains(n, "ramen") || strings.Contains(n, "pizza") || strings.Contains(n, "burger") || strings.Contains(n, "taco") || strings.Contains(n, "food") || strings.Contains(n, "sushi"):
		return "coffee"
	case strings.Contains(n, "soccer") || strings.Contains(n, "basketball") || strings.Contains(n, "tennis") || strings.Contains(n, "ball") || strings.Contains(n, "sports") || strings.Contains(n, "boxing") || strings.Contains(n, "strike"):
		return "sports"
	case strings.Contains(n, "tux") || strings.Contains(n, "linux") || strings.Contains(n, "penguin"):
		return "tux"
	case strings.Contains(n, "pokeball") || strings.Contains(n, "pokemon") || strings.Contains(n, "pokéball"):
		return "pokeball"
	case strings.Contains(n, "tornado") || strings.Contains(n, "cyclone") || strings.Contains(n, "typhoon") || strings.Contains(n, "dust-devil"):
		return "tornado"
	case strings.Contains(n, "disco") || strings.Contains(n, "party"):
		return "disco"
	case strings.Contains(n, "pacman"):
		return "pacman"
	case strings.Contains(n, "cat"):
		return "cat"
	case strings.Contains(n, "stickman") || strings.Contains(n, "rick") || strings.Contains(n, "dance"):
		return "dancer"
	case strings.Contains(n, "skull") || strings.Contains(n, "ghost"):
		return "ghost"
	case strings.Contains(n, "invader") || strings.Contains(n, "galaga") || strings.Contains(n, "asteroid"):
		return "space-invaders"
	case strings.Contains(n, "pong"):
		return "pong"
	case strings.Contains(n, "tetris"):
		return "tetris"
	case strings.Contains(n, "snake"):
		return "snake"
	case strings.Contains(n, "donut") || strings.Contains(n, "torus"):
		return "donut"
	case strings.Contains(n, "cube") || strings.Contains(n, "tesseract") || strings.Contains(n, "minecraft"):
		return "cube"
	case strings.Contains(n, "dna") || strings.Contains(n, "helix"):
		return "dna"
	case strings.Contains(n, "matrix") || strings.Contains(n, "terminal") || strings.Contains(n, "binary") || strings.Contains(n, "code"):
		return "matrix"
	case strings.Contains(n, "equalizer") || strings.Contains(n, "music") || strings.Contains(n, "audio") || strings.Contains(n, "spectrum") || strings.Contains(n, "soundwave") || strings.Contains(n, "boombox") || strings.Contains(n, "drum") || strings.Contains(n, "guitar") || strings.Contains(n, "headphone"):
		return "equalizer"
	case strings.Contains(n, "rocket") || strings.Contains(n, "spaceship") || strings.Contains(n, "shuttle") || strings.Contains(n, "probe") || strings.Contains(n, "saturn-v") || strings.Contains(n, "apollo"):
		return "rocket"
	case strings.Contains(n, "earth") || strings.Contains(n, "globe") || strings.Contains(n, "sphere") || strings.Contains(n, "moon") || strings.Contains(n, "saturn") || strings.Contains(n, "planet") || strings.Contains(n, "mars") || strings.Contains(n, "jupiter"):
		return "earth"
	case strings.Contains(n, "car") || strings.Contains(n, "truck") || strings.Contains(n, "vehicle") || strings.Contains(n, "locomotive") || strings.Contains(n, "train") || strings.Contains(n, "motorcycle") || strings.Contains(n, "submarine") || strings.Contains(n, "flight") || strings.Contains(n, "helicopter") || strings.Contains(n, "ship") || strings.Contains(n, "racer") || strings.Contains(n, "bomber"):
		return "vehicle"
	case strings.Contains(n, "neural") || strings.Contains(n, "brain") || strings.Contains(n, "network") || strings.Contains(n, "agent") || strings.Contains(n, "model") || strings.Contains(n, "connect") || strings.Contains(n, "synapse") || strings.Contains(n, "omni") || strings.Contains(n, "apex"):
		return "neural"
	case strings.Contains(n, "radar") || strings.Contains(n, "scan"):
		return "radar"
	case strings.Contains(n, "blackhole") || strings.Contains(n, "wormhole") || strings.Contains(n, "vortex") || strings.Contains(n, "singularity") || strings.Contains(n, "portal"):
		return "blackhole"
	case strings.Contains(n, "hs-") || strings.Contains(n, "hsai") || strings.Contains(n, "gravity") || strings.Contains(n, "vision"):
		return "hsai"
	case strings.Contains(n, "dragon") || strings.Contains(n, "phoenix") || strings.Contains(n, "bird") || strings.Contains(n, "bat") || strings.Contains(n, "butterfly") || strings.Contains(n, "wing") || strings.Contains(n, "eagle"):
		return "wings"
	case strings.Contains(n, "yin-yang") || strings.Contains(n, "peace") || strings.Contains(n, "hazard") || strings.Contains(n, "biohazard") || strings.Contains(n, "infinity") || strings.Contains(n, "lambda") || strings.Contains(n, "triforce"):
		return "symbol"
	case strings.Contains(n, "heart"):
		return "heart"
	case strings.Contains(n, "fire") || strings.Contains(n, "volcano") || strings.Contains(n, "doom") || strings.Contains(n, "flame") || strings.Contains(n, "wildfire"):
		return "fire"
	case strings.Contains(n, "water") || strings.Contains(n, "ocean") || strings.Contains(n, "wave") || strings.Contains(n, "river") || strings.Contains(n, "waterfall") || strings.Contains(n, "sea") || strings.Contains(n, "surf"):
		return "waves"
	case strings.Contains(n, "rain") || strings.Contains(n, "storm") || strings.Contains(n, "deluge") || strings.Contains(n, "monsoon") || strings.Contains(n, "hail"):
		return "rain"
	case strings.Contains(n, "snow") || strings.Contains(n, "blizzard") || strings.Contains(n, "frost") || strings.Contains(n, "iceberg"):
		return "particles"
	case strings.Contains(n, "spiral") || strings.Contains(n, "tunnel"):
		return "spiral"
	case strings.Contains(n, "galaxy") || strings.Contains(n, "milky") || strings.Contains(n, "star") || strings.Contains(n, "supernova") || strings.Contains(n, "nebula") || strings.Contains(n, "constellation") || strings.Contains(n, "meteor") || strings.Contains(n, "pulsar") || strings.Contains(n, "quasar") || strings.Contains(n, "eclipse") || strings.Contains(n, "telescope") || strings.Contains(n, "hubble") || strings.Contains(n, "andromeda"):
		return "stars"
	case strings.Contains(n, "clock") || strings.Contains(n, "countdown") || strings.Contains(n, "time") || strings.Contains(n, "hourglass") || strings.Contains(n, "metronome"):
		return "clock"
	}

	patterns := []string{
		"pacman", "cat", "dancer", "ghost", "space-invaders", "pong", "tetris", "snake",
		"cube", "donut", "dna", "earth", "spiral", "matrix", "equalizer", "rocket",
		"neural", "radar", "blackhole", "vehicle", "hsai", "wings", "symbol",
		"heart", "fire", "waves", "rain", "particles", "stars", "clock", "rotate", "bounce",
		"piano", "vinyl", "city", "coffee", "sports", "tux", "pokeball", "tornado", "disco",
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
	case "piano":
		drawPiano(grid, frameNum, seed)
	case "vinyl":
		drawVinyl(grid, frameNum, seed)
	case "city":
		drawCitySkyline(grid, frameNum, seed)
	case "coffee":
		drawCoffee(grid, frameNum, seed)
	case "sports":
		drawSports(grid, frameNum, seed)
	case "tux":
		drawTux(grid, frameNum, seed)
	case "pokeball":
		drawPokeball(grid, frameNum, seed)
	case "tornado":
		drawTornado(grid, frameNum, seed)
	case "disco":
		drawDiscoBall(grid, frameNum, seed)
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
