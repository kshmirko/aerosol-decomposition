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
	//dbfile := flag.String("dbfile", "db.json", "Имя json файла с базой данных оптических свойств частиц")
	inpfile := flag.String("in", "data.txt", "Имя файла с данными")
	algorithm := flag.String("alg", "spso", "Название используемого алгоритма spso или de или denm. \nspso - метод роя частиц, \nde - метод дифференциальной эволюции, \ndenm - метод дифференциальной эфолюции + уточнение симплекс-методом.")
	mustlog := flag.Bool("log", false, "Показывать лог работы алгоритма?")
	use_aggls := flag.Bool("aggls", false, "Использовать агломераты обломков для моделирования минерального аэрозоля?")
	keysfile := flag.String("keys", "KEYS.txt", "Файл с наименованием исспользуемых компонентов")
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
	log.Printf("\nВ качестве решателя используется алгоритм %s\n", *algorithm)

	// Выводим информацию об используемых типах
	db.PrintTable()

	keys, err := components.LoadKeys(*keysfile)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Filter(keys)
	fmt.Println(len(keys))
	// Загружаем информацию об измерениях
	meas, err := measurements.LoadFromFile(*inpfile)

	if err != nil {
		log.Fatal(err)
	}

	meas.Print()

	avg := meas.MakeAverage()
	avg.Print1()
	sols := solver.NewSolutions(meas.Len() + 1)

	sols = DoSolve(avg, sol, db, mustlog, sols)

	plots.PlotY(avg.Data, sols[0].Yh, "#pts", "f(x)", "Optical coefs", "Average.pdf")
	new_db := db
	//sols[0].GetOptDb()

	for _, mi := range meas {

		sols = DoSolve(mi, sol, new_db, mustlog, sols)
	}

	fmt.Println("Len = ", len(sols))
}

func DoSolve(mi measurements.Measurement, sol func(db *components.OpticalDB, y0 utlis.Vector, mustlog bool) solver.SolutionType, db components.OpticalDB, mustlog *bool, sols solver.Solutions) solver.Solutions {
	//fmt.Printf("\nОбработка данных с нзванием %s\n", mi.Title)

	ret := sol(&db, mi.Data, *mustlog)
	sols = append(sols, ret)
	ret.Print(mi.Title)
	// fmt.Printf("Resulting error: %.2f%%\n", ret.Err*100.0)
	// fmt.Printf("X = %.2f\n", ret.Xsol)
	// ret.Mix.PrintComponents()

	// fmt.Printf("Rmean = %.3f, Reff = %.3f, Mre = %7.2f, Mim=%8.3f\n",
	// 	ret.Mix.MeanRadius(),
	// 	ret.Mix.EffectiveRadius(),
	// 	ret.Mix.RefrReIdx()[1],
	// 	ret.Mix.RefrImIdx()[1])
	return sols
}
