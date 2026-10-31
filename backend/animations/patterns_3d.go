package animations

import (
	"fmt"
	"math"
	"strings"
)

type point3D struct {
	x, y, z float64
}

func rotateX(p point3D, angle float64) point3D {
	rad := angle * math.Pi / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return point3D{
		x: p.x,
		y: p.y*cos - p.z*sin,
		z: p.y*sin + p.z*cos,
	}
}

func rotateY(p point3D, angle float64) point3D {
	rad := angle * math.Pi / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return point3D{
		x: p.x*cos + p.z*sin,
		y: p.y,
		z: -p.x*sin + p.z*cos,
	}
}

func rotateZ(p point3D, angle float64) point3D {
	rad := angle * math.Pi / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return point3D{
		x: p.x*cos - p.y*sin,
		y: p.x*sin + p.y*cos,
		z: p.z,
	}
}

func drawLine(grid [][]rune, x0, y0, x1, y1 int, ch rune) {
	h := len(grid)
	w := len(grid[0])
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy

	for {
		if x0 >= 0 && x0 < w && y0 >= 0 && y0 < h {
			grid[y0][x0] = ch
		}
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func drawCube(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	if strings.Contains(animName, "minecraft") {
		swing := mod(frameNum, 4)
		pickRune := '/'
		if swing >= 2 {
			pickRune = '\\'
		}
		steve := []string{
			"    .-----.    ",
			"    | o o |    ",
			"    |  -  |    ",
			fmt.Sprintf("  ==[###]==%c", pickRune),
			"    |  |  |    ",
			"   [#]   [#]   ",
		}
		startY := int(cy) - len(steve)/2
		for i, line := range steve {
			y := startY + i
			startX := int(cx) - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	// 8 vertices of a 3D cube
	vertices := []point3D{
		{-1, -1, -1},
		{1, -1, -1},
		{1, 1, -1},
		{-1, 1, -1},
		{-1, -1, 1},
		{1, -1, 1},
		{1, 1, 1},
		{-1, 1, 1},
	}

	// 12 edges
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
		{4, 5}, {5, 6}, {6, 7}, {7, 4},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
	}

	angle := float64(frameNum * 6)
	proj := make([]struct{ x, y int }, len(vertices))

	scaleX := float64(w) * 0.4
	scaleY := float64(h) * 0.4

	for i, v := range vertices {
		r := rotateX(v, angle)
		r = rotateY(r, angle*0.8)
		r = rotateZ(r, angle*0.4)

		dist := 3.0 + r.z
		if dist < 0.1 {
			dist = 0.1
		}
		px := int(cx + (r.x/dist)*scaleX)
		py := int(cy + (r.y/dist)*scaleY)
		proj[i] = struct{ x, y int }{px, py}
	}

	for _, e := range edges {
		p1 := proj[e[0]]
		p2 := proj[e[1]]
		drawLine(grid, p1.x, p1.y, p2.x, p2.y, '#')
	}
}

func drawDonut(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	A := float64(frameNum) * 0.08
	B := float64(frameNum) * 0.04

	cosA, sinA := math.Cos(A), math.Sin(A)
	cosB, sinB := math.Cos(B), math.Sin(B)

	zbuffer := make([]float64, w*h)
	for i := range zbuffer {
		zbuffer[i] = 0
	}

	chars := []rune(".,-~:;=!*#$@")

	for theta := 0.0; theta < 2*math.Pi; theta += 0.14 {
		costheta := math.Cos(theta)
		sintheta := math.Sin(theta)

		for phi := 0.0; phi < 2*math.Pi; phi += 0.05 {
			cosphi := math.Cos(phi)
			sinphi := math.Sin(phi)

			circlex := 2.0 + costheta
			circley := sintheta

			x := circlex*(cosB*cosphi+sinA*sinB*sinphi) - circley*cosA*sinB
			y := circlex*(sinB*cosphi-sinA*cosB*sinphi) + circley*cosA*cosB
			z := 3.5 + cosA*circlex*sinphi + circley*sinA
			ooz := 1.0 / z

			xp := int(cx + 30.0*ooz*x*0.6)
			yp := int(cy + 15.0*ooz*y*0.6)

			L := cosphi*costheta*sinB - cosA*costheta*sinphi - sinA*sintheta + cosB*(cosA*sintheta-costheta*sinA*sinphi)

			if yp >= 0 && yp < h && xp >= 0 && xp < w {
				idx := yp*w + xp
				if ooz > zbuffer[idx] {
					zbuffer[idx] = ooz
					luminanceIndex := int(L * 8)
					if luminanceIndex > 0 {
						if luminanceIndex >= len(chars) {
							luminanceIndex = len(chars) - 1
						}
						grid[yp][xp] = chars[luminanceIndex]
					} else {
						grid[yp][xp] = '.'
					}
				}
			}
		}
	}
}

func drawDNA(grid [][]rune, frameNum, seed int) {
	h := len(grid)
	w := len(grid[0])
	cx := w / 2

	amplitude := float64(w) * 0.25
	if amplitude < 4 {
		amplitude = 4
	}

	bases := []string{"A-T", "T-A", "C-G", "G-C"}

	for y := 0; y < h; y++ {
		phase := float64(frameNum)*0.2 + float64(y)*0.4
		s := math.Sin(phase)

		x1 := int(float64(cx) + s*amplitude)
		x2 := int(float64(cx) - s*amplitude)

		if x1 > x2 {
			x1, x2 = x2, x1
		}

		// Draw connecting rung at intervals
		if y%2 == 0 && x2-x1 > 2 {
			basePair := bases[(y/2+frameNum)%len(bases)]
			bpIdx := 0
			for x := x1 + 1; x < x2; x++ {
				if x >= 0 && x < w {
					if bpIdx < len(basePair) {
						grid[y][x] = rune(basePair[bpIdx])
						bpIdx++
					} else {
						grid[y][x] = '-'
					}
				}
			}
		}

		// Draw strands
		if x1 >= 0 && x1 < w {
			if s >= 0 {
				grid[y][x1] = '('
			} else {
				grid[y][x1] = ')'
			}
		}
		if x2 >= 0 && x2 < w {
			if s >= 0 {
				grid[y][x2] = ')'
			} else {
				grid[y][x2] = '('
			}
		}
	}
}

func drawEarth(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := float64(h) * 0.45
	if float64(w)*0.25 < r {
		r = float64(w) * 0.25
	}
	if r < 3 {
		r = 3
	}

	rot := float64(frameNum)*0.15 + float64(seed)*0.2

	isJupiter := strings.Contains(animName, "jupiter")
	isMoon := strings.Contains(animName, "moon")
	isMars := strings.Contains(animName, "mars")

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := (float64(x) - cx) * 0.55
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r {
				lon := math.Asin(dx/r) + rot
				lat := math.Asin(dy / r)

				if isJupiter {
					// Cloud bands & Great Red Spot storm
					bands := math.Sin(lat * 8.0)
					if math.Abs(lat+0.3) < 0.2 && math.Abs(math.Sin(lon)) > 0.7 {
						grid[y][x] = '@' // Great Red Spot storm
					} else if bands > 0.2 {
						grid[y][x] = '='
					} else {
						grid[y][x] = '~'
					}
				} else if isMoon {
					// Moon craters
					crater := math.Sin(lon*5) * math.Cos(lat*5)
					if crater > 0.4 {
						grid[y][x] = 'o'
					} else if dx > 0 {
						grid[y][x] = '.'
					} else {
						grid[y][x] = ' '
					}
				} else if isMars {
					// Mars terrain
					m := math.Sin(lon*4+lat*2) + math.Cos(lon*2)
					if m > 0.3 {
						grid[y][x] = '#'
					} else {
						grid[y][x] = '.'
					}
				} else {
					v := math.Sin(lon*3) * math.Cos(lat*2)
					if math.Abs(dist-r) < 0.6 {
						grid[y][x] = 'O'
					} else if v > 0.3 {
						grid[y][x] = '#'
					} else if v > -0.1 {
						grid[y][x] = '~'
					} else {
						grid[y][x] = '.'
					}
				}
			}
		}
	}
}

func drawSpiral(grid [][]rune, frameNum, seed int, animName string) {
	h := len(grid)
	w := len(grid[0])
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	if strings.Contains(animName, "megaman") {
		megaman := []string{
			"    .---.    ",
			"   / [o] \\   ",
			"  ==( [O] )=====> (BUSTER SHOT)",
			"   /  |  \\   ",
		}
		startY := int(cy) - len(megaman)/2
		for i, line := range megaman {
			y := startY + i
			startX := int(cx) - 8
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "portal") {
		for a := 0.0; a < 2*math.Pi; a += 0.2 {
			px := int(cx + 8.0*math.Cos(a)/0.55)
			py := int(cy + 4.0*math.Sin(a))
			setCell(grid, px, py, '@')
		}
		portal := []string{
			"  [APERTURE]  ",
			"  < PORTAL >  ",
		}
		startY := int(cy) - len(portal)/2
		for i, line := range portal {
			y := startY + i
			startX := int(cx) - len(line)/2
			for j, r := range line {
				setCell(grid, startX+j, y, r)
			}
		}
		return
	}

	if strings.Contains(animName, "mandelbrot") || strings.Contains(animName, "fractal") {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				cr := (float64(x)-cx)/(float64(w)*0.25) - 0.5
				ci := (float64(y)-cy)/(float64(h)*0.25)
				zr, zi := 0.0, 0.0
				iter := 0
				for iter < 12 && zr*zr+zi*zi < 4.0 {
					tmp := zr*zr - zi*zi + cr
					zi = 2.0*zr*zi + ci
					zr = tmp
					iter++
				}
				if iter == 12 {
					setCell(grid, x, y, '#')
				} else if iter > 6 {
					setCell(grid, x, y, '+')
				} else if iter > 2 {
					setCell(grid, x, y, '.')
				}
			}
		}
		return
	}

	maxPoints := 70
	chars := []rune("*+o°.~-")

	for i := 0; i < maxPoints; i++ {
		angle := float64(i)*0.25 + float64(frameNum)*0.15
		radius := float64(i) * 0.35
		x := int(cx + radius*math.Cos(angle)*1.6)
		y := int(cy + radius*math.Sin(angle)*0.8)

		if x >= 0 && x < w && y >= 0 && y < h {
			grid[y][x] = chars[i%len(chars)]
		}
	}
}
