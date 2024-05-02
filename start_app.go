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

func main() {
	inpfile := flag.String("in", "data.txt", "Имя файла с данными")
	algorithm := flag.String("alg", "denm", "Название используемого алгоритма spso или de или denm. \nspso - метод роя частиц, \nde - метод дифференциальной эволюции, \ndenm - метод дифференциальной эфолюции + уточнение симплекс-методом.")
	mustlog := flag.Bool("log", false, "Показывать лог работы алгоритма?")
	use_aggls := flag.Bool("aggls", false, "Использовать агломераты обломков для моделирования минерального аэрозоля?")
	keysfile := flag.String("keys", "KEYS.txt", "Файл с наименованием исспользуемых компонентов")
	plot_psd := flag.Bool("psdplot", false, "Создавать графики функции распределения")
	plot_test := flag.Bool("testplot", false, "Создавать графики сравнения измеренных данных и восстановленных")
	dep_scale := flag.Float64("dep-scale-fact", 1.0, "Весовой коэффициент. Изменяет роль деполяризации в суммарной невязке.\nПараметр может принимать значения на отрезке [0.0, 1.0]")
	//rel_hum := flag.Int("hum", 0, "Относительная влажность для компонентов")
	flag.Parse()

	var db components.OpticalDB
	var err error

	//if db, err = components.LoadDB(*dbfile); err != nil {
	if *use_aggls {
		db = components.GenerateDBAggl()
	} else {
		db = components.GenerateDB()
	}
	//}

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
	} else {
		log.Fatal("Неизвстный алгоритм")
	}
	fmt.Printf("\nВ качестве решателя используется алгоритм %s\n", *algorithm)

	// Выводим информацию об используемых типах
	db.PrintTable()

	keys, err := components.LoadKeys(*keysfile)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Filter(keys)
	// Загружаем информацию об измерениях
	meas, err := measurements.LoadFromFile(*inpfile)

	if err != nil {
		log.Fatal(err)
	}

	meas.Print()

	avg := meas.MakeAverage()
	avg.Print1()
	sols := solver.NewSolutions(meas.Len() + 1)

	fmt.Printf("     Title       Err  [        X1         X2         X3]      C1      C2      C3      Rmean       Reff        Mre        Mim\n")
	fmt.Printf("     -----       ---  [        --         --         --]      --      --      --      -----       ----        ---        ---\n")
	sols = DoSolve(avg, sol, db, mustlog, dep_scale, sols)

	_ = *plot_psd
	if *plot_test {
		plots.PlotY(avg.Data, sols[0].Yh, "#pts", "f(x)", "Optical coefs", avg.Title+".pdf")
	}

	new_db := db
	R, _ := utlis.LogSpace(0.005, 15.0, 30)
	for i, mi := range meas {

		sols = DoSolve(mi, sol, new_db, mustlog, dep_scale, sols)
		if *plot_test {
			plots.Scatter(mi.Data, sols[i+1].Yh, "#pts", "f(x)", "Optical coefs", mi.Title+".pdf")
		}

		if *plot_psd {
			Y := sols[i+1].Mix.ValueVol(R)
			plots.PlotXY(R, Y, "Radius, um", "dV/dr", "Volume distribution", "psd-"+mi.Title+".pdf")
		}
	}

	fmt.Println()
	meas = append(measurements.Measurements{avg}, meas...)
	//fmt.Printf("%8s ", "Average")
	fmt.Println("Исходные данные:")
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
	//fmt.Printf("%8s ", "Average")
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

	sols = append(sols, ret)
	ret.Print(mi.Title)
	return sols
}
