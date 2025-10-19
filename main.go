package main

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"slices"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ElectricField struct {
	Rect                Rect
	PotentialDifference float64
}

func (e ElectricField) FieldStrength() Vec2 {
	E := e.PotentialDifference / e.Rect.Width() //Uniform Electric Field Strength
	return Vec2{E, 0}
	//points from + to - and I've set that to the useful way of my sim
}

func (e ElectricField) Draw(screen *ebiten.Image) {
	accelRegion := ebiten.NewImage(int(e.Rect.Width()*float64(PXPM)), int(e.Rect.Height()*float64(PXPM)))
	accelRegion.Fill(color.RGBA{250, 50, 50, 100})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(e.Rect.Min.X*float64(PXPM)), float64(e.Rect.Min.Y*float64(PXPM)))

	screen.DrawImage(accelRegion, &drawOps)

	OVERHANG := 30.0
	WIDTH := 20.0
	COLOR := color.RGBA{50, 50, 50, 255}

	platePlus := ebiten.NewImage(int(WIDTH), int(e.Rect.Height()*float64(PXPM)+OVERHANG))
	platePlus.Fill(COLOR)

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(e.Rect.Min.X*float64(PXPM))-WIDTH, float64(e.Rect.Min.Y*float64(PXPM))-OVERHANG/2)

	screen.DrawImage(platePlus, &ops)

	GAP := 50.0

	plateNeg := ebiten.NewImage(int(WIDTH), int(e.Rect.Height()*float64(PXPM)+OVERHANG))
	part := ebiten.NewImage(int(WIDTH), int(e.Rect.Height()*float64(PXPM)+OVERHANG)/2-int(GAP)/2)
	part.Fill(COLOR)
	partOp := ebiten.DrawImageOptions{}
	plateNeg.DrawImage(part, &partOp)
	partOp.GeoM.Translate(0, (e.Rect.Height()*float64(PXPM)+OVERHANG)/2+GAP)
	plateNeg.DrawImage(part, &partOp)

	ops = ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(e.Rect.Max.X*float64(PXPM)), float64(e.Rect.Min.Y*float64(PXPM))-OVERHANG/2)

	screen.DrawImage(plateNeg, &ops)

	TEXT_OFFSET := 5.0
	TEXT_SIZE := 38.0

	w, _ := text.Measure("+", &text.GoTextFace{
		Source: fontSource,
		Size:   TEXT_SIZE,
	}, 0)

	op := &text.DrawOptions{}
	op.GeoM.Translate(
		float64(e.Rect.Min.X*float64(PXPM))-WIDTH+(WIDTH-w)/2,
		float64(e.Rect.Min.Y*float64(PXPM)-OVERHANG/2+e.Rect.Height()*float64(PXPM)+OVERHANG+TEXT_OFFSET),
	)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "+", &text.GoTextFace{
		Source: fontSource,
		Size:   TEXT_SIZE,
	}, op)

	w, _ = text.Measure("—", &text.GoTextFace{
		Source: fontSource,
		Size:   TEXT_SIZE,
	}, 0)

	op = &text.DrawOptions{}
	op.GeoM.Translate(
		float64(e.Rect.Max.X*float64(PXPM))+(WIDTH-w)/2,
		float64(e.Rect.Min.Y*float64(PXPM)-OVERHANG/2+e.Rect.Height()*float64(PXPM)+OVERHANG+TEXT_OFFSET),
	)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "—", &text.GoTextFace{
		Source: fontSource,
		Size:   TEXT_SIZE,
	}, op)
}

type Detector struct {
	Rect              Rect
	AcellerationField ElectricField
	DataLogger        DataLogger
}

func (d *Detector) Update(molecules []*Molecule) {
	for _, m := range molecules {
		if d.Rect.Contains(m.Pos) {
			d.TakeReading(m)
		}
	}
}

func (d *Detector) TakeReading(molecule *Molecule) {
	molecule.Active = false
	z := molecule.Charge //simulate reading the charge
	molecule.Charge = 0
	t := float64(molecule.DriftTicks) * DT

	L := d.Rect.Min.X - d.AcellerationField.Rect.Max.X

	E := float64(z) * d.AcellerationField.PotentialDifference // Electrical energy
	v := L / t                                                // Constant velocity
	m := 2 * E / (v * v)                                      // Kinetic energy

	mpz := m / float64(z)

	// fmt.Println(molecule.Mass(), mpz)
	d.DataLogger.LogData(mpz)
}

func (d Detector) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(d.Rect.Width()*float64(PXPM)), int(d.Rect.Height()*float64(PXPM)))
	img.Fill(color.RGBA{60, 75, 75, 255})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(d.Rect.Min.X*float64(PXPM)), float64(d.Rect.Min.Y*float64(PXPM)))

	screen.DrawImage(img, &drawOps)
}

type LoggerEntry struct {
	mz        float64
	abundance int
}

type DataLogger struct {
	data         []LoggerEntry
	totalEntries int
}

func (d DataLogger) String() string {
	ordered := make([]LoggerEntry, len(d.data))
	copy(ordered, d.data)
	sort.Slice(ordered, func(i, j int) bool {
		return d.data[i].abundance < d.data[j].abundance
	})

	str := "================================\n"
	str += "Abundance: mass (amount)\n"
	for _, e := range ordered {
		str += fmt.Sprintf("|%-6d |%-6.6f|\n", e.abundance, e.mz)
	}
	str += "\nAbundance: mass (%)\n"
	for _, e := range ordered {
		percentage := float64(e.abundance) / float64(d.totalEntries) * 100
		str += fmt.Sprintf("|%-6.3f%% |%-6.6f|\n", percentage, e.mz)
	}
	return str
}

func (d *DataLogger) LogData(mz float64) {
	d.totalEntries++
	index := slices.IndexFunc(d.data, func(e LoggerEntry) bool {
		return e.mz == mz
	})

	if index >= 0 {
		d.data[index].abundance++
		return
	}

	d.data = append(d.data, LoggerEntry{
		mz:        mz,
		abundance: 1,
	})
}

type Simulation struct {
	selector Selector

	accelerationRegion ElectricField
	detector           Detector

	buttons []*Button

	molecules         []*Molecule
	drawableMolecules []RenderMolecule
}

func NewSimulation() *Simulation {
	s := &Simulation{
		accelerationRegion: ElectricField{
			Rect: NewRect(
				float64(100/PXPM), float64(150/PXPM),
				float64(300/PXPM), float64(750/PXPM),
			),
			PotentialDifference: 16_000,
		},
		detector: Detector{
			Rect: NewRect(
				float64(1400/PXPM), float64(150/PXPM),
				float64(1500/PXPM), float64(750/PXPM),
			),
			DataLogger: DataLogger{
				data:         make([]LoggerEntry, 0),
				totalEntries: 0,
			},
		},

		molecules:         []*Molecule{},
		drawableMolecules: []RenderMolecule{},
	}

	s.selector = NewSelector(NewRect(0, 0, 1600, 80), s)

	s.buttons = []*Button{
		{
			Text: "Ionise",
			TextColor: ButtonColor{
				Primary: color.White,
			},
			TextSize: 30,
			Rect:     NewRect(100, 800, 300, 850),
			ButtonColor: ButtonColor{
				Primary:   color.RGBA{0, 70, 25, 255},
				Hover:     color.RGBA{0, 63, 22, 255},
				Secondary: color.RGBA{0, 94, 35, 255},
			},
			Fuction:      s.IoniseMolecules,
			MaxClickTime: 10,
		},
		{
			Text: "Clean",
			TextColor: ButtonColor{
				Primary: color.White,
			},
			TextSize: 30,
			Rect:     NewRect(1350, 800, 1500, 850),
			ButtonColor: ButtonColor{
				Primary:   color.RGBA{0, 70, 25, 255},
				Hover:     color.RGBA{0, 63, 22, 255},
				Secondary: color.RGBA{0, 94, 35, 255},
			},
			Fuction:      s.CleanSimulation,
			MaxClickTime: 10,
		},
		{
			Text: "Get Output",
			TextColor: ButtonColor{
				Primary: color.White,
			},
			TextSize: 30,
			Rect:     NewRect(325, 800, 575, 850),
			ButtonColor: ButtonColor{
				Primary:   color.RGBA{0, 70, 25, 255},
				Hover:     color.RGBA{0, 63, 22, 255},
				Secondary: color.RGBA{0, 94, 35, 255},
			},
			Fuction:      s.GetOutput,
			MaxClickTime: 10,
		},
	}
	s.detector.AcellerationField = s.accelerationRegion

	return s
}

func (s *Simulation) IoniseMolecules() {
	for _, m := range s.molecules {
		if s.accelerationRegion.Rect.Contains(m.Pos) {
			m.Charge = 1
		}
	}
}

func (s *Simulation) CleanSimulation() {
	s.molecules = []*Molecule{}
	s.drawableMolecules = []RenderMolecule{}
}

func (s *Simulation) GetOutput() {
	fmt.Print(s.detector.DataLogger)
}

func (s Simulation) GetSpawn() Vec2 {
	x := float64(100) / float64(PXPM)

	TOP := 391.0

	pixY := 48 * rand.Float64()
	y := (TOP + pixY) / float64(PXPM)

	return Vec2{x, y}

}

func (s *Simulation) Update() error {
	s.selector.Update()

	for _, b := range s.buttons {
		b.Update()
	}

	activeMolecules := []*Molecule{}

	for _, m := range s.molecules {
		if m.Active {
			m.Update(s.accelerationRegion)
			activeMolecules = append(activeMolecules, m)
		}
	}

	s.detector.Update(activeMolecules)

	return nil
}

func (s Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 100, 120, 255})

	s.accelerationRegion.Draw(screen)
	s.detector.Draw(screen)

	for _, m := range s.drawableMolecules {
		m.Draw(screen)
	}

	s.selector.Draw(screen)

	for _, b := range s.buttons {
		b.Draw(screen)
	}
}

func (s Simulation) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Mass Spectrometry")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	sim := NewSimulation()
	if err := ebiten.RunGame(sim); err != nil {
		panic(err)
	}
}
