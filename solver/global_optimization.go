package solver

import (
	"fmt"
	"log"
	"math"
	"sort"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	"github.com/MaxHalford/eaopt"
	combinations "github.com/mxschmitt/golang-combinations"
)

const MAX_COMPS int = 3

func GetNormL2(yh, y0 utlis.Vector) float64 {
	tot := 0.0
	for i := range y0 {
		tot += math.Pow((y0[i]-yh[i])/y0[i], 2)
	}
	return math.Sqrt(tot) / float64(len(y0))
}

func GetNormL1(yh, y0 utlis.Vector) float64 {
	tot := 0.0
	for i := range y0 {
		tot += math.Abs((y0[i] - yh[i]) / y0[i])
	}
	return tot / float64(len(y0))
}

func FindSolution(db *components.OpticalDB, y0 utlis.Vector) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)
	fmt.Println(len(combs), len(combs[0]))

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {
		spso, err := eaopt.NewSPSO(200, 150, 0, 10000, 0.8, false, nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Номер тройки: %d, %d\n", i, len(cmb))

		//x0 := make(utlis.Vector, MAX_COMPS)
		mix := components.AerosolModeMix{
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[0],
				N:            1.0,
			},
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[1],
				N:            1.0,
			},
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[2],
				N:            1.0,
			},
		}

		// Функция для минимизации
		F := func(x []float64) float64 {
			penalty := 0
			for i := range x {
				if x[i] < 0 {
					penalty += 1000
				}
			}
			yh := mix.F(x)
			return GetNormL2(yh, y0) + float64(penalty)
		}

		xsol, yerr, err := spso.Minimize(F, uint(MAX_COMPS))
		score[i].Mix = mix
		score[i].Xsol = xsol
		score[i].Mix.SetCoefs(xsol)
		score[i].Err = yerr
		score[i].Yh = mix.F(xsol)
		log.Printf("%7.2f | %.2f \n", xsol, yerr*100)
		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDE(db *components.OpticalDB, y0 utlis.Vector) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)
	fmt.Println(len(combs), len(combs[0]))

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {
		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Номер тройки: %d, %d\n", i, len(cmb))

		//x0 := make(utlis.Vector, MAX_COMPS)
		mix := components.AerosolModeMix{
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[0],
				N:            1.0,
			},
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[1],
				N:            1.0,
			},
			components.AerosolModeMixItem{
				OpticalCoefs: cmb[2],
				N:            1.0,
			},
		}

		// Функция для минимизации
		F := func(x []float64) float64 {
			penalty := 0
			for i := range x {
				if x[i] < 0 {
					penalty += 1000
				}
			}
			yh := mix.F(x)
			return GetNormL2(yh, y0) + float64(penalty)
		}

		xsol, yerr, err := spso.Minimize(F, uint(MAX_COMPS))
		score[i].Mix = mix
		score[i].Xsol = xsol
		score[i].Mix.SetCoefs(xsol)
		score[i].Err = yerr
		score[i].Yh = mix.F(xsol)
		log.Printf("%7.2f | %.2f \n", xsol, yerr*100)
		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}
