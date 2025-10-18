package main

const (
	DT   float64 = 0.001                //Seconds per tick
	PXPM int     = 100                  //pixels per metre
	L    float64 = 1100 / float64(PXPM) //length of drift region (m)
)
