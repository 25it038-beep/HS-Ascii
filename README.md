# HS ASCII LIVE
### By HS AI Solutions

A production-ready, publicly accessible ASCII animation streaming platform with **340+ continuously animated ASCII experiences** across 15 specialized categories.

## Features
- **Real HTTP Streaming**: ANSI escape sequence streaming with in-place overwriting and zero flicker.
- **Full-Screen Terminal Optimized**: Standard full CMD screen (80×24 default), maximized terminal mode (`?full=true` for 120×35), or exact custom dimensions (`?w=...&h=...`).
- **340+ Procedural Animations**: High-performance continuous generators across 15 specialized categories.
- **Rich Rendering Engines**:
  - **Retro Gaming**: Pac-Man, Space Invaders, Pong, Tetris, Snake slithering & food eating, Galaga, Asteroids.
  - **3D Geometry & Mathematics**: True 3D wireframe rotating cube, 3D rotating torus donut (donut.c homage), 3D double helix DNA, and rotating globe sphere.
  - **Tech & Cyberpunk**: Neon city skyline with synthwave sun, Matrix digital rain, multi-band audio equalizer, rocket launch, neural network synapse pulses, radar sweep, and accretion disk blackhole.
  - **Music & Lifestyle**: Piano keyboard with floating notes, spinning vinyl record with tonearm, disco ball light beams, steaming coffee cup with rising vapour.
  - **Sports & Nature**: Parabolic soccer ball kick & goal net, Tux penguin on ice, Pokéball capture, twisting tornado, cellular flames, and official HS AI logo.
- **Concurrent Viewers**: Non-blocking goroutine streaming with graceful context cancellation.
- **REST & Query APIs**: Dynamic `/list`, `/api/animations`, `/api/stats`, and `/random` endpoints.
- **Customizable Parameters**: Speed (`slow`, `normal`, `fast`), FPS (`1-60`), color (`red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`), width, and height.
- **Automated Test Suite**: Unit and integration test coverage for all 340+ animations, pattern generators, and HTTP endpoints.

## Architecture
```
backend/
├── engine/        # Core streaming, frame scheduling, ANSI renderer, registry
├── handlers/      # HTTP streaming handlers, REST metadata, stats, health
├── animations/    # Procedural generators, 3D math projections, retro patterns
├── middleware/    # CORS, rate limiting, access logging
└── config/        # Environment configurations
```

## Local Installation
```bash
cd backend
go mod tidy
go test -v ./...
go run main.go
```

## Endpoints
- `GET /<animation>` - stream continuous animation
- `GET /random` - stream random animation
- `GET /list` - list all animations (JSON)
- `GET /help` - terminal usage instructions
- `GET /health` - health check
- `GET /api/animations` - full JSON catalogue
- `GET /api/stats` - usage and registry statistics

### Examples
```bash
# Retro Gaming & Characters
curl https://hs-ascii-1.onrender.com/pacman
curl https://hs-ascii-1.onrender.com/space-invaders
curl https://hs-ascii-1.onrender.com/pokeball-catch
curl https://hs-ascii-1.onrender.com/minecraft-steve

# Full-Screen Maximized Mode (?full=true)
curl "https://hs-ascii-1.onrender.com/neon-city-skyline?full=true"
curl "https://hs-ascii-1.onrender.com/matrix-rain?full=true&color=green"
curl "https://hs-ascii-1.onrender.com/rotating-cube?full=true"
curl "https://hs-ascii-1.onrender.com/spinning-donut?full=true"

# Music, Lifestyle & Sports
curl https://hs-ascii-1.onrender.com/piano-keys
curl https://hs-ascii-1.onrender.com/vinyl-record
curl https://hs-ascii-1.onrender.com/disco-ball
curl https://hs-ascii-1.onrender.com/steaming-coffee
curl https://hs-ascii-1.onrender.com/soccer-kick

# Sci-Fi & HS AI Exclusive
curl https://hs-ascii-1.onrender.com/hs-ai-solutions-logo?fps=15
curl https://hs-ascii-1.onrender.com/hs-quantum-leap
curl https://hs-ascii-1.onrender.com/dna-helix
```

## Categories (349 Total)
1. **Characters** (40): rickroll, dancing-rick, pikachu, cat, dog, among-us, pacman, mario, sonic, skull, ghost, dancing-stickman, robot, alien, spider, ninja, samurai, viking, pirate, wizard, batman, superman, ironman, spiderman, yoda, darth-vader, gandalf, cyber-ninja, naruto-run, goku-kamehameha, luffy-gear, pikachu-thunder, super-mario-jump, sonic-spin, megaman-buster, ironman-flight, batman-signal, spider-web-swing, darth-vader-choke, yoda-force
2. **Space** (36): galaxy, milky-way, blackhole, solar-system, rotating-earth, moon-phases, saturn, shooting-stars, meteor-shower, supernova, wormhole, spaceship, rocket-launch, astronaut, ufo, comet, asteroid-belt, nebula, constellation, milky-way2, jupiter-storm, pulsar, quasars, rings-of-uranus, solar-eclipse, deep-space-probe, james-webb-telescope, mars-rover-curiosity, iss-spacestation, hubble-deepfield, trappist-system, andromeda-collision, europa-geysers, titan-methane-lakes, oort-cloud, pulsar-beacon
3. **Nature** (34): fire, waterfall, ocean-waves, rain, thunderstorm, tornado, snowfall, forest, blooming-flowers, sunrise, sunset, volcano, lightning, wind, geyser, aurora, cave, river, blizzard, autumn-leaves, desert-dunes, coral-reef, bamboo-forest, meteor-impact, coral-spawning, whale-breech, geyser-eruption, sequoia-tree, bioluminescence-sea, crystal-cavern, iceberg-calving, desert-oasis, wildfire-sweep, canyon-river
4. **Technology** (36): matrix-rain, hacking-terminal, computer-boot, digital-glitch, network-connections, binary-rain, ai-neural-network, circuit-board, cpu-processing, gpu-rendering, data-streams, cyber-grid, digital-face, quantum-computing, ai-brain, encryption, quantum-teleport, hologram, server-farm, drone-swarm, blockchain-mining, cyber-skull, radar-scanner, synthwave-sun, laser-grid, terminal-matrix, docker-whale, linux-tux-penguin, kubernetes-wheel, git-branch-merge, vim-editor, crypto-wallet, microchip-die, fiber-optic-pulse, mainframe-tape, satellite-uplink
5. **Abstract** (23): rotating-cube, 3d-sphere, spinning-donut, dna-helix, fractal, tunnel, particle-explosion, kaleidoscope, infinity-loop, waves, spiral, morphing-shapes, vortex, optical-illusion, meta-cube, fractal-2, spiral-2, hypercube-tesseract, lissajous-curve, torus-knot, sine-wave-field, sacred-geometry, quantum-foam
6. **Symbols** (19): animated-heart, broken-heart, smiley, crying-face, fire-heart, music-notes, equalizer, hourglass, loading-animation, countdown, digital-clock, pulse, blink, yin-yang, peace-sign, radioactive-hazard, biohazard, infinity-symbol, audio-bars
7. **HS AI Exclusive** (27): hs-ai-solutions-logo, hs-ai-studio, gravity-ai, galaxy-vision, hs-ai-neural-network, ai-model-training, ai-agent, hs-ai-core, hs-digital-universe, hs-ascii-live-intro, hs-core-v2, vision-alpha, hs-superintelligence, hs-agentic-flow, hs-neural-matrix, hs-deep-thought, hs-galaxy-core, hs-quantum-leap, hs-autonomous-agent, hs-omni-intelligence, hs-vision-core-v3, hs-neural-synapse, hs-sentient-stream, hs-apex-model, hs-matrix-titan, hs-infinite-scale, hs-hyper-drive
8. **Gaming** (25): space-invaders, pong, tetris, flappy-bird, dino-runner, snake-game, doom-fire, arcade-joystick, gameboy, zelda-triforce, minecraft-steve, portal-aperture, half-life-lambda, doom-slayer, street-fighter-hadouken, mortal-kombat-subzero, zelda-sword-master, pokeball-catch, donkey-kong-barrels, contra-spreadgun, galaga-fighter, asteroids-ship, dig-dug, frogger-crossing, mega-man
9. **Vehicles** (20): cybertruck, speeding-car, steam-locomotive, airplane-flight, submarine, motorcycle, helicopter, hovercraft, sailing-ship, space-shuttle, formula1-racer, bullet-train-shinkansen, stealth-bomber, apollo-saturn-v, harley-chopper, monster-truck, zeppelin-airship, pirate-galleon, u-boat-periscope, hover-bike
10. **Creatures** (20): roaring-lion, dragon-flight, phoenix-rising, flying-bat, swimming-whale, hopping-frog, galloping-horse, slithering-snake, flapping-butterfly, howling-wolf, majestic-eagle, stealth-panther, charging-rhino, leaping-dolphin, chameleon-blend, octopus-ink, owl-stare, peacock-feather, mantis-strike, polar-bear
11. **Music** (15): piano-keys, vinyl-record, drum-kit, electric-guitar, saxophone-jazz, synthesizer-wave, cassette-tape, boombox, audio-spectrum, turntable-scratch, soundwave-pulse, disco-ball, metronome, microphone-glow, headphones-beat
12. **Sports** (15): soccer-kick, basketball-dunk, tennis-match, baseball-swing, skater-kickflip, surfer-barrel, skiing-downhill, boxing-punch, swimming-freestyle, bowling-strike, archery-bullseye, cycling-sprint, golf-swing, gymnast-flip, weightlifter-snatch
13. **Cyberpunk** (15): neon-city-skyline, cyber-samurai, synthwave-highway, hologram-girl, glitch-skull, neon-sunset, cyber-monolith, wireframe-landscape, cyberspace-portal, megacity-spire, data-core-breach, quantum-flux, cyber-visor, neural-sync, orbital-elevator
14. **Food** (12): steaming-coffee, sizzling-pizza, burger-stack, ramen-noodles, popping-corn, ice-cream-cone, sushi-roll, boba-tea, taco-tuesday, donut-glaze, champagne-toast, birthday-cake
15. **Weather** (12): typhoon-cyclone, sandstorm-desert, monsoon-deluge, rainbow-arc, hailstorm, solar-flare, foggy-morning, lightning-strike, autumn-gale, frost-crystals, dust-devil, spring-breeze

## Verification & Testing
Run the comprehensive test suite verifying all 349 animations, render engines, and HTTP endpoints:
```bash
cd backend
go test -v ./...
```

## License
HS AI Solutions
