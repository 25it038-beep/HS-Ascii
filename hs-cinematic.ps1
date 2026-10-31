# ============================================================
#  HS AI SOLUTIONS  |  CINEMATIC EXPERIENCE  (premium edition)
#  Run in Windows Terminal (or any VT / true-colour terminal):
#      powershell -ExecutionPolicy Bypass -File .\hs-cinematic.ps1
#  Press any key to exit.
# ============================================================
$Host.UI.RawUI.WindowTitle = "HS AI SOLUTIONS | CINEMATIC EXPERIENCE"
try { [Console]::OutputEncoding = [System.Text.Encoding]::UTF8 } catch {}

$src = @'
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;

public static partial class HSCinema
{
    [DllImport("kernel32.dll")] static extern IntPtr GetStdHandle(int n);
    [DllImport("kernel32.dll")] static extern bool GetConsoleMode(IntPtr h, out uint m);
    [DllImport("kernel32.dll")] static extern bool SetConsoleMode(IntPtr h, uint m);

    const string ESC = "\u001b";
    const float PI = 3.14159265f;
    static int Cols, Rows, W, Hc, H, OX, OY;
    static float[] fr, fg, fb;
    static int[] pt, pb;
    static float Fade = 1f;
    static StringBuilder sb = new StringBuilder(1 << 18);
    static Stream so;
    static UTF8Encoding utf8 = new UTF8Encoding(false);
    struct Ov { public int row, col; public char ch; public float r, g, b; }
    static List<Ov> ovs = new List<Ov>();
    delegate void Scene(float t);

    // ---------- particles / stars ----------
    static float[] stx, sty, stp, stb;
    static float[] ex0, ey0, esp, eph;
    static void InitRand()
    {
        Random rnd = new Random(11);
        int n = 150;
        stx = new float[n]; sty = new float[n]; stp = new float[n]; stb = new float[n];
        for (int i = 0; i < n; i++)
        {
            stx[i] = (float)rnd.NextDouble(); sty[i] = (float)rnd.NextDouble();
            stp[i] = (float)rnd.NextDouble(); stb[i] = 0.15f + 0.55f * (float)rnd.NextDouble();
        }
        int m = 70;
        ex0 = new float[m]; ey0 = new float[m]; esp = new float[m]; eph = new float[m];
        for (int i = 0; i < m; i++)
        {
            ex0[i] = (float)rnd.NextDouble(); ey0[i] = (float)rnd.NextDouble();
            esp[i] = 0.4f + (float)rnd.NextDouble(); eph[i] = (float)rnd.NextDouble() * 6.28f;
        }
    }

    // ---------- helpers ----------
    static float Clamp01(float x) { return x < 0f ? 0f : (x > 1f ? 1f : x); }
    static float Sm(float a, float b, float x) { float t = Clamp01((x - a) / (b - a)); return t * t * (3f - 2f * t); }
    static float Lerp(float a, float b, float t) { return a + (b - a) * t; }
    static float Sin(float x) { return (float)Math.Sin(x); }
    static float Cos(float x) { return (float)Math.Cos(x); }
    static float Exp(float x) { return (float)Math.Exp(x); }
    static float Sqrt(float x) { return (float)Math.Sqrt(x); }
    static float Pow(float a, float b) { return (float)Math.Pow(a, b); }
    static float Frac(float x) { return x - (float)Math.Floor(x); }
    static float OutBack(float p) { float c1 = 1.4f, c3 = c1 + 1f; float q = p - 1f; return 1f + c3 * q * q * q + c1 * q * q; }

    static void Bl(int x, int y, float r, float g, float b, float a)
    {
        if (x < 0 || y < 0 || x >= W || y >= H || a <= 0f) return;
        if (a > 1f) a = 1f;
        int i = y * W + x;
        fr[i] += (r - fr[i]) * a; fg[i] += (g - fg[i]) * a; fb[i] += (b - fb[i]) * a;
    }
    static void Ad(int x, int y, float r, float g, float b)
    {
        if (x < 0 || y < 0 || x >= W || y >= H) return;
        int i = y * W + x;
        fr[i] += r; fg[i] += g; fb[i] += b;
    }
    static void Backdrop(float gr, float gg, float gb, float amt, float cyf)
    {
        float hy = H * cyf, ix = 1f / (W * 0.5f), iy = 1f / (H * 0.5f);
        for (int y = 0; y < H; y++)
        {
            float dy = (y - hy) * iy;
            for (int x = 0; x < W; x++)
            {
                float dx = (x - W * 0.5f) * ix;
                float d2 = dx * dx + dy * dy;
                float g = amt * Exp(-d2 * 2.4f);
                int i = y * W + x;
                fr[i] = 0.006f + gr * g; fg[i] = 0.007f + gg * g; fb[i] = 0.011f + gb * g;
            }
        }
    }
    static void Stars(float t, float alpha)
    {
        for (int k = 0; k < stx.Length; k++)
        {
            int x = (int)(stx[k] * W), y = (int)(sty[k] * H);
            float tw = 0.55f + 0.45f * Sin(t * (0.8f + stp[k] * 2f) + stp[k] * 20f);
            float v = stb[k] * tw * alpha;
            Ad(x, y, v * 0.8f, v * 0.9f, v);
        }
    }
    static void Embers(float t, float alpha, float cr, float cg, float cb)
    {
        for (int k = 0; k < ex0.Length; k++)
        {
            float y = H - Frac(ey0[k] + t * 0.05f * esp[k]) * H;
            float x = ex0[k] * W + Sin(t * esp[k] + eph[k]) * 3f;
            float fl = 0.45f + 0.55f * Sin(t * 5f * esp[k] + eph[k]);
            float edge = Sm(0f, 0.15f, y / H) * Sm(1f, 0.85f, y / H);
            float v = alpha * fl * edge * (0.5f + 0.5f * esp[k] * 0.7f);
            int xi = (int)x, yi = (int)y;
            Ad(xi, yi, cr * v, cg * v, cb * v);
            Ad(xi + 1, yi, cr * v * 0.25f, cg * v * 0.25f, cb * v * 0.25f);
            Ad(xi - 1, yi, cr * v * 0.25f, cg * v * 0.25f, cb * v * 0.25f);
            Ad(xi, yi + 1, cr * v * 0.25f, cg * v * 0.25f, cb * v * 0.25f);
            Ad(xi, yi - 1, cr * v * 0.25f, cg * v * 0.25f, cb * v * 0.25f);
        }
    }

    // ---------- text overlay ----------
    static string Sp(string s)
    {
        StringBuilder o = new StringBuilder();
        for (int i = 0; i < s.Length; i++)
        {
            if (s[i] == ' ') o.Append("  "); else o.Append(s[i]).Append(' ');
        }
        return o.ToString().TrimEnd();
    }
    static void Put(int row, int col, char ch, float r, float g, float b, float a)
    {
        if (ch == ' ' || a <= 0.01f) return;
        Ov o = new Ov(); o.row = row; o.col = col; o.ch = ch; o.r = r * a; o.g = g * a; o.b = b * a;
        ovs.Add(o);
    }
    static void TextC(int row, string s, float r, float g, float b, float a)
    {
        int c0 = (Cols - s.Length) / 2;
        for (int i = 0; i < s.Length; i++) Put(row, c0 + i, s[i], r, g, b, a);
    }

    // ---------- renderer ----------
    static void Raw(string s)
    {
        byte[] bt = utf8.GetBytes(s); so.Write(bt, 0, bt.Length); so.Flush();
    }
    static void Alloc()
    {
        Cols = Math.Max(30, Console.WindowWidth); Rows = Math.Max(12, Console.WindowHeight);
        W = Math.Min(Cols, 150);
        Hc = Math.Min(Rows - 1, 42);
        H = Hc * 2;
        OX = (Cols - W) / 2; OY = (Rows - Hc) / 2;
        fr = new float[W * H]; fg = new float[W * H]; fb = new float[W * H];
        pt = new int[W * Hc]; pb = new int[W * Hc];
        for (int i = 0; i < pt.Length; i++) { pt[i] = -1; pb[i] = -1; }
        Raw(ESC + "[0m" + ESC + "[2J");
    }
    static int Q(float v)
    {
        int i = (int)(v * Fade * 255f + 0.5f);
        if (i < 0) i = 0; if (i > 255) i = 255;
        return i & 0xFC;
    }
    static void SetFg(int c) { sb.Append("\u001b[38;2;").Append((c >> 16) & 255).Append(';').Append((c >> 8) & 255).Append(';').Append(c & 255).Append('m'); }
    static void SetBg(int c) { sb.Append("\u001b[48;2;").Append((c >> 16) & 255).Append(';').Append((c >> 8) & 255).Append(';').Append(c & 255).Append('m'); }
    static void Render()
    {
        sb.Length = 0;
        int cF = -2, cB = -2;
        for (int cy = 0; cy < Hc; cy++)
        {
            bool move = true;
            for (int cx = 0; cx < W; cx++)
            {
                int i1 = (2 * cy) * W + cx, i2 = i1 + W, ci = cy * W + cx;
                int t = (Q(fr[i1]) << 16) | (Q(fg[i1]) << 8) | Q(fb[i1]);
                int b = (Q(fr[i2]) << 16) | (Q(fg[i2]) << 8) | Q(fb[i2]);
                if (t == pt[ci] && b == pb[ci]) { move = true; continue; }
                pt[ci] = t; pb[ci] = b;
                if (move) { sb.Append(ESC).Append('[').Append(OY + cy + 1).Append(';').Append(OX + cx + 1).Append('H'); move = false; }
                if (t == b) { if (cB != b) { SetBg(b); cB = b; } sb.Append(' '); }
                else
                {
                    if (cF != t) { SetFg(t); cF = t; }
                    if (cB != b) { SetBg(b); cB = b; }
                    sb.Append('\u2580');
                }
            }
        }
        for (int k = 0; k < ovs.Count; k++)
        {
            Ov o = ovs[k];
            int cx = o.col - OX, cy = o.row;
            int bg = 0;
            if (cx >= 0 && cx < W && cy >= 0 && cy < Hc)
            {
                int ci = cy * W + cx, t = pt[ci], b = pb[ci];
                if (t >= 0 && b >= 0)
                {
                    int rr = (((t >> 16) & 255) + ((b >> 16) & 255)) / 2;
                    int gg = (((t >> 8) & 255) + ((b >> 8) & 255)) / 2;
                    int bb = ((t & 255) + (b & 255)) / 2;
                    bg = (rr << 16) | (gg << 8) | bb;
                }
                pt[ci] = -1;
            }
            int fgc = (Q(o.r) << 16) | (Q(o.g) << 8) | Q(o.b);
            sb.Append(ESC).Append('[').Append(OY + cy + 1).Append(';').Append(o.col + 1).Append('H');
            SetFg(fgc); SetBg(bg); sb.Append(o.ch);
        }
        sb.Append(ESC).Append("[0m");
        Raw(sb.ToString());
    }
    static bool KeyHit()
    {
        try { if (Console.KeyAvailable) { Console.ReadKey(true); return true; } } catch (Exception) { }
        return false;
    }
    static bool Run(float dur, Scene draw)
    {
        Stopwatch sw = Stopwatch.StartNew();
        while (true)
        {
            float t = (float)sw.Elapsed.TotalSeconds;
            if (t >= dur) break;
            if (Console.WindowWidth != Cols || Console.WindowHeight != Rows) Alloc();
            Fade = 1f; ovs.Clear();
            draw(t);
            Render();
            if (KeyHit()) return false;
            int wait = 33 - (int)((sw.Elapsed.TotalSeconds - t) * 1000.0);
            if (wait > 1) Thread.Sleep(wait);
        }
        return true;
    }

    public static void Play()
    {
        try
        {
            try { IntPtr h = GetStdHandle(-11); uint m; if (GetConsoleMode(h, out m)) SetConsoleMode(h, m | 0x0004); } catch (Exception) { }
            so = Console.OpenStandardOutput();
            try { Console.TreatControlCAsInput = true; } catch (Exception) { }
            InitRand();
            Raw(ESC + "[?1049h" + ESC + "[?25l" + ESC + "[2J");
            Alloc();
            if (Run(7.0f, Intro) && Run(8.4f, EyeScene) && Run(12.6f, RoseScene)) Run(360000f, HeartScene);
        }
        finally
        {
            try { Raw(ESC + "[0m" + ESC + "[?25h" + ESC + "[?1049l"); } catch (Exception) { }
            try { Console.TreatControlCAsInput = false; } catch (Exception) { }
        }
    }

    // =====================================================
    //  SCENE 1 : INTRO
    // =====================================================
    static void Intro(float t)
    {
        Backdrop(0.03f, 0.05f, 0.08f, 0.6f * Sm(0f, 2f, t), 0.45f);
        Stars(t, Sm(0f, 1.5f, t));
        int tr = Hc / 2 - 3;
        int py = tr * 2 + 1;
        float len = W * 0.5f * Sm(0.4f, 2.4f, t);
        float inten = 0.9f * Sm(0.4f, 1.4f, t) * (1f - 0.8f * Sm(2.4f, 3.6f, t));
        for (int x = 0; x < W; x++)
        {
            float dx = Math.Abs(x - W / 2f);
            if (dx > len) continue;
            float fall = Pow(1f - dx / len, 1.6f);
            for (int dy = -4; dy <= 4; dy++)
            {
                float v = inten * fall * Exp(-dy * dy * 0.5f);
                Ad(x, py + dy, v * 0.55f, v * 0.78f, v);
            }
        }
        string title = Sp("HS AI SOLUTIONS");
        int c0 = (Cols - title.Length) / 2;
        float rev = (t - 1.5f) / 0.09f;
        float sw = (t - 3.0f) / 1.4f * title.Length * 1.3f - 6f;
        for (int i = 0; i < title.Length; i++)
        {
            float la = Clamp01(rev - i);
            float sh = Exp(-((i - sw) * (i - sw)) / 18f);
            Put(tr, c0 + i, title[i], Lerp(0.70f, 1f, sh), Lerp(0.74f, 1f, sh), Lerp(0.80f, 1f, sh), la);
        }
        TextC(tr + 2, Sp("PRESENTS"), 0.50f, 0.55f, 0.62f, Sm(3.4f, 4.4f, t));
        TextC(tr + 5, Sp("A CINEMATIC EXPERIENCE"), 0.25f, 0.80f, 0.95f, Sm(4.5f, 5.5f, t));
        Fade = 1f - Sm(6.3f, 7.0f, t);
    }

    // =====================================================
    //  SCENE 2 : THE EYE  (+ zoom into the pupil)
    // =====================================================
    static void EyeScene(float t)
    {
        Backdrop(0.0f, 0.22f, 0.30f, 0.55f * Sm(0.2f, 1.6f, t) * (1f - Sm(5f, 6.5f, t)), 0.5f);
        Stars(t, 0.5f * (1f - Sm(4.5f, 5.5f, t)));
        float cx = W * 0.5f, cy = H * 0.5f;
        float a = Math.Min(W * 0.40f, H * 1.15f);
        float b = a * 0.36f;
        float open = Sm(0.25f, 1.7f, t);
        if (t > 3.6f && t < 3.98f) open *= 1f - 0.96f * Sin((t - 3.6f) / 0.38f * PI);
        float look = Sm(1.6f, 2.6f, t) * (1f - Sm(4.2f, 5.0f, t));
        float lx = a * 0.16f * Sin(t * 1.25f) * look;
        float ly = b * 0.10f * Sin(t * 0.9f + 1f) * look;
        float ri = b * 0.92f;
        float rp = ri * (0.36f + 0.05f * Sin(t * 2.1f) + 0.30f * Sm(5.0f, 7.0f, t));
        float z = 1f;
        if (t > 5.0f) { float p = Math.Min(1f, (t - 5.0f) / 3.2f); z = Exp(4.4f * Pow(p, 1.5f)); }
        float pcx = cx + lx, pcy = cy + ly;
        float glintK = 1f - Sm(5.0f, 5.7f, t);
        float aa = 1f / z;
        float glowK = Clamp01(open * 1.5f) * (1f - Sm(5.0f, 6.0f, t));
        for (int y = 0; y < H; y++)
        {
            for (int x = 0; x < W; x++)
            {
                float ex = pcx + (x - cx) / z, ey = pcy + (y - cy) / z;
                float u = (ex - cx) / a, dy = ey - cy, au = Math.Abs(u);
                float lidK = au < 1f ? Pow(1f - u * u, 0.85f) : 0f;
                float lh = b * open * lidK * (dy < 0f ? 1.0f : 0.72f);
                float edge = lh - Math.Abs(dy);
                int i = y * W + x;
                // soft cyan aura around the lids
                float dist = au < 1f ? -edge : Sqrt(((au - 1f) * a) * ((au - 1f) * a) + dy * dy);
                if (edge <= 0f && glowK > 0f)
                {
                    float gl = 0.42f * glowK * Exp(-dist * z * 0.20f);
                    fr[i] += 0.05f * gl; fg[i] += 0.62f * gl; fb[i] += 0.85f * gl;
                }
                if (edge > 0f)
                {
                    float ia0 = Clamp01(edge / aa);
                    float depth = Math.Abs(dy) / Math.Max(lh, 0.001f);
                    float shade = 1f - 0.62f * Pow(depth, 2.2f) - 0.40f * Pow(au, 3f);
                    if (dy < 0f) shade -= 0.18f * Sm(0.3f, 1f, depth);
                    if (shade < 0.05f) shade = 0.05f;
                    float sr = 0.80f * shade, sg = 0.85f * shade, sb2 = 0.93f * shade;
                    float ddx = ex - pcx, ddy = ey - pcy;
                    float d = Sqrt(ddx * ddx + ddy * ddy);
                    float ia = Clamp01((ri - d) / aa);
                    if (ia > 0f)
                    {
                        float nd = d / ri;
                        float ang = (float)Math.Atan2(ddy, ddx);
                        float fib = 0.5f + 0.5f * Sin(ang * 26f + 2.2f * Sin(ang * 5f + nd * 3f));
                        float fib2 = 0.5f + 0.5f * Sin(ang * 71f + nd * 9f);
                        float f = 0.55f * fib + 0.45f * fib2;
                        float q = (nd - 0.52f) * 7f;
                        float ring = Exp(-q * q);
                        float rr = Lerp(0.02f, 0.22f, f * (1f - 0.5f * nd)) + 0.25f * ring;
                        float rg = Lerp(0.16f, 0.78f, f * (1f - 0.35f * nd)) + 0.35f * ring;
                        float rb = Lerp(0.30f, 0.95f, f * (1f - 0.2f * nd)) + 0.30f * ring;
                        float limb = 1f - 0.78f * Sm(0.78f, 1.0f, nd);
                        float ish = Clamp01(shade * 1.15f + 0.2f);
                        rr *= limb * ish; rg *= limb * ish; rb *= limb * ish;
                        float qq = (d - rp) / (ri * 0.12f);
                        float pg = 0.55f * Exp(-qq * qq);
                        rg += pg * 0.7f; rb += pg; rr += pg * 0.1f;
                        float pa = Clamp01((rp - d) / aa);
                        rr = Lerp(rr, 0f, pa); rg = Lerp(rg, 0f, pa); rb = Lerp(rb, 0f, pa);
                        float gx = ex - (pcx - ri * 0.36f), gy = ey - (pcy - ri * 0.40f);
                        float g1 = Clamp01((ri * 0.15f - Sqrt(gx * gx + gy * gy)) / aa) * glintK;
                        float hx = ex - (pcx + ri * 0.30f), hy = ey - (pcy + ri * 0.30f);
                        float g2 = Clamp01((ri * 0.07f - Sqrt(hx * hx + hy * hy)) / aa) * 0.6f * glintK;
                        float gm = Math.Max(g1, g2);
                        rr = Lerp(rr, 1f, gm); rg = Lerp(rg, 1f, gm); rb = Lerp(rb, 1f, gm);
                        sr = Lerp(sr, rr, ia); sg = Lerp(sg, rg, ia); sb2 = Lerp(sb2, rb, ia);
                    }
                    fr[i] = Lerp(fr[i], sr, ia0); fg[i] = Lerp(fg[i], sg, ia0); fb[i] = Lerp(fb[i], sb2, ia0);
                }
                if (au < 1f)
                {
                    float rq = edge * z;
                    float rim = Exp(-rq * rq / 1.4f) * Clamp01((1f - au) * 6f) * Clamp01(open * 3f) * (1f - Sm(5f, 5.8f, t));
                    fr[i] += 0.35f * rim; fg[i] += 0.85f * rim; fb[i] += 1.0f * rim;
                }
            }
        }
        float ca = Sm(2.0f, 3.0f, t) * (1f - Sm(4.6f, 5.2f, t));
        TextC(Hc - 4, Sp("LOOK CLOSER"), 0.35f, 0.85f, 1f, ca);
    }

    // =====================================================
    //  SCENE 3 : DARKNESS -> RED ROSE -> burst
    // =====================================================
    static float StemX(float y, float rcx, float yTop) { return rcx + 3.2f * Sin((y - yTop) * 0.07f); }

    static void Leaf(float cx, float cy, float ang, float len, float wid, float grow)
    {
        if (grow <= 0.01f) return;
        len *= grow; wid *= grow;
        float ca = Cos(ang), sa = Sin(ang);
        int ext = (int)len + 2;
        for (int y = (int)cy - ext; y <= (int)cy + ext; y++)
            for (int x = (int)cx - ext; x <= (int)cx + ext; x++)
            {
                float dx = x - cx, dy = y - cy;
                float uu = (dx * ca + dy * sa) / len, vv = (-dx * sa + dy * ca) / wid;
                float d = uu * uu + vv * vv;
                if (d >= 1f) continue;
                float lens = 1f - Math.Abs(uu) * 0.0f;
                float al = Clamp01((1f - d) * 5f);
                float tt = (uu + 1f) * 0.5f;
                float rr = Lerp(0.04f, 0.10f, tt), gg = Lerp(0.28f, 0.58f, tt), bb = Lerp(0.12f, 0.22f, tt);
                if (Math.Abs(vv) < 0.07f) { rr += 0.08f; gg += 0.15f; bb += 0.06f; }
                if (vv < 0f) { rr *= 1.2f; gg *= 1.2f; bb *= 1.2f; }
                Bl(x, y, rr * lens, gg * lens, bb * lens, al);
            }
    }

    static void RoseHead(float rcx, float rcy, float Rc, float rot, float alpha)
    {
        const float N = 3f, K = 1.8f;
        int ext = (int)(Rc * 1.6f) + 2;
        for (int y = (int)rcy - ext; y <= (int)rcy + ext; y++)
            for (int x = (int)rcx - ext; x <= (int)rcx + ext; x++)
            {
                if (x < 0 || y < 0 || x >= W || y >= H) continue;
                float dx = x - rcx, dy = y - rcy;
                float r = Sqrt(dx * dx + dy * dy);
                float th = (float)Math.Atan2(dy, dx);
                float rn = r / Rc;
                float s1 = th * N / (2f * PI) + K * (float)Math.Log(1.04f) + rot;
                float fe = Frac(s1);
                float edgeMul = 0.88f + 0.13f * fe;
                float ia = Clamp01((edgeMul - rn) * Rc) * alpha;
                if (ia <= 0f) continue;
                float s = th * N / (2f * PI) + K * (float)Math.Log(rn + 0.05f) + rot;
                float f = Frac(s);
                float lit = Pow(f, 1.4f);
                float shade = 0.30f + 0.70f * lit;
                shade += 0.35f * Sm(0.88f, 0.99f, f);
                shade *= 0.55f + 0.45f * Sm(0.0f, 0.12f, f);
                shade *= 0.30f + 0.70f * Sm(0.04f, 0.5f, rn);
                shade *= 1f - 0.28f * Sm(0.78f, 1.0f, rn);
                float hl = Sm(0.9f, 0.99f, f) * 0.18f;
                float rr = 0.90f * shade + hl + 0.04f, gg = 0.035f * shade + hl * 0.25f, bb = 0.08f * shade + 0.015f + hl * 0.2f;
                float cw = 1f - Sm(0.03f, 0.16f, rn);
                rr = Lerp(rr, 0.22f, cw); gg = Lerp(gg, 0.0f, cw); bb = Lerp(bb, 0.03f, cw);
                Bl(x, y, rr, gg, bb, ia);
            }
    }

    static void RoseScene(float t)
    {
        float gl = Sm(0.3f, 3f, t) * 0.6f + 0.4f * Sm(4f, 6f, t);
        Backdrop(0.38f, 0.02f, 0.05f, gl * 0.45f, 0.42f);
        Embers(t, Sm(0.2f, 1.4f, t) * 0.9f, 1f, 0.25f, 0.10f);

        float rcx = W * 0.5f, rcy = H * 0.34f;
        float R = Math.Min(W * 0.19f, H * 0.28f);
        float yTop = rcy + R * 0.5f, yBot = H * 0.97f;
        float g = Sm(0f, 1f, Clamp01((t - 1.3f) / 3.0f));
        float yEnd = Lerp(yBot, yTop, g);
        float expl = Sm(10.4f, 12.3f, t);
        float vis = 1f - Sm(10.6f, 11.8f, t);

        if (g > 0f && vis > 0f)
        {
            for (int y = (int)yEnd; y < (int)yBot; y++)
            {
                float xc = StemX(y, rcx, yTop);
                for (int x = (int)xc - 2; x <= (int)xc + 2; x++)
                {
                    float cov = Clamp01(1.5f - Math.Abs(x - xc));
                    float lightSide = x < xc ? 1.25f : 0.8f;
                    Bl(x, y, 0.05f * lightSide, 0.40f * lightSide, 0.16f * lightSide, cov * vis);
                }
            }
            float[] thf = new float[] { 0.30f, 0.52f, 0.74f };
            for (int k = 0; k < 3; k++)
            {
                float ty = yTop + (yBot - yTop) * thf[k];
                if (ty < yEnd) continue;
                float xc = StemX(ty, rcx, yTop);
                int dir = (k % 2 == 0) ? 1 : -1;
                for (int j = 0; j < 3; j++) Bl((int)(xc + dir * (2 + j)), (int)(ty - j), 0.12f, 0.45f, 0.18f, vis);
            }
            float y1 = yTop + (yBot - yTop) * 0.42f, y2 = yTop + (yBot - yTop) * 0.64f;
            float l1 = Sm(0f, 1f, Clamp01((y1 - yEnd) / (R * 1.4f))) * vis;
            float l2 = Sm(0f, 1f, Clamp01((y2 - yEnd) / (R * 1.4f))) * vis;
            float sx1 = StemX(y1, rcx, yTop), sx2 = StemX(y2, rcx, yTop);
            float LL = R * 0.95f, LW = R * 0.30f;
            Leaf(sx1 + Cos(-0.55f) * LL * l1 * 0.95f, y1 + Sin(-0.55f) * LL * l1 * 0.95f, -0.55f, LL, LW, l1);
            Leaf(sx2 + Cos(PI + 0.55f) * LL * l2 * 0.95f, y2 + Sin(PI + 0.55f) * LL * l2 * 0.95f, PI + 0.55f, LL, LW, l2);
        }

        float p = Clamp01((t - 4.0f) / 3.2f);
        if (p > 0f && vis > 0f)
        {
            float grow = OutBack(p);
            float Rc = R * grow * (1f + 0.22f * expl);
            // red halo behind the bloom
            float halo = 0.5f * Sm(0f, 1f, p) * vis;
            for (int y = (int)(rcy - Rc * 2.2f); y <= (int)(rcy + Rc * 2.2f); y++)
                for (int x = (int)(rcx - Rc * 2.2f); x <= (int)(rcx + Rc * 2.2f); x++)
                {
                    float dx = x - rcx, dy = y - rcy;
                    float q = Sqrt(dx * dx + dy * dy) / Math.Max(Rc, 1f);
                    if (q < 2.2f) { float hv = halo * Exp(-q * q * 0.9f) * 0.35f; Ad(x, y, hv, hv * 0.03f, hv * 0.06f); }
                }
            RoseHead(rcx, rcy, Rc, t * 0.16f, vis);
        }

        // falling petals
        if (t > 7.4f)
        {
            for (int k = 0; k < 16; k++)
            {
                float sp = 2.0f + (k % 5) * 0.9f;
                float yy = Frac((k * 0.137f) + (t - 7.4f) * sp / H * 1.0f) * H;
                float xx = ((k * 37) % 100) / 100f * W + Sin(t * 0.9f + k) * 6f;
                float al = Sm(7.4f, 8.4f, t) * vis;
                for (int j = 0; j < 3; j++) { Bl((int)xx + j, (int)yy, 0.85f, 0.05f, 0.10f, 0.9f * al); Bl((int)xx + j, (int)yy + 1, 0.55f, 0.02f, 0.06f, 0.9f * al); }
            }
        }

        // burst of sparks
        if (expl > 0f)
        {
            for (int k = 0; k < 140; k++)
            {
                float ang = k * 2.39996f;
                float spd = (0.25f + 0.75f * Frac(k * 0.618f)) * R * 3.4f;
                float rad = spd * (float)Math.Pow(expl, 0.6);
                float x = rcx + Cos(ang) * rad, y = rcy + Sin(ang) * rad * 0.8f + expl * expl * 6f;
                float v = (1f - expl) * (0.6f + 0.4f * Sin(k + t * 20f));
                Ad((int)x, (int)y, 1.0f * v, 0.22f * v, 0.10f * v);
                Ad((int)x + 1, (int)y, 0.4f * v, 0.05f * v, 0.02f * v);
            }
        }

        float c1 = Sm(4.6f, 5.6f, t) * (1f - Sm(7.2f, 7.8f, t));
        TextC(Hc - 5, Sp("A RED ROSE"), 1f, 0.35f, 0.40f, c1);
        float c2 = Sm(7.6f, 8.6f, t) * (1f - Sm(11.0f, 11.8f, t));
        TextC(Hc - 5, Sp("BEAUTY IN DARKNESS"), 0.95f, 0.40f, 0.45f, c2);
        Fade = 1f - Sm(12.1f, 12.6f, t);
    }

    // =====================================================
    //  SCENE 4 : THE BEATING HEART (loops)
    // =====================================================
    static float Wave(float c)
    {
        if (c < 0.10f) return 0f;
        if (c < 0.20f) return 0.12f * Sin((c - 0.10f) / 0.10f * PI);
        if (c < 0.30f) return 0f;
        if (c < 0.34f) return -0.15f * Sin((c - 0.30f) / 0.04f * PI);
        if (c < 0.40f) { float k = (c - 0.34f) / 0.06f; return k < 0.5f ? k * 2f : (1f - k) * 2f; }
        if (c < 0.44f) return -0.35f * Sin((c - 0.40f) / 0.04f * PI);
        if (c < 0.52f) return 0f;
        if (c < 0.72f) return 0.22f * Sin((c - 0.52f) / 0.20f * PI);
        return 0f;
    }

    static void HeartScene(float t)
    {
        const float T = 1.1f;
        float ph = t % T;
        float pulse = Exp(-Pow((ph - 0.06f) / 0.07f, 2f)) + 0.7f * Exp(-Pow((ph - 0.30f) / 0.09f, 2f));
        float app = t < 1.2f ? OutBack(Clamp01(t / 1.1f)) : 1f;
        float vis = Sm(0f, 0.7f, t);
        Backdrop(0.40f, 0.02f, 0.06f, (0.30f + 0.25f * pulse) * vis, 0.40f);
        Embers(t, 0.8f * vis, 1f, 0.2f, 0.12f);

        float cx = W * 0.5f, cy = H * 0.40f;
        float s0 = Math.Min(W * 0.17f, H * 0.26f);
        float s = s0 * (1f + 0.10f * pulse) * app;
        if (s < 1f) s = 1f;
        float gs = 0.45f + 0.8f * pulse;

        int ext = (int)(s * 2.6f);
        for (int y = (int)cy - ext; y <= (int)cy + ext; y++)
            for (int x = (int)cx - ext; x <= (int)cx + ext; x++)
            {
                if (x < 0 || y < 0 || x >= W || y >= H) continue;
                float X = (x - cx) / s, Y = -(y - cy) / s + 0.15f;
                float a = X * X + Y * Y - 1f;
                float f = a * a * a - X * X * Y * Y * Y;
                if (f > 0f)
                {
                    float gl = Exp(-f * 3.2f) * gs * 0.55f * vis;
                    if (gl > 0.004f) Ad(x, y, gl, gl * 0.08f, gl * 0.14f);
                }
                else
                {
                    float d = Clamp01(-f);
                    float aaE = Clamp01(Pow(-f, 0.34f) * s * 0.55f);
                    float ink = 0.60f + 0.40f * Pow(d, 0.45f);
                    ink += 0.18f * (-X * 0.5f + Y * 0.5f);
                    float hx = X + 0.52f, hy = Y - 0.58f;
                    float hl = Exp(-(hx * hx * 14f + hy * hy * 30f)) * 0.55f;
                    float rim = Exp(-d * 9f) * 0.25f;
                    float rr = 0.95f * ink + hl + rim + 0.05f * pulse;
                    float gg = 0.05f * ink + hl * 0.55f + rim * 0.3f;
                    float bb = 0.10f * ink + hl * 0.55f + rim * 0.25f;
                    Bl(x, y, rr * vis, gg * vis, bb * vis, aaE);
                }
            }

        // expanding heartbeat ring
        if (ph > 0.06f)
        {
            float q = (ph - 0.06f) / 0.95f;
            float rad = s0 * (1.05f + 1.7f * q);
            float al = (1f - q) * 0.5f * vis;
            int steps = (int)(rad * 9f);
            for (int k = 0; k < steps; k++)
            {
                float an = k * 2f * PI / steps;
                Ad((int)(cx + Cos(an) * rad * 1.15f), (int)(cy + 0.12f * s0 + Sin(an) * rad * 0.95f), al * 0.8f, al * 0.05f, al * 0.12f);
            }
        }

        // ECG trace
        float yb = H * 0.80f, amp = H * 0.07f;
        float prevY = yb;
        for (int x = 0; x < W; x++)
        {
            float sc = x / (float)W * 3f + (t - 0.06f) / T + 0.37f;
            float yy = yb - Wave(Frac(sc)) * amp;
            float v = (0.12f + 0.88f * Pow(x / (float)W, 1.6f)) * vis;
            int y0 = (int)Math.Min(prevY, yy), y1 = (int)Math.Max(prevY, yy);
            for (int y = y0; y <= y1; y++) Ad(x, y, v * 1.0f, v * 0.14f, v * 0.22f);
            Ad(x, y0 - 1, v * 0.15f, v * 0.02f, v * 0.04f);
            Ad(x, y1 + 1, v * 0.15f, v * 0.02f, v * 0.04f);
            prevY = yy;
        }

        float tp = 0.75f + 0.25f * Sin(t * 2.2f);
        TextC(Hc - 4, Sp("FOREVER IN MOTION"), 1f, 0.40f, 0.45f, Sm(1.2f, 2.4f, t) * tp);
        TextC(Hc - 2, "HS  AI  SOLUTIONS", 0.55f, 0.58f, 0.64f, Sm(2.0f, 3.2f, t) * 0.8f);
        TextC(Hc - 1, "press any key to exit", 0.28f, 0.30f, 0.34f, Sm(4.0f, 5.0f, t));
    }
}

'@

if (-not ("HSCinema" -as [type])) { Add-Type -TypeDefinition $src -Language CSharp }

try {
    [HSCinema]::Play()
}
finally {
    [Console]::ResetColor()
    try { [Console]::CursorVisible = $true } catch {}
    Write-Host ""
    Write-Host "          H S   A I   S O L U T I O N S   |   END OF EXPERIENCE" -ForegroundColor DarkRed
    Write-Host ""
}
