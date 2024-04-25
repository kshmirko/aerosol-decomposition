package utlis

import "math"

func LogSpace(r0, r1 float64, N int) ([]float64, float64) {
	r := make([]float64, N)
	lgr0 := math.Log10(r0)
	lgr1 := math.Log10(r1)
	dlgr := (lgr1 - lgr0) / float64(N-1.0)

	for i := range r {
		r[i] = math.Pow(10, lgr0+dlgr*float64(i))
	}

	return r, dlgr
}

func Trapz(x, y []float64) float64 {
	sum := 0.0
	for i := 0; i < len(x)-1; i++ {
		sum = sum + 0.5*(y[i]+y[i+1])*(x[i+1]-x[i])
	}
	return sum
}

func Scale(a, b float64, x []float64) []float64 {
	ret := make([]float64, len(x))
	for i := range ret {
		ret[i] = a*x[i] + b
	}
	return ret
}

func Add(a, b []float64) []float64 {
	ret := make([]float64, len(a))
	for i := range a {
		ret[i] = a[i] + b[i]
	}
	return ret
}
