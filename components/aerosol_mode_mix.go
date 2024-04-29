package components

import (
	"fmt"
	"log"
	"math"

	//"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	utils "gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// AerosolModeMix - тип аэрозольной смеси
type AerosolModeMix []AerosolModeMixItem

func (amx *AerosolModeMix) SetCoefs(x utils.Vector) {
	if len(*amx) != len(x) {
		log.Fatal("Number of components in mixtuer and length of the x array should be equal")
	}

	for i := range x {
		(*amx)[i].N = x[i]
	}
}

// MeanRadius - возвращает средний радиус аэрозольной смеси
func (am AerosolModeMix) MeanRadius() float64 {
	total := 0.0
	rmean := 0.0
	for i := range am {
		rmean += am[i].MeanRadius() * am[i].N
		total += am[i].N
	}
	rmean /= total
	return rmean
}

// Area -  возвращает площадь поверхности частиц аэрозольного распределения
func (am AerosolModeMix) Area() float64 {
	area := 0.0
	for i := range am {
		area += am[i].Area()
	}
	return area
}

func (am AerosolModeMix) Number() float64 {
	number := 0.0
	for i := range am {
		number += am[i].Number()
	}
	return number
}

// Volume - возвращает объем частиц аэрозольного распределения
func (am AerosolModeMix) Volume() float64 {
	volume := 0.0
	for i := range am {
		volume += am[i].Volume()
	}

	return volume
}

// EffectiveRadius - возвращает эффективный радиус аэрозольного распределения
func (am AerosolModeMix) EffectiveRadius() float64 {

	// total := 0.0
	// rmean := 0.0
	// for i := range am {
	// 	rmean += am[i].EffectiveRadius() * am[i].N
	// 	total += am[i].N
	// }
	// rmean /= total
	return 3 * am.Volume() / am.Area()
}

// RefrReIdx - действительная часть показателя преломления смеси
func (am AerosolModeMix) RefrReIdx() utils.Vector {
	ret := make(utils.Vector, len(am[0].Ext))

	vol := am.Volume()
	for _, ami := range am {
		utils.AddVScale(ret, ami.RefrReIdx(), ami.Volume()/vol)
	}
	return ret
}

// RefrImIdx - мнимая часть показателя преломления смеси
func (am AerosolModeMix) RefrImIdx() utils.Vector {
	ret := make(utils.Vector, len(am[0].Ext))

	vol := am.Volume()
	for _, ami := range am {
		utils.AddVScale(ret, ami.RefrImIdx(), ami.Volume()/vol)
	}
	return ret
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

func (am AerosolModeMix) ValueVol(r []float64) []float64 {
	ret := make([]float64, len(r))
	psd := am.Value(r)

	for j := range r {
		ret[j] = psd[j] * 4.189 * math.Pow(r[j], 3.0)
	}

	return ret
}

// Ext - коэффициенты ослабления смеси
func (am AerosolModeMix) Ext() utils.Vector {
	ret := make(utils.Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utils.Add(ret, v.Extinction())
	}
	return ret
}

// Bck -  коэффициеты обратного рассеяния смеси
func (am AerosolModeMix) Bck() utils.Vector {
	ret := make(utils.Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utils.Add(ret, v.Backscatter())
	}
	return ret
}

// B22 - коэффициенты b22 смеси
func (am AerosolModeMix) B22() utils.Vector {
	ret := make(utils.Vector, len(am[0].Ext))
	for _, v := range am {
		ret = utils.Add(ret, v.Backscatter22())
	}
	return ret
}

// F - Целевая функция подгона
func (am AerosolModeMix) F(x utils.Vector) utils.Vector {
	if len(am) != len(x) {
		log.Fatal("Количество компонентов смеси должно быть равно числу параметров x")
	}
	for i := range x {
		am[i].N = x[i]
	}
	b := am.Bck()
	e := am.Ext()
	b22 := am.B22()
	re := am.RefrReIdx()
	im := am.RefrImIdx()
	dep := utils.CalcDep(b, b22)

	ret := utils.Vector{b[0], b[1], b[2], e[0], e[1], dep[1], re[1], im[1]}
	return ret
}

func (am AerosolModeMix) PrintComponents() {
	for i := range am {
		fmt.Printf("Comp #%2d = %7s\t", i, am[i].Title)
	}
	fmt.Println()
}
