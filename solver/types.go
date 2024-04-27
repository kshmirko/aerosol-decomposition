package solver

import (
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
