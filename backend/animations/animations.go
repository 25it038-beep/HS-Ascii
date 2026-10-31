package animations

import (
	"fmt"
	"hs-ascii-live/engine"
	"strings"
)

func RegisterAll(r *engine.Registry) {
	categories := map[string][]string{
		"characters": {
			"rickroll", "disco-dancer", "pikachu", "cat", "dog", "among-us", "pacman", "mario", "sonic", "skull", "ghost", "dancing-stickman", "robot", "alien", "spider", "ninja", "samurai", "viking", "pirate", "wizard",
			"batman", "superman", "ironman", "spiderman", "yoda", "darth-vader", "gandalf", "cyber-ninja",
			"naruto-run", "goku-kamehameha", "luffy-gear", "charizard-flame", "bowser-fire", "knuckles-punch", "megaman-buster", "thor-hammer", "joker-laugh", "venom-symbiote", "death-star-laser", "millennium-falcon",
		},
		"space": {
			"galaxy", "milky-way", "blackhole", "solar-system", "rotating-earth", "moon-phases", "saturn", "shooting-stars", "meteor-shower", "supernova", "wormhole", "spaceship", "rocket-launch", "astronaut", "ufo", "comet", "asteroid-belt", "nebula", "constellation", "crab-nebula",
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
			"rotating-cube", "3d-sphere", "spinning-donut", "dna-helix", "fractal", "tunnel", "particle-explosion", "kaleidoscope", "infinity-loop", "waves", "spiral", "morphing-shapes", "vortex", "optical-illusion", "meta-cube", "mandelbrot-set", "golden-spiral",
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
			"minecraft-steve", "portal-aperture", "half-life-lambda", "doom-slayer", "street-fighter-hadouken", "mortal-kombat-subzero", "zelda-sword-master", "pokeball-catch", "donkey-kong-barrels", "contra-spreadgun", "galaga-fighter", "asteroids-ship", "dig-dug", "frogger-crossing", "metroid-samus",
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

	// Register Cinematic Experience as official intro and cinematic animation
	cineNames := []string{"intro", "cinematic", "cinema", "hs-cinematic", "cinematic-intro", "hs-ascii-live-intro"}
	for _, cn := range cineNames {
		r.Register(NewCinematicAnimation(cn))
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

func setCell(grid [][]rune, x, y int, ch rune) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x] = ch
	}
}

func patternForName(name string) string {
	n := strings.ToLower(name)
	switch {
	// Unique Character & Pop Culture engines
	case strings.Contains(n, "mario") || strings.Contains(n, "bowser"):
		return "mario"
	case strings.Contains(n, "sonic") || strings.Contains(n, "knuckles") || strings.Contains(n, "donkey") || strings.Contains(n, "saxophone"):
		return "sonic"
	case strings.Contains(n, "pikachu") || strings.Contains(n, "charizard") || strings.Contains(n, "gameboy"):
		return "pikachu"
	case strings.Contains(n, "dog"):
		return "dog"
	case strings.Contains(n, "among-us"):
		return "among-us"
	case strings.Contains(n, "batman") || strings.Contains(n, "joker") || strings.Contains(n, "gandalf") || strings.Contains(n, "wizard"):
		return "batman"
	case strings.Contains(n, "superman") || strings.Contains(n, "ironman") || strings.Contains(n, "thor") || strings.Contains(n, "hero") || strings.Contains(n, "flight"):
		return "hero-flight"
	case strings.Contains(n, "darth") || (strings.Contains(n, "vader") && !strings.Contains(n, "invader")) || strings.Contains(n, "yoda") || strings.Contains(n, "death-star") || strings.Contains(n, "falcon") || strings.Contains(n, "samurai") || strings.Contains(n, "ninja") || strings.Contains(n, "sword") || strings.Contains(n, "saber"):
		return "lightsaber"
	case strings.Contains(n, "spider") || strings.Contains(n, "venom"):
		return "spider"
	case strings.Contains(n, "skull"):
		return "skull"
	case strings.Contains(n, "ghost"):
		return "ghost"
	case strings.Contains(n, "pacman"):
		return "pacman"
	case strings.Contains(n, "cat"):
		return "cat"
	case strings.Contains(n, "stickman") || strings.Contains(n, "rick") || strings.Contains(n, "dance") || strings.Contains(n, "naruto") || strings.Contains(n, "goku") || strings.Contains(n, "luffy"):
		return "dancer"

	// Sports engines
	case strings.Contains(n, "basketball"):
		return "basketball"
	case strings.Contains(n, "tennis"):
		return "tennis"
	case strings.Contains(n, "bowling"):
		return "bowling"
	case strings.Contains(n, "soccer") || strings.Contains(n, "kick") || strings.Contains(n, "sports") || strings.Contains(n, "baseball") || strings.Contains(n, "golf") || strings.Contains(n, "boxing") || strings.Contains(n, "weightlifter"):
		return "sports"

	// Food engines
	case strings.Contains(n, "burger") || strings.Contains(n, "taco"):
		return "burger"
	case strings.Contains(n, "pizza"):
		return "pizza"
	case strings.Contains(n, "sushi"):
		return "sushi"
	case strings.Contains(n, "boba") || strings.Contains(n, "ice-cream"):
		return "boba"
	case strings.Contains(n, "coffee") || strings.Contains(n, "boba-tea") || strings.HasSuffix(n, "-tea") || strings.Contains(n, "ramen") || strings.Contains(n, "toast") || strings.Contains(n, "champagne"):
		return "coffee"

	// Vehicles engines
	case strings.Contains(n, "formula1") || strings.Contains(n, "f1"):
		return "f1-racer"
	case strings.Contains(n, "helicopter"):
		return "helicopter"
	case strings.Contains(n, "submarine") || strings.Contains(n, "u-boat"):
		return "submarine"
	case strings.Contains(n, "locomotive") || strings.Contains(n, "train"):
		return "train"
	case strings.Contains(n, "sailing") || strings.Contains(n, "galleon") || strings.Contains(n, "ship") || strings.Contains(n, "pirate") || strings.Contains(n, "viking") || strings.Contains(n, "zeppelin"):
		return "sailing-ship"
	case strings.Contains(n, "motorcycle") || strings.Contains(n, "chopper") || strings.Contains(n, "bike") || strings.Contains(n, "cycling") || strings.Contains(n, "skater"):
		return "motorcycle"
	case strings.Contains(n, "airplane") || strings.Contains(n, "bomber"):
		return "airplane"
	case strings.Contains(n, "rocket") || strings.Contains(n, "shuttle") || strings.Contains(n, "saturn-v") || strings.Contains(n, "apollo") || strings.Contains(n, "spaceship") || strings.Contains(n, "astronaut"):
		return "rocket"
	case strings.Contains(n, "cybertruck") || strings.Contains(n, "car") || strings.Contains(n, "truck") || strings.Contains(n, "racer") || strings.Contains(n, "rover"):
		return "vehicle"

	// Music engines
	case strings.Contains(n, "piano") || strings.Contains(n, "keyboard"):
		return "piano"
	case strings.Contains(n, "vinyl") || strings.Contains(n, "record") || strings.Contains(n, "turntable"):
		return "vinyl"
	case strings.Contains(n, "guitar"):
		return "guitar"
	case strings.Contains(n, "boombox") || strings.Contains(n, "headphone"):
		return "boombox"
	case strings.Contains(n, "cassette"):
		return "cassette"
	case strings.Contains(n, "disco"):
		return "disco"

	// Animals & Nature engines
	case strings.Contains(n, "wolf") || strings.Contains(n, "bear") || strings.Contains(n, "panther"):
		return "wolf"
	case strings.Contains(n, "frog"):
		return "frog"
	case strings.Contains(n, "whale") || strings.Contains(n, "dolphin") || strings.Contains(n, "docker"):
		return "whale"
	case strings.Contains(n, "horse"):
		return "horse"
	case strings.Contains(n, "saturn"):
		return "saturn"

	// Retro & Sci-Fi engines
	case strings.Contains(n, "crying") || strings.Contains(n, "face"):
		return "ghost"
	case strings.Contains(n, "mantis"):
		return "spider"
	case strings.Contains(n, "panther"):
		return "cat"
	case strings.Contains(n, "hologram") || strings.Contains(n, "visor"):
		return "radar"
	case strings.Contains(n, "hadouken") || strings.Contains(n, "subzero"):
		return "hero-flight"
	case strings.Contains(n, "dig-dug"):
		return "snake"
	case strings.Contains(n, "city") || strings.Contains(n, "skyline") || strings.Contains(n, "synthwave") || strings.Contains(n, "megacity"):
		return "city"
	case strings.Contains(n, "terminal") || strings.Contains(n, "hacking") || strings.Contains(n, "boot") || strings.Contains(n, "vim") || strings.Contains(n, "computer"):
		return "terminal"
	case strings.Contains(n, "invader") || strings.Contains(n, "galaga") || strings.Contains(n, "asteroid") || strings.Contains(n, "fighter"):
		return "space-invaders"
	case strings.Contains(n, "pong"):
		return "pong"
	case strings.Contains(n, "tetris"):
		return "tetris"
	case strings.Contains(n, "snake"):
		return "snake"
	case strings.Contains(n, "pokeball"):
		return "pokeball"
	case strings.Contains(n, "tux") || strings.Contains(n, "linux"):
		return "tux"
	case strings.Contains(n, "matrix") || strings.Contains(n, "binary") || strings.Contains(n, "cyber-grid"):
		return "matrix"
	case strings.Contains(n, "radar") || strings.Contains(n, "scan"):
		return "radar"
	case strings.Contains(n, "blackhole") || strings.Contains(n, "singularity"):
		return "blackhole"
	case strings.Contains(n, "wormhole") || strings.Contains(n, "portal") || strings.Contains(n, "cyberspace"):
		return "wormhole"
	case strings.Contains(n, "tornado") || strings.Contains(n, "cyclone") || strings.Contains(n, "typhoon") || strings.Contains(n, "gale") || strings.Contains(n, "wind"):
		return "tornado"
	case strings.Contains(n, "volcano") || strings.Contains(n, "cave"):
		return "volcano"
	case strings.Contains(n, "waterfall") || strings.Contains(n, "geyser"):
		return "waterfall"
	case strings.Contains(n, "rainbow"):
		return "rainbow"
	case strings.Contains(n, "aurora"):
		return "aurora"
	case strings.Contains(n, "firework") || strings.Contains(n, "supernova"):
		return "firework"
	case strings.Contains(n, "eclipse") || strings.Contains(n, "sunset") || strings.Contains(n, "sunrise"):
		return "eclipse"
	case strings.Contains(n, "hourglass"):
		return "hourglass"
	case strings.Contains(n, "countdown"):
		return "countdown"

	// 3D & Math engines
	case strings.Contains(n, "donut") || strings.Contains(n, "torus"):
		return "donut"
	case strings.Contains(n, "cube") || strings.Contains(n, "tesseract") || strings.Contains(n, "minecraft") || strings.Contains(n, "meta-cube") || strings.Contains(n, "wireframe"):
		return "cube"
	case strings.Contains(n, "dna") || strings.Contains(n, "helix"):
		return "dna"
	case strings.Contains(n, "earth") || strings.Contains(n, "globe") || strings.Contains(n, "sphere") || strings.Contains(n, "planet") || strings.Contains(n, "mars") || strings.Contains(n, "jupiter") || strings.Contains(n, "moon") || strings.Contains(n, "uranus") || strings.Contains(n, "europa") || strings.Contains(n, "titan") || strings.Contains(n, "trappist") || strings.Contains(n, "solar"):
		return "earth"
	case strings.Contains(n, "spiral") || strings.Contains(n, "tunnel") || strings.Contains(n, "vortex") || strings.Contains(n, "fractal") || strings.Contains(n, "mandelbrot") || strings.Contains(n, "kaleidoscope") || strings.Contains(n, "optical"):
		return "spiral"
	case strings.Contains(n, "equalizer") || strings.Contains(n, "spectrum") || strings.Contains(n, "audio") || strings.Contains(n, "soundwave") || strings.Contains(n, "drum"):
		return "equalizer"
	case strings.Contains(n, "neural") || strings.Contains(n, "brain") || strings.Contains(n, "network") || strings.Contains(n, "synapse") || strings.Contains(n, "omni") || strings.Contains(n, "ai-"):
		return "neural"
	case strings.Contains(n, "hs-") || strings.Contains(n, "hsai") || strings.Contains(n, "gravity") || strings.Contains(n, "vision") || strings.Contains(n, "gpu") || strings.Contains(n, "microchip") || strings.Contains(n, "quantum") || strings.Contains(n, "crypto") || strings.Contains(n, "blockchain") || strings.Contains(n, "server") || strings.Contains(n, "mainframe") || strings.Contains(n, "satellite") || strings.Contains(n, "kubernetes") || strings.Contains(n, "git"):
		return "hsai"
	case strings.Contains(n, "dragon") || strings.Contains(n, "phoenix") || strings.Contains(n, "bird") || strings.Contains(n, "bat") || strings.Contains(n, "butterfly") || strings.Contains(n, "eagle") || strings.Contains(n, "wing"):
		return "wings"
	case strings.Contains(n, "yin-yang") || strings.Contains(n, "peace") || strings.Contains(n, "hazard") || strings.Contains(n, "biohazard") || strings.Contains(n, "infinity") || strings.Contains(n, "triforce") || strings.Contains(n, "lambda"):
		return "symbol"
	case strings.Contains(n, "heart") || strings.Contains(n, "pulse"):
		return "heart"
	case strings.Contains(n, "fire") || strings.Contains(n, "flame") || strings.Contains(n, "doom"):
		return "fire"
	case strings.Contains(n, "water") || strings.Contains(n, "ocean") || strings.Contains(n, "wave") || strings.Contains(n, "river") || strings.Contains(n, "sea") || strings.Contains(n, "surf") || strings.Contains(n, "swimming") || strings.Contains(n, "octopus"):
		return "waves"
	case strings.Contains(n, "rain") || strings.Contains(n, "storm") || strings.Contains(n, "monsoon") || strings.Contains(n, "hail") || strings.Contains(n, "lightning"):
		return "rain"
	case strings.Contains(n, "snow") || strings.Contains(n, "blizzard") || strings.Contains(n, "frost") || strings.Contains(n, "iceberg") || strings.Contains(n, "leaves") || strings.Contains(n, "tree") || strings.Contains(n, "forest") || strings.Contains(n, "sand") || strings.Contains(n, "dunes") || strings.Contains(n, "oasis") || strings.Contains(n, "crystal") || strings.Contains(n, "coral") || strings.Contains(n, "pop") || strings.Contains(n, "corn"):
		return "particles"
	case strings.Contains(n, "clock") || strings.Contains(n, "time") || strings.Contains(n, "metronome"):
		return "clock"
	case strings.Contains(n, "galaxy") || strings.Contains(n, "milky") || strings.Contains(n, "star") || strings.Contains(n, "nebula") || strings.Contains(n, "constellation") || strings.Contains(n, "pulsar") || strings.Contains(n, "quasar") || strings.Contains(n, "telescope") || strings.Contains(n, "hubble") || strings.Contains(n, "andromeda") || strings.Contains(n, "deep-space") || strings.Contains(n, "comet") || strings.Contains(n, "meteor") || strings.Contains(n, "oort"):
		return "stars"
	}

	patterns := []string{
		"mario", "sonic", "pikachu", "dog", "among-us", "skull", "batman", "hero-flight", "lightsaber", "spider",
		"basketball", "tennis", "bowling", "sports", "burger", "pizza", "coffee", "helicopter", "submarine", "train",
		"sailing-ship", "motorcycle", "airplane", "rocket", "vehicle", "piano", "vinyl", "guitar", "disco", "city",
		"terminal", "space-invaders", "pong", "tetris", "snake", "pokeball", "tux", "matrix", "radar", "blackhole",
		"wormhole", "tornado", "volcano", "waterfall", "rainbow", "aurora", "firework", "eclipse", "hourglass", "countdown",
		"donut", "cube", "dna", "earth", "spiral", "equalizer", "neural", "hsai", "wings", "symbol",
		"heart", "fire", "waves", "rain", "particles", "clock", "stars", "rotate", "bounce",
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
	case "f1-racer":
		drawF1Racer(grid, frameNum, seed)
	case "boombox":
		drawBoombox(grid, frameNum, seed)
	case "cassette":
		drawCassette(grid, frameNum, seed)
	case "boba":
		drawBoba(grid, frameNum, seed, animName)
	case "sushi":
		drawSushi(grid, frameNum, seed)
	case "saturn":
		drawSaturn(grid, frameNum, seed)
	case "wolf":
		drawWolf(grid, frameNum, seed, animName)
	case "frog":
		drawFrog(grid, frameNum, seed)
	case "whale":
		drawWhale(grid, frameNum, seed, animName)
	case "horse":
		drawHorse(grid, frameNum, seed)
	case "mario":
		drawMario(grid, frameNum, seed, animName)
	case "sonic":
		drawSonic(grid, frameNum, seed, animName)
	case "pikachu":
		drawPikachu(grid, frameNum, seed, animName)
	case "dog":
		drawDog(grid, frameNum, seed)
	case "among-us":
		drawAmongUs(grid, frameNum, seed)
	case "skull":
		drawSkull(grid, frameNum, seed, animName)
	case "batman":
		drawBatman(grid, frameNum, seed, animName)
	case "hero-flight":
		drawHeroFlight(grid, frameNum, seed, animName)
	case "lightsaber":
		drawLightsaber(grid, frameNum, seed, animName)
	case "spider":
		drawSpider(grid, frameNum, seed, animName)
	case "basketball":
		drawBasketball(grid, frameNum, seed)
	case "tennis":
		drawTennis(grid, frameNum, seed)
	case "bowling":
		drawBowling(grid, frameNum, seed)
	case "burger":
		drawBurger(grid, frameNum, seed)
	case "pizza":
		drawPizza(grid, frameNum, seed)
	case "helicopter":
		drawHelicopter(grid, frameNum, seed)
	case "submarine":
		drawSubmarine(grid, frameNum, seed)
	case "train":
		drawTrain(grid, frameNum, seed)
	case "sailing-ship":
		drawSailingShip(grid, frameNum, seed, animName)
	case "motorcycle":
		drawMotorcycle(grid, frameNum, seed, animName)
	case "airplane":
		drawHeroFlight(grid, frameNum, seed, animName)
	case "eclipse":
		drawEclipse(grid, frameNum, seed)
	case "hourglass":
		drawHourglass(grid, frameNum, seed)
	case "countdown":
		drawCountdown(grid, frameNum, seed)
	case "volcano":
		drawVolcano(grid, frameNum, seed)
	case "waterfall":
		drawWaterfall(grid, frameNum, seed)
	case "rainbow":
		drawRainbow(grid, frameNum, seed)
	case "terminal":
		drawTerminalHacker(grid, frameNum, seed)
	case "guitar":
		drawGuitar(grid, frameNum, seed)
	case "firework":
		drawFirework(grid, frameNum, seed)
	case "aurora":
		drawAurora(grid, frameNum, seed)
	case "piano":
		drawPiano(grid, frameNum, seed)
	case "vinyl":
		drawVinyl(grid, frameNum, seed)
	case "city":
		drawCitySkyline(grid, frameNum, seed)
	case "coffee":
		drawCoffee(grid, frameNum, seed)
	case "sports":
		drawSports(grid, frameNum, seed, animName)
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
		drawCat(grid, frameNum, seed, animName)
	case "dancer":
		drawDancer(grid, frameNum, seed, animName)
	case "ghost":
		drawGhost(grid, frameNum, seed, animName)
	case "space-invaders":
		drawSpaceInvaders(grid, frameNum, seed)
	case "pong":
		drawPong(grid, frameNum, seed)
	case "tetris":
		drawTetris(grid, frameNum, seed)
	case "snake":
		drawSnake(grid, frameNum, seed)
	case "cube":
		drawCube(grid, frameNum, seed, animName)
	case "donut":
		drawDonut(grid, frameNum, seed)
	case "dna":
		drawDNA(grid, frameNum, seed)
	case "earth":
		drawEarth(grid, frameNum, seed, animName)
	case "spiral", "wormhole":
		drawSpiral(grid, frameNum, seed, animName)
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
		drawVehicle(grid, frameNum, seed, animName)
	case "hsai":
		drawHsAiLogo(grid, frameNum, seed, animName)
	case "wings":
		drawWings(grid, frameNum, seed, animName)
	case "symbol":
		drawSymbol(grid, frameNum, seed, animName)
	case "heart":
		drawHeart(grid, frameNum, seed, animName)
	case "rain":
		drawRain(grid, frameNum, seed, animName)
	case "fire":
		drawFire(grid, frameNum, seed, animName)
	case "waves":
		drawWaves(grid, frameNum, seed, animName)
	case "particles":
		drawParticles(grid, frameNum, seed, animName)
	case "rotate":
		drawRotate(grid, frameNum, seed)
	case "stars":
		drawStars(grid, frameNum, seed, animName)
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
