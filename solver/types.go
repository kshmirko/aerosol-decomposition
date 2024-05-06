package solver

import (
	"fmt"
	"os"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	"github.com/olekukonko/tablewriter"
)

// SolutionType - структура - содержит результат расчетов
type SolutionType struct {
	TitleSol string
	Mix      components.AerosolModeMix
	Xsol     utlis.Vector
	Xfrac    utlis.Vector
	Vfrac    utlis.Vector
	Yh       utlis.Vector
	Err      float64
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

func (st *SolutionType) MakeVFractions() {

	(*st).Vfrac = make(utlis.Vector, len(st.Xsol))
	total := 0.0
	for _, x := range (*st).Mix {
		total += x.Volume()
	}

	for i, x := range (*st).Mix {
		st.Vfrac[i] = x.Volume() / total * 100
	}
}

func (st SolutionType) Print() {

	st.MakeFractions()
	st.MakeVFractions()
	fmt.Printf("%10s %10.1f %10.2f %5.2f %7s %7s %7s %10.3f %10.3f %10.3f %10.3f %10.2f %10.2f\n",
		st.TitleSol, st.Err*100, st.Xsol, st.Xfrac, st.Mix[0].Title, st.Mix[1].Title, st.Mix[2].Title,
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

func (st *Solutions) Print() {
	tbl := tablewriter.NewWriter(os.Stdout)
	tbl.SetHeader([]string{
		"Title",
		"Err",
		"X1", "X2", "X3",
		"X1%", "X2%", "X3%",
		"C1", "C2", "C3",
		"Rmean", "Reff", "Mre", "MIm", "V1%", "V2%", "V3%", "Vol", "Area",
	})
	tbl.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tbl.SetCenterSeparator("|")

	for _, sti := range *st {
		sti.MakeFractions()
		sti.MakeVFractions()
		tbl.Append([]string{
			sti.TitleSol,
			fmt.Sprintf("%.1f", sti.Err*100.0),
			fmt.Sprintf("%.2f", sti.Xsol[0]),
			fmt.Sprintf("%.2f", sti.Xsol[1]),
			fmt.Sprintf("%.2f", sti.Xsol[2]),
			fmt.Sprintf("%.2f", sti.Xfrac[0]),
			fmt.Sprintf("%.2f", sti.Xfrac[1]),
			fmt.Sprintf("%.2f", sti.Xfrac[2]),
			sti.Mix[0].Title,
			sti.Mix[1].Title,
			sti.Mix[2].Title,
			fmt.Sprintf("%.3f", sti.Mix.MeanRadius()),
			fmt.Sprintf("%.3f", sti.Mix.EffectiveRadius()),
			fmt.Sprintf("%.3f", sti.Mix.RefrReIdx()[1]),
			fmt.Sprintf("%.3f", sti.Mix.RefrImIdx()[1]),
			fmt.Sprintf("%.3f", sti.Vfrac[0]),
			fmt.Sprintf("%.3f", sti.Vfrac[1]),
			fmt.Sprintf("%.3f", sti.Vfrac[2]),
			fmt.Sprintf("%.3f", sti.Mix.Volume()),
			fmt.Sprintf("%.3f", sti.Mix.Area()),
		})
	}
	tbl.SetCaption(true, "Результаты расчетов")
	tbl.Render()
	fmt.Println()
}
