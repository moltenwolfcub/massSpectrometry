package main

const (
	DT   Second = 0.001              //Seconds per tick
	PXPM int    = 100                //pixels per metre
	L    Metre  = 1100 / Metre(PXPM) //length of drift region (m)
)

type Second float64

func (s Second) ToTick() Tick {
	return Tick(s / DT)
}

type Tick int

func (t Tick) ToSecond() Second {
	return Second(t) * DT
}

type Metre float64

func (m Metre) ToPixel() Pixel {
	return Pixel(float64(m) * float64(PXPM))
}

type Pixel int

func (p Pixel) ToMetre() Metre {
	return Metre(float64(p) / float64(PXPM))
}
