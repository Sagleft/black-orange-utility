package main

import (
	"math/rand"
)

const (
	machineMaxParts  = 100
	machineMaxCycles = 50

	partMaxTeeth    = 100
	partMaxEyes     = 100
	partMaxDiameter = 10.5
	partMaxPower    = 1000
)

// WeirdMachin do weird things!
type WeirdMachin struct {
	Cycles int
}

// WeirdGear - part of a mad machine
type WeirdGear struct {
	NumberOfTeeth int
	NumberOfEyes  int
	Diameter      float32
	Power         float32
}

// WeirdResult - wtf is this?
type WeirdResult struct{}

func newWeirdMachine() *WeirdMachin {
	machine := WeirdMachin{
		Cycles: rand.Intn(machineMaxCycles),
	}

	partsCount := rand.Intn(machineMaxParts)
	madParts := map[int]*WeirdGear{}
	for partIndex := 0; partIndex < partsCount; partIndex++ {
		madParts[partIndex] = &WeirdGear{
			NumberOfTeeth: rand.Intn(partMaxTeeth),
			NumberOfEyes:  rand.Intn(partMaxEyes),
			Diameter:      rand.Float32() * partMaxDiameter,
			Power:         rand.Float32() * partMaxPower,
		}
	}

	return &machine
}

func (machine *WeirdMachin) work() *WeirdResult {
	// TODO
	return nil
}
