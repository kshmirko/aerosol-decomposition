package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/measurements"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/plots"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/solver"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

const R_MIN = 0.005
const R_MAX = 15.0
const NPTS = 30

func main() {

	inpfile := flag.String("in", "data.txt", "Имя файла с данными")
	algorithm := flag.String("alg", "DENM", "Название используемого алгоритма spso или de или denm. "+
		"\nSPSO - метод роя частиц, "+
		"\nDE - метод дифференциальной эволюции, "+
		"\nDENM - метод дифференциальной эволюции + уточнение симплекс-методом, "+
		"\nDELBFGS - метод дифференциальной эволюции + уточнение LBFGS (более оптимально использует память),"+
		"\nDEBFGS - метод дифференциальной эволюции + уточнение BFGS"+
		"\nDEGD - метод дифференциальной эволюции + уточнение градиентным спуском")
	mustlog := flag.Bool("log", false, "Показывать лог работы алгоритма?")
	use_aggls := flag.Bool("aggls", false, "Использовать агломераты обломков для моделирования минерального аэрозоля?")
	keysfile := flag.String("keys", "KEYS.txt", "Файл с наименованием исспользуемых компонентов")
	plot_psd := flag.Bool("psdplot", false, "Создавать графики функции распределения")
	plot_test := flag.Bool("testplot", false, "Создавать графики сравнения измеренных данных и восстановленных")
	dep_scale := flag.Float64("dep-scale-fact", 1.0, "Весовой коэффициент. Изменяет роль деполяризации в суммарной невязке.\nПараметр может принимать значения на отрезке [0.0, 1.0]")
	no_b1064 := flag.Bool("no-b1064", false, "Не учитывать коэффициент обратного рассеяния на 1064 нм")
	flag.Parse()

	var db components.OpticalDB
	var err error

	if *use_aggls {
		db = components.GenerateDBAggl()
	} else {
		db = components.GenerateDB()
	}

	if *dep_scale < 0 {
		*dep_scale = 0.0
	} else if *dep_scale > 1.0 {
		*dep_scale = 1.0
	}

	// Определяем выбор алгоритма глобальной оптимизации
	sol := solver.FindSolution
	if strings.ToLower(*algorithm) == "spso" {
		sol = solver.FindSolution
	} else if strings.ToLower(*algorithm) == "de" {
		sol = solver.FindSolutionDE
	} else if strings.ToLower(*algorithm) == "denm" {
		sol = solver.FindSolutionDENM
	} else if strings.ToLower(*algorithm) == "delbfgs" {
		sol = solver.FindSolutionDELBFGS
	} else if strings.ToLower(*algorithm) == "debfgs" {
		sol = solver.FindSolutionDEBFGS
	} else if strings.ToLower(*algorithm) == "degd" {
		sol = solver.FindSolutionDEBFGS
	} else {
		log.Fatal("Неизвстный алгоритм, читайте внимательно документацию")
	}
	fmt.Printf("\nВ качестве решателя используется алгоритм %s\n\n", *algorithm)

	// Выводим информацию об используемых типах
	db.PrintTable()

	// Загружаем файл с компонентами
	keys, err := components.LoadKeys(*keysfile)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Filter(keys)

	// Загружаем информацию об измерениях
	meas, err := measurements.LoadFromFile(*inpfile)

	if *no_b1064 {
		for i := range meas {
			meas[i].Data[2] = 0.0
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	// Осредняем измерения
	avg := meas.MakeAverage()
	meas = append(measurements.Measurements{avg}, meas...)
	meas.Print()
	sols := solver.NewSolutions(meas.Len())

	//displayHeader()

	new_db := db
	R, _ := utlis.LogSpace(R_MIN, R_MAX, NPTS)
	for i, mi := range meas {

		sols = DoSolve(mi, sol, new_db, mustlog, dep_scale, sols)

		if *plot_test {
			plots.Scatter(mi.Data, sols[i].Yh, "#pts", "f(x)", "Optical coefs", mi.Title+".pdf")
		}

		if *plot_psd {
			Y := sols[i].Mix.ValueVol(R)
			plots.PlotXY(R, Y, "Radius, um", "dV/dr", "Volume distribution", "psd-"+mi.Title+".pdf")
		}
	}
	fmt.Println()
	sols.Print()

	fmt.Println("Исходные данные:")
	fmt.Println("----------------")
	for i := range meas {
		fmt.Printf("%8s  ", meas[i].Title)
	}
	fmt.Println()
	for i := range avg.Data {
		for j := range meas {
			fmt.Printf("%.2e  ", meas[j].Data[i])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Восстановленные данные:")
	fmt.Println("-----------------------")
	for i := range meas {
		fmt.Printf("%8s  ", meas[i].Title)
	}
	fmt.Println()
	for i := range avg.Data {
		for j := range sols {
			fmt.Printf("%.2e  ", sols[j].Yh[i])
		}
		fmt.Println()
	}

}

func DoSolve(mi measurements.Measurement,
	sol func(db *components.OpticalDB, y0 utlis.Vector, mustlog bool, dep_scale float64) solver.SolutionType,
	db components.OpticalDB,
	mustlog *bool,
	dep_scale *float64,
	sols solver.Solutions) solver.Solutions {

	ret := sol(&db, mi.Data, *mustlog, *dep_scale)
	ret.TitleSol = mi.Title

	sols = append(sols, ret)
	//ret.Print()
	return sols
}
