package solver

import (
	"log"
	"math"
	"sort"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	"github.com/MaxHalford/eaopt"
	"github.com/crhntr/neldermead"
	combinations "github.com/mxschmitt/golang-combinations"
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize"
)

const MAX_COMPS int = 3

func GetNormL2(yh, y0 utlis.Vector, dep_scale float64) float64 {
	tot := 0.0
	cnt := 0
	for i := range y0 {
		if i == 5 {
			tot += dep_scale * math.Pow((y0[i]-yh[i])/y0[i], 2)
			cnt += 1
		} else {
			if y0[i] > 1.0e-9 {
				tot += math.Pow((y0[i]-yh[i])/y0[i], 2)
				cnt += 1
			}
		}
	}

	return math.Sqrt(tot / float64(cnt))
}

func GetNormL1(yh, y0 utlis.Vector, dep_scale float64) float64 {
	tot := 0.0
	cnt := 0
	for i := range y0 {
		if i == 5 {
			tot += dep_scale * math.Abs((y0[i]-yh[i])/y0[i])
			cnt += 1
		} else {
			if y0[i] > 1.0e-9 {
				tot += math.Abs((y0[i] - yh[i]) / y0[i])
				cnt += 1
			}
		}
	}

	// tot := 0.0
	// for i := range y0 {
	// 	tot += math.Abs((y0[i] - yh[i]) / y0[i])
	// }
	return tot / float64(cnt)
}

func FindSolution(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {
		spso, err := eaopt.NewSPSO(200, 150, 0, 10000, 0.8, false, nil)
		if err != nil {
			log.Println(err)
		}

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
			return GetNormL2(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, err := spso.Minimize(F, uint(MAX_COMPS))
		score[i].Mix = mix
		score[i].Xsol = xsol
		score[i].Mix.SetCoefs(xsol)
		score[i].Err = yerr
		score[i].Yh = mix.F(xsol)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("%7.2f | %.2f \n", xsol, yerr*100)
		}
		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDE(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {
		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}

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
			return GetNormL2(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, err := spso.Minimize(F, uint(MAX_COMPS))
		score[i].Mix = mix
		score[i].Xsol = xsol
		score[i].Mix.SetCoefs(xsol)
		score[i].Err = yerr
		score[i].Yh = mix.F(xsol)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("%7.2f | %.2f \n", xsol, yerr*100)
		}
		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDENM(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {

		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}

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
			return GetNormL1(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, _ := spso.Minimize(F, uint(MAX_COMPS))

		opts := neldermead.NewOptions()
		opts.Constraints = []neldermead.Constraint{
			{
				Min: 0.0,
				Max: 1000000.0,
			},
			{
				Min: 0.0,
				Max: 1000000.0,
			},
			{
				Min: 0.0,
				Max: 1000000.0,
			},
		}

		xsol1, err := neldermead.Run(F, xsol, opts)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("Global solution: Err=%5.2f, %.2e\n", yerr*100, xsol)
			log.Printf("Refinement sol.: Err=%5.2f, %.2e\n", xsol1.F*100, xsol1.X)
		}
		score[i].Mix = mix
		score[i].Xsol = xsol1.X
		score[i].Mix.SetCoefs(xsol1.X)
		score[i].Err = xsol1.F
		score[i].Yh = mix.F(xsol1.X)

		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDELBFGS(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {

		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}

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
					penalty += int(1000 * math.Exp(-x[i]/100))
					x[i] = 0
				}
			}
			yh := mix.F(x)
			return GetNormL1(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, _ := spso.Minimize(F, uint(MAX_COMPS))

		p := optimize.Problem{
			Func: F,
			Grad: func(grad, x []float64) {
				fd.Gradient(grad, F, x, nil)
			},
			Hess: func(hess *mat.SymDense, x []float64) {
				fd.Hessian(hess, F, x, nil)
			},
		}

		var meth = &optimize.LBFGS{}
		xsol1, err := optimize.Minimize(p, xsol, nil, meth)

		// opts := neldermead.NewOptions()
		// opts.Constraints = []neldermead.Constraint{
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// }

		// xsol1, err := neldermead.Run(F, xsol, opts)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("Global solution: Err=%5.2f, %.2e\n", yerr*100, xsol)
			log.Printf("Refinement sol.: Err=%5.2f, %.2e\n", xsol1.F*100, xsol1.X)
		}
		score[i].Mix = mix
		score[i].Xsol = xsol1.X
		score[i].Mix.SetCoefs(xsol1.X)
		score[i].Err = xsol1.F
		score[i].Yh = mix.F(xsol1.X)

		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDEBFGS(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {

		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}

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
					penalty += int(1000 * math.Exp(-x[i]/100))
					x[i] = 0
				}
			}
			yh := mix.F(x)
			return GetNormL1(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, _ := spso.Minimize(F, uint(MAX_COMPS))

		p := optimize.Problem{
			Func: F,
			Grad: func(grad, x []float64) {
				fd.Gradient(grad, F, x, nil)
			},
			Hess: func(hess *mat.SymDense, x []float64) {
				fd.Hessian(hess, F, x, nil)
			},
		}

		var meth = &optimize.BFGS{}
		xsol1, err := optimize.Minimize(p, xsol, nil, meth)

		// opts := neldermead.NewOptions()
		// opts.Constraints = []neldermead.Constraint{
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// 	{
		// 		Min: 0.0,
		// 		Max: 1000000.0,
		// 	},
		// }

		// xsol1, err := neldermead.Run(F, xsol, opts)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("Global solution: Err=%5.2f, %.2e\n", yerr*100, xsol)
			log.Printf("Refinement sol.: Err=%5.2f, %.2e\n", xsol1.F*100, xsol1.X)
		}
		score[i].Mix = mix
		score[i].Xsol = xsol1.X
		score[i].Mix.SetCoefs(xsol1.X)
		score[i].Err = xsol1.F
		score[i].Yh = mix.F(xsol1.X)

		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}

func FindSolutionDEGD(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) SolutionType {
	combs := combinations.Combinations(*db, MAX_COMPS)

	score := make([]SolutionType, len(combs))
	for i, cmb := range combs {

		spso, err := eaopt.NewDiffEvo(40, 30, 0, 10000, 0.5, 0.2, false, nil)
		if err != nil {
			log.Println(err)
		}

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
					penalty += int(1000 * math.Exp(-x[i]/100))
					x[i] = 0
				}
			}
			yh := mix.F(x)
			return GetNormL1(yh, y0, dep_scale) + float64(penalty)
		}

		xsol, yerr, _ := spso.Minimize(F, uint(MAX_COMPS))

		p := optimize.Problem{
			Func: F,
			Grad: func(grad, x []float64) {
				fd.Gradient(grad, F, x, nil)
			},
			Hess: func(hess *mat.SymDense, x []float64) {
				fd.Hessian(hess, F, x, nil)
			},
		}

		var meth = &optimize.GradientDescent{}
		xsol1, err := optimize.Minimize(p, xsol, nil, meth)

		// xsol1, err := neldermead.Run(F, xsol, opts)
		if mustlog {
			log.Printf("Номер тройки: %d, %d\n", i, len(cmb))
			log.Printf("Global solution: Err=%5.2f, %.2e\n", yerr*100, xsol)
			log.Printf("Refinement sol.: Err=%5.2f, %.2e\n", xsol1.F*100, xsol1.X)
		}
		score[i].Mix = mix
		score[i].Xsol = xsol1.X
		score[i].Mix.SetCoefs(xsol1.X)
		score[i].Err = xsol1.F
		score[i].Yh = mix.F(xsol1.X)

		if err != nil {
			log.Println(err)
		}

	}

	sort.Slice(score, func(i, j int) bool {
		return score[i].Err < score[j].Err
	})

	return score[0]

}
