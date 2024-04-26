package components

import (
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// Описывает оптические коэффициенты обной моды вместе в параметрами распределения
type OpticalCoefs struct {
	AerosolMode `json:"aerosol_mode,omitempty"`
	Rh          int          `json:"rh"`
	Wvl         utlis.Vector `json:"wvl,omitempty"`
	Ext         utlis.Vector `json:"ext,omitempty"`
	Bck         utlis.Vector `json:"bck,omitempty"`
	B22         utlis.Vector `json:"b_22,omitempty"`
	MRe         utlis.Vector `json:"m_re,omitempty"`
	MIm         utlis.Vector `json:"m_im,omitempty"`
}

type OpticalDB []OpticalCoefs

// AerosolModeMixItem - элемент смеси частиц
type AerosolModeMixItem struct {
	OpticalCoefs `json:"optical_coefs,omitempty"`
	N            float64 `json:"n,omitempty"`
}

// Value  вычисление функции распределения по отдельному компоненту
func (am AerosolModeMixItem) Value(r utlis.Vector) utlis.Vector {
	ret := am.AerosolMode.Value(r)
	ret = utlis.Scale(am.N, 0, ret)
	return ret
}

// MeanRadius вычисление среднего радиуса по отдельному компоненту
func (am AerosolModeMixItem) MeanRadius() float64 {
	return am.AerosolMode.MeanRadius()
}

// Area - вычисление площади поверхности  отдельного компонента
func (am AerosolModeMixItem) Area() float64 {
	return am.N * am.AerosolMode.Area()
}

// Volume - вычисление объема отдельного компонента
func (am AerosolModeMixItem) Volume() float64 {
	return am.N * am.AerosolMode.Volume()
}

// EffectiveRadius - вычисление эффективного радиуса  отдельного компонента
func (am AerosolModeMixItem) EffectiveRadius() float64 {
	return am.AerosolMode.EffectiveRadius()
}

// Extinction - коэффициенты ослабления для отдельного компонента
func (am AerosolModeMixItem) Extinction() utlis.Vector {
	ret := make(utlis.Vector, DIM_SIZE)
	for i, ext := range am.Ext {
		ret[i] = am.N * ext
	}
	return ret
}

// Backscatter - коэффициенты обратного рассеяния для отдельного компонента
func (am AerosolModeMixItem) Backscatter() utlis.Vector {
	ret := make(utlis.Vector, DIM_SIZE)
	for i, bck := range am.Bck {
		ret[i] = am.N * bck
	}
	return ret
}

// Backscatter22 - B22 для отдельного компонента
func (am AerosolModeMixItem) Backscatter22() utlis.Vector {
	ret := make(utlis.Vector, DIM_SIZE)
	for i, b22 := range am.B22 {
		ret[i] = am.N * b22
	}
	return ret
}

// RefrReIdx - действительная часть коэффициента преломления
func (am AerosolModeMixItem) RefrReIdx() utlis.Vector {
	ret := make(utlis.Vector, DIM_SIZE)
	vol := am.Volume()
	for i, m := range am.MRe {
		ret[i] = vol * m
	}
	return ret
}

// RefrImIdx - мнимая часть коэффициента преломления
func (am AerosolModeMixItem) RefrImIdx() utlis.Vector {
	ret := make(utlis.Vector, DIM_SIZE)
	vol := am.Volume()
	for i, m := range am.MIm {
		ret[i] = vol * m
	}
	return ret
}

func (db OpticalDB) PrintTable() {
	fmt.Printf("%4s %10s %10s %10s %7s %7s\n", "#", "Title", "Reff", "Rmean", "MRe", "MIm")
	fmt.Printf("%4s %10s %10s %10s %7s %7s\n", "-", "-----", "----", "-----", "---", "---")

	for i, dbi := range db {
		fmt.Printf("%4d %10s %10.4f %10.4f %7.3f %7.4f\n", i, dbi.Title, dbi.AerosolMode.EffectiveRadius(), dbi.AerosolMode.MeanRadius(), dbi.MRe[1], dbi.MIm[1])
	}
}
