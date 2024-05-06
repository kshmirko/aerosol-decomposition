package components

import (
	"fmt"
	"os"
	"slices"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	"github.com/olekukonko/tablewriter"
)

// Описывает оптические коэффициенты одной моды вместе в параметрами распределения
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

	return am.OpticalCoefs.MRe
}

// RefrImIdx - мнимая часть коэффициента преломления
func (am AerosolModeMixItem) RefrImIdx() utlis.Vector {
	return am.OpticalCoefs.MIm
}

func (am AerosolModeMixItem) Number() float64 {
	return am.N
}

// PrintTable - Печатает таблицу с данными аэрозольных компонентов в стантартное устройство вывода
func (db OpticalDB) PrintTable() {

	tbl := tablewriter.NewWriter(os.Stdout)
	tbl.SetHeader([]string{"#", "Title", "RH", "Reff", "Rmean", "MRe", "Mim"})
	tbl.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tbl.SetCenterSeparator("|")

	tbl.SetAlignment(tablewriter.ALIGN_RIGHT)
	for i, dbi := range db {
		tbl.Append([]string{
			fmt.Sprintf("%4d", i),
			dbi.Title,
			fmt.Sprintf("%d", dbi.Rh),
			fmt.Sprintf("%.4f", dbi.AerosolMode.EffectiveRadius()),
			fmt.Sprintf("%.4f", dbi.AerosolMode.MeanRadius()),
			fmt.Sprintf("%.3f", dbi.MRe[1]),
			fmt.Sprintf("%.4f", dbi.MIm[1]),
		})
	}
	tbl.SetCaption(true, "База данных")
	tbl.Render()
	fmt.Println()
}

func (db OpticalDB) Filter(keys ElemKeys) OpticalDB {
	ret := make(OpticalDB, 0, len(keys))

	for _, k := range keys {
		idx := slices.IndexFunc(db, func(key OpticalCoefs) bool {
			return (key.Title == k.Title) && (key.Rh == k.Rh)
		})
		if idx != -1 {
			ret = append(ret, db[idx])
		}
	}

	return ret

}
