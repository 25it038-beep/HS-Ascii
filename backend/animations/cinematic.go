package animations

import (
	"fmt"
	"hs-ascii-live/engine"
	"math"
	"math/rand"
	"strings"
	"sync"
)

const (
	cinePI = float32(math.Pi)
)

// Random tables for stars and embers
var (
	cineOnce sync.Once
	cineStX  []float32
	cineStY  []float32
	cineStP  []float32
	cineStB  []float32
	cineEx0  []float32
	cineEy0  []float32
	cineEsp  []float32
	cineEph  []float32
)

func initCineRand() {
	cineOnce.Do(func() {
		rnd := rand.New(rand.NewSource(11))
		n := 150
		cineStX = make([]float32, n)
		cineStY = make([]float32, n)
		cineStP = make([]float32, n)
		cineStB = make([]float32, n)
		for i := 0; i < n; i++ {
			cineStX[i] = rnd.Float32()
			cineStY[i] = rnd.Float32()
			cineStP[i] = rnd.Float32()
			cineStB[i] = 0.15 + 0.55*rnd.Float32()
		}

		m := 70
		cineEx0 = make([]float32, m)
		cineEy0 = make([]float32, m)
		cineEsp = make([]float32, m)
		cineEph = make([]float32, m)
		for i := 0; i < m; i++ {
			cineEx0[i] = rnd.Float32()
			cineEy0[i] = rnd.Float32()
			cineEsp[i] = 0.4 + rnd.Float32()
			cineEph[i] = rnd.Float32() * 6.28
		}
	})
}

// Math helper functions
func cClamp01(x float32) float32 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}

func cSm(a, b, x float32) float32 {
	t := cClamp01((x - a) / (b - a))
	return t * t * (3 - 2*t)
}

func cLerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

func cSin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func cCos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func cExp(x float32) float32 {
	return float32(math.Exp(float64(x)))
}

func cSqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func cPow(a, b float32) float32 {
	return float32(math.Pow(float64(a), float64(b)))
}

func cFrac(x float32) float32 {
	return x - float32(math.Floor(float64(x)))
}

func cOutBack(p float32) float32 {
	c1 := float32(1.4)
	c3 := c1 + 1.0
	q := p - 1.0
	return 1.0 + c3*q*q*q + c1*q*q
}

func cSp(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if r == ' ' {
			sb.WriteString("  ")
		} else {
			sb.WriteRune(r)
			if i < len(s)-1 {
				sb.WriteByte(' ')
			}
		}
	}
	return sb.String()
}

type cineOverlay struct {
	row, col int
	ch       rune
	r, g, b  float32
}

type cineContext struct {
	w, hc, h int
	fr       []float32
	fg       []float32
	fb       []float32
	ovs      []cineOverlay
	fade     float32
}

func (c *cineContext) bl(x, y int, r, g, b, a float32) {
	if x < 0 || y < 0 || x >= c.w || y >= c.h || a <= 0 {
		return
	}
	if a > 1 {
		a = 1
	}
	i := y*c.w + x
	c.fr[i] += (r - c.fr[i]) * a
	c.fg[i] += (g - c.fg[i]) * a
	c.fb[i] += (b - c.fb[i]) * a
}

func (c *cineContext) ad(x, y int, r, g, b float32) {
	if x < 0 || y < 0 || x >= c.w || y >= c.h {
		return
	}
	i := y*c.w + x
	c.fr[i] += r
	c.fg[i] += g
	c.fb[i] += b
}

func (c *cineContext) backdrop(gr, gg, gb, amt, cyf float32) {
	hy := float32(c.h) * cyf
	ix := 1.0 / (float32(c.w) * 0.5)
	iy := 1.0 / (float32(c.h) * 0.5)
	for y := 0; y < c.h; y++ {
		dy := (float32(y) - hy) * iy
		for x := 0; x < c.w; x++ {
			dx := (float32(x) - float32(c.w)*0.5) * ix
			d2 := dx*dx + dy*dy
			g := amt * cExp(-d2*2.4)
			i := y*c.w + x
			c.fr[i] = 0.006 + gr*g
			c.fg[i] = 0.007 + gg*g
			c.fb[i] = 0.011 + gb*g
		}
	}
}

func (c *cineContext) stars(t, alpha float32) {
	for k := 0; k < len(cineStX); k++ {
		x := int(cineStX[k] * float32(c.w))
		y := int(cineStY[k] * float32(c.h))
		tw := 0.55 + 0.45*cSin(t*(0.8+cineStP[k]*2.0)+cineStP[k]*20.0)
		v := cineStB[k] * tw * alpha
		c.ad(x, y, v*0.8, v*0.9, v)
	}
}

func (c *cineContext) embers(t, alpha, cr, cg, cb float32) {
	for k := 0; k < len(cineEx0); k++ {
		y := float32(c.h) - cFrac(cineEy0[k]+t*0.05*cineEsp[k])*float32(c.h)
		x := cineEx0[k]*float32(c.w) + cSin(t*cineEsp[k]+cineEph[k])*3.0
		fl := 0.45 + 0.55*cSin(t*5.0*cineEsp[k]+cineEph[k])
		edge := cSm(0, 0.15, y/float32(c.h)) * cSm(1, 0.85, y/float32(c.h))
		v := alpha * fl * edge * (0.5 + 0.5*cineEsp[k]*0.7)
		xi := int(x)
		yi := int(y)
		c.ad(xi, yi, cr*v, cg*v, cb*v)
		c.ad(xi+1, yi, cr*v*0.25, cg*v*0.25, cb*v*0.25)
		c.ad(xi-1, yi, cr*v*0.25, cg*v*0.25, cb*v*0.25)
		c.ad(xi, yi+1, cr*v*0.25, cg*v*0.25, cb*v*0.25)
		c.ad(xi, yi-1, cr*v*0.25, cg*v*0.25, cb*v*0.25)
	}
}

func (c *cineContext) put(row, col int, ch rune, r, g, b, a float32) {
	if ch == ' ' || a <= 0.01 {
		return
	}
	c.ovs = append(c.ovs, cineOverlay{
		row: row,
		col: col,
		ch:  ch,
		r:   r * a,
		g:   g * a,
		b:   b * a,
	})
}

func (c *cineContext) textC(row int, s string, r, g, b, a float32) {
	runes := []rune(s)
	c0 := (c.w - len(runes)) / 2
	for i, ch := range runes {
		c.put(row, c0+i, ch, r, g, b, a)
	}
}

// Scene 1: Intro (duration 7.0s)
func (c *cineContext) drawIntro(t float32) {
	c.backdrop(0.03, 0.05, 0.08, 0.6*cSm(0, 2, t), 0.45)
	c.stars(t, cSm(0, 1.5, t))

	tr := c.hc/2 - 2
	if tr < 1 {
		tr = 1
	}
	py := tr*2 + 1
	length := float32(c.w) * 0.5 * cSm(0.4, 2.4, t)
	inten := 0.9 * cSm(0.4, 1.4, t) * (1.0 - 0.8*cSm(2.4, 3.6, t))

	for x := 0; x < c.w; x++ {
		dx := float32(math.Abs(float64(float32(x) - float32(c.w)/2.0)))
		if dx > length {
			continue
		}
		fall := cPow(1.0-dx/length, 1.6)
		for dy := -4; dy <= 4; dy++ {
			v := inten * fall * cExp(-float32(dy*dy)*0.5)
			c.ad(x, py+dy, v*0.55, v*0.78, v)
		}
	}

	title := cSp("HS AI SOLUTIONS")
	titleRunes := []rune(title)
	c0 := (c.w - len(titleRunes)) / 2
	rev := (t - 1.5) / 0.09
	sw := (t-3.0)/1.4*float32(len(titleRunes))*1.3 - 6.0

	for i, ch := range titleRunes {
		la := cClamp01(rev - float32(i))
		sh := cExp(-((float32(i) - sw) * (float32(i) - sw)) / 18.0)
		c.put(tr, c0+i, ch, cLerp(0.70, 1.0, sh), cLerp(0.74, 1.0, sh), cLerp(0.80, 1.0, sh), la)
	}

	c.textC(tr+2, cSp("PRESENTS"), 0.50, 0.55, 0.62, cSm(3.4, 4.4, t))
	c.textC(tr+4, cSp("A CINEMATIC EXPERIENCE"), 0.25, 0.80, 0.95, cSm(4.5, 5.5, t))
	c.fade = 1.0 - cSm(6.3, 7.0, t)
}

// Scene 2: The Eye (duration 8.4s)
func (c *cineContext) drawEye(t float32) {
	c.backdrop(0.0, 0.22, 0.30, 0.55*cSm(0.2, 1.6, t)*(1.0-cSm(5.0, 6.5, t)), 0.5)
	c.stars(t, 0.5*(1.0-cSm(4.5, 5.5, t)))

	cx := float32(c.w) * 0.5
	cy := float32(c.h) * 0.5
	a := float32(math.Min(float64(float32(c.w)*0.40), float64(float32(c.h)*1.15)))
	b := a * 0.36
	open := cSm(0.25, 1.7, t)
	if t > 3.6 && t < 3.98 {
		open *= 1.0 - 0.96*cSin((t-3.6)/0.38*cinePI)
	}

	look := cSm(1.6, 2.6, t) * (1.0 - cSm(4.2, 5.0, t))
	lx := a * 0.16 * cSin(t*1.25) * look
	ly := b * 0.10 * cSin(t*0.9+1.0) * look
	ri := b * 0.92
	rp := ri * (0.36 + 0.05*cSin(t*2.1) + 0.30*cSm(5.0, 7.0, t))

	z := float32(1.0)
	if t > 5.0 {
		p := float32(math.Min(1.0, float64((t-5.0)/3.2)))
		z = cExp(4.4 * cPow(p, 1.5))
	}

	pcx := cx + lx
	pcy := cy + ly
	glintK := 1.0 - cSm(5.0, 5.7, t)
	aa := 1.0 / z
	glowK := cClamp01(open*1.5) * (1.0 - cSm(5.0, 6.0, t))

	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			ex := pcx + (float32(x)-cx)/z
			ey := pcy + (float32(y)-cy)/z
			u := (ex - cx) / a
			dy := ey - cy
			au := float32(math.Abs(float64(u)))

			lidK := float32(0.0)
			if au < 1.0 {
				lidK = cPow(1.0-u*u, 0.85)
			}
			topFactor := float32(0.72)
			if dy < 0 {
				topFactor = 1.0
			}
			lh := b * open * lidK * topFactor
			edge := lh - float32(math.Abs(float64(dy)))
			i := y*c.w + x

			dist := -edge
			if au >= 1.0 {
				d1 := (au - 1.0) * a
				dist = cSqrt(d1*d1 + dy*dy)
			}

			if edge <= 0 && glowK > 0 {
				gl := 0.42 * glowK * cExp(-dist*z*0.20)
				c.fr[i] += 0.05 * gl
				c.fg[i] += 0.62 * gl
				c.fb[i] += 0.85 * gl
			}

			if edge > 0 {
				ia0 := cClamp01(edge / aa)
				denom := lh
				if denom < 0.001 {
					denom = 0.001
				}
				depth := float32(math.Abs(float64(dy))) / denom
				shade := 1.0 - 0.62*cPow(depth, 2.2) - 0.40*cPow(au, 3.0)
				if dy < 0 {
					shade -= 0.18 * cSm(0.3, 1.0, depth)
				}
				if shade < 0.05 {
					shade = 0.05
				}
				sr := 0.80 * shade
				sg := 0.85 * shade
				sb2 := 0.93 * shade

				ddx := ex - pcx
				ddy := ey - pcy
				d := cSqrt(ddx*ddx + ddy*ddy)
				ia := cClamp01((ri - d) / aa)
				if ia > 0 {
					nd := d / ri
					ang := float32(math.Atan2(float64(ddy), float64(ddx)))
					fib := 0.5 + 0.5*cSin(ang*26.0+2.2*cSin(ang*5.0+nd*3.0))
					fib2 := 0.5 + 0.5*cSin(ang*71.0+nd*9.0)
					f := 0.55*fib + 0.45*fib2
					q := (nd - 0.52) * 7.0
					ring := cExp(-q * q)
					rr := cLerp(0.02, 0.22, f*(1.0-0.5*nd)) + 0.25*ring
					rg := cLerp(0.16, 0.78, f*(1.0-0.35*nd)) + 0.35*ring
					rb := cLerp(0.30, 0.95, f*(1.0-0.2*nd)) + 0.30*ring
					limb := 1.0 - 0.78*cSm(0.78, 1.0, nd)
					ish := cClamp01(shade*1.15 + 0.2)
					rr *= limb * ish
					rg *= limb * ish
					rb *= limb * ish

					qq := (d - rp) / (ri * 0.12)
					pg := 0.55 * cExp(-qq*qq)
					rg += pg * 0.7
					rb += pg
					rr += pg * 0.1

					pa := cClamp01((rp - d) / aa)
					rr = cLerp(rr, 0, pa)
					rg = cLerp(rg, 0, pa)
					rb = cLerp(rb, 0, pa)

					gx := ex - (pcx - ri*0.36)
					gy := ey - (pcy - ri*0.40)
					g1 := cClamp01((ri*0.15-cSqrt(gx*gx+gy*gy))/aa) * glintK

					hx := ex - (pcx + ri*0.30)
					hy := ey - (pcy + ri*0.30)
					g2 := cClamp01((ri*0.07-cSqrt(hx*hx+hy*hy))/aa) * 0.6 * glintK
					gm := float32(math.Max(float64(g1), float64(g2)))

					rr = cLerp(rr, 1.0, gm)
					rg = cLerp(rg, 1.0, gm)
					rb = cLerp(rb, 1.0, gm)

					sr = cLerp(sr, rr, ia)
					sg = cLerp(sg, rg, ia)
					sb2 = cLerp(sb2, rb, ia)
				}
				c.fr[i] = cLerp(c.fr[i], sr, ia0)
				c.fg[i] = cLerp(c.fg[i], sg, ia0)
				c.fb[i] = cLerp(c.fb[i], sb2, ia0)
			}

			if au < 1.0 {
				rq := edge * z
				rim := cExp(-rq*rq/1.4) * cClamp01((1.0-au)*6.0) * cClamp01(open*3.0) * (1.0 - cSm(5.0, 5.8, t))
				c.fr[i] += 0.35 * rim
				c.fg[i] += 0.85 * rim
				c.fb[i] += 1.00 * rim
			}
		}
	}

	ca := cSm(2.0, 3.0, t) * (1.0 - cSm(4.6, 5.2, t))
	c.textC(c.hc-3, cSp("LOOK CLOSER"), 0.35, 0.85, 1.0, ca)
}

// Scene 3: Red Rose & Stem (duration 12.6s)
func cStemX(y, rcx, yTop float32) float32 {
	return rcx + 3.2*cSin((y-yTop)*0.07)
}

func (c *cineContext) leaf(cx, cy, ang, length, wid, grow float32) {
	if grow <= 0.01 {
		return
	}
	length *= grow
	wid *= grow
	ca := cCos(ang)
	sa := cSin(ang)
	ext := int(length) + 2

	for y := int(cy) - ext; y <= int(cy)+ext; y++ {
		for x := int(cx) - ext; x <= int(cx)+ext; x++ {
			dx := float32(x) - cx
			dy := float32(y) - cy
			uu := (dx*ca + dy*sa) / length
			vv := (-dx*sa + dy*ca) / wid
			d := uu*uu + vv*vv
			if d >= 1.0 {
				continue
			}
			al := cClamp01((1.0 - d) * 5.0)
			tt := (uu + 1.0) * 0.5
			rr := cLerp(0.04, 0.10, tt)
			gg := cLerp(0.28, 0.58, tt)
			bb := cLerp(0.12, 0.22, tt)
			if float32(math.Abs(float64(vv))) < 0.07 {
				rr += 0.08
				gg += 0.15
				bb += 0.06
			}
			if vv < 0 {
				rr *= 1.2
				gg *= 1.2
				bb *= 1.2
			}
			c.bl(x, y, rr, gg, bb, al)
		}
	}
}

func (c *cineContext) roseHead(rcx, rcy, Rc, rot, alpha float32) {
	const N = 3.0
	const K = 1.8
	ext := int(Rc*1.6) + 2

	for y := int(rcy) - ext; y <= int(rcy)+ext; y++ {
		for x := int(rcx) - ext; x <= int(rcx)+ext; x++ {
			if x < 0 || y < 0 || x >= c.w || y >= c.h {
				continue
			}
			dx := float32(x) - rcx
			dy := float32(y) - rcy
			r := cSqrt(dx*dx + dy*dy)
			th := float32(math.Atan2(float64(dy), float64(dx)))
			rn := r / Rc
			s1 := th*N/(2.0*cinePI) + K*float32(math.Log(1.04)) + rot
			fe := cFrac(s1)
			edgeMul := 0.88 + 0.13*fe
			ia := cClamp01((edgeMul-rn)*Rc) * alpha
			if ia <= 0 {
				continue
			}

			s := th*N/(2.0*cinePI) + K*float32(math.Log(float64(rn+0.05))) + rot
			f := cFrac(s)
			lit := cPow(f, 1.4)
			shade := 0.30 + 0.70*lit
			shade += 0.35 * cSm(0.88, 0.99, f)
			shade *= 0.55 + 0.45*cSm(0.0, 0.12, f)
			shade *= 0.30 + 0.70*cSm(0.04, 0.5, rn)
			shade *= 1.0 - 0.28*cSm(0.78, 1.0, rn)
			hl := cSm(0.9, 0.99, f) * 0.18
			rr := 0.90*shade + hl + 0.04
			gg := 0.035*shade + hl*0.25
			bb := 0.08*shade + 0.015 + hl*0.2
			cw := 1.0 - cSm(0.03, 0.16, rn)
			rr = cLerp(rr, 0.22, cw)
			gg = cLerp(gg, 0.0, cw)
			bb = cLerp(bb, 0.03, cw)
			c.bl(x, y, rr, gg, bb, ia)
		}
	}
}

func (c *cineContext) drawRose(t float32) {
	gl := cSm(0.3, 3.0, t)*0.6 + 0.4*cSm(4.0, 6.0, t)
	c.backdrop(0.38, 0.02, 0.05, gl*0.45, 0.42)
	c.embers(t, cSm(0.2, 1.4, t)*0.9, 1.0, 0.25, 0.10)

	rcx := float32(c.w) * 0.5
	rcy := float32(c.h) * 0.34
	R := float32(math.Min(float64(float32(c.w)*0.19), float64(float32(c.h)*0.28)))
	yTop := rcy + R*0.5
	yBot := float32(c.h) * 0.97
	g := cSm(0, 1, cClamp01((t-1.3)/3.0))
	yEnd := cLerp(yBot, yTop, g)
	expl := cSm(10.4, 12.3, t)
	vis := 1.0 - cSm(10.6, 11.8, t)

	if g > 0 && vis > 0 {
		for y := int(yEnd); y < int(yBot); y++ {
			xc := cStemX(float32(y), rcx, yTop)
			for x := int(xc) - 2; x <= int(xc)+2; x++ {
				cov := cClamp01(1.5 - float32(math.Abs(float64(float32(x)-xc))))
				lightSide := float32(0.8)
				if float32(x) < xc {
					lightSide = 1.25
				}
				c.bl(x, y, 0.05*lightSide, 0.40*lightSide, 0.16*lightSide, cov*vis)
			}
		}

		thf := []float32{0.30, 0.52, 0.74}
		for k := 0; k < 3; k++ {
			ty := yTop + (yBot-yTop)*thf[k]
			if ty < yEnd {
				continue
			}
			xc := cStemX(ty, rcx, yTop)
			dir := float32(1.0)
			if k%2 != 0 {
				dir = -1.0
			}
			for j := 0; j < 3; j++ {
				c.bl(int(xc+dir*float32(2+j)), int(ty-float32(j)), 0.12, 0.45, 0.18, vis)
			}
		}

		y1 := yTop + (yBot-yTop)*0.42
		y2 := yTop + (yBot-yTop)*0.64
		l1 := cSm(0, 1, cClamp01((y1-yEnd)/(R*1.4))) * vis
		l2 := cSm(0, 1, cClamp01((y2-yEnd)/(R*1.4))) * vis
		sx1 := cStemX(y1, rcx, yTop)
		sx2 := cStemX(y2, rcx, yTop)
		ll := R * 0.95
		lw := R * 0.30
		c.leaf(sx1+cCos(-0.55)*ll*l1*0.95, y1+cSin(-0.55)*ll*l1*0.95, -0.55, ll, lw, l1)
		c.leaf(sx2+cCos(cinePI+0.55)*ll*l2*0.95, y2+cSin(cinePI+0.55)*ll*l2*0.95, cinePI+0.55, ll, lw, l2)
	}

	p := cClamp01((t - 4.0) / 3.2)
	if p > 0 && vis > 0 {
		grow := cOutBack(p)
		rc := R * grow * (1.0 + 0.22*expl)
		halo := 0.5 * cSm(0, 1, p) * vis
		for y := int(rcy - rc*2.2); y <= int(rcy+rc*2.2); y++ {
			for x := int(rcx - rc*2.2); x <= int(rcx+rc*2.2); x++ {
				dx := float32(x) - rcx
				dy := float32(y) - rcy
				q := cSqrt(dx*dx+dy*dy) / float32(math.Max(float64(rc), 1.0))
				if q < 2.2 {
					hv := halo * cExp(-q*q*0.9) * 0.35
					c.ad(x, y, hv, hv*0.03, hv*0.06)
				}
			}
		}
		c.roseHead(rcx, rcy, rc, t*0.16, vis)
	}

	// Falling petals
	if t > 7.4 {
		for k := 0; k < 16; k++ {
			sp := 2.0 + float32(k%5)*0.9
			yy := cFrac(float32(k)*0.137+(t-7.4)*sp/float32(c.h)*1.0) * float32(c.h)
			xx := float32((k*37)%100)/100.0*float32(c.w) + cSin(t*0.9+float32(k))*6.0
			al := cSm(7.4, 8.4, t) * vis
			for j := 0; j < 3; j++ {
				c.bl(int(xx)+j, int(yy), 0.85, 0.05, 0.10, 0.9*al)
				c.bl(int(xx)+j, int(yy)+1, 0.55, 0.02, 0.06, 0.9*al)
			}
		}
	}

	// Explosion of sparks
	if expl > 0 {
		for k := 0; k < 140; k++ {
			ang := float32(k) * 2.39996
			spd := (0.25 + 0.75*cFrac(float32(k)*0.618)) * R * 3.4
			rad := spd * float32(math.Pow(float64(expl), 0.6))
			x := rcx + cCos(ang)*rad
			y := rcy + cSin(ang)*rad*0.8 + expl*expl*6.0
			v := (1.0 - expl) * (0.6 + 0.4*cSin(float32(k)+t*20.0))
			c.ad(int(x), int(y), 1.0*v, 0.22*v, 0.10*v)
			c.ad(int(x)+1, int(y), 0.4*v, 0.05*v, 0.02*v)
		}
	}

	c1 := cSm(4.6, 5.6, t) * (1.0 - cSm(7.2, 7.8, t))
	c.textC(c.hc-4, cSp("A RED ROSE"), 1.0, 0.35, 0.40, c1)
	c2 := cSm(7.6, 8.6, t) * (1.0 - cSm(11.0, 11.8, t))
	c.textC(c.hc-4, cSp("BEAUTY IN DARKNESS"), 0.95, 0.40, 0.45, c2)
	c.fade = 1.0 - cSm(12.1, 12.6, t)
}

// Scene 4: Beating Heart & ECG Trace (loops)
func cWave(c float32) float32 {
	if c < 0.10 {
		return 0
	}
	if c < 0.20 {
		return 0.12 * cSin((c-0.10)/0.10*cinePI)
	}
	if c < 0.30 {
		return 0
	}
	if c < 0.34 {
		return -0.15 * cSin((c-0.30)/0.04*cinePI)
	}
	if c < 0.40 {
		k := (c - 0.34) / 0.06
		if k < 0.5 {
			return k * 2.0
		}
		return (1.0 - k) * 2.0
	}
	if c < 0.44 {
		return -0.35 * cSin((c-0.40)/0.04*cinePI)
	}
	if c < 0.52 {
		return 0
	}
	if c < 0.72 {
		return 0.22 * cSin((c-0.52)/0.20*cinePI)
	}
	return 0
}

func (c *cineContext) drawHeart(t float32) {
	const T = float32(1.1)
	ph := float32(math.Mod(float64(t), float64(T)))
	pulse := cExp(-cPow((ph-0.06)/0.07, 2.0)) + 0.7*cExp(-cPow((ph-0.30)/0.09, 2.0))
	app := float32(1.0)
	if t < 1.2 {
		app = cOutBack(cClamp01(t / 1.1))
	}
	vis := cSm(0, 0.7, t)
	c.backdrop(0.40, 0.02, 0.06, (0.30+0.25*pulse)*vis, 0.40)
	c.embers(t, 0.8*vis, 1.0, 0.2, 0.12)

	cx := float32(c.w) * 0.5
	cy := float32(c.h) * 0.40
	s0 := float32(math.Min(float64(float32(c.w)*0.17), float64(float32(c.h)*0.26)))
	s := s0 * (1.0 + 0.10*pulse) * app
	if s < 1.0 {
		s = 1.0
	}
	gs := 0.45 + 0.8*pulse

	ext := int(s * 2.6)
	for y := int(cy) - ext; y <= int(cy)+ext; y++ {
		for x := int(cx) - ext; x <= int(cx)+ext; x++ {
			if x < 0 || y < 0 || x >= c.w || y >= c.h {
				continue
			}
			X := (float32(x) - cx) / s
			Y := -(float32(y)-cy)/s + 0.15
			a := X*X + Y*Y - 1.0
			f := a*a*a - X*X*Y*Y*Y
			if f > 0 {
				gl := cExp(-f*3.2) * gs * 0.55 * vis
				if gl > 0.004 {
					c.ad(x, y, gl, gl*0.08, gl*0.14)
				}
			} else {
				d := cClamp01(-f)
				aaE := cClamp01(cPow(-f, 0.34) * s * 0.55)
				ink := 0.60 + 0.40*cPow(d, 0.45)
				ink += 0.18 * (-X*0.5 + Y*0.5)
				hx := X + 0.52
				hy := Y - 0.58
				hl := cExp(-(hx*hx*14.0 + hy*hy*30.0)) * 0.55
				rim := cExp(-d*9.0) * 0.25
				rr := 0.95*ink + hl + rim + 0.05*pulse
				gg := 0.05*ink + hl*0.55 + rim*0.3
				bb := 0.10*ink + hl*0.55 + rim*0.25
				c.bl(x, y, rr*vis, gg*vis, bb*vis, aaE)
			}
		}
	}

	// Expanding heartbeat ring
	if ph > 0.06 {
		q := (ph - 0.06) / 0.95
		rad := s0 * (1.05 + 1.7*q)
		al := (1.0 - q) * 0.5 * vis
		steps := int(rad * 9.0)
		for k := 0; k < steps; k++ {
			an := float32(k) * 2.0 * cinePI / float32(steps)
			c.ad(int(cx+cCos(an)*rad*1.15), int(cy+0.12*s0+cSin(an)*rad*0.95), al*0.8, al*0.05, al*0.12)
		}
	}

	// Dynamic ECG trace
	yb := float32(c.h) * 0.80
	amp := float32(c.h) * 0.07
	prevY := yb
	for x := 0; x < c.w; x++ {
		sc := float32(x)/float32(c.w)*3.0 + (t-0.06)/T + 0.37
		yy := yb - cWave(cFrac(sc))*amp
		v := (0.12 + 0.88*cPow(float32(x)/float32(c.w), 1.6)) * vis
		y0 := int(math.Min(float64(prevY), float64(yy)))
		y1 := int(math.Max(float64(prevY), float64(yy)))
		for y := y0; y <= y1; y++ {
			c.ad(x, y, v*1.0, v*0.14, v*0.22)
		}
		c.ad(x, y0-1, v*0.15, v*0.02, v*0.04)
		c.ad(x, y1+1, v*0.15, v*0.02, v*0.04)
		prevY = yy
	}

	tp := 0.75 + 0.25*cSin(t*2.2)
	c.textC(c.hc-4, cSp("FOREVER IN MOTION"), 1.0, 0.40, 0.45, cSm(1.2, 2.4, t)*tp)
	c.textC(c.hc-2, "HS  AI  SOLUTIONS", 0.55, 0.58, 0.64, cSm(2.0, 3.2, t)*0.8)
	c.fade = 1.0 - cSm(30.0, 32.0, t)
}

func cQuant(v, fade float32) int {
	i := int(v*fade*255.0 + 0.5)
	if i < 0 {
		i = 0
	}
	if i > 255 {
		i = 255
	}
	return i & 0xFC
}

// CinematicAnimation is the official HS AI Solutions Cinematic Experience
type CinematicAnimation struct {
	name string
}

func NewCinematicAnimation(name string) *CinematicAnimation {
	initCineRand()
	return &CinematicAnimation{name: name}
}

func (ca *CinematicAnimation) Name() string {
	return ca.name
}

func (ca *CinematicAnimation) Category() string {
	return "hsai"
}

func (ca *CinematicAnimation) Description() string {
	return "HS AI Solutions Cinematic Experience (Intro, The Eye, Red Rose, Beating Heart & ECG)"
}

func (ca *CinematicAnimation) NextFrame(opts engine.AnimationOptions, frameNum int) engine.Frame {
	w := opts.Width
	if w <= 0 {
		w = 80
	}
	h := opts.Height
	if h <= 0 {
		h = 24
	}
	if w < 10 {
		w = 10
	}
	if h < 4 {
		h = 4
	}

	fps := opts.FPS
	if fps <= 0 {
		fps = 10
	}

	// 1 row reserved for header
	hc := h - 1
	if hc < 2 {
		hc = 2
	}
	pixelH := hc * 2

	totalTime := float32(frameNum) / float32(fps)
	// 60-second seamless full loop
	cycleT := float32(math.Mod(float64(totalTime), 60.0))

	ctx := &cineContext{
		w:    w,
		hc:   hc,
		h:    pixelH,
		fr:   make([]float32, w*pixelH),
		fg:   make([]float32, w*pixelH),
		fb:   make([]float32, w*pixelH),
		ovs:  make([]cineOverlay, 0, 64),
		fade: 1.0,
	}

	if cycleT < 7.0 {
		ctx.drawIntro(cycleT)
	} else if cycleT < 15.4 {
		ctx.drawEye(cycleT - 7.0)
	} else if cycleT < 28.0 {
		ctx.drawRose(cycleT - 15.4)
	} else {
		ctx.drawHeart(cycleT - 28.0)
	}

	// Map overlays into a lookup grid for instant O(1) access
	type ovCell struct {
		ch      rune
		r, g, b float32
	}
	ovGrid := make(map[int]ovCell, len(ctx.ovs))
	for _, o := range ctx.ovs {
		if o.row >= 0 && o.row < hc && o.col >= 0 && o.col < w {
			ovGrid[o.row*w+o.col] = ovCell{ch: o.ch, r: o.r, g: o.g, b: o.b}
		}
	}

	// Render header row containing required identifier
	title := "── HS ASCII LIVE: HS AI SOLUTIONS CINEMATIC EXPERIENCE ──"
	if len(title) > w {
		title = "── HS ASCII LIVE: CINEMATIC ──"
	}
	if len(title) > w && w >= 13 {
		title = "HS ASCII LIVE"
	} else if len(title) > w {
		title = title[:w]
	}

	headerRunes := make([]rune, w)
	for i := range headerRunes {
		headerRunes[i] = ' '
	}
	startCol := (w - len([]rune(title))) / 2
	if startCol < 0 {
		startCol = 0
	}
	for i, r := range []rune(title) {
		if startCol+i < w {
			headerRunes[startCol+i] = r
		}
	}

	var sb strings.Builder
	sb.Grow(w * hc * 16)

	// Cyan header style
	sb.WriteString("\x1b[36m")
	sb.WriteString(string(headerRunes))
	sb.WriteString("\x1b[0m")

	curFg := -1
	curBg := -1

	for cy := 0; cy < hc; cy++ {
		sb.WriteString("\r\n")
		for cx := 0; cx < w; cx++ {
			// Check if overlay text exists here
			if ov, ok := ovGrid[cy*w+cx]; ok {
				fgCol := (cQuant(ov.r, ctx.fade) << 16) | (cQuant(ov.g, ctx.fade) << 8) | cQuant(ov.b, ctx.fade)
				// Background is black or slight dim
				if curBg != 0 {
					sb.WriteString("\x1b[40m")
					curBg = 0
				}
				if curFg != fgCol {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", (fgCol>>16)&255, (fgCol>>8)&255, fgCol&255))
					curFg = fgCol
				}
				sb.WriteRune(ov.ch)
				continue
			}

			// Pixel row
			i1 := (2*cy)*w + cx
			i2 := i1 + w
			tCol := (cQuant(ctx.fr[i1], ctx.fade) << 16) | (cQuant(ctx.fg[i1], ctx.fade) << 8) | cQuant(ctx.fb[i1], ctx.fade)
			bCol := (cQuant(ctx.fr[i2], ctx.fade) << 16) | (cQuant(ctx.fg[i2], ctx.fade) << 8) | cQuant(ctx.fb[i2], ctx.fade)

			if tCol == bCol {
				if tCol == 0 {
					if curBg != 0 {
						sb.WriteString("\x1b[40m")
						curBg = 0
					}
					sb.WriteByte(' ')
				} else {
					if curBg != bCol {
						sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", (bCol>>16)&255, (bCol>>8)&255, bCol&255))
						curBg = bCol
					}
					sb.WriteByte(' ')
				}
			} else {
				if curFg != tCol {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", (tCol>>16)&255, (tCol>>8)&255, tCol&255))
					curFg = tCol
				}
				if curBg != bCol {
					if bCol == 0 {
						sb.WriteString("\x1b[40m")
						curBg = 0
					} else {
						sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", (bCol>>16)&255, (bCol>>8)&255, bCol&255))
						curBg = bCol
					}
				}
				sb.WriteRune('▀')
			}
		}
	}
	sb.WriteString("\x1b[0m")

	return engine.Frame{
		Content: sb.String(),
	}
}
