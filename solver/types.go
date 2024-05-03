package solver

import (
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// SolutionType - структура - содержит результат расчетов
type SolutionType struct {
	Mix   components.AerosolModeMix
	Xsol  utlis.Vector
	Xfrac utlis.Vector
	Yh    utlis.Vector
	Err   float64
}

type Solutions []SolutionType

func NewSolutions(n int) Solutions {
	ret := make(Solutions, 0, n)
	return ret
}

func (st *SolutionType) MakeFractions() {

	(*st).Xfrac = make(utlis.Vector, len(st.Xsol))
	total := 0.0
	for _, x := range (*st).Xsol {
		total += x
	}

	for i, x := range (*st).Xsol {
		st.Xfrac[i] = x / total * 100
	}
}

func (st SolutionType) Print(title string) {

	st.MakeFractions()
	fmt.Printf("%10s %10.1f %10.2f %5.2f %7s %7s %7s %10.3f %10.3f %10.3f %10.3f %10.2f %10.2f\n",
		title, st.Err*100, st.Xsol, st.Xfrac, st.Mix[0].Title, st.Mix[1].Title, st.Mix[2].Title,
		st.Mix.MeanRadius(), st.Mix.EffectiveRadius(), st.Mix.RefrReIdx()[1], st.Mix.RefrImIdx()[1],
		st.Mix.Volume(), st.Mix.Area(),
	)

}

func (st SolutionType) GetOptDb() components.OpticalDB {
	db := make(components.OpticalDB, 0, len(st.Mix))
	for _, m := range st.Mix {
		db = append(db, m.OpticalCoefs)
	}
	return db
}
