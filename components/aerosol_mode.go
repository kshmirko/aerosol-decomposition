package components

import (
	"math"
)

// Lognormal size distribution
// f(x) = 1/(sqrt(2pi)*x*ln(s))exp(-0.5*(ln(x)-ln(xm))**2/(ln(s)**2))
const DIM_SIZE int = 3

type RVector3D [DIM_SIZE]float64
type CVector3D [DIM_SIZE]complex128
type Vector []float64
type CVector []complex128

// AerosolMode - описывает форму распределения по размерам
type AerosolMode struct {
	Title string  `json:"title,omitempty"`
	Sigma float64 `json:"sigma,omitempty"`
	Rm    float64 `json:"rm,omitempty"`
}

func (am AerosolMode) MeanRadius() float64 {
	return am.Rm * math.Exp(0.5*math.Pow(math.Log(am.Sigma), 2))
}

func (am AerosolMode) Area() float64 {
	return 4 * math.Pi * math.Pow(am.Rm, 2) * math.Exp(2*math.Pow(math.Log(am.Sigma), 2))
}

func (am AerosolMode) Volume() float64 {
	return 4.0 / 3.0 * math.Pi * math.Pow(am.Rm, 3) * math.Exp(4.5*math.Pow(math.Log(am.Sigma), 2))
}

func (am AerosolMode) EffectiveRadius() float64 {
	return am.Rm * math.Exp(2.5*math.Pow(math.Log(am.Sigma), 2))
}

func (am AerosolMode) Value(r []float64) []float64 {
	ret := make([]float64, len(r))
	A := 1.0 / (math.Sqrt(math.Pi*2) * math.Log(am.Sigma))
	for i := range r {

		B := -0.5 * math.Pow(math.Log(r[i])-math.Log(am.Rm), 2) / math.Pow(math.Log(am.Sigma), 2)
		ret[i] = A / r[i] * math.Exp(B)
	}
	return ret
}
