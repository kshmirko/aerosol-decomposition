package components

import (
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// AerosolModeMix - тип аэрозольной смеси
type AerosolModeMix []AerosolModeMixItem

// MeanRadius - возвращает средний радиус аэрозольной смеси
func (am AerosolModeMix) MeanRadius() float64 {
	total := 0.0
	rmean := 0.0
	for i := range am {
		rmean += am[i].MeanRadius() * am[i].N
		total += am[i].N
	}
	rmean /= total
	return total
}

// Area -  возвращает площадь поверхности частиц аэрозольного распределения
func (am AerosolModeMix) Area() float64 {
	area := 0.0
	for i := range am {
		area += am[i].Area()
	}
	return area
}

// Volume - возвращает объем частиц аэрозольного распределения
func (am AerosolModeMix) Volume() float64 {
	volume := 0.0
	for i := range am {
		volume += am[i].Volume()
	}

	return float64(volume)
}

// EffectiveRadius - возвращает эффективный радиус аэрозольного распределения
func (am AerosolModeMix) EffectiveRadius() float64 {
	total := 0.0
	rmean := 0.0
	for i := range am {
		rmean += am[i].EffectiveRadius() * am[i].N
		total += am[i].N
	}
	rmean /= total
	return total
}

// Value - возвращает функцию распределения по заданной смеси
func (am AerosolModeMix) Value(r []float64) []float64 {
	ret := make([]float64, len(r))

	for _, v := range am {
		tmp := v.Value(r)
		for j := range r {
			ret[j] += tmp[j]
		}
	}

	return ret
}

func (am AerosolModeMix) Ext() Vector {
	ret := make(Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utlis.Add(ret, v.Extinction())
	}
	return ret
}

func (am AerosolModeMix) Bck() Vector {
	ret := make(Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utlis.Add(ret, v.Backscatter())
	}
	return ret
}

func (am AerosolModeMix) B22() Vector {
	ret := make(Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utlis.Add(ret, v.Backscatter22())
	}
	return ret
}
