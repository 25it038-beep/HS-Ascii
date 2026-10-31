# HS ASCII LIVE
### By HS AI Solutions

A production-ready, publicly accessible ASCII animation streaming platform with **190+ continuously animated ASCII experiences** across 10 categories.

## Features
- **Real HTTP Streaming**: ANSI escape sequence streaming with clear-screen and cursor-homing.
- **190+ Procedural Animations**: High-performance continuous generators across 10 specialized categories.
- **Rich Rendering Engines**:
  - **Retro Gaming**: Pac-Man, Space Invaders, Pong, Tetris, Snake slithering and food eating.
  - **3D Geometry & Mathematics**: True 3D wireframe rotating cube, 3D rotating torus donut (donut.c homage), 3D double helix DNA, and rotating globe sphere.
  - **Tech & Sci-Fi**: Matrix digital rain, multi-band audio equalizer, rocket launch, neural network synapse pulses, radar sweep, and accretion disk blackhole.
  - **Nature & Symbols**: Pulsing anatomical hearts, cellular flame simulation, sinusoidal ocean waves, dancing stickman, flying dragons, and official HS AI logo.
- **Concurrent Viewers**: Non-blocking goroutine streaming with graceful context cancellation.
- **REST & Query APIs**: Dynamic `/list`, `/api/animations`, `/api/stats`, and `/random` endpoints.
- **Customizable Parameters**: Speed (`slow`, `normal`, `fast`), FPS (`1-60`), color (`red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`), width, and height.
- **Automated Test Suite**: Unit and integration test coverage for all animations, pattern generators, and HTTP endpoints.

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
# Classic character and gaming animations
curl https://hs-ascii-1.onrender.com/pacman
curl https://hs-ascii-1.onrender.com/space-invaders
curl https://hs-ascii-1.onrender.com/rickroll

# High-tech & sci-fi streams
curl https://hs-ascii-1.onrender.com/matrix-rain?color=green&fps=20
curl https://hs-ascii-1.onrender.com/hs-ai-solutions-logo?fps=15
curl https://hs-ascii-1.onrender.com/cybertruck?speed=fast

# 3D Math & Geometry
curl https://hs-ascii-1.onrender.com/rotating-cube
curl https://hs-ascii-1.onrender.com/spinning-donut
curl https://hs-ascii-1.onrender.com/dna-helix
```

## Categories (193 Total)
1. **Characters** (28): rickroll, dancing-rick, pikachu, cat, dog, among-us, pacman, mario, sonic, skull, ghost, dancing-stickman, robot, alien, spider, ninja, samurai, viking, pirate, wizard, batman, superman, ironman, spiderman, yoda, darth-vader, gandalf, cyber-ninja
2. **Space** (26): galaxy, milky-way, blackhole, solar-system, rotating-earth, moon-phases, saturn, shooting-stars, meteor-shower, supernova, wormhole, spaceship, rocket-launch, astronaut, ufo, comet, asteroid-belt, nebula, constellation, milky-way2, jupiter-storm, pulsar, quasars, rings-of-uranus, solar-eclipse, deep-space-probe
3. **Nature** (24): fire, waterfall, ocean-waves, rain, thunderstorm, tornado, snowfall, forest, blooming-flowers, sunrise, sunset, volcano, lightning, wind, geyser, aurora, cave, river, blizzard, autumn-leaves, desert-dunes, coral-reef, bamboo-forest, meteor-impact
4. **Technology** (26): matrix-rain, hacking-terminal, computer-boot, digital-glitch, network-connections, binary-rain, ai-neural-network, circuit-board, cpu-processing, gpu-rendering, data-streams, cyber-grid, digital-face, quantum-computing, ai-brain, encryption, quantum-teleport, hologram, server-farm, drone-swarm, blockchain-mining, cyber-skull, radar-scanner, synthwave-sun, laser-grid, terminal-matrix
5. **Abstract** (23): rotating-cube, 3d-sphere, spinning-donut, dna-helix, fractal, tunnel, particle-explosion, kaleidoscope, infinity-loop, waves, spiral, morphing-shapes, vortex, optical-illusion, meta-cube, fractal-2, spiral-2, hypercube-tesseract, lissajous-curve, torus-knot, sine-wave-field, sacred-geometry, quantum-foam
6. **Symbols** (19): animated-heart, broken-heart, smiley, crying-face, fire-heart, music-notes, equalizer, hourglass, loading-animation, countdown, digital-clock, pulse, blink, yin-yang, peace-sign, radioactive-hazard, biohazard, infinity-symbol, audio-bars
7. **HS AI Exclusive** (17): hs-ai-solutions-logo, hs-ai-studio, gravity-ai, galaxy-vision, hs-ai-neural-network, ai-model-training, ai-agent, hs-ai-core, hs-digital-universe, hs-ascii-live-intro, hs-core-v2, vision-alpha, hs-superintelligence, hs-agentic-flow, hs-neural-matrix, hs-deep-thought, hs-galaxy-core
8. **Gaming** (10): space-invaders, pong, tetris, flappy-bird, dino-runner, snake-game, doom-fire, arcade-joystick, gameboy, zelda-triforce
9. **Vehicles** (10): cybertruck, speeding-car, steam-locomotive, airplane-flight, submarine, motorcycle, helicopter, hovercraft, sailing-ship, space-shuttle
10. **Creatures** (10): roaring-lion, dragon-flight, phoenix-rising, flying-bat, swimming-whale, hopping-frog, galloping-horse, slithering-snake, flapping-butterfly, howling-wolf

## Verification & Testing
Run the comprehensive test suite verifying all 193 animations, render engines, and HTTP endpoints:
```bash
cd backend
go test -v ./...
```

## License
HS AI Solutions
