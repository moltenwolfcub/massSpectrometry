package main

const (
	DT   Second  = 0.001                //Seconds per tick
	PXPM int     = 100                  //pixels per metre
	L    float64 = float64(1100 / PXPM) //length of drift region (m)
)

type Second float64

func (s Second) ToTick() Tick {
	return Tick(s / DT)
}

type Tick int

func (t Tick) ToSecond() Second {
	return Second(t) * DT
}
