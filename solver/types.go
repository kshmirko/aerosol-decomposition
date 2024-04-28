package solver

import (
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// SolutionType - структура - содержит результат расчетов
type SolutionType struct {
	Mix  components.AerosolModeMix
	Xsol utlis.Vector
	Yh   utlis.Vector
	Err  float64
}

type Solutions []SolutionType

func NewSolutions(n int) Solutions {
	ret := make(Solutions, 0, n)
	return ret
}

func (st SolutionType) Print(title string) {

	fmt.Printf("%10s %10.1f %10.2f %7s %7s %7s %10.3f %10.3f %10.3f %10.3f\n",
		title, st.Err*100, st.Xsol, st.Mix[0].Title, st.Mix[1].Title, st.Mix[2].Title,
		st.Mix.MeanRadius(), st.Mix.EffectiveRadius(), st.Mix.RefrReIdx()[1], st.Mix.RefrImIdx()[1],
	)

}

func (st SolutionType) GetOptDb() components.OpticalDB {
	db := make(components.OpticalDB, 0, len(st.Mix))
	for _, m := range st.Mix {
		db = append(db, m.OpticalCoefs)
	}
	return db
}
