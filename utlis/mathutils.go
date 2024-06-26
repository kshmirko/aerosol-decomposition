package utlis

import (
	"log"
	"math"
)

// LogSpace - создание логарифмически эквидистантного вектора диной N на интервале [r0;r1]
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

// Trapz - вычисление интеграла методом трапеций
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

func AddVScale(y, x []float64, a float64) {
	if len(y) != len(x) {
		log.Fatal("len(x) should be equal to len(y)")
	}

	for i := range y {
		y[i] = y[i] + (a * x[i])
	}
}

// CalcDep - вычисление деполяризационного отношения
func CalcDep(b11, b22 []float64) []float64 {
	if len(b11) != len(b22) {
		log.Fatalln("Should be len(b11)==len(b22)")
	}

	ret := make([]float64, len(b11))

	for i := range ret {
		ret[i] = (b11[i] - b22[i]) / (b11[i] + b22[i])
	}
	return ret
}

// CalcStatistics вычисляет статистические параметры на основе
// распределения числа частиц по размерам y(x)
// rmean, reff, vol
func CalcStatistics(x, y []float64) (float64, float64, float64) {
	sum := 0.0
	sum2 := 0.0

	for i := 1; i < len(x); i++ {
		sum += 0.5 * (y[i] + y[i-1]) * (x[i] - x[i-1])
		sum2 += 0.5 * (y[i]*x[i] + y[i-1]*x[i-1]) * (x[i] - x[i-1])
	}
	rmean := sum2 / sum

	vol := 0.0
	area := 0.0

	for i := 1; i < len(x); i++ {
		vol += 0.5 * (4.189*y[i]*math.Pow(x[i], 3) + 4.189*y[i-1]*math.Pow(x[i-1], 3)) * (x[i] - x[i-1])
		area += 0.5 * (12.57*y[i]*math.Pow(x[i], 2) + 12.57*y[i-1]*math.Pow(x[i-1], 2)) * (x[i] - x[i-1])
	}

	return rmean, vol / (3.0 * area), vol
}
