package models

import "gonum.org/v1/gonum/stat/distuv"

type PoissonDist struct {
}

func NewPoissonDist() *PoissonDist {
	return &PoissonDist{}
}

func (pd *PoissonDist) Generate(lambda float64) float64 {
	poisson := distuv.Poisson{Lambda: lambda, Src: nil}
	return poisson.Rand()
}
